package main

import(
	"fmt"
	kec "github.com/ethereum/go-ethereum/crypto"
	"os"
	"io"
	"log"
	"reflect"
	"context"
	"time"
	"encoding/json"
	//"github.com/ethereum/go-ethereum/core"
	mathlib "math"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	import_contracts "./contracts"

)

const buffer_size=64
const maxLinesToGate=2
const noOfInputGates=8
const totalNumberOfGates=16
const keyLength=32
const hashFunctionOutputBitSize=32
const gateTupleSize=32
const waitingTime=10

var contractAddress common.Address

type SenderToReceiverStruct struct {
		EncryptedOutputOfGates [totalNumberOfGates][]byte
		ReceiverEntryKey [keyLength]byte
}

type SenderToContractStruct struct{
	Price big.Int
	Keycommit [hashFunctionOutputBitSize]byte
	MerkleRootOfCircuit [hashFunctionOutputBitSize]byte
	MerkleRootOfEncInp [hashFunctionOutputBitSize]byte
	ConReceiverEntryKey [keyLength]byte
	TotalNumberOfGates big.Int
	MaxInputLinesToAGate big.Int
}

func main(){

		//channels
		channel_SenToRec:=make(chan []byte)				//sender sends the byterarray(of
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
		alloc[authSender.From] = core.GenesisAccount{Balance: big.NewInt(11000000000000000)}  //1.1 ether
		alloc[authReceiver.From] = core.GenesisAccount{Balance: big.NewInt(250000000000000000)} //2.5 ether
		//client
		client := backends.NewSimulatedBackend(alloc,60000000000)

		//Calculating sender and receiver initial account balances
		senderInitialBalance, err := client.BalanceAt(context.Background(),senderAddress, nil)
		if err != nil{
				log.Fatal(err)
		}
		fmt.Println("-->Main\nSenderInitialAccaountBalance : ",senderInitialBalance,"\n")

		receiverInitialBalance, err := client.BalanceAt(context.Background(),receiverAddress, nil)
		if err != nil{
				log.Fatal(err)
		}
		fmt.Println("-->Main\nReceiverInitialAccountBalance : ",receiverInitialBalance,"\n")


		//go routines
		go Sender(authSender,client,channel_SenToRec,channel_SenToMain)
		go maliciousReceiver_WrongComplain(authReceiver,client,channel_SenToRec,channel_RecToMainIni,channel_RecToMainRev)


		//channels... Parties communicate to the main function
		ReceicverInittilaizationMessage:=<-channel_RecToMainIni
		fmt.Println("\n->Main fn\n",ReceicverInittilaizationMessage)

		SenderMessageToMain:=<-channel_SenToMain
		fmt.Println("\n->Main fn\n",SenderMessageToMain)

		ReceiverRevealMessage:=<-channel_RecToMainRev
		fmt.Println("\n->Main fn\n",ReceiverRevealMessage)

		//getting the parties account balance at the end of protocol
		senderFinalBalance, err := client.BalanceAt(context.Background(),senderAddress, nil)
		if err != nil{
				log.Fatal(err)
		}

		senderDifferenceInFinalAndInitialBalance:=new(big.Int).Sub(senderInitialBalance,senderFinalBalance)
		fmt.Println("\n-->Main\nSenderFinalAccountBalance : ",senderFinalBalance,"\nSenderDifferenceInFinalAndInitialBalance : ",senderDifferenceInFinalAndInitialBalance)

		receiverFinalBalance, err := client.BalanceAt(context.Background(),receiverAddress, nil)
		if err != nil{
				log.Fatal(err)
		}
		receiverDifferenceInFinalAndInitialBalance:=new(big.Int).Sub(receiverInitialBalance,receiverFinalBalance)
		fmt.Println("\n-->Main\nReceiverFinalAccountBalance : ",receiverFinalBalance,"\nReceiverDifferenceInFinalAndInitialBalance : ",receiverDifferenceInFinalAndInitialBalance ,"\n")

}


