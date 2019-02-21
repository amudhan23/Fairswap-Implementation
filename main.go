package main

import(
	"fmt"
	kec "github.com/ethereum/go-ethereum/crypto"
	"os"
	"io"
	"log"
	"encoding/json"
	//"reflect"



  "context"
  "crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core"
	"crypto/rand"
  "github.com/ethereum/go-ethereum/common"
	"math/big"
   import_contracts "./contracts"
   //"reflect"
)

const NoOfInputGates=8
const BUFFER_SIZE=32
const TotalNumberOfgates=16
const MaxLinesToGate=2
const TotNumOfEncOutVecToMer=16


//these struct and SetCircuitTuples function is common to both sender and receiver
type circuit struct{
	Index int
	Operation int
	InputToTheGate []int
}

type sendertoreciever struct {
		Id int                                        //id
		Keycommit [32]byte														//key commitment
		EncodedOutputOfGates [TotNumOfEncOutVecToMer][32]byte		//Encoded output of gates
}

type sendertocontract struct{
	Id int
	Price big.Int
	Keycommit [32]byte
	MerkleRootOfCircuit [32]byte
	MerkleRootOfEncInp [32]byte
}

func main(){

	//reading from a file "a.txt" and storing bytearray of 32 bytes in inputvectors of size equal to NoOfInputGates
	var currentByte int64 = 0
	fileBuffer := make([]byte, BUFFER_SIZE)
	file, err := os.Open("a.txt") // For read access.
	if err != nil {
			log.Fatal(err)
	}
	var inputvectors [NoOfInputGates][32]byte
	count:=0
	for {
			n,err2 := file.ReadAt(fileBuffer, currentByte)
			currentByte += BUFFER_SIZE
			fileBuffer=fileBuffer[:n]
			if count==NoOfInputGates{
					break
			}
			for j:=0;j<n;j++{
				inputvectors[count][j]=fileBuffer[j]
			}
			if err2 == io.EOF {
					break
			}
			count++
	}
	file.Close()

	//address(sender,receiver) generation
	//sender private key generation
	privatekeysender,err := crypto.GenerateKey()
	authsender := bind.NewKeyedTransactor(privatekeysender)

	//reciever private key generation
	privatekeyreceiver,err:=crypto.GenerateKey()
	authreceiver :=bind.NewKeyedTransactor(privatekeyreceiver)

	//genesisalloc
	alloc := make(core.GenesisAlloc)
	alloc[authsender.From] = core.GenesisAccount{Balance: big.NewInt(1100000000000000000)}  //1.1 ether
	alloc[authreceiver.From] = core.GenesisAccount{Balance: big.NewInt(2500000000000000000)} //2.5 ether

	//client
	client := backends.NewSimulatedBackend(alloc,60000000000)

	//sender address generation
	publickeysender := privatekeysender.Public()
	publickeyECDSAsender, ok := publickeysender.(*ecdsa.PublicKey)
	if !ok {
			log.Fatal("error casting public key to ECDSA")
	}
	senderaddress := crypto.PubkeyToAddress(*publickeyECDSAsender)
	sendernonce, err := client.PendingNonceAt(context.Background(), senderaddress)
	if err != nil {
			log.Fatal(err)
	}
	authsender.Nonce = big.NewInt(int64(sendernonce))
	authsender.Value = big.NewInt(0)

	//receiveraddress generation
	publickeyreceiver := privatekeyreceiver.Public()
	publickeyECDSAreceiver, ok := publickeyreceiver.(*ecdsa.PublicKey)
	if !ok {
			log.Fatal("error casting public key to ECDSA")
	}
	receiveraddress := crypto.PubkeyToAddress(*publickeyECDSAreceiver)
	receivernonce, err := client.PendingNonceAt(context.Background(), receiveraddress)
	if err != nil {
			log.Fatal(err)
	}
	authreceiver.Nonce = big.NewInt(int64(receivernonce))
	authreceiver.Value = big.NewInt(0)

	//objects
	var SenToConObject sendertocontract
	var SenToRecObject sendertoreciever
	var circuitobjects [TotalNumberOfgates]circuit //array of circuit object

	//sets the circuit
	setCircuitTuples(&circuitobjects)

	/*sender fns call*/
	key:=keyGenerate()
	keyCommit:=fnKeycommit(key)
	MRinp:=createMerkleTree(inputvectors)                          //MRinp (merkle root for input vectors)
	MRcir:=createMerkleTreeForCir(circuitobjects)									//MRcir (merkle root for ciruit)
	encodedGateOutputs:=Encode(inputvectors,key,circuitobjects,MRinp)
	MRencout:=createMerkleTreeForEncInp(encodedGateOutputs)

	/*sentoconobject initialize*/
	idInitializingForContract(&SenToConObject)
	setPrice(&SenToConObject)
	SenToConObject.Keycommit=keyCommit
	SenToConObject.MerkleRootOfCircuit=MRcir
	SenToConObject.MerkleRootOfEncInp=MRencout


	/*SenToRecObject initialize*/
	idInitializingForRec(&SenToRecObject)
	SenToRecObject.Keycommit=keyCommit
	SenToRecObject.EncodedOutputOfGates=encodedGateOutputs

	/*Deploy contract*/
	contractAddress, _, _, err:=import_contracts.DeployFairswap(
		authsender,
		client,
		&SenToConObject.Price,
		SenToConObject.Keycommit,
		SenToConObject.MerkleRootOfCircuit,
		SenToConObject.MerkleRootOfEncInp,
	)
	if err!=nil{
		log.Fatal(err)
	}
	// commit all pending transactions
	client.Commit()


	//channel creation
	channel_SenToRec:=make(chan []byte)
	channel_RecToMain:=make(chan string)
	//fmt.Println("sender address inside main :",senderaddress)
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}

	go Receiver(contractAddress,receiveraddress,client,authreceiver,circuitobjects,channel_SenToRec,channel_RecToMain)
	channel_SenToRec<-byteArrayFromSenToRec

	fmt.Println(<-channel_RecToMain)

}


