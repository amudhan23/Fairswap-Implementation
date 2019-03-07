package main

import(
	"fmt"
	kec "github.com/ethereum/go-ethereum/crypto"
	"os"
	"io"
	"log"
	"encoding/json"
	//"reflect"

	"time"
  //"context"
  "crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core"
	"crypto/rand"
  "github.com/ethereum/go-ethereum/common"
	"math/big"
	mathlib "math"
   import_contracts "./contracts"
)

const NoOfInputGates=8
const BUFFER_SIZE=32
const TotalNumberOfgates=16
const MaxLinesToGate=2
const TotNumOfEncOutVecToMer=16
const depth=5

var contractAddress common.Address

//these struct and SetCircuitTuples function is common to both sender and receiver
type circuit struct{
	Index int
	Operation int
	InputToTheGate []int
}

type SenderToReceiverStruct struct {
		Id int                                        //id
		Keycommit [32]byte														//key commitment
		EncodedOutputOfGates [TotNumOfEncOutVecToMer][32]byte		//Encoded output of gates
}

type SenderToContractStruct struct{
	Id int
	Price big.Int
	Keycommit [32]byte
	MerkleRootOfCircuit [32]byte
	MerkleRootOfEncInp [32]byte
}

func main(){

	//channels
	channel_SenToRec:=make(chan []byte)				//sender sends the byterarray(of Id,Keycommit,EncodedOutputOfGates)
	channel_RecToMainIni:=make(chan string)		//Receiver sends a notificaion once he has finished his entire tasks
	channel_RecToMainRev:=make(chan string)
	channel_SenToMain:=make(chan string)

	//sender address generation
	authSender,senderAddress := AuthAndAddressGeneration()
	fmt.Println("\n->Sender\nSender Address :",senderAddress)

	//receiver address generation
	authReceiver,receiverAddress := AuthAndAddressGeneration()
	fmt.Println("\n->Receiver\nReceiver Address :",receiverAddress)

	//blockchain creation
	alloc := make(core.GenesisAlloc)
	alloc[authSender.From] = core.GenesisAccount{Balance: big.NewInt(1100000000000000000)}  //1.1 ether
	alloc[authReceiver.From] = core.GenesisAccount{Balance: big.NewInt(2500000000000000000)} //2.5 ether
	//client
	client := backends.NewSimulatedBackend(alloc,60000000000)

	//go routines
	go Sender(authSender,client,channel_SenToRec,channel_SenToMain)
	go Receiver(authReceiver,client,channel_SenToRec,channel_RecToMainIni,channel_RecToMainRev)

	SenderMessageToMain:=<-channel_SenToMain
	fmt.Println("\n->Main fn\n",SenderMessageToMain)

	ReceicverInittilaizationMessage:=<-channel_RecToMainIni
	fmt.Println("\n->Main fn\n",ReceicverInittilaizationMessage)

	ReceiverRevealMessage:=<-channel_RecToMainRev
	fmt.Println("\n->Main fn\n",ReceiverRevealMessage)

}