func Sender(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,channel_SenToMain chan string){


	var currentByte int64 = 0
	fileBuffer := make([]byte, buffer_size)
	file, err := os.Open("a.txt") // For read access.
	if err != nil {
			log.Fatal(err)
	}
	var inputVectors [noOfInputGates][buffer_size]byte
	count:=0
	for {
			n,err := file.ReadAt(fileBuffer, currentByte)
			currentByte += buffer_size
			fileBuffer=fileBuffer[:n]
			if count==noOfInputGates{
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

	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct

	gateTuples:=initializeCircuitTuples_asAnArray()
	MRgateTuples:=createMerkleRootForGateTuples(gateTuples)
	//maliciousGateTuples:=maliciousInitializeCircuitTuples_asAnArray()
	key:=keyGenerate()
	keyCommit:=fnKeycommit(key)
	MRinp:=createMerkleRootForInputVectors(inputVectors)
	encodedGateOutputs,_:=Encode(inputVectors,key,gateTuples,MRinp)
	MRencout:=createMerkleRootForEncInp(encodedGateOutputs)
	recEntryKey:=keyGenerate()

	setPrice(&SenToConObject)
	setTotalNumberOfGates(&SenToConObject)
	setMaxLinesToAGate(&SenToConObject)
	SenToConObject.Keycommit=keyCommit
	SenToConObject.MerkleRootOfEncInp=MRencout
	SenToConObject.MerkleRootOfCircuit=MRgateTuples
	SenToConObject.ConReceiverEntryKey=recEntryKey

	contractAddress, _,_, err=import_contracts.DeployFairswap(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.Keycommit,
		SenToConObject.MerkleRootOfEncInp,
		SenToConObject.MerkleRootOfCircuit,
		SenToConObject.ConReceiverEntryKey,
		&SenToConObject.TotalNumberOfGates,
		&SenToConObject.MaxInputLinesToAGate,
	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()

	SenToRecObject.EncryptedOutputOfGates=encodedGateOutputs
	SenToRecObject.ReceiverEntryKey=recEntryKey
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}

	channel_SenToRec<-byteArrayFromSenToRec
	fmt.Println("\n->Sender \nByteArray sent to Receiver")
	fmt.Println("\n->Sender\n Initialization Finished")

	//creating a contract instance
	instance,err:= import_contracts.NewFairswap(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	RevealedBool:=false
	for i:=0;i<waitingTime;i++ {
		Phase,err:=instance.Phase(nil)
		if err!=nil{
			log.Fatal(err)
		}
		if(Phase==2){
			RevealedBool=true
			break;
		}
		if(Phase==6){
			fmt.Println("\n->Sender\nReceiver has not paid enough ether")
			channel_SenToMain<-"Malicious Receiver-Dind't pay enough ether "
			break;
		}
		time.Sleep(time.Second *1)

	}

	Phase,err:=instance.Phase(nil)
	if err!=nil{
			log.Fatal(err)
	}
	if(Phase==1){
		fmt.Println("\n->Sender\nReceiver hasnt accepted within time")
		channel_SenToMain<-"Either the reciever is not interested now"
	}else{

			if(!RevealedBool){
				channel_SenToMain<-"Receiver has not accepted within the timeout, Key not Revealed"

			}else{

							_,err=instance.RevealKey(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},key)
							if err!=nil{
								log.Fatal(err)
							}
							client.Commit()
					}

			//The contract is in keyRevealed stage
			//Once the sender has revealed the key, the receiver has time until waitingTime to either register a complaint or
			//walk away(not do any execution or transaction--got the correct file)
			//case 1 : If receiver gives a wrong complaint then the contract stage goes to finishedAndMaliciousReciever. The sender sends
			//				"Receiver Wrong complaint" to the main function and finishes the protocol.
			//case 2 : If the receiver doesn't respond, then the stage contract stage is still in keyRevealed and the sender gets the ether
			//				 after the waitingTime
			for i:=0;i<waitingTime;i++ {

					Phase,err:=instance.Phase(nil)
					if err!=nil{
						log.Fatal(err)
					}
					if(Phase==6){
						fmt.Println("\n->Sender\nReceiver has made a wrong complaint")
						channel_SenToMain<-"Receiver Wrong complaint "
						break;
					}
					if(Phase==5){
						fmt.Println("\n->Sender\nIt seems I have made a mistake")
						channel_SenToMain<-"Sender has wrongly encrypted filechunks"
						break
					}
					time.Sleep(time.Second *1)
			}

			gasLimit := uint64(2100000000)

			//the sender getting the ether
			Phase,err:=instance.Phase(nil)
			if err!=nil{
				log.Fatal(err)
			}

			if(Phase==3){
			_,err=instance.SenderGetEther(&bind.TransactOpts{
					From:authSender.From,
					Signer:authSender.Signer,
					Value: big.NewInt(10000000000),
					GasLimit:gasLimit,
				})
			if err!=nil{
						log.Fatal(err)
			}

			//the protocol gets finished and the sender sends the message to the main function
			channel_SenToMain<-"Protocol2 completed.Sender Got the money"
			}
	}
}

//channel_RecToMainRev(chan string) - to notify the main function for the completion of reveal phase
func Receiver(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,channel_RecToMainIni chan string,channel_RecToMainRev chan string){

	byteArrayFromSenToRec:=<-channel_SenToRec

	var Receiver_SenToRecObj SenderToReceiverStruct
	err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
	if err != nil {
			log.Fatal(err)
	}

	instance,err := import_contracts.NewFairswap(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	Receiver_MerkleRootOfEncInp,err:=instance.EncInputRoot(nil)
	if err!=nil{
		log.Fatal(err)
	}

	Receiver_MerkleRootOfGateTuples,err:=instance.EncCircuitRoot(nil)
	if err!=nil{
		log.Fatal(err)
	}

	Receiver_EncryptedOutputOfGates:=Receiver_SenToRecObj.EncryptedOutputOfGates
	gateTuples:=initializeCircuitTuples_asAnArray()

	boolValueEncryptedGateOutput:=checkMerkleRootForEncInp(Receiver_EncryptedOutputOfGates,Receiver_MerkleRootOfEncInp)
	boolValueEncryptedGateTuples:=checkMerkleRootForGateTuples(gateTuples,Receiver_MerkleRootOfGateTuples)

	fmt.Println("\n->Receiver\nboolValueEncryptedGateOutput(comparision between merkleroot of encryptedGateOutputs from contract and calulated encryptedGateOutputs merkleroot): ",boolValueEncryptedGateOutput)
	fmt.Println("\n->Receiver\nboolValueEncryptedGateTuples(comparision between merkleroot of encryptedGateTuples from contract and calulated encryptedGateTuples merkleroot): ",boolValueEncryptedGateTuples)

	if(boolValueEncryptedGateOutput && boolValueEncryptedGateTuples){

				//the boolValueEncode was true, so the receiver calls the accept contract function and pays the required price
			_,err=instance.InitializeRecieverAddress(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value: nil,
			},Receiver_SenToRecObj.ReceiverEntryKey)
			if err!=nil{
					log.Fatal(err)
			}

			client.Commit()

			value:=big.NewInt(100)
			gasLimit := uint64(230000000)

			_	,err=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
				GasLimit:gasLimit,
			})
			if err != nil {
					log.Fatal(err)
			}

			client.Commit()

			fmt.Println("\n->Receiver\n Finished Accepting")
			channel_RecToMainIni<-"Accepted stage got over"


			//The contract is in accepted stage
			//The sender
			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
			//case 2 : The sender doesn't reveal the key within time. AcceptedBool remains false after the for loop. Receiver later calls the contract function
			//			   ReceiverGetEther to get the ether.
			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and AcceptedBool is true after the for loop.
			AcceptedBool:=false
			var contract_Key [keyLength]byte
			for i:=0;i<waitingTime;i++ {
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

				if(Phase==5){
					fmt.Println("\n->Receiver\nSender Has revealed the wrong key")
					channel_RecToMainRev<-"Wrong key revealed"

				}
				time.Sleep(time.Second *1)

			}


			if(!AcceptedBool){

					fmt.Println("\n->Receiver\nSender has not revealed the key...ReceiverGetEther function called")
					gasLimit := uint64(2100000000)

					//the receiver getting the ether
					_,err=instance.ReceiverGetEther(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value: nil,
							GasLimit:gasLimit,
						})
					if err!=nil{
								log.Fatal(err)
					}
					client.Commit()

					channel_RecToMainRev<-"Sender hasn't revealed the key within the Timeout"

			} else{

				fmt.Println("\n->Receiver\nSender has revealed the key....Proceeding further")

				merkleTreeForGateTuples:=createMerkleTreeForGateTuples(gateTuples)
			 	merkleTreeForEncInp:=createMerkleTreeForEncInp(Receiver_SenToRecObj.EncryptedOutputOfGates)

			 	complain,decodedOutputs,gateComplain,encodedVectors_ThisGate :=Extract(Receiver_SenToRecObj.EncryptedOutputOfGates,contract_Key,gateTuples,merkleTreeForEncInp,merkleTreeForGateTuples,Receiver_MerkleRootOfEncInp,Receiver_MerkleRootOfGateTuples)

				fmt.Println("\n->Receiver\nDecoded Ouptuts : ",decodedOutputs)

				if(complain==nil){
					fmt.Println("\n->Receiver\nNo complain")
					channel_RecToMainRev<-"Receiver Finished Protocol "
				} else{

						//complain function being called
						_,err=instance.Complain(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value:nil,
							GasLimit:gasLimit,
						},complain,gateComplain,encodedVectors_ThisGate)
						if err!=nil{
							log.Fatal(err)
						}
						client.Commit()

						Phase,err:=instance.Phase(nil)
						if err!=nil{
							log.Fatal(err)
						}
						if(Phase==5){
							fmt.Println("\n->Receiver\nSender Has wrongly calulated the encrypted inputs")
							channel_RecToMainRev<-"Wrong encrypted inputs"

						}
				}
			}
			channel_RecToMainRev<-"Reveal Phase Finished"
		}	else{

				channel_RecToMainIni<-"Sender sent wrong encodedGateOutputs,not equal to contract encodedGateOutputs merkleroot,Not gone into accepted stage"
				channel_RecToMainRev<-"Sender sent wrong encodedGateOutputs,not equal to contract encodedGateOutputs merkleroot,Not gone into revealed phase"
		}
}