//function to model the circuit given....
func setCircuitTuples(circuitobjects *[TotalNumberOfgates]circuit){

	//Giving tuple value to input gates
	for i:=0;i<NoOfInputGates;i++{
			circuitobjects[i].Index = i
			circuitobjects[i].Operation =1
			circuitobjects[i].InputToTheGate = nil
	}

	//Giving tuple values to !inputgates && !lastgate
	k:=0
	for i:=NoOfInputGates;i<TotalNumberOfgates-1;i++ {
			circuitobjects[i].Index=i
			circuitobjects[i].Operation = 2
			for j:=0;j<MaxLinesToGate;j++{
				circuitobjects[i].InputToTheGate=append(circuitobjects[i].InputToTheGate,k)
				k++
			}
	}

	//Last gate
	circuitobjects[TotalNumberOfgates-1].Index = TotalNumberOfgates-1
	circuitobjects[TotalNumberOfgates-1].Operation = 3
	circuitobjects[TotalNumberOfgates-1].InputToTheGate=append(circuitobjects[TotalNumberOfgates-1].InputToTheGate,TotalNumberOfgates-2)

}


/*Sender Functions*/

//key generate
func keyGenerate() ([32]byte){

		keyGen := make([]byte, 32)
		var key [32]byte

	  _, err := rand.Read(keyGen)
    if err != nil {
        // handle error here
    }

		for i:=0;i<len(keyGen);i++{
			key[i]=keyGen[i]
		}
		return key
}

//fn fnKeycommit
//commits(hash's) the key
//input parameters
//key : key to be commited
//returns the commitment of key
func fnKeycommit(key [32]byte) ([32]byte){

			var toBEHashedKey []byte
			var HashedKey [32]byte

			for j:=0;j<32;j++{
				toBEHashedKey=append(toBEHashedKey,key[j])
			}

			HashedAfter:=kec.Keccak256(toBEHashedKey)

			for j:=0;j<32;j++{
					HashedKey[j]=HashedAfter[j]
			}

			return HashedKey
}

//fn createMerkleTree
//creates a merkle tree and returns the merkleroot
func createMerkleTree(inputvectors [NoOfInputGates][32]byte) ([32]byte){
	var merkletree [2*NoOfInputGates-1][32]byte

	//input gates
	for i:=0;i<NoOfInputGates;i++{
		for j:=0;j<32;j++{
				merkletree[i][j]=inputvectors[i][j]
		}
	}

	//!inputgates
	k:=0
	for i:=0;i<NoOfInputGates-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<32;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<32;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed)

		for j:=0;j<32;j++{
			merkletree[i+NoOfInputGates][j]=Hashed[j]
		}
		k=k+2
	}

	return merkletree[2*NoOfInputGates-2]
}