func Sender(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,channel_SenToMain chan string){

	//reading from a file "a.txt" and storing bytearray of 32 bytes in inputVectors of size equal to NoOfInputGates
	fmt.Println("\n->Sender\nBegins")
	var currentByte int64 = 0
	fileBuffer := make([]byte, BUFFER_SIZE)
	file, err := os.Open("a.txt") // For read access.
	if err != nil {
			log.Fatal(err)
	}
	var inputVectors [NoOfInputGates][32]byte
	count:=0
	for {
			n,err := file.ReadAt(fileBuffer, currentByte)
			currentByte += BUFFER_SIZE
			fileBuffer=fileBuffer[:n]
			if count==NoOfInputGates{
					break
			}
			for j:=0;j<n;j++{
				inputVectors[count][j]=fileBuffer[j]
			}
			if err == io.EOF {
					log.Fatal(err)
			}
			count++
	}
	file.Close()



	fmt.Println("->Sender\nInputVectors:",inputVectors)


	//objects
	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct
	var circuitObjects [TotalNumberOfgates]circuit //array of circuit object

	//sets the circuit
	setCircuitTuples(&circuitObjects)

	/*sender fns call*/
	key:=keyGenerate()
	keyCommit:=fnKeycommit(key)
	MRinp:=createMerkleTree(inputVectors)                          //MRinp (merkle root for input vectors)
	MRcir:=createMerkleTreeForCir(circuitObjects)									//MRcir (merkle root for ciruit)
	encodedGateOutputs,_:=Encode(inputVectors,key,circuitObjects,MRinp)
	MRencout:=createMerkleTreeForEncInp(encodedGateOutputs)					//MRencout(merkle root for encoded outputs)

	fmt.Println("\n->Sender \n Key:",key,"\n\nkeyCommit : ",keyCommit,"\n\nMRinp (merkle root for input vectors): ",MRinp,"\n\nMRcir (merkle root for ciruit)",MRcir,"\n\n encodedGateOutputs : ",encodedGateOutputs,"\n\n MRencout(merkle root for encoded outputs) : ",MRencout)

	/*sentoconobject initialize*/
	idInitializingForContract(&SenToConObject)
	setPrice(&SenToConObject)
	SenToConObject.Keycommit=keyCommit
	SenToConObject.MerkleRootOfCircuit=MRcir
	SenToConObject.MerkleRootOfEncInp=MRencout

	a:=[]*big.Int{}
	for i:=0;i<TotalNumberOfgates;i++{
		b:=big.NewInt(int64(circuitObjects[i].Operation))
		a=append(a,b)
	}


	/*SenToRecObject initialize*/
	idInitializingForRec(&SenToRecObject)
	SenToRecObject.Keycommit=keyCommit
	SenToRecObject.EncodedOutputOfGates=encodedGateOutputs

	/*Deploy contract*/
	contractAddress, _,_, err=import_contracts.DeployFairswap(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.Keycommit,
		SenToConObject.MerkleRootOfCircuit,
		SenToConObject.MerkleRootOfEncInp,
		a,

	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()

	fmt.Println("\n->Sender \n Contract Deployed",)

	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("\n->Sender \nByteArray sent to Receiver")

	channel_SenToRec<-byteArrayFromSenToRec

	fmt.Println("\n->Sender\n Initialization Finished")

	instance,err:= import_contracts.NewFairswap(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}


	RevealedBool:=false

	for i:=0;i<360;i++ {
		Phase,err:=instance.Phase(nil)
		if err!=nil{
			log.Fatal(err)
		}
		if(Phase==2){
			_,err=instance.RevealKey(&bind.TransactOpts{
					From:authSender.From,
					Signer:authSender.Signer,
					Value: nil,
			},key)
			if err!=nil{
				log.Fatal(err)
			}
			client.Commit()
			RevealedBool=true
			break;
		}
		time.Sleep(time.Second *1)

	}

	if(!RevealedBool){
		fmt.Println("\n->Sender\nReceiver have not accepted within the timeout, Key not Revealed")

	}else{
			fmt.Println("\n->Sender\nReceiver have accepted within the timeout, Key Revealed")
	}

	channel_SenToMain<-"Key Workout"

}


func Receiver(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,channel_RecToMainIni chan string,channel_RecToMainRev chan string){

	//objects
	var circuitObjects [TotalNumberOfgates]circuit  //array of circuit object
	//sets the circuit
	setCircuitTuples(&circuitObjects)

	byteArrayFromSenToRec:=<-channel_SenToRec

	fmt.Println("\n->Receiver\nByteArray got from Sender ")

	var Receiver_SenToRecObj SenderToReceiverStruct
	err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
	if err != nil {
			log.Fatal(err)
	}

	instance,err := import_contracts.NewFairswap(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	_,err=instance.Initializerecieveraddrs(&bind.TransactOpts{
		From:authReceiver.From,
		Signer:authReceiver.Signer,
		Value: nil,
	})
	if err!=nil{
			log.Fatal(err)
	}

	client.Commit()

	Receiver_MerkleRootOfCircuit,err:=instance.CircuitMerkleRoot(nil)
	if err!=nil{
		log.Fatal(err)
	}

	Receiver_MerkleRootOfEncInp,err:=instance.EncInputRoot(nil)
	if err!=nil{
		log.Fatal(err)
	}





	boolvalueCircuit:=CheckcreateMerkleTreeForCir(circuitObjects,Receiver_MerkleRootOfCircuit)

	fmt.Println("\n->Receiver\nboolvalueCircuit(comparision between merkleroot of cicuit from contract and calulated circuit merkleroot): ",boolvalueCircuit)

	Receiver_EncodedOutputOfGates:=Receiver_SenToRecObj.EncodedOutputOfGates

	fmt.Println("\n->Receiver\n Receiver_EncodedOutputOfGates",Receiver_EncodedOutputOfGates)

	boolvalueEncode:=CheckcreateMerkleTreeForEncInp(Receiver_EncodedOutputOfGates,Receiver_MerkleRootOfEncInp)

	fmt.Println("\n->Receiver\boolvalueEncode(comparision between merkleroot of encodedGateOutputs from contract and calulated encodedGateOutputs merkleroot): ",boolvalueEncode)

	value:=big.NewInt(9000)
	gasLimit := uint64(2100000000)

	if(boolvalueEncode && boolvalueCircuit){
	_,err=instance.Accept(&bind.TransactOpts{
		From:authReceiver.From,
		Signer:authReceiver.Signer,
		Value:value,
		GasLimit:gasLimit,
	})
		if err != nil {
        log.Fatal(err)
  	}
	}

	client.Commit()
	fmt.Println("\n->Receiver\nFinished")
	channel_RecToMainIni<-"Initialization Phase Finished"

	receiverAddress,_:=instance.Recieveraddrs(nil)
	fmt.Println("\n->Receiver\nreceiverAddress",receiverAddress,"\n")

	fmt.Println("\n->Receiver\nInitialization ends")


	AcceptedBool:=false
	var contract_Key [32]byte

	for i:=0;i<360;i++ {
		Phase,err:=instance.Phase(nil)
		if err!=nil{
			log.Fatal(err)
		}
		if(Phase==3){
			contract_Key,err=instance.Key(nil)
			if err!=nil{
				log.Fatal(err)
			}
			fmt.Println("\n->Receiver\ncontract_Key",contract_Key)
			AcceptedBool=true
			break;
		}
		time.Sleep(time.Second *1)

	}


	if(!AcceptedBool){

			fmt.Println("\n->Receiver\nSender has not revealed the key")

	} else{

		fmt.Println("\n->Receiver\nSender has revealed the key....Proceeding further")

	// for error		Receiver_EncodedOutputOfGates[9][0]=3
		merkletree:=ReceiverMerkleTreeCreate(Receiver_EncodedOutputOfGates)

		complain,_,index:=Extract(Receiver_EncodedOutputOfGates,contract_Key,circuitObjects,merkletree,Receiver_MerkleRootOfEncInp)


		if(index==nil){
			fmt.Println("\n->Receiver\nNo complain")
		} else{
			indexes:=[]*big.Int{}
			for i:=0;i<len(index);i++{
				b:=big.NewInt(int64(index[i]))
				indexes=append(indexes,b)
			}

				_,err=instance.Complain(&bind.TransactOpts{
					From:authReceiver.From,
					Signer:authReceiver.Signer,
					Value:nil,
					GasLimit:gasLimit,
				},complain,indexes)
				if err!=nil{
					log.Fatal(err)
				}

				client.Commit()

		}
	}
/*	ComplianBool,err:=instance.ComplianBool(nil)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("\n\n\n\nComplianBool : ",ComplianBool,"\n\n\n\n")*/


	channel_RecToMainRev<-"Reveal Phase Finished"
}


func AuthAndAddressGeneration() (*bind.TransactOpts,common.Address){
  privateKey,err:=crypto.GenerateKey()
  if err!=nil{
    log.Fatal(err)
  }
  auth:=bind.NewKeyedTransactor(privateKey)

  publicKey := privateKey.Public()
  publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
  if !ok {
      log.Fatal("error casting public key to ECDSA")
  }
  address := crypto.PubkeyToAddress(*publicKeyECDSA)
  return auth,address
}



//function to model the circuit given....
func setCircuitTuples(circuitObjects *[TotalNumberOfgates]circuit){

	//Giving tuple value to input gates
	for i:=0;i<NoOfInputGates;i++{
			circuitObjects[i].Index = i
			circuitObjects[i].Operation =1
			circuitObjects[i].InputToTheGate = nil
	}

	//Giving tuple values to !inputgates && !lastgate
	k:=0
	for i:=NoOfInputGates;i<TotalNumberOfgates-1;i++ {
			circuitObjects[i].Index=i
			circuitObjects[i].Operation = 2
			for j:=0;j<MaxLinesToGate;j++{
				circuitObjects[i].InputToTheGate=append(circuitObjects[i].InputToTheGate,k)
				k++
			}
	}

	//Last gate
	circuitObjects[TotalNumberOfgates-1].Index = TotalNumberOfgates-1
	circuitObjects[TotalNumberOfgates-1].Operation = 3
	circuitObjects[TotalNumberOfgates-1].InputToTheGate=append(circuitObjects[TotalNumberOfgates-1].InputToTheGate,TotalNumberOfgates-2)

}


/*Sender Functions*/

//key generate
func keyGenerate() ([32]byte){

		keyGen := make([]byte, 32)
		var key [32]byte

		//Here
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
func createMerkleTree(inputVectors [NoOfInputGates][32]byte) ([32]byte){
	var merkletree [2*NoOfInputGates-1][32]byte

	//input gates
	for i:=0;i<NoOfInputGates;i++{
		for j:=0;j<32;j++{
				merkletree[i][j]=inputVectors[i][j]
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

	  //returns the merkleroot
	return merkletree[2*TotNumOfEncOutVecToMer-2]
}

//fn createMerkleTreeForCir
//the function creates a merkle tree for the circuit with each gate's byte array(got after doing json.Marshal to each ciruit gate) as input
//input parameters
//cicuitobjects : array of ciruit gate object
//returns merkleroot
//Unlike the merkletreeconstruction for encryptedGateOutputs,where the input vectors were not hashed(in the sense they were stored as it is), here the inputVectors
//to the circuit merkletreeconstruction is hashed and stored to maintain the uniformity in number of bytes in each gate array
func createMerkleTreeForCir(circuitObjects [TotalNumberOfgates]circuit) ([32]byte){

	var circuitByteArray [TotalNumberOfgates][]byte  //contains the json.marshal byte array of each gate

	for i:=0;i<TotalNumberOfgates;i++ {
		b, err := json.Marshal(circuitObjects[i])
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
// inputVectors  :  The set of all input vectors
// key 					 :  The key with which the encryption function will be performed
//circuitObjects :  An array of circuit objects that contains info related to eacg gate
//MRx 					 :  Merkle root of input vectors.Specific to this particular circuit.
//Returns : An array of encrypted vectors
func Encode(inputVectors [NoOfInputGates][32]byte,key [32]byte,circuitObjects [TotalNumberOfgates]circuit,MRx [32]byte) ([TotalNumberOfgates][32]byte,[TotalNumberOfgates][32]byte)  {
	var out [TotalNumberOfgates][32]byte     //stores the output of each gate
	var z [TotalNumberOfgates][32]byte			//stores the encryption of each gate's output

	//Input Gates-No operation performed.Direct encryption of input vectors
	for i:=0;i<NoOfInputGates;i++ {
			out[i]=inputVectors[i]
			z[i]=Enc(key,out[i])  //calls the encryption fn
	}

	//Not Input Gates
	//For each gate in this range,the input vectors to a particular gate are processed according to the operation and the ouput and encryption of this output is stored
	for i:=NoOfInputGates;i<TotalNumberOfgates;i++{
			Op:=circuitObjects[i].Operation  //Operation of gate by the index "i"

			//Operation 2: The gate takes in the output of two of its children, concatenates the two inputs and hashes it.Later the output is encrypted.
			if(Op==2) {
					leftsibling:=circuitObjects[i].InputToTheGate[0]
					rightsibling:=circuitObjects[i].InputToTheGate[1]
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

	return z,out
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
func idInitializingForRec(object *SenderToReceiverStruct){
	object.Id = 1
}

func idInitializingForContract(object *SenderToContractStruct){
	object.Id = 1
}

//set the price of the good
func setPrice(object *SenderToContractStruct){

	object.Price=*big.NewInt(80)     //800 wei
}


/*Receiver associated functions*/

//fn CheckcreateMerkleTreeForCir
//constructs merkletree and compares between the merkleroot got from construction and merkleroot passed as function parameter
//input parameters
//cicuitobjects : array of ciruit gate object
//merklerootforcircuit : merkleroot of circuit
//Returns bool value,true if merklerootforcircuit matches with the merkleroot got from construction,else false
func CheckcreateMerkleTreeForCir(circuitObjects [TotalNumberOfgates]circuit,merklerootforcircuit [32]byte) (bool){

		var circuitByteArray [TotalNumberOfgates][]byte     //contains the json.marshal byte array of each gate

		for i:=0;i<TotalNumberOfgates;i++ {
			b, err := json.Marshal(circuitObjects[i])
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



func Extract(encodedGateOutputs [TotNumOfEncOutVecToMer][32]byte,key [32]byte,circuitObjects [TotalNumberOfgates]circuit,merkletree [2*TotNumOfEncOutVecToMer-1][32]byte,Receiver_MerkleRootOfEncInp [32]byte) ([][depth][32]byte,[][32]byte,[]int) {

	var decodedOutputs [][32]byte
	var complain [][depth][32]byte
	var out [32]byte
	var indices []int
	for i:=0;i<NoOfInputGates;i++ {
			out =Dec(key,encodedGateOutputs[i])
			decodedOutputs = append(decodedOutputs,out)
	}

	for i:=NoOfInputGates;i<TotalNumberOfgates;i++{
		noOFinputToThisParticularGate:=len(circuitObjects[i].InputToTheGate)

				var operationInputs [][32]byte
				for j:=0;j<noOFinputToThisParticularGate;j++{

								operationInputs=append(operationInputs,decodedOutputs[circuitObjects[i].InputToTheGate[j]])
				}

		out = Dec(key,encodedGateOutputs[i])
		decodedOutputs=append(decodedOutputs,out)

						if decodedOutputs[i]!=Operation(circuitObjects[i].Operation,operationInputs,decodedOutputs,Receiver_MerkleRootOfEncInp){
							index:=i
							indices=append(indices,index)
							MproofTree:=Mproof(i,merkletree)
							complain=append(complain,MproofTree)
							for k:=0;k<noOFinputToThisParticularGate;k++{
									MproofTree:=Mproof(circuitObjects[i].InputToTheGate[k],merkletree)
									complain=append(complain,MproofTree)
									indices=append(indices,circuitObjects[i].InputToTheGate[k])
							}
							return complain,decodedOutputs,indices
					}


	}
	return complain,decodedOutputs,indices
}


func Dec(key [32]byte,encryptedtext [32]byte) ([32]byte){
			var keyPlusIndex []byte

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

			var plaintext [32]byte
			for i:=0;i<32;i++{
				plaintext[i]=key032[i]^encryptedtext[i]
			}

			return plaintext
}

func Mproof(index int,merkletree [2*TotNumOfEncOutVecToMer-1][32]byte) ([depth][32]byte){

	depth:=int(mathlib.Log2(16))+1
	var tree [5][32]byte
	tree[0]=merkletree[index]
	for i:=1;i<depth;i++{
		if(index%2==0){
			tree[i]=merkletree[index+1]
			index=index/2+TotNumOfEncOutVecToMer
		} else {
			tree[i]=merkletree[index-1]
			index=index/2+TotNumOfEncOutVecToMer
		}
	}
	return tree

}


func Operation(Op int,operationInputs [][32]byte,decodedOutputs [][32]byte,Receiver_MerkleRootOfEncInp [32]byte)  ([32]byte){


		var result [32]byte
		if(Op==2) {
				leftsibling:=operationInputs[0]
				rightsibling:=operationInputs[1]
				var toBeHashed []byte
				for j:=0;j<32;j++ {
						toBeHashed= append(toBeHashed,leftsibling[j])
				}
				for j:=0;j<32;j++ {
						toBeHashed= append(toBeHashed,rightsibling[j])
				}

				hashed := kec.Keccak256(toBeHashed)

				for j:=0;j<32;j++{
					result[j]=hashed[j]
				}
		}

		//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
		if(Op==3) {
			var merkletreeofdecodedinput1_8 [NoOfInputGates][32]byte
			for k:=0;k<NoOfInputGates;k++{
					merkletreeofdecodedinput1_8[k]=decodedOutputs[k]
			}
			merklerootofdecodedinput1_8:=createMerkleTree(merkletreeofdecodedinput1_8)
			for j:=0;j<32;j++{
				if (j==0){
						if(operationInputs[0][j]==merklerootofdecodedinput1_8[j]){
							result[j]=1
						} else {
							result[j]=0
									 }
						} else {
							if(operationInputs[0][j]==merklerootofdecodedinput1_8[j]){
								result[j]=0
							}	else{
								result[j]=1
							}
				}
			}
		}
		return result

}

func ReceiverMerkleTreeCreate(encryptedGateOutputs [TotNumOfEncOutVecToMer][32]byte) ([2*TotNumOfEncOutVecToMer-1][32]byte){
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

	return merkletree  //returns the merkleroot
}


func Mverify(complaint [depth][32]byte,_index int,ciphertextRoot [32]byte) (bool){
	_value:=complaint[0]

	var Hashed []byte
	var i uint
	for i=0; i<depth-1; i++{
		if ((_index & (1<<i))>>i == 1)	{

			var toBeHashed []byte
			for j:=0;j<32;j++{
				toBeHashed=append(toBeHashed,complaint[i+1][j])
			}
			for j:=0;j<32;j++{
				toBeHashed=append(toBeHashed, _value[j])
			}

			Hashed = kec.Keccak256(toBeHashed)

			for j:=0;j<32;j++{
				_value[j]=Hashed[j]

			}

		}else{

			var toBeHashed []byte
			for j:=0;j<32;j++{
				toBeHashed=append(toBeHashed,_value[j])
			}
			for j:=0;j<32;j++{
				toBeHashed=append(toBeHashed,complaint[i+1][j])
			}

			Hashed = kec.Keccak256(toBeHashed)

			for j:=0;j<32;j++{
				_value[j]=Hashed[j]
			}

		}

	}
	return (_value==ciphertextRoot)
}