func AuthAndAddressGeneration() (*bind.TransactOpts,common.Address){

	privateKey,err:=crypto.GenerateKey()
  if err!=nil{
    log.Fatal(err)
  }
  auth:=bind.NewKeyedTransactor(privateKey)

	//fmt.Println(reflect.TypeOf(auth))

  publicKey := privateKey.Public()
  publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
  if !ok {
      log.Fatal("error casting public key to ECDSA")
  }
  address := crypto.PubkeyToAddress(*publicKeyECDSA)
  return auth,address
}


func fnKeycommit(key [keyLength]byte) ([hashFunctionOutputBitSize]byte){

			var toBEHashedKey []byte
			var HashedKey [hashFunctionOutputBitSize]byte

			for j:=0;j<keyLength;j++{
				toBEHashedKey=append(toBEHashedKey,key[j])
			}

			HashedAfter:=kec.Keccak256(toBEHashedKey)

			for j:=0;j<hashFunctionOutputBitSize;j++{
					HashedKey[j]=HashedAfter[j]
			}

			return HashedKey
}

func setPrice(object *SenderToContractStruct){

	object.Price=*big.NewInt(80)     //800 wei
}

func setTotalNumberOfGates(object *SenderToContractStruct){

	object.TotalNumberOfGates=*big.NewInt(totalNumberOfGates)     //800 wei
}

func setMaxLinesToAGate(object *SenderToContractStruct){

	object.MaxInputLinesToAGate=*big.NewInt(maxLinesToGate)     //800 wei
}

func Encode(inputVectors [noOfInputGates][buffer_size]byte,key [32]byte,gateTuples [totalNumberOfGates][gateTupleSize]byte,MRx [hashFunctionOutputBitSize]byte) ([totalNumberOfGates][]byte,[totalNumberOfGates][]byte)  {

	var out [totalNumberOfGates][]byte     //stores the output of each gate
	var z [totalNumberOfGates][]byte

	for i:=0;i<noOfInputGates;i++ {
			for j:=0;j<buffer_size;j++{
					out[i]=append(out[i],inputVectors[i][j])
			}

			temp:=Enc(key,out[i])
			for j:=0;j<buffer_size;j++{
				z[i]=append(z[i],temp[j])
			}
	}

	//Not Input Gates
	//For each gate in this range,the input vectors to a particular gate are processed according to the operation and the ouput and encryption of this output is stored
	for i:=noOfInputGates;i<noOfInputGates+noOfInputGates/2;i++{
			Op:=int(gateTuples[i][1])  //Operation of gate by the index "i"

			//Operation 2: The gate takes in the output of two of its children, concatenates the two inputs and hashes it.Later the output is encrypted.
			if(Op==2) {
					leftsibling:=int(gateTuples[i][2])
					rightsibling:=int(gateTuples[i][3])
					var toBeHashed []byte
					for j:=0;j<buffer_size;j++ {
							toBeHashed= append(toBeHashed,out[leftsibling][j])
					}
					for j:=0;j<buffer_size;j++ {
							toBeHashed= append(toBeHashed,out[rightsibling][j])
					}
					hash := kec.Keccak256(toBeHashed)
					for j:=0;j<hashFunctionOutputBitSize;j++{
						out[i]=append(out[i],hash[j])
					}
			}

			//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
			if(Op==3) {
				for j:=0;j<hashFunctionOutputBitSize;j++{
					if (j==0){
							if(out[i-1][j]==MRx[j]){
								out[i]=append(out[i],1)
							} else {
								out[i]=append(out[i],0)
										 }
							} else {
								if(out[i-1][j]==MRx[j]){
									out[i]=append(out[i],0)
								}	else{
								out[i]=append(out[i],1)
								}
					}
				}
			}

			if(Op==1) {
					for j:=0;j<buffer_size;j++{
						out[i]=append(out[i],byte(0))
					}
			}

			temp:=Enc(key,out[i])
			for j:=0;j<hashFunctionOutputBitSize;j++{
				z[i]=append(z[i],temp[j])
			}
		}


	for i:=noOfInputGates+noOfInputGates/2;i<totalNumberOfGates;i++{
			Op:=int(gateTuples[i][1])  //Operation of gate by the index "i"

			//Operation 2: The gate takes in the output of two of its children, concatenates the two inputs and hashes it.Later the output is encrypted.
			if(Op==2) {
					leftsibling:=int(gateTuples[i][2])
					rightsibling:=int(gateTuples[i][3])
					var toBeHashed []byte
					for j:=0;j<hashFunctionOutputBitSize;j++ {
							toBeHashed= append(toBeHashed,out[leftsibling][j])
					}
					for j:=0;j<hashFunctionOutputBitSize;j++ {
							toBeHashed= append(toBeHashed,out[rightsibling][j])
					}
					hash := kec.Keccak256(toBeHashed)
					for j:=0;j<hashFunctionOutputBitSize;j++{
						out[i]=append(out[i],hash[j])
					}
			}

			//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
			if(Op==3) {
				for j:=0;j<hashFunctionOutputBitSize;j++{
					if (j==0){
							if(out[i-1][j]==MRx[j]){
								out[i]=append(out[i],1)
							} else {
								out[i]=append(out[i],0)
										 }
							} else {
								if(out[i-1][j]==MRx[j]){
									out[i]=append(out[i],0)
								}	else{
								out[i]=append(out[i],1)
								}
					}
				}
			}

			if(Op==1) {
					for j:=0;j<hashFunctionOutputBitSize;j++{
						out[i]=append(out[i],byte(0))
					}
			}

			temp:=Enc(key,out[i])

			for j:=0;j<hashFunctionOutputBitSize;j++{
				z[i]=append(z[i],temp[j])
			}
		}

	return z,out
}