//Function  createMerkleTreeForEncInp
//constructs the merkle tree for the enoded inputs and returns the merkleroot
//Input Parameters :
//encryptedGateOutputs : Vectors of encrypted output of the circuit gates
//Returns merkleroot
func createMerkleTreeForEncInp(encryptedGateOutputs [TotNumOfEncOutVecToMer][32]byte) ([32]byte){
	var merkletree [2*TotNumOfEncOutVecToMer-1][32]byte

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<TotNumOfEncOutVecToMer;i++{
		for j:=0;j<32;j++{
				merkletree[i][j]=encryptedGateOutputs[i][j]
		}
	}

	//!inputgates
	//merkle tree construction
	k:=0
	for i:=0;i<TotNumOfEncOutVecToMer-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<32;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<32;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input

		for j:=0;j<32;j++{
			merkletree[i+TotNumOfEncOutVecToMer][j]=Hashed[j]
		}
		k=k+2
	}

	return merkletree[2*TotNumOfEncOutVecToMer-2]  //returns the merkleroot
}

//fn createMerkleTreeForCir
//the function creates a merkle tree for the circuit with each gate's byte array(got after doing json.Marshal to each ciruit gate) as input
//input parameters
//cicuitobjects : array of ciruit gate object
//returns merkleroot
//Unlike the merkletreeconstruction for encryptedGateOutputs,where the input vectors were not hashed(in the sense they were stored as it is), here the inputvectors
//to the circuit merkletreeconstruction is hashed and stored to maintain the uniformity in number of bytes in each gate array
func createMerkleTreeForCir(circuitobjects [TotalNumberOfgates]circuit) ([32]byte){

	var circuitByteArray [TotalNumberOfgates][]byte  //contains the json.marshal byte array of each gate

	for i:=0;i<TotalNumberOfgates;i++ {
		b, err := json.Marshal(circuitobjects[i])
		if err != nil {
			fmt.Println("error:", err)
		}
		for j:=0;j<len(b);j++{
			circuitByteArray[i]=append(circuitByteArray[i],b[j])
		}
	}


	//merkletreeconstruction
	var merkletree [2*TotalNumberOfgates-1][32]byte

	//input gates
	for i:=0;i<TotalNumberOfgates;i++{
		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<len(circuitByteArray[i]);j++{
			toBeHashed=append(toBeHashed,circuitByteArray[i][j])
		}
		Hashed=kec.Keccak256(toBeHashed)
		for j:=0;j<len(Hashed);j++{
				merkletree[i][j]=Hashed[j]
		}
	}

	//!inputgates
	k:=0
	for i:=0;i<TotalNumberOfgates-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<32;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<32;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed)

		for j:=0;j<32;j++{
			merkletree[i+TotalNumberOfgates][j]=Hashed[j]
		}
		k=k+2
	}

	return merkletree[2*TotalNumberOfgates-2]
}


//Encode function
//This fn is used to encrypt the processed input vectors
//For input gates,no operation is performed by the gate on the input vector, it is just a "pass through" gate. So the input vectors are directly encrypted and stored.
//For the remaining higher hierarchical gates the output is calulated according to characteristic operation of a particular gate.And the output is then encrypted.