func initializeCircuitTuples_asAnArray() ([totalNumberOfGates][32]byte){


			var gateTuples [totalNumberOfGates][32]byte

			paddingNumber:=totalNumberOfGates
			//inputGates
			for i:=0; i<noOfInputGates; i++ {
				gateTuples[i][0]=byte(i)
				gateTuples[i][1]=1
				noOfInputsToTheGate:=0
					for j:=2; j<2+noOfInputsToTheGate; j++ {

					}
					for k:=2+noOfInputsToTheGate; k<32; k++ {
						gateTuples[i][k]=byte(paddingNumber)
					}
			}

			//middleGates
			l:=0
			for i:=noOfInputGates; i<totalNumberOfGates-1; i++ {
				gateTuples[i][0]=byte(i)
				gateTuples[i][1]=2
				noOfInputsToTheGate:=2
					for j:=2; j<2+noOfInputsToTheGate; j++ {
							gateTuples[i][j]=byte(l)
							l++
					}
					for k:=2+noOfInputsToTheGate;  k<32; k++ {
						gateTuples[i][k]=byte(paddingNumber)
					}
			}

			//Lastgate
			gateTuples[totalNumberOfGates-1][0]=totalNumberOfGates-1;
			gateTuples[totalNumberOfGates-1][1]=3
			noOfInputsToTheGate:=1;
				for j:=2; j<2+noOfInputsToTheGate; j++ {
						gateTuples[totalNumberOfGates-1][j]=totalNumberOfGates-2
				}
				for k:=2+noOfInputsToTheGate; k<32; k++ {
					gateTuples[totalNumberOfGates-1][k]=byte(paddingNumber)
				}

			return gateTuples

}


func maliciousInitializeCircuitTuples_asAnArray() ([totalNumberOfGates][32]byte){


			var gateTuples [totalNumberOfGates][32]byte

			paddingNumber:=totalNumberOfGates
			//inputGates
			for i:=0; i<noOfInputGates; i++ {
				gateTuples[i][0]=byte(i)
				gateTuples[i][1]=1
				noOfInputsToTheGate:=0
					for j:=2; j<2+noOfInputsToTheGate; j++ {

					}
					for k:=2+noOfInputsToTheGate; k<32; k++ {
						gateTuples[i][k]=byte(paddingNumber)
					}
			}

			//middleGates
			l:=0
			for i:=noOfInputGates; i<totalNumberOfGates-1; i++ {
				gateTuples[i][0]=byte(i)
				gateTuples[i][1]=2
				noOfInputsToTheGate:=2
					for j:=2; j<2+noOfInputsToTheGate; j++ {
							gateTuples[i][j]=byte(l)
							l++
					}
					for k:=2+noOfInputsToTheGate;  k<32; k++ {
						gateTuples[i][k]=byte(paddingNumber)
					}
			}

			//Lastgate
			gateTuples[totalNumberOfGates-1][0]=totalNumberOfGates-1;
			gateTuples[totalNumberOfGates-1][1]=3
			noOfInputsToTheGate:=1;
				for j:=2; j<2+noOfInputsToTheGate; j++ {
						gateTuples[totalNumberOfGates-1][j]=totalNumberOfGates-2
				}
				for k:=2+noOfInputsToTheGate; k<32; k++ {
					gateTuples[totalNumberOfGates-1][k]=byte(paddingNumber)
				}

			gateTuples[noOfInputGates+1][1]=1

			return gateTuples

}


func keyGenerate() ([32]byte){

		keyGen := make([]byte, 32)
		var key [32]byte

	  _, err := rand.Read(keyGen)
    if err != nil {
        log.Fatal(err)
    }

		for i:=0;i<len(keyGen);i++{
			key[i]=keyGen[i]
		}
		return key
}

func generateFinalKey(key [32]byte,keySizeNeeded int) ([]byte) {
	var keyHashed []byte
	var temp [32]byte
	var finalKey []byte
	finalKey = make([]byte, keySizeNeeded)

	for i:=0;i<keySizeNeeded/32;i++ {

		if(i==0){
			var keyPlusIndex []byte

			k:=0
			for j:=0;j<32;j++{
				keyPlusIndex=append(keyPlusIndex,key[j])
			}
			keyPlusIndex=append(keyPlusIndex,byte(k))
			keyHashed = kec.Keccak256(keyPlusIndex)

			for j:=0;j<32;j++{
				temp[j]=keyHashed[j]
				finalKey[j]=temp[j]
			}

		}else{

			keyHashed=nil

			for j:=0;j<32;j++{
				keyHashed=append(keyHashed,temp[j])
			}
			keyHashed=kec.Keccak256(keyHashed)

			for j:=0;j<32;j++{
				temp[j]=keyHashed[j]
				finalKey[i*32+j]=temp[j]
			}
		}
	}
	return finalKey
}

func Enc(key [32]byte,plainText []byte) ([]byte) {

			finalKey:=generateFinalKey(key,len(plainText))
			var encryptedtext []byte
			for i:=0;i<len(plainText);i++{
				encryptedtext=append(encryptedtext,finalKey[i]^plainText[i])
			}
			return encryptedtext
}

func Decrypt(key [32]byte,encryptedtext []byte) ([]byte){

			finalKey:=generateFinalKey(key,len(encryptedtext))
			var plainText []byte
			for i:=0;i<len(encryptedtext);i++{
				plainText=append(plainText,finalKey[i]^encryptedtext[i])
			}

			return plainText
}

func Hash(_toBeHashed []byte) ([32]byte){

	var toBeHashed []byte
	var Hashed [hashFunctionOutputBitSize]byte

	for j:=0;j<len(_toBeHashed);j++{
		toBeHashed=append(toBeHashed,_toBeHashed[j])
	}

	HashedAfter:=kec.Keccak256(toBeHashed)

	for j:=0;j<hashFunctionOutputBitSize;j++{
			Hashed[j]=HashedAfter[j]
	}

	return Hashed

}


func createMerkleRootForInputVectors(inputVectors [noOfInputGates][buffer_size]byte) ([hashFunctionOutputBitSize]byte){

	var merkletree [2*noOfInputGates-1][]byte

	//input gates
	for i:=0;i<noOfInputGates;i++{
		for j:=0;j<buffer_size;j++{
				merkletree[i]=append(merkletree[i],inputVectors[i][j])
		}
	}

	//!inputgates
	k:=0

	for i:=noOfInputGates;i<noOfInputGates+noOfInputGates/2;i++{

		var toBeHashed []byte
		var Hashed []byte

		for j:=0;j<buffer_size;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<buffer_size;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed)

		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i]=append(merkletree[i],Hashed[j])
		}

		k=k+2

	}


	for i:=noOfInputGates+noOfInputGates/2;i<2*noOfInputGates-1;i++{

		var toBeHashed []byte
		var Hashed []byte

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed)

		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i]=append(merkletree[i],Hashed[j])
		}
		k=k+2
	}


	var merkleRoot [hashFunctionOutputBitSize]byte
	for j:=0;j<hashFunctionOutputBitSize;j++{
		merkleRoot[j]=merkletree[2*noOfInputGates-2][j]
	}
	return merkleRoot
}



func createMerkleRootForEncInp(encryptedGateOutputs [totalNumberOfGates][]byte) ([hashFunctionOutputBitSize]byte){

	var merkletree [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
		for j:=0;j<buffer_size;j++{
				merkletree[i]=Hash(encryptedGateOutputs[i])
		}
	}

	//!inputgates
	//merkle tree construction
	k:=0
	for i:=0;i<totalNumberOfGates-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input

		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i+totalNumberOfGates][j]=Hashed[j]
		}
		k=k+2
	}

	  //returns the merkleroot
	return merkletree[2*totalNumberOfGates-2]
}


func createMerkleTreeForEncInp(encryptedGateOutputs [totalNumberOfGates][]byte) ([2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte){

	var merkletree [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
		for j:=0;j<buffer_size;j++{
				merkletree[i]=Hash(encryptedGateOutputs[i])
		}
	}

	//!inputgates
	//merkle tree construction
	k:=0
	for i:=0;i<totalNumberOfGates-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input

		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i+totalNumberOfGates][j]=Hashed[j]
		}
		k=k+2
	}

	  //returns the merkleroot
	return merkletree
}



func checkMerkleRootForEncInp(encryptedGateOutputs [totalNumberOfGates][]byte,merklerootforEncInput [hashFunctionOutputBitSize]byte) (bool){

	var merkletree [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
		for j:=0;j<buffer_size;j++{
				merkletree[i]=Hash(encryptedGateOutputs[i])
		}
	}

	//!inputgates
	//merkle tree construction
	k:=0
	for i:=0;i<totalNumberOfGates-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input

		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i+totalNumberOfGates][j]=Hashed[j]
		}
		k=k+2
	}

	for j:=0;j<hashFunctionOutputBitSize;j++{
		if(merkletree[2*totalNumberOfGates-2][j]!=merklerootforEncInput[j]){
			return false
		}
	}

	return true

}

func checkMerkleRootForGateTuples(gateTuples [totalNumberOfGates][gateTupleSize]byte,merkleRootOfGateTuples [hashFunctionOutputBitSize]byte) (bool){

	var merkletree [2*totalNumberOfGates-1][gateTupleSize]byte

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
		for j:=0;j<gateTupleSize;j++{
				merkletree[i][j]=gateTuples[i][j]
		}
	}

	//!inputgates
	//merkle tree construction
	k:=0
	for i:=0;i<totalNumberOfGates-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input

		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i+totalNumberOfGates][j]=Hashed[j]
		}
		k=k+2
	}

	for j:=0;j<hashFunctionOutputBitSize;j++{
	 if(merkletree[2*totalNumberOfGates-2][j]!=merkleRootOfGateTuples[j]){
		 return false
	 }
 }

 	return true
}

func createMerkleRootForGateTuples(gateTuples [totalNumberOfGates][gateTupleSize]byte) ([hashFunctionOutputBitSize]byte){

	var merkletree [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
		for j:=0;j<gateTupleSize;j++{
				merkletree[i][j]=gateTuples[i][j]
		}
	}

	//!inputgates
	//merkle tree construction
	k:=0
	for i:=0;i<totalNumberOfGates-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input

		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i+totalNumberOfGates][j]=Hashed[j]
		}
		k=k+2
	}

	  //returns the merkleroot
	return merkletree[2*totalNumberOfGates-2]
}

func createMerkleTreeForGateTuples(gateTuples [totalNumberOfGates][gateTupleSize]byte) ([2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte){

	var merkletree [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
		for j:=0;j<gateTupleSize;j++{
				merkletree[i][j]=gateTuples[i][j]
		}
	}

	//!inputgates
	//merkle tree construction
	k:=0
	for i:=0;i<totalNumberOfGates-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input

		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i+totalNumberOfGates][j]=Hashed[j]
		}
		k=k+2
	}

	  //returns the merkleroot
	return merkletree
}