//The Encode function input parameters :
// inputvectors  :  The set of all input vectors
// key 					 :  The key with which the encryption function will be performed
//circuitobjects :  An array of circuit objects that contains info related to eacg gate
//MRx 					 :  Merkle root of input vectors.Specific to this particular circuit.
//Returns : An array of encrypted vectors
func Encode(inputvectors [NoOfInputGates][32]byte,key [32]byte,circuitobjects [TotalNumberOfgates]circuit,MRx [32]byte) ([TotalNumberOfgates][32]byte)  {
	var out [TotalNumberOfgates][32]byte     //stores the output of each gate
	var z [TotalNumberOfgates][32]byte			//stores the encryption of each gate's output

	//Input Gates-No operation performed.Direct encryption of input vectors
	for i:=0;i<NoOfInputGates;i++ {
			out[i]=inputvectors[i]
			z[i]=Enc(key,out[i])  //calls the encryption fn
	}

	//Not Input Gates
	//For each gate in this range,the input vectors to a particular gate are processed according to the operation and the ouput and encryption of this output is stored
	for i:=NoOfInputGates;i<TotalNumberOfgates;i++{
			Op:=circuitobjects[i].Operation  //Operation of gate by the index "i"

			//Operation 2: The gate takes in the output of two of its children, concatenates the two inputs and hashes it.Later the output is encrypted.
			if(Op==2) {
					leftsibling:=circuitobjects[i].InputToTheGate[0]
					rightsibling:=circuitobjects[i].InputToTheGate[1]
					var toBeHashed []byte
					for j:=0;j<32;j++ {
							toBeHashed= append(toBeHashed,out[leftsibling][j])
					}
					for j:=0;j<32;j++ {
							toBeHashed= append(toBeHashed,out[rightsibling][j])
					}
					hash := kec.Keccak256(toBeHashed)
					for j:=0;j<32;j++{
						out[i][j]=hash[j]
					}
			}

			//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
			if(Op==3) {
				for j:=0;j<32;j++{
					if (j==0){
							if(out[i-1][j]==MRx[j]){
								out[i][j]=1
							} else {
								out[i][j]=0
										 }
							} else {
								if(out[i-1][j]==MRx[j]){
									out[i][j]=0
								}	else{
									out[i][j]=1
								}
					}
				}
			}

			z[i]=Enc(key,out[i])  //encryption function is called

		}
	return z
}

//Encrypt function
//The function is used to encrypt(XOR) the given vector.
//Input Parameters
//key : The key with which the XORing is done.
//plaintext : the input vector
//Returns the encrypt output
//Operation performed : plaintext (XOR) key
func Enc(key [32]byte,plaintext [32]byte) ([32]byte) {

			var keyPlusIndex []byte //appends the key(input parameter) and index "i"(range : 0 to Total Number of gates-1 )

			//In this specific case we are solving,the index("0") always remains the same.The funcion can be easily extended to accomadate varying index
			k:=0
			for i:=0;i<len(key);i++{
				keyPlusIndex=append(keyPlusIndex,key[i])
			}
			keyPlusIndex=append(keyPlusIndex,byte(k))
			key0 := kec.Keccak256(keyPlusIndex)


			var key032 [32]byte
			for i:=0;i<len(key0);i++{
				key032[i]=key0[i]
			}

			var encryptedtext [32]byte
			for i:=0;i<32;i++{
				encryptedtext[i]=key032[i]^plaintext[i]  //xor operation
			}
			return encryptedtext
}

//initialize id
func idInitializingForRec(object *sendertoreciever){
	object.Id = 1
}

func idInitializingForContract(object *sendertocontract){
	object.Id = 1
}

//set the price of the good
func setPrice(object *sendertocontract){

	object.Price=*big.NewInt(80)     //800 wei
}



/*Receiver and associated functions*/
func Receiver(contract_address common.Address,receiveraddress common.Address,_client *backends.SimulatedBackend,authreceiver *bind.TransactOpts ,circuitobjects [TotalNumberOfgates]circuit,channel_SenToRec chan []byte,channel_RecToMain chan string){

	instance,err := import_contracts.NewFairswap(contract_address,_client)
	if err != nil {
			log.Fatal(err)
	}

	_,err=instance.Initializerecieveraddrs(&bind.TransactOpts{
		From:authreceiver.From,
		Signer:authreceiver.Signer,
		Value: nil,
	})
	if err!=nil{
			log.Fatal(err)
	}

	_client.Commit()

	byteArrayFromSenToRec:=<-channel_SenToRec

	var Receiver_SenToRecObj sendertoreciever
	_ = json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)

	Receiver_MerkleRootOfCircuit,err:=instance.CircuitMerkleRoot(nil)
	if err!=nil{
		log.Fatal(err)
	}

	Receiver_MerkleRootOfEncInp,err:=instance.EncInputRoot(nil)
	if err!=nil{
		log.Fatal(err)
	}

	boolvalueCircuit:=CheckcreateMerkleTreeForCir(circuitobjects,Receiver_MerkleRootOfCircuit)
	boolvalueEncode:=CheckcreateMerkleTreeForEncInp(Receiver_SenToRecObj.EncodedOutputOfGates,Receiver_MerkleRootOfEncInp)

	value:=big.NewInt(9000)
	gasLimit := uint64(2100000)


	if(boolvalueEncode && boolvalueCircuit){
	_,err=instance.Accept(&bind.TransactOpts{
		From:authreceiver.From,
		Signer:authreceiver.Signer,
		Value:value,
		GasLimit:gasLimit,
	})
		if err != nil {
        log.Fatal(err)
  	}
	}

	_client.Commit()

	channel_RecToMain<-"Initialized"
}


//fn CheckcreateMerkleTreeForCir
//constructs merkletree and compares between the merkleroot got from construction and merkleroot passed as function parameter
//input parameters
//cicuitobjects : array of ciruit gate object
//merklerootforcircuit : merkleroot of circuit
//Returns bool value,true if merklerootforcircuit matches with the merkleroot got from construction,else false
func CheckcreateMerkleTreeForCir(circuitobjects [TotalNumberOfgates]circuit,merklerootforcircuit [32]byte) (bool){

		var circuitByteArray [TotalNumberOfgates][]byte     //contains the json.marshal byte array of each gate

		for i:=0;i<TotalNumberOfgates;i++ {
			b, err := json.Marshal(circuitobjects[i])
			if err != nil {
				fmt.Println("error:", err)
			}
			for j:=0;j<len(b);j++{
				circuitByteArray[i]=append(circuitByteArray[i],b[j])
			}
		}

		//merkletreeconstruction
		var merkletree [2*TotalNumberOfgates-1][32]byte

		//input gates
		for i:=0;i<TotalNumberOfgates;i++{
			var toBeHashed []byte
			var Hashed []byte
			for j:=0;j<len(circuitByteArray[i]);j++{
				toBeHashed=append(toBeHashed,circuitByteArray[i][j])
			}
			Hashed=kec.Keccak256(toBeHashed)
			for j:=0;j<len(Hashed);j++{
					merkletree[i][j]=Hashed[j]
			}
		}


		//!inputgates
		k:=0
		for i:=0;i<TotalNumberOfgates-1;i++{

			var toBeHashed []byte
			var Hashed []byte
			for j:=0;j<32;j++{
				toBeHashed=append(toBeHashed,merkletree[k][j])
			}

			for j:=0;j<32;j++{
				toBeHashed=append(toBeHashed,merkletree[k+1][j])
			}
			Hashed=kec.Keccak256(toBeHashed)
			for j:=0;j<32;j++{
				merkletree[i+TotalNumberOfgates][j]=Hashed[j]
			}
			k=k+2
		}

	for j:=0;j<32;j++{
		if(merkletree[2*TotalNumberOfgates-2][j]!=merklerootforcircuit[j]){
			return false
			break
		}
	}

	return true
}


//fn CheckcreateMerkleTreeForEncInp
//constructs merkletree and compares between the merkleroot got from construction and merkleroot passed as function parameter
//input parameters
//encryptedGateOutputs : Vectors of encrypted output of the circuit gates
//merklerootforEncInput : merkle root of Encrypted Input
//Returns bool value,true if merklerootforcircuit matches with the merkleroot got from construction,else false
func CheckcreateMerkleTreeForEncInp(encryptedGateOutputs [TotNumOfEncOutVecToMer][32]byte, merklerootforEncInput [32]byte) (bool){
	var merkletree [2*TotNumOfEncOutVecToMer-1][32]byte


	//inputgates
	for i:=0;i<TotNumOfEncOutVecToMer;i++{
		for j:=0;j<32;j++{
				merkletree[i][j]=encryptedGateOutputs[i][j]
		}
	}


	//!inputgates
	k:=0
	for i:=0;i<TotNumOfEncOutVecToMer-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<32;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<32;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}
		Hashed=kec.Keccak256(toBeHashed)
		for j:=0;j<32;j++{
			merkletree[i+TotNumOfEncOutVecToMer][j]=Hashed[j]
		}
		k=k+2
	}


	for j:=0;j<32;j++{
		if(merkletree[2*TotNumOfEncOutVecToMer-2][j]!=merklerootforEncInput[j]){
			return false
		}
	}
	return true
}