func Extract(encodedGateOutputs [totalNumberOfGates][]byte,key [32]byte,gateTuples [totalNumberOfGates][gateTupleSize]byte,merkleTreeForEncGateOutput [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte,merkleTreeForGateTuples [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte,Receiver_MerkleRootOfEncInp [hashFunctionOutputBitSize]byte,Receiver_MerkleRootOfGateTuples [hashFunctionOutputBitSize]byte) ([][][hashFunctionOutputBitSize]byte,[totalNumberOfGates][]byte,[][hashFunctionOutputBitSize]byte,[][]byte) {

	var decodedOutputs [totalNumberOfGates][]byte
	var encodedVectors_ThisGate  [][]byte
	 //to store all the decodedOutputs. Dynamic in size due to the fact that a wrong gate operation can lead to stopping the decoding process and
	//generating the complaint
	var complain [][][hashFunctionOutputBitSize]byte
	var out []byte
	var gateComplain [][hashFunctionOutputBitSize]byte

	//Decryption of input vectors. Have taken the consumption that an error cannot be generated in the input gates. Even if the sender has something fishy in the
	//input gates, this effect will be refelected in the next level of gates.So, a relevant complaint can be given to the contract.
	for i:=0;i<noOfInputGates;i++ {
			out =Decrypt(key,encodedGateOutputs[i])
			for j:=0;j<len(out);j++{
				decodedOutputs[i] = append(decodedOutputs[i],out[j])
			}
	}

	//Decryption of not input gates - gates above the base level.
	for i:=noOfInputGates;i<totalNumberOfGates;i++{
		out = Decrypt(key,encodedGateOutputs[i])

		for j:=0;j<len(out);j++{
			decodedOutputs[i] = append(decodedOutputs[i],out[j])
		}

		operation:=int(gateTuples[i][1])

		var operationInputs [][]byte

		for k:=2;k<2+maxLinesToGate;k++{

			if(gateTuples[i][k]==totalNumberOfGates){
				break
			}
			operationInputs=append(operationInputs,decodedOutputs[gateTuples[i][k]])
		}

						if(!reflect.DeepEqual(decodedOutputs[i],Operation(operation,operationInputs,decodedOutputs,Receiver_MerkleRootOfEncInp)))	 {
						encodedVectors_ThisGate=append(encodedVectors_ThisGate,encodedGateOutputs[i])
						gateComplain=Mproof(i,merkleTreeForGateTuples)
						//merkle proof for the output of a complaint gate
						MproofTree:=Mproof(i,merkleTreeForEncGateOutput)
						complain=append(complain,MproofTree)

						//merkle proof for the inputs of a complaint gate and appending to the complain vector
						for k:=2;k<2+maxLinesToGate;k++{

							if gateTuples[i][k]==totalNumberOfGates {
								break
							}
								encodedVectors_ThisGate=append(encodedVectors_ThisGate,encodedGateOutputs[(gateTuples[i][k])])
								MproofTree:=Mproof(int(gateTuples[i][k]),merkleTreeForEncGateOutput)
								complain=append(complain,MproofTree)
						}
						return complain,decodedOutputs,gateComplain,encodedVectors_ThisGate

					}
			}
	return complain,decodedOutputs,gateComplain,encodedVectors_ThisGate
}


func Mproof(index int,merkletree [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte) ([][hashFunctionOutputBitSize]byte){

	depth:=int(mathlib.Log2(totalNumberOfGates))+1
	tree:=make([][32]byte,depth)
	tree[0]=merkletree[index]
	for i:=1;i<depth;i++{
		if(index%2==0){
			tree[i]=merkletree[index+1]
			index=index/2+totalNumberOfGates
		} else {
			tree[i]=merkletree[index-1]
			index=index/2+totalNumberOfGates
		}
	}
	return tree

}


func Operation(Op int,operationInputs [][]byte,decodedOutputs [totalNumberOfGates][]byte,Receiver_MerkleRootOfEncInp [hashFunctionOutputBitSize]byte)  ([]byte){

		var result []byte
		if(Op==2) {
				leftsibling:=operationInputs[0]
				rightsibling:=operationInputs[1]
				var toBeHashed []byte
				for j:=0;j<len(leftsibling);j++ {
						toBeHashed= append(toBeHashed,leftsibling[j])
				}
				for j:=0;j<len(rightsibling);j++ {
						toBeHashed= append(toBeHashed,rightsibling[j])
				}

				hashed := kec.Keccak256(toBeHashed)

				for j:=0;j<hashFunctionOutputBitSize;j++{
					result=append(result,hashed[j])
				}
			}

		//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
		if(Op==3) {
			var merkleTreeOfDecryptedInputVectors [noOfInputGates][buffer_size]byte
			for k:=0;k<noOfInputGates;k++{
					for j:=0;j<buffer_size;j++{
						merkleTreeOfDecryptedInputVectors[k][j]=decodedOutputs[k][j]
					}
			}
			merkleRootofDecodedInputVectors:=createMerkleRootForInputVectors(merkleTreeOfDecryptedInputVectors)
			for j:=0;j<hashFunctionOutputBitSize;j++{
				if (j==0){
						if(operationInputs[0][j]==merkleRootofDecodedInputVectors[j]){
							result=append(result,1)
						} else {
							result=append(result,0)
									 }
						} else {
							if(operationInputs[0][j]==merkleRootofDecodedInputVectors[j]){
								result=append(result,0)
							}	else{
								result=append(result,1)
							}
				}
			}
		}
		return result
}

func Mverify(complaint [][32]byte,_index int,ciphertextRoot [32]byte) (bool){

	_value:=complaint[0]
	depth:=uint(len(complaint))
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

func maliciousSender_WrongEncryption(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,channel_SenToMain chan string){


	var currentByte int64 = 0
	fileBuffer := make([]byte, buffer_size)
	file, err := os.Open("a.txt") // For read access.
	if err != nil {
			log.Fatal(err)
	}
	var inputVectors [noOfInputGates][buffer_size]byte
	count:=0
	for {
			n,err := file.ReadAt(fileBuffer, currentByte)
			currentByte += buffer_size
			fileBuffer=fileBuffer[:n]
			if count==noOfInputGates{
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

	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct

	gateTuples:=initializeCircuitTuples_asAnArray()
	MRgateTuples:=createMerkleRootForGateTuples(gateTuples)
	maliciousGateTuples:=maliciousInitializeCircuitTuples_asAnArray()
	key:=keyGenerate()
	keyCommit:=fnKeycommit(key)
	MRinp:=createMerkleRootForInputVectors(inputVectors)
	encodedGateOutputs,_:=Encode(inputVectors,key,maliciousGateTuples,MRinp)
	MRencout:=createMerkleRootForEncInp(encodedGateOutputs)
	recEntryKey:=keyGenerate()

	setPrice(&SenToConObject)
	setTotalNumberOfGates(&SenToConObject)
	setMaxLinesToAGate(&SenToConObject)
	SenToConObject.Keycommit=keyCommit
	SenToConObject.MerkleRootOfEncInp=MRencout
	SenToConObject.MerkleRootOfCircuit=MRgateTuples
	SenToConObject.ConReceiverEntryKey=recEntryKey

	contractAddress, _,_, err=import_contracts.DeployFairswap(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.Keycommit,
		SenToConObject.MerkleRootOfEncInp,
		SenToConObject.MerkleRootOfCircuit,
		SenToConObject.ConReceiverEntryKey,
		&SenToConObject.TotalNumberOfGates,
		&SenToConObject.MaxInputLinesToAGate,
	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()

	SenToRecObject.EncryptedOutputOfGates=encodedGateOutputs
	SenToRecObject.ReceiverEntryKey=recEntryKey
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}

	channel_SenToRec<-byteArrayFromSenToRec
	fmt.Println("\n->Sender \nByteArray sent to Receiver")
	fmt.Println("\n->Sender\n Initialization Finished")

	//creating a contract instance
	instance,err:= import_contracts.NewFairswap(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	RevealedBool:=false
	for i:=0;i<waitingTime;i++ {
		Phase,err:=instance.Phase(nil)
		if err!=nil{
			log.Fatal(err)
		}
		if(Phase==2){
			RevealedBool=true
			break;
		}
		if(Phase==6){
			fmt.Println("\n->Sender\nReceiver has not paid enough ether")
			channel_SenToMain<-"Malicious Receiver-Dind't pay enough ether "
			break;
		}
		time.Sleep(time.Second *1)

	}

	Phase,err:=instance.Phase(nil)
	if err!=nil{
			log.Fatal(err)
	}
	if(Phase==1){
		fmt.Println("\n->Sender\nReceiver hasnt accepted within time")
		channel_SenToMain<-"Either the reciever is not interested now"
	}else{

			if(!RevealedBool){
				channel_SenToMain<-"Receiver has not accepted within the timeout, Key not Revealed"

			}else{

							_,err=instance.RevealKey(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},key)
							if err!=nil{
								log.Fatal(err)
							}
							client.Commit()
					}

			//The contract is in keyRevealed stage
			//Once the sender has revealed the key, the receiver has time until waitingTime to either register a complaint or
			//walk away(not do any execution or transaction--got the correct file)
			//case 1 : If receiver gives a wrong complaint then the contract stage goes to finishedAndMaliciousReciever. The sender sends
			//				"Receiver Wrong complaint" to the main function and finishes the protocol.
			//case 2 : If the receiver doesn't respond, then the stage contract stage is still in keyRevealed and the sender gets the ether
			//				 after the waitingTime
			for i:=0;i<waitingTime;i++ {

					Phase,err:=instance.Phase(nil)
					if err!=nil{
						log.Fatal(err)
					}
					if(Phase==6){
						fmt.Println("\n->Sender\nReceiver has made a wrong complaint")
						channel_SenToMain<-"Receiver Wrong complaint "
						break;
					}
					if(Phase==5){
						fmt.Println("\n->Sender\nIt seems I have made a mistake")
						channel_SenToMain<-"Sender has wrongly encrypted filechunks"
						break
					}
					time.Sleep(time.Second *1)
			}

			gasLimit := uint64(2100000000)

			//the sender getting the ether
			Phase,err:=instance.Phase(nil)
			if err!=nil{
				log.Fatal(err)
			}

			if(Phase==3){
			_,err=instance.SenderGetEther(&bind.TransactOpts{
					From:authSender.From,
					Signer:authSender.Signer,
					Value: big.NewInt(10000000000),
					GasLimit:gasLimit,
				})
			if err!=nil{
						log.Fatal(err)
			}

			//the protocol gets finished and the sender sends the message to the main function
			channel_SenToMain<-"Protocol2 completed.Sender Got the money"
			}
	}
}

func maliciousSender_WrongKeyReveal(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,channel_SenToMain chan string){


	var currentByte int64 = 0
	fileBuffer := make([]byte, buffer_size)
	file, err := os.Open("a.txt") // For read access.
	if err != nil {
			log.Fatal(err)
	}
	var inputVectors [noOfInputGates][buffer_size]byte
	count:=0
	for {
			n,err := file.ReadAt(fileBuffer, currentByte)
			currentByte += buffer_size
			fileBuffer=fileBuffer[:n]
			if count==noOfInputGates{
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

	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct

	gateTuples:=initializeCircuitTuples_asAnArray()
	MRgateTuples:=createMerkleRootForGateTuples(gateTuples)
	maliciousGateTuples:=maliciousInitializeCircuitTuples_asAnArray()
	key:=keyGenerate()
	keyCommit:=fnKeycommit(key)
	MRinp:=createMerkleRootForInputVectors(inputVectors)
	encodedGateOutputs,_:=Encode(inputVectors,key,maliciousGateTuples,MRinp)
	MRencout:=createMerkleRootForEncInp(encodedGateOutputs)
	recEntryKey:=keyGenerate()

	setPrice(&SenToConObject)
	setTotalNumberOfGates(&SenToConObject)
	setMaxLinesToAGate(&SenToConObject)
	SenToConObject.Keycommit=keyCommit
	SenToConObject.MerkleRootOfEncInp=MRencout
	SenToConObject.MerkleRootOfCircuit=MRgateTuples
	SenToConObject.ConReceiverEntryKey=recEntryKey

	contractAddress, _,_, err=import_contracts.DeployFairswap(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.Keycommit,
		SenToConObject.MerkleRootOfEncInp,
		SenToConObject.MerkleRootOfCircuit,
		SenToConObject.ConReceiverEntryKey,
		&SenToConObject.TotalNumberOfGates,
		&SenToConObject.MaxInputLinesToAGate,
	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()

	SenToRecObject.EncryptedOutputOfGates=encodedGateOutputs
	SenToRecObject.ReceiverEntryKey=recEntryKey
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}

	channel_SenToRec<-byteArrayFromSenToRec
	fmt.Println("\n->Sender \nByteArray sent to Receiver")
	fmt.Println("\n->Sender\n Initialization Finished")

	//creating a contract instance
	instance,err:= import_contracts.NewFairswap(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	RevealedBool:=false
	for i:=0;i<waitingTime;i++ {
		Phase,err:=instance.Phase(nil)
		if err!=nil{
			log.Fatal(err)
		}
		if(Phase==2){
			RevealedBool=true
			break;
		}
		if(Phase==6){
			fmt.Println("\n->Sender\nReceiver has not paid enough ether")
			channel_SenToMain<-"Malicious Receiver-Dind't pay enough ether "
			break;
		}
		time.Sleep(time.Second *1)

	}

	Phase,err:=instance.Phase(nil)
	if err!=nil{
			log.Fatal(err)
	}
	if(Phase==1){
		fmt.Println("\n->Sender\nReceiver hasnt accepted within time")
		channel_SenToMain<-"Either the reciever is not interested now"
	}else{

			if(!RevealedBool){
				channel_SenToMain<-"Receiver has not accepted within the timeout, Key not Revealed"

			}else{

							var _maliciousKey [32]byte
							_,err=instance.RevealKey(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},_maliciousKey)
							if err!=nil{
								log.Fatal(err)
							}
							client.Commit()
							fmt.Println("\n->Sender :\n I have revealed the wrong key")
					}
			channel_SenToMain<-"Sender Revealed the wrong key"
	}
}


func maliciousReceiver_WrongComplain(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,channel_RecToMainIni chan string,channel_RecToMainRev chan string){

	byteArrayFromSenToRec:=<-channel_SenToRec

	var Receiver_SenToRecObj SenderToReceiverStruct
	err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
	if err != nil {
			log.Fatal(err)
	}

	instance,err := import_contracts.NewFairswap(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	Receiver_MerkleRootOfEncInp,err:=instance.EncInputRoot(nil)
	if err!=nil{
		log.Fatal(err)
	}

	Receiver_MerkleRootOfGateTuples,err:=instance.EncCircuitRoot(nil)
	if err!=nil{
		log.Fatal(err)
	}

	Receiver_EncryptedOutputOfGates:=Receiver_SenToRecObj.EncryptedOutputOfGates
	gateTuples:=initializeCircuitTuples_asAnArray()

	boolValueEncryptedGateOutput:=checkMerkleRootForEncInp(Receiver_EncryptedOutputOfGates,Receiver_MerkleRootOfEncInp)
	boolValueEncryptedGateTuples:=checkMerkleRootForGateTuples(gateTuples,Receiver_MerkleRootOfGateTuples)

	fmt.Println("\n->Receiver\nboolValueEncryptedGateOutput(comparision between merkleroot of encryptedGateOutputs from contract and calulated encryptedGateOutputs merkleroot): ",boolValueEncryptedGateOutput)
	fmt.Println("\n->Receiver\nboolValueEncryptedGateTuples(comparision between merkleroot of encryptedGateTuples from contract and calulated encryptedGateTuples merkleroot): ",boolValueEncryptedGateTuples)

	if(boolValueEncryptedGateOutput && boolValueEncryptedGateTuples){

				//the boolValueEncode was true, so the receiver calls the accept contract function and pays the required price
			_,err=instance.InitializeRecieverAddress(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value: nil,
			},Receiver_SenToRecObj.ReceiverEntryKey)
			if err!=nil{
					log.Fatal(err)
			}

			client.Commit()

			value:=big.NewInt(100)
			gasLimit := uint64(230000000)

			_	,err=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
				GasLimit:gasLimit,
			})
			if err != nil {
					log.Fatal(err)
			}

			client.Commit()

			fmt.Println("\n->Receiver\n Finished Accepting")
			channel_RecToMainIni<-"Accepted stage got over"


			//The contract is in accepted stage
			//The sender
			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
			//case 2 : The sender doesn't reveal the key within time. AcceptedBool remains false after the for loop. Receiver later calls the contract function
			//			   ReceiverGetEther to get the ether.
			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and AcceptedBool is true after the for loop.
			AcceptedBool:=false
			var contract_Key [keyLength]byte
			for i:=0;i<waitingTime;i++ {
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

				if(Phase==5){
					fmt.Println("\n->Receiver\nSender Has revealed the wrong key")
					channel_RecToMainRev<-"Wrong key revealed"

				}
				time.Sleep(time.Second *1)

			}


			if(!AcceptedBool){

					fmt.Println("\n->Receiver\nSender has not revealed the key...ReceiverGetEther function called")
					gasLimit := uint64(2100000000)

					//the receiver getting the ether
					_,err=instance.ReceiverGetEther(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value: nil,
							GasLimit:gasLimit,
						})
					if err!=nil{
								log.Fatal(err)
					}
					client.Commit()

					channel_RecToMainRev<-"Sender hasn't revealed the key within the Timeout"

			} else{

				fmt.Println("\n->Receiver\nSender has revealed the key....Proceeding further")

				merkleTreeForGateTuples:=createMerkleTreeForGateTuples(gateTuples)
			 	merkleTreeForEncInp:=createMerkleTreeForEncInp(Receiver_SenToRecObj.EncryptedOutputOfGates)

				maliciousGateTuples:=maliciousInitializeCircuitTuples_asAnArray()

			 	complain,decodedOutputs,gateComplain,encodedVectors_ThisGate :=Extract(Receiver_SenToRecObj.EncryptedOutputOfGates,contract_Key,maliciousGateTuples,merkleTreeForEncInp,merkleTreeForGateTuples,Receiver_MerkleRootOfEncInp,Receiver_MerkleRootOfGateTuples)

				fmt.Println("\n->Receiver\nDecoded Ouptuts : ",decodedOutputs)

				if(complain==nil){
					fmt.Println("\n->Receiver\nNo complain")
					channel_RecToMainRev<-"Receiver Finished Protocol "
				} else{

					var malicious_encodedVectors_ThisGate [][]byte
					for j:=0;j<len(encodedVectors_ThisGate);j++{
						malicious_encodedVectors_ThisGate=append(malicious_encodedVectors_ThisGate,encodedVectors_ThisGate[len(encodedVectors_ThisGate)-1-j])
					}

						//complain function being called
						_,err=instance.Complain(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value:nil,
							GasLimit:gasLimit,
						},complain,gateComplain,encodedVectors_ThisGate)
						if err!=nil{
							log.Fatal(err)
						}
						client.Commit()

						Phase,err:=instance.Phase(nil)
						if err!=nil{
							log.Fatal(err)
						}
						if(Phase==5){
							fmt.Println("\n->Receiver\nSender Has wrongly calulated the encrypted inputs")
							channel_RecToMainRev<-"Wrong encrypted inputs"

						}
						channel_RecToMainRev<-"Receiver Made a wrong complaint"
				}
			}
			channel_RecToMainRev<-"Reveal Phase Finished"
		}	else{

				channel_RecToMainIni<-"Sender sent wrong encodedGateOutputs,not equal to contract encodedGateOutputs merkleroot,Not gone into accepted stage"
				channel_RecToMainRev<-"Sender sent wrong encodedGateOutputs,not equal to contract encodedGateOutputs merkleroot,Not gone into revealed phase"
		}
}
