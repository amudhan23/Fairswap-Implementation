package main

import(
	"fmt"
	kec "github.com/ethereum/go-ethereum/crypto"
	"os"
	"io"
	"log"
	"reflect"
	"time"
	"encoding/json"
	"sync"
	"strconv"
	//"github.com/ethereum/go-ethereum/core"
	mathlib "math"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/accounts/keystore"
	//"github.com/ethereum/go-ethereum/ethclient"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	import_contracts "./contracts"

)

//the following parameters are got as command line arguments
// var buffer_size int
// var noOfFileChunks int


const waitingTime=40  //max time to complete a stage
const hashFunctionOutputBitSize=32 //hash function's output size(32 bytes in this case)
const gateTupleSize=32
const keySize=32
const maxLinesToGate=2

var totalNumberOfGates int
var noOfInputGates int
var buffer_size int
var noOfFileChunks int

var contractAddress []common.Address

type SenderToReceiverStruct struct {
		EncryptedOutputOfGates [][]byte
}

type SenderToContractStruct struct{
	Price big.Int
	Keycommit [hashFunctionOutputBitSize]byte
	MerkleRootOfCircuit [hashFunctionOutputBitSize]byte
	MerkleRootOfEncInp [hashFunctionOutputBitSize]byte
	TotalNumberOfGates big.Int
	MaxInputLinesToAGate big.Int
}

func main(){

	var noOfPairsOfBuyersAndSellers int
	var logFilePath string
	logFilePath="/home/amudhan/protocols/Fairswap/text6.log"
	if(len(os.Args)<6){
		buffer_size=32
		noOfFileChunks=8
		noOfInputGates=noOfFileChunks
		totalNumberOfGates=noOfInputGates*2
		logFilePath="./text5.log"
		noOfPairsOfBuyersAndSellers=1
	}else{
		var err error
		buffer_size, err = strconv.Atoi(os.Args[1])
		if(err!=nil){
			log.Fatal(err)
		}
		noOfFileChunks, err = strconv.Atoi(os.Args[2])
		if(err!=nil){
			log.Fatal(err)
		}
		noOfPairsOfBuyersAndSellers,err=strconv.Atoi(os.Args[3])
		if(err!=nil){
			log.Fatal(err)
		}
		noOfInputGates=noOfFileChunks
		totalNumberOfGates=noOfInputGates*2
		logFilePath=os.Args[4]
	}

	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
			fmt.Println("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)


	contractAddress=make([]common.Address,noOfPairsOfBuyersAndSellers)

	alloc := make(core.GenesisAlloc)

	authSenders:=make([]*bind.TransactOpts,noOfPairsOfBuyersAndSellers)
	authReceivers:=make([]*bind.TransactOpts,noOfPairsOfBuyersAndSellers)
	//log.Println("\n->Sender\nSender Address :",senderAddress)

	for i:=0;i<noOfPairsOfBuyersAndSellers;i++{
		authSenders[i],_ = AuthAndAddressGeneration()
		authReceivers[i],_= AuthAndAddressGeneration()
		alloc[authSenders[i].From] = core.GenesisAccount{Balance: big.NewInt(11000000000000000)}  //1.1 ether
		alloc[authReceivers[i].From] = core.GenesisAccount{Balance: big.NewInt(250000000000000000)} //2.5 ether
	}

	//client
	client := backends.NewSimulatedBackend(alloc,600000000000000000)

	//go routines

	channel_SenToRec:=make([]chan []byte,noOfPairsOfBuyersAndSellers)
		for i := range channel_SenToRec {
	 		channel_SenToRec[i] = make(chan []byte)
	 }

	var wg sync.WaitGroup
	wg.Add(2*noOfPairsOfBuyersAndSellers)
	for i:=0;i<noOfPairsOfBuyersAndSellers;i++{
		pairIndex:=i
		go Sender(authSenders[i],client,channel_SenToRec[pairIndex],&wg,pairIndex)
		go Receiver(authReceivers[i],client,channel_SenToRec[pairIndex],&wg,pairIndex)
	}

	wg.Wait()  //Wait blocks until the WaitGroup counter is zero
}


func Sender(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup,pairIndex int){

	sender_round1_computingTime_calculatingVariables_start := time.Now()

	inputTextFilePath:=os.Args[5]
	inputVectors:=FileRead(inputTextFilePath)

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

	setPrice(&SenToConObject)
	setTotalNumberOfGates(&SenToConObject)
	setMaxLinesToAGate(&SenToConObject)
	SenToConObject.Keycommit=keyCommit
	SenToConObject.MerkleRootOfEncInp=MRencout
	SenToConObject.MerkleRootOfCircuit=MRgateTuples

	sender_round1_computingTime_calculatingVariables_stop:=time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round1_computingTime_calculatingVariables_time : %v \n", sender_round1_computingTime_calculatingVariables_stop.Sub(sender_round1_computingTime_calculatingVariables_start))

	var err error
	var tx *types.Transaction
	sender_round1_commitTime_deployingContract_start := time.Now()
  contractAddress[pairIndex],tx,_, err=import_contracts.DeployFairswap(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.Keycommit,
		SenToConObject.MerkleRootOfEncInp,
		SenToConObject.MerkleRootOfCircuit,
		&SenToConObject.TotalNumberOfGates,
		&SenToConObject.MaxInputLinesToAGate,
	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()

	instance,err:= import_contracts.NewFairswap(contractAddress[pairIndex],client)
	if err != nil {
			log.Fatal(err)
	}


	sender_round1_commitTime_deployingContract_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round1_commitTime_deployingContract_time : %v \n", sender_round1_commitTime_deployingContract_stop.Sub(sender_round1_commitTime_deployingContract_start))
	log.SetPrefix("\nFairswap - Gas Cost ")
	log.Println("Sender Round 1 : ",tx.Gas())



	sender_round1_communicationTime_sendingEncryptedGateOutputs_start := time.Now()


	SenToRecObject.EncryptedOutputOfGates=encodedGateOutputs

	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		//log.Println("error:", err)
	}

	channel_SenToRec<-byteArrayFromSenToRec

	sender_round1_communicationTime_sendingEncryptedGateOutputs_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round1_communicationTime_sendingEncryptedGateOutputs_time : %v \n", sender_round1_communicationTime_sendingEncryptedGateOutputs_stop.Sub(sender_round1_communicationTime_sendingEncryptedGateOutputs_start))


	sender_round2_waitingTime_forReceiverToAccept_start := time.Now()



	receiverAcceptedStatus:=false
	for i:=0;i<waitingTime;i++ {

		Phase,err:=instance.Phase(nil)
		if err!=nil{
			log.Fatal(err)
		}
		if(Phase==2){
			receiverAcceptedStatus=true
			break;
		}
		time.Sleep(time.Second *1)

	}


	sender_round2_waitingTime_forReceiverToAccept_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round2_waitingTime_forReceiverToAccept_time : %v \n", sender_round2_waitingTime_forReceiverToAccept_stop.Sub(sender_round2_waitingTime_forReceiverToAccept_start))


	if(!receiverAcceptedStatus){
			//Receiver has not accepted within time
			}else{

						 sender_round3_commitTime_revealKey_start := time.Now()
						 tx,err=instance.RevealKey(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},key)
							if err!=nil{
								log.Fatal(err)
							}
							client.Commit()
							//log.Println("\n->Sender\nKeyRevealed")
							sender_round3_commitTime_revealKey_stop := time.Now()
							log.SetPrefix("\nFairswap - Time ")
							log.Printf("sender_round3_commitTime_revealKey_time : %v \n", sender_round3_commitTime_revealKey_stop.Sub(sender_round3_commitTime_revealKey_start))
							//client.Commit()
							log.SetPrefix("\nFairswap - Gas Cost ")
							log.Println("Sender Round 3 : ",tx.Gas())


			sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_start:=time.Now()

			for i:=0;i<waitingTime;i++{

					Phase,err:=instance.Phase(nil)
					if err!=nil{
						log.Fatal(err)
					}
					if(Phase==4){
						break
					}
					time.Sleep(time.Second *1)
			}
			sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_stop:=time.Now()
			log.SetPrefix("\nFairswap - Time ")
			log.Printf("sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_time : %v \n", sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_stop.Sub(sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_start))

			//gasLimit := uint64(10000000000)
			//the sender getting the ether
			Phase,err:=instance.Phase(nil)
			if err!=nil{
				log.Fatal(err)
			}

			if(Phase==3){

					timeout,err:=instance.Timeout(nil)
					if err!=nil{
									log.Fatal(err)
					}


					for{
						_,err:=instance.Nowy(&bind.TransactOpts{
							 From:authSender.From,
							 Signer:authSender.Signer,
						 })
						 if err!=nil{
									 log.Fatal(err)
						 }
						 client.Commit()

						 Now,err:=instance.Now(nil)
						 if err!=nil{
									 log.Fatal(err)
						 }

						 if(Now.Cmp(timeout)==1){
								break
						 }

						 time.Sleep(time.Second*1)
					}

					sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start:=time.Now()
					tx,err=instance.SenderGetEtherReceiverNoComplainMessage(&bind.TransactOpts{
							From:authSender.From,
							Signer:authSender.Signer,
						})
					if err!=nil{
								log.Fatal(err)
					}
					client.Commit()

					log.SetPrefix("\nFairswap - Gas Cost ")
					log.Println("Sender Round 4 ",tx.Gas())
					sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop:=time.Now()
					log.SetPrefix("\nFairswap - Time ")
					log.Printf("sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_time : %v \n", sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop.Sub(sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start))

			}
	}
	wg.Done()
}


func maliciousSender_WrongKeyReveal(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup,pairIndex int){

	sender_round1_computingTime_calculatingVariables_start := time.Now()

	inputTextFilePath:=os.Args[5]
	inputVectors:=FileRead(inputTextFilePath)

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

	setPrice(&SenToConObject)
	setTotalNumberOfGates(&SenToConObject)
	setMaxLinesToAGate(&SenToConObject)
	SenToConObject.Keycommit=keyCommit
	SenToConObject.MerkleRootOfEncInp=MRencout
	SenToConObject.MerkleRootOfCircuit=MRgateTuples

	sender_round1_computingTime_calculatingVariables_stop:=time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round1_computingTime_calculatingVariables_time : %v \n", sender_round1_computingTime_calculatingVariables_stop.Sub(sender_round1_computingTime_calculatingVariables_start))

	var err error
	var tx *types.Transaction
	sender_round1_commitTime_deployingContract_start := time.Now()
  contractAddress[pairIndex],tx,_, err=import_contracts.DeployFairswap(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.Keycommit,
		SenToConObject.MerkleRootOfEncInp,
		SenToConObject.MerkleRootOfCircuit,
		&SenToConObject.TotalNumberOfGates,
		&SenToConObject.MaxInputLinesToAGate,
	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()

	instance,err:= import_contracts.NewFairswap(contractAddress[pairIndex],client)
	if err != nil {
			log.Fatal(err)
	}

	sender_round1_commitTime_deployingContract_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round1_commitTime_deployingContract_time : %v \n", sender_round1_commitTime_deployingContract_stop.Sub(sender_round1_commitTime_deployingContract_start))
	log.SetPrefix("\nFairswap - Gas Cost ")
	log.Println("Sender Round 1 : ",tx.Gas())

	sender_round1_communicationTime_sendingEncryptedGateOutputs_start := time.Now()
	SenToRecObject.EncryptedOutputOfGates=encodedGateOutputs

	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		//log.Println("error:", err)
	}

	channel_SenToRec<-byteArrayFromSenToRec

	sender_round1_communicationTime_sendingEncryptedGateOutputs_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round1_communicationTime_sendingEncryptedGateOutputs_time : %v \n", sender_round1_communicationTime_sendingEncryptedGateOutputs_stop.Sub(sender_round1_communicationTime_sendingEncryptedGateOutputs_start))


	sender_round2_waitingTime_forReceiverToAccept_start := time.Now()



	receiverAcceptedStatus:=false
	for i:=0;i<waitingTime;i++ {

		Phase,err:=instance.Phase(nil)
		if err!=nil{
			log.Fatal(err)
		}
		if(Phase==2){
			receiverAcceptedStatus=true
			break;
		}
		time.Sleep(time.Second *1)

	}


	sender_round2_waitingTime_forReceiverToAccept_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round2_waitingTime_forReceiverToAccept_time : %v \n", sender_round2_waitingTime_forReceiverToAccept_stop.Sub(sender_round2_waitingTime_forReceiverToAccept_start))


	if(!receiverAcceptedStatus){
			//Receiver has not accepted within time
			}else{

						malicious_key:=keyGenerate()
						 sender_round3_commitTime_revealKey_start := time.Now()
						 tx,err=instance.RevealKey(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},malicious_key)
							if err!=nil{
								log.Fatal(err)
							}
							client.Commit()
							//log.Println("\n->Sender\nKeyRevealed")
							sender_round3_commitTime_revealKey_stop := time.Now()
							log.SetPrefix("\nFairswap - Time ")
							log.Printf("sender_round3_commitTime_revealKey_time : %v \n", sender_round3_commitTime_revealKey_stop.Sub(sender_round3_commitTime_revealKey_start))
							//client.Commit()
							log.SetPrefix("\nFairswap - Gas Cost ")
							log.Println("Sender Round 3 : ",tx.Gas())


			sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_start:=time.Now()

			for i:=0;i<waitingTime;i++{

					Phase,err:=instance.Phase(nil)
					if err!=nil{
						log.Fatal(err)
					}
					if(Phase==4){
						break
					}
					time.Sleep(time.Second *1)
			}
			sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_stop:=time.Now()
			log.SetPrefix("\nFairswap - Time ")
			log.Printf("sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_time : %v \n", sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_stop.Sub(sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_start))


	}
	wg.Done()
}

//channel_RecToMainRev(chan string) - to notify the main function for the completion of reveal phase
func Receiver(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup,pairIndex int){

	receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start := time.Now()
	//log.Println("\n->Receiver\n Begins")
	byteArrayFromSenToRec:=<-channel_SenToRec

	receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_time : %v \n", receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop.Sub(receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start))

	receiver_round1_computingTime_merkleRootCheck_start := time.Now()


	var Receiver_SenToRecObj SenderToReceiverStruct
	err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
	if err != nil {
			log.Fatal(err)
	}


	instance,err:= import_contracts.NewFairswap(contractAddress[pairIndex],client)
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

	receiver_round1_computingTime_merkleRootCheck_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("receiver_round1_computingTime_merkleRootCheck_time : %v \n", receiver_round1_computingTime_merkleRootCheck_stop.Sub(receiver_round1_computingTime_merkleRootCheck_start))

	//log.Println("\n->Receiver\nboolValueEncryptedGateOutput(comparision between merkleroot of encryptedGateOutputs from contract and calulated encryptedGateOutputs merkleroot): ",boolValueEncryptedGateOutput)
	//log.Println("\n->Receiver\nboolValueEncryptedGateTuples(comparision between merkleroot of encryptedGateTuples from contract and calulated encryptedGateTuples merkleroot): ",boolValueEncryptedGateTuples)



	if(boolValueEncryptedGateOutput && boolValueEncryptedGateTuples){

		 value:=big.NewInt(100)

		 receiver_round2_commitTime_Accept_start := time.Now()

			tx,err:=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
			})
			if err != nil {
					log.Fatal(err)
			}
			client.Commit()



			receiver_round2_commitTime_Accept_stop := time.Now()
			log.SetPrefix("\nFairswap - Time ")
			log.Printf("receiver_round2_commitTime_Accept_time : %v \n", receiver_round2_commitTime_Accept_stop.Sub(receiver_round2_commitTime_Accept_start))

			log.SetPrefix("\nFairswap - Gas Cost ")
			log.Println("Receiver Round 2 : ",tx.Gas())


			//log.Println("\n->Receiver\n Accepted")

			//The contract is in accepted stage
			//The sender
			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
			//case 2 : The sender doesn't reveal the key within time. senderKeyRevealedStatus remains false after the for loop. Receiver later calls the contract function
			//			   ReceiverGetEther to get the ether.
			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and senderKeyRevealedStatus is true after the for loop.


		receiver_round3_waitingTime_forSenderToRevealKey_start := time.Now()
			senderKeyRevealedStatus:=false
			protocolFinishedStatus:=false
			var contract_Key [keySize]byte
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
					senderKeyRevealedStatus=true
					break;
				}
				if(Phase==4){
					protocolFinishedStatus=true
					//wrong key Revealed
					break;
				}
				time.Sleep(time.Second *1)

			}
			receiver_round3_waitingTime_forSenderToRevealKey_stop := time.Now()
			log.SetPrefix("\nFairswap - Time ")
			log.Printf("receiver_round3_waitingTime_forSenderToRevealKey_time : %v \n", receiver_round3_waitingTime_forSenderToRevealKey_stop.Sub(receiver_round3_waitingTime_forSenderToRevealKey_start))

			if(protocolFinishedStatus){
				wg.Done()
			}

			if(!senderKeyRevealedStatus){

				 timeout,err:=instance.Timeout(nil)
				 if err!=nil{
							 log.Fatal(err)
				 }

				 for{

						_,err:=instance.Nowy(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
						})
						if err!=nil{
									log.Fatal(err)
						}
						client.Commit()
						Now,err:=instance.Now(nil)
						if err!=nil{
									log.Fatal(err)
						}
						if(Now.Cmp(timeout)==1){
							 break
							}
						time.Sleep(time.Second*1)
				 }

					//the receiver getting the ether
					receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start := time.Now()
					tx,err=instance.ReceiverGetEtherSenderNotRevealedKey(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value: nil,
						})
					if err!=nil{
								log.Fatal(err)
					}
					client.Commit()
					receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop := time.Now()
					log.SetPrefix("\nFairswap - Gas Cost ")
					log.Println("Receiver Round 3 (Sender Not Revealed Key) : ",tx.Cost())
					log.Printf("\nreceiver_round3_commitTime_senderKeyNotRevealedGettingEther_time : %v \n", receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop.Sub(receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start))
					//client.Commit()

			} else{


				receiver_round4_computingTime_extraction_start := time.Now()

				merkleTreeForGateTuples:=createMerkleTreeForGateTuples(gateTuples)
			 	merkleTreeForEncInp:=createMerkleTreeForEncInp(Receiver_SenToRecObj.EncryptedOutputOfGates)

			 	complain,_,gateComplain,encodedVectors_ThisGate :=Extract(Receiver_SenToRecObj.EncryptedOutputOfGates,contract_Key,gateTuples,merkleTreeForEncInp,merkleTreeForGateTuples,Receiver_MerkleRootOfEncInp,Receiver_MerkleRootOfGateTuples)

				receiver_round4_computingTime_extraction_stop := time.Now()
				log.SetPrefix("\nFairswap - Time ")
				log.Printf("receiver_round4_computingTime_extraction_time : %v \n", receiver_round4_computingTime_extraction_stop.Sub(receiver_round4_computingTime_extraction_start))
				//log.Println("\n->Receiver\nDecoded Ouptuts : ",decodedOutputs)

				if(complain==nil){
						log.SetPrefix("\nFairswap - Gas Cost ")
						log.Println("Receiver Round 4 : ",0)
					//log.Println("\n\n->Receiver\n Finished Protocol")
				} else{

						//complain function being called
						receiver_round4_commitTime_complain_start := time.Now()
						tx,err:=instance.Complain(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value:nil,
							//GasLimit:gasLimit,
						},complain,gateComplain,encodedVectors_ThisGate)
						if err!=nil{
							log.Fatal(err)
						}
						client.Commit()

						receiver_round4_commitTime_complain_stop := time.Now()
						log.SetPrefix("\nFairswap - Time ")
						log.Printf("\nreceiver_round4_commitTime_complain_time : %v \n", receiver_round4_commitTime_complain_stop.Sub(receiver_round4_commitTime_complain_start))

						log.SetPrefix("\nFairswap - Gas Cost ")
						log.Println("Receiver Round 4 : ",tx.Cost())

					}
				}
		}	else{

				if(!boolValueEncryptedGateOutput){
						//log.Println("\n->Receiver\nThe calculated merkle root of encrypted gate outputs and merkle root of encrypted gate outputs got from the contract didn't match")
					} else{
							//log.Println("\n->Receiver\nThe calculated merkle root of gate tuples and merkle root of egate tuples got from the contract didn't match")
					}
		}
		wg.Done()
}

func maliciousReceiver_WrongComplain(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup,pairIndex int){

	receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start := time.Now()
	//log.Println("\n->Receiver\n Begins")
	byteArrayFromSenToRec:=<-channel_SenToRec

	receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_time : %v \n", receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop.Sub(receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start))

	receiver_round1_computingTime_merkleRootCheck_start := time.Now()


	var Receiver_SenToRecObj SenderToReceiverStruct
	err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
	if err != nil {
			log.Fatal(err)
	}

	instance,err := import_contracts.NewFairswap(contractAddress[pairIndex],client)
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

	receiver_round1_computingTime_merkleRootCheck_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("receiver_round1_computingTime_merkleRootCheck_time : %v \n", receiver_round1_computingTime_merkleRootCheck_stop.Sub(receiver_round1_computingTime_merkleRootCheck_start))

	//log.Println("\n->Receiver\nboolValueEncryptedGateOutput(comparision between merkleroot of encryptedGateOutputs from contract and calulated encryptedGateOutputs merkleroot): ",boolValueEncryptedGateOutput)
	//log.Println("\n->Receiver\nboolValueEncryptedGateTuples(comparision between merkleroot of encryptedGateTuples from contract and calulated encryptedGateTuples merkleroot): ",boolValueEncryptedGateTuples)

	if(boolValueEncryptedGateOutput && boolValueEncryptedGateTuples){

		 value:=big.NewInt(100)

		 receiver_round2_commitTime_Accept_start := time.Now()

			tx,err:=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
			})
			if err != nil {
					log.Fatal(err)
			}
			client.Commit()

			receiver_round2_commitTime_Accept_stop := time.Now()
			log.SetPrefix("\nFairswap - Time ")
			log.Printf("receiver_round2_commitTime_Accept_time : %v \n", receiver_round2_commitTime_Accept_stop.Sub(receiver_round2_commitTime_Accept_start))

			log.SetPrefix("\nFairswap - Gas Cost ")
			log.Println("Receiver Round 2 : ",tx.Gas())


			//log.Println("\n->Receiver\n Accepted")

			//The contract is in accepted stage
			//The sender
			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
			//case 2 : The sender doesn't reveal the key within time. senderKeyRevealedStatus remains false after the for loop. Receiver later calls the contract function
			//			   ReceiverGetEther to get the ether.
			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and senderKeyRevealedStatus is true after the for loop.


		receiver_round3_waitingTime_forSenderToRevealKey_start := time.Now()
			senderKeyRevealedStatus:=false
			protocolFinishedStatus:=false
			var contract_Key [keySize]byte
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
					senderKeyRevealedStatus=true
					break;
				}
				if(Phase==4){
					protocolFinishedStatus=true
					//wrong key Revealed
					break;
				}
				time.Sleep(time.Second *1)

			}
			receiver_round3_waitingTime_forSenderToRevealKey_stop := time.Now()
			log.SetPrefix("\nFairswap - Time ")
			log.Printf("receiver_round3_waitingTime_forSenderToRevealKey_time : %v \n", receiver_round3_waitingTime_forSenderToRevealKey_stop.Sub(receiver_round3_waitingTime_forSenderToRevealKey_start))

			if(protocolFinishedStatus){
				wg.Done()
			}

			if(!senderKeyRevealedStatus){

				 timeout,err:=instance.Timeout(nil)
				 if err!=nil{
							 log.Fatal(err)
				 }

				 for{

						_,err:=instance.Nowy(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
						})
						if err!=nil{
									log.Fatal(err)
						}
						client.Commit()
						Now,err:=instance.Now(nil)
						if err!=nil{
									log.Fatal(err)
						}
						if(Now.Cmp(timeout)==1){
							 break
							}
						time.Sleep(time.Second*1)
				 }

					//the receiver getting the ether
					receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start := time.Now()
					tx,err=instance.ReceiverGetEtherSenderNotRevealedKey(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value: nil,
						})
					if err!=nil{
								log.Fatal(err)
					}
					client.Commit()
					receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop := time.Now()
					log.SetPrefix("\nFairswap - Gas Cost ")
					log.Println("Receiver Round 3 (Sender Not Revealed Key) : ",tx.Cost())
					log.Printf("\nreceiver_round3_commitTime_senderKeyNotRevealedGettingEther_time : %v \n", receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop.Sub(receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start))
					//client.Commit()

			} else{

				receiver_round4_computingTime_extraction_start := time.Now()

				maliciousGateTuples:=maliciousInitializeCircuitTuples_asAnArray()

				merkleTreeForGateTuples:=createMerkleTreeForGateTuples(gateTuples)
			 	merkleTreeForEncInp:=createMerkleTreeForEncInp(Receiver_SenToRecObj.EncryptedOutputOfGates)

			 	complain,_,gateComplain,encodedVectors_ThisGate :=Extract(Receiver_SenToRecObj.EncryptedOutputOfGates,contract_Key,maliciousGateTuples,merkleTreeForEncInp,merkleTreeForGateTuples,Receiver_MerkleRootOfEncInp,Receiver_MerkleRootOfGateTuples)

				receiver_round4_computingTime_extraction_stop := time.Now()
				log.SetPrefix("\nFairswap - Time ")
				log.Printf("receiver_round4_computingTime_extraction_time : %v \n", receiver_round4_computingTime_extraction_stop.Sub(receiver_round4_computingTime_extraction_start))
				//log.Println("\n->Receiver\nDecoded Ouptuts : ",decodedOutputs)

				if(complain==nil){
						log.SetPrefix("\nFairswap - Gas Cost ")
						log.Println("Receiver Round 4 : ",0)
					//log.Println("\n\n->Receiver\n Finished Protocol")
				} else{

						//complain function being called
						receiver_round4_commitTime_complain_start := time.Now()
						tx,err:=instance.Complain(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value:nil,
							//GasLimit:gasLimit,
						},complain,gateComplain,encodedVectors_ThisGate)
						if err!=nil{
							log.Fatal(err)
						}
						client.Commit()

						receiver_round4_commitTime_complain_stop := time.Now()
						log.SetPrefix("\nFairswap - Time ")
						log.Printf("\nreceiver_round4_commitTime_complain_time : %v \n", receiver_round4_commitTime_complain_stop.Sub(receiver_round4_commitTime_complain_start))

						log.SetPrefix("\nFairswap - Gas Cost ")
						log.Println("Receiver Round 4 : ",tx.Cost())

					}
				}
		}	else{

				if(!boolValueEncryptedGateOutput){
						//log.Println("\n->Receiver\nThe calculated merkle root of encrypted gate outputs and merkle root of encrypted gate outputs got from the contract didn't match")
					} else{
							//log.Println("\n->Receiver\nThe calculated merkle root of gate tuples and merkle root of egate tuples got from the contract didn't match")
					}
		}
		wg.Done()
}




//generating a  computing party in the blockchain
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



	func keyGenerate() ([keySize]byte){

		keyGen := make([]byte, keySize)
		var key [keySize]byte

	  _, err := rand.Read(keyGen)
    if err != nil {
        log.Fatal(err)
    }

		for i:=0;i<len(keyGen);i++{
			key[i]=keyGen[i]
		}
		return key
}

func generateFinalKey(key [keySize]byte,keySizeNeeded int) ([]byte) {
	var keyHashed []byte
	var temp [32]byte
	var finalKey []byte
	finalKey = make([]byte, keySizeNeeded)

	for i:=0;i<keySizeNeeded/keySize;i++ {

		if(i==0){
			var keyPlusIndex []byte

			k:=0
			for j:=0;j<keySize;j++{
				keyPlusIndex=append(keyPlusIndex,key[j])
			}
			keyPlusIndex=append(keyPlusIndex,byte(k))
			keyHashed = kec.Keccak256(keyPlusIndex)

			for j:=0;j<hashFunctionOutputBitSize;j++{
				temp[j]=keyHashed[j]
				finalKey[j]=temp[j]
			}

		}else{

			keyHashed=nil

			for j:=0;j<hashFunctionOutputBitSize;j++{
				keyHashed=append(keyHashed,temp[j])
			}
			keyHashed=kec.Keccak256(keyHashed)

			for j:=0;j<hashFunctionOutputBitSize;j++{
				temp[j]=keyHashed[j]
				finalKey[i*32+j]=temp[j]
			}
		}
	}
	return finalKey
}

func Encrypt(key [keySize]byte,plainText []byte) ([]byte) {

			finalKey:=generateFinalKey(key,len(plainText))
			var encryptedtext []byte
			for i:=0;i<len(plainText);i++{
				encryptedtext=append(encryptedtext,finalKey[i]^plainText[i])
			}
			return encryptedtext
}

func Decrypt(key [keySize]byte,encryptedtext []byte) ([]byte){

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

func fnKeycommit(key [keySize]byte) ([hashFunctionOutputBitSize]byte){

			var toBEHashedKey []byte
			var HashedKey [hashFunctionOutputBitSize]byte

			for j:=0;j<keySize;j++{
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

	object.TotalNumberOfGates=*big.NewInt(int64(totalNumberOfGates))     //800 wei
}

func setMaxLinesToAGate(object *SenderToContractStruct){

	object.MaxInputLinesToAGate=*big.NewInt(maxLinesToGate)     //800 wei
}


func FileRead(filePath string) ([][]byte){

	var currentByte int64 = 0
	fileBuffer := make([]byte, buffer_size)
	FileChunks := make([][]byte, noOfFileChunks)
	file, err := os.Open(filePath)
	if err != nil {
			log.Fatal(err)
	}
	count:=0
	for {
			n,err := file.ReadAt(fileBuffer, currentByte)
			if err!=nil{
				log.Fatal(err)
			}
			currentByte += int64(buffer_size)
			fileBuffer=fileBuffer[:n]
			if count==noOfFileChunks{
					break
			}
			FileChunks[count] = make([]byte, buffer_size)
			copy(FileChunks[count],fileBuffer)
			if err == io.EOF {
					log.Fatal(err)
			}
			count++
	}
	file.Close()
	return FileChunks
}


func initializeCircuitTuples_asAnArray() ([][gateTupleSize]byte){


			var gateTuples [][gateTupleSize]byte
			gateTuples = make([][gateTupleSize]byte,totalNumberOfGates)

			paddingNumber:=totalNumberOfGates
			//inputGates
			for i:=0; i<noOfInputGates; i++ {
				gateTuples[i][0]=byte(i)
				gateTuples[i][1]=1
				noOfInputsToTheGate:=0
					for j:=2; j<2+noOfInputsToTheGate; j++ {

					}
					for k:=2+noOfInputsToTheGate; k<gateTupleSize; k++ {
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
					for k:=2+noOfInputsToTheGate;  k<gateTupleSize; k++ {
						gateTuples[i][k]=byte(paddingNumber)
					}
			}

			//Lastgate
			gateTuples[totalNumberOfGates-1][0]=byte(totalNumberOfGates-1)
			gateTuples[totalNumberOfGates-1][1]=3
			noOfInputsToTheGate:=1;
				for j:=2; j<2+noOfInputsToTheGate; j++ {
						gateTuples[totalNumberOfGates-1][j]=byte(totalNumberOfGates-2)
				}
				for k:=2+noOfInputsToTheGate; k<gateTupleSize; k++ {
					gateTuples[totalNumberOfGates-1][k]=byte(paddingNumber)
				}

			return gateTuples
}

func maliciousInitializeCircuitTuples_asAnArray() ([][gateTupleSize]byte){


			var gateTuples [][gateTupleSize]byte
			gateTuples = make([][gateTupleSize]byte,totalNumberOfGates)

			paddingNumber:=totalNumberOfGates
			//inputGates
			for i:=0; i<noOfInputGates; i++ {
				gateTuples[i][0]=byte(i)
				gateTuples[i][1]=1
				noOfInputsToTheGate:=0
					for j:=2; j<2+noOfInputsToTheGate; j++ {

					}
					for k:=2+noOfInputsToTheGate; k<gateTupleSize; k++ {
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
					for k:=2+noOfInputsToTheGate;  k<gateTupleSize; k++ {
						gateTuples[i][k]=byte(paddingNumber)
					}
			}

			//Lastgate
			gateTuples[totalNumberOfGates-1][0]=byte(totalNumberOfGates-1)
			gateTuples[totalNumberOfGates-1][1]=3
			noOfInputsToTheGate:=1;
				for j:=2; j<2+noOfInputsToTheGate; j++ {
						gateTuples[totalNumberOfGates-1][j]=byte(totalNumberOfGates-2)
				}
				for k:=2+noOfInputsToTheGate; k<gateTupleSize; k++ {
					gateTuples[totalNumberOfGates-1][k]=byte(paddingNumber)
				}

			gateTuples[noOfInputGates][1]=1

			return gateTuples
}


func createMerkleRootForInputVectors(inputVectors [][]byte) ([hashFunctionOutputBitSize]byte){

	var merkletree [][]byte
	merkletree = make([][]byte,2*noOfInputGates-1)

	//input gates
	for i:=0;i<noOfInputGates;i++{
				merkletree[i]=inputVectors[i]
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

func createMerkleRootForGateTuples(gateTuples [][gateTupleSize]byte) ([hashFunctionOutputBitSize]byte){

	var merkletree [][gateTupleSize]byte
	merkletree = make([][gateTupleSize]byte,2*totalNumberOfGates-1)
	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
			//merkletree[i]=make([]byte,gateTupleSize)
			merkletree[i]=gateTuples[i]
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


func Encode(inputVectors [][]byte,key [keySize]byte,gateTuples [][gateTupleSize]byte,MRx [hashFunctionOutputBitSize]byte) ([][]byte,[][]byte)  {

	out:=make([][]byte,totalNumberOfGates)
	z:=make([][]byte,totalNumberOfGates)

	for i:=0;i<noOfInputGates;i++ {
			out[i]=inputVectors[i]
			temp:=Encrypt(key,out[i])
			z[i]=temp
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
					out[i]=hash
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

			temp:=Encrypt(key,out[i])
			z[i]=temp
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

			temp:=Encrypt(key,out[i])
			z[i]=temp
		}

	return z,out
}

func maliciousSender_WrongEncryption(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup,pairIndex int){

	sender_round1_computingTime_calculatingVariables_start := time.Now()

	inputTextFilePath:=os.Args[5]
	inputVectors:=FileRead(inputTextFilePath)


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

	setPrice(&SenToConObject)
	setTotalNumberOfGates(&SenToConObject)
	setMaxLinesToAGate(&SenToConObject)
	SenToConObject.Keycommit=keyCommit
	SenToConObject.MerkleRootOfEncInp=MRencout
	SenToConObject.MerkleRootOfCircuit=MRgateTuples

	sender_round1_computingTime_calculatingVariables_stop:=time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round1_computingTime_calculatingVariables_time : %v \n", sender_round1_computingTime_calculatingVariables_stop.Sub(sender_round1_computingTime_calculatingVariables_start))

	var err error
	var tx *types.Transaction
	sender_round1_commitTime_deployingContract_start := time.Now()
  contractAddress[pairIndex],tx,_, err=import_contracts.DeployFairswap(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.Keycommit,
		SenToConObject.MerkleRootOfEncInp,
		SenToConObject.MerkleRootOfCircuit,
		&SenToConObject.TotalNumberOfGates,
		&SenToConObject.MaxInputLinesToAGate,
	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()

	instance,err:= import_contracts.NewFairswap(contractAddress[pairIndex],client)
	if err != nil {
			log.Fatal(err)
	}

	sender_round1_commitTime_deployingContract_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round1_commitTime_deployingContract_time : %v \n", sender_round1_commitTime_deployingContract_stop.Sub(sender_round1_commitTime_deployingContract_start))
	log.SetPrefix("\nFairswap - Gas Cost ")
	log.Println("Sender Round 1 : ",tx.Gas())



	sender_round1_communicationTime_sendingEncryptedGateOutputs_start := time.Now()


	SenToRecObject.EncryptedOutputOfGates=encodedGateOutputs

	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		//log.Println("error:", err)
	}

	channel_SenToRec<-byteArrayFromSenToRec

	sender_round1_communicationTime_sendingEncryptedGateOutputs_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round1_communicationTime_sendingEncryptedGateOutputs_time : %v \n", sender_round1_communicationTime_sendingEncryptedGateOutputs_stop.Sub(sender_round1_communicationTime_sendingEncryptedGateOutputs_start))


	sender_round2_waitingTime_forReceiverToAccept_start := time.Now()



	receiverAcceptedStatus:=false
	for i:=0;i<waitingTime;i++ {

		Phase,err:=instance.Phase(nil)
		if err!=nil{
			log.Fatal(err)
		}
		if(Phase==2){
			receiverAcceptedStatus=true
			break;
		}
		time.Sleep(time.Second *1)

	}


	sender_round2_waitingTime_forReceiverToAccept_stop := time.Now()
	log.SetPrefix("\nFairswap - Time ")
	log.Printf("sender_round2_waitingTime_forReceiverToAccept_time : %v \n", sender_round2_waitingTime_forReceiverToAccept_stop.Sub(sender_round2_waitingTime_forReceiverToAccept_start))


	if(!receiverAcceptedStatus){
			//Receiver has not accepted within time
			}else{

						 sender_round3_commitTime_revealKey_start := time.Now()
						 tx,err=instance.RevealKey(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},key)
							if err!=nil{
								log.Fatal(err)
							}
							client.Commit()
							//log.Println("\n->Sender\nKeyRevealed")
							sender_round3_commitTime_revealKey_stop := time.Now()
							log.SetPrefix("\nFairswap - Time ")
							log.Printf("sender_round3_commitTime_revealKey_time : %v \n", sender_round3_commitTime_revealKey_stop.Sub(sender_round3_commitTime_revealKey_start))
							//client.Commit()
							log.SetPrefix("\nFairswap - Gas Cost ")
							log.Println("Sender Round 3 : ",tx.Gas())


			sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_start:=time.Now()

			for i:=0;i<waitingTime;i++{

					Phase,err:=instance.Phase(nil)
					if err!=nil{
						log.Fatal(err)
					}
					if(Phase==4){
						break
					}
					time.Sleep(time.Second *1)
			}
			sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_stop:=time.Now()
			log.SetPrefix("\nFairswap - Time ")
			log.Printf("sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_time : %v \n", sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_stop.Sub(sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_start))

			//gasLimit := uint64(10000000000)
			//the sender getting the ether
			Phase,err:=instance.Phase(nil)
			if err!=nil{
				log.Fatal(err)
			}

			if(Phase==3){

					timeout,err:=instance.Timeout(nil)
					if err!=nil{
									log.Fatal(err)
					}


					for{
						_,err:=instance.Nowy(&bind.TransactOpts{
							 From:authSender.From,
							 Signer:authSender.Signer,
						 })
						 if err!=nil{
									 log.Fatal(err)
						 }
						 client.Commit()

						 Now,err:=instance.Now(nil)
						 if err!=nil{
									 log.Fatal(err)
						 }

						 if(Now.Cmp(timeout)==1){
								break
						 }

						 time.Sleep(time.Second*1)
					}

					sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start:=time.Now()
					tx,err=instance.SenderGetEtherReceiverNoComplainMessage(&bind.TransactOpts{
							From:authSender.From,
							Signer:authSender.Signer,
						})
					if err!=nil{
								log.Fatal(err)
					}
					client.Commit()

					log.SetPrefix("\nFairswap - Gas Cost ")
					log.Println("Sender Round 4 ",tx.Gas())
					sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop:=time.Now()
					log.SetPrefix("\nFairswap - Time ")
					log.Printf("sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_time : %v \n", sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop.Sub(sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start))

			}
	}
	wg.Done()
}


func createMerkleRootForEncInp(encryptedGateOutputs [][]byte) ([hashFunctionOutputBitSize]byte){

	// var merkletree [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte
	merkletree:=make([][hashFunctionOutputBitSize]byte,2*totalNumberOfGates-1)

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
				merkletree[i]=Hash(encryptedGateOutputs[i])
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


func checkMerkleRootForEncInp(encryptedGateOutputs [][]byte,merklerootforEncInput [hashFunctionOutputBitSize]byte) (bool){

	// var merkletree [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte
	merkletree:=make([][hashFunctionOutputBitSize]byte,2*totalNumberOfGates-1)

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
				merkletree[i]=Hash(encryptedGateOutputs[i])
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

func checkMerkleRootForGateTuples(gateTuples [][gateTupleSize]byte,merkleRootOfGateTuples [hashFunctionOutputBitSize]byte) (bool){

	var merkletree [][gateTupleSize]byte
	merkletree = make([][gateTupleSize]byte,2*totalNumberOfGates-1)
	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
			//merkletree[i]=make([]byte,gateTupleSize)
			merkletree[i]=gateTuples[i]
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

func createMerkleTreeForGateTuples(gateTuples [][gateTupleSize]byte) ([][hashFunctionOutputBitSize]byte){

	merkletree :=make([][hashFunctionOutputBitSize]byte,2*totalNumberOfGates-1)

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
				merkletree[i]=gateTuples[i]
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

	return merkletree
}

func createMerkleTreeForEncInp(encryptedGateOutputs [][]byte) ([][hashFunctionOutputBitSize]byte){

	merkletree:=make([][hashFunctionOutputBitSize]byte,2*totalNumberOfGates-1)

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totalNumberOfGates;i++{
				merkletree[i]=Hash(encryptedGateOutputs[i])
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

func Extract(encodedGateOutputs [][]byte,key [32]byte,gateTuples [][gateTupleSize]byte,merkleTreeForEncGateOutput [][hashFunctionOutputBitSize]byte,merkleTreeForGateTuples [][hashFunctionOutputBitSize]byte,Receiver_MerkleRootOfEncInp [hashFunctionOutputBitSize]byte,Receiver_MerkleRootOfGateTuples [hashFunctionOutputBitSize]byte) ([][][hashFunctionOutputBitSize]byte,[][]byte,[][hashFunctionOutputBitSize]byte,[][]byte) {

	// var decodedOutputs [totalNumberOfGates][]byte
	decodedOutputs:=make([][]byte,totalNumberOfGates)
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

			if(int(gateTuples[i][k])==totalNumberOfGates){
				break
			}
			operationInputs=append(operationInputs,decodedOutputs[gateTuples[i][k]])
		}
				 if(!reflect.DeepEqual(decodedOutputs[i],Operation(operation,operationInputs,decodedOutputs,Receiver_MerkleRootOfEncInp)))	 {
				// OperationVal:=Operation(operation,operationInputs,decodedOutputs,Receiver_MerkleRootOfEncInp)
				// fmt.Println("OperationVal : ",OperationVal)
						encodedVectors_ThisGate=append(encodedVectors_ThisGate,encodedGateOutputs[i])
						gateComplain=Mproof(i,merkleTreeForGateTuples)
						//merkle proof for the output of a complaint gate
						MproofTree:=Mproof(i,merkleTreeForEncGateOutput)
						complain=append(complain,MproofTree)

						//merkle proof for the inputs of a complaint gate and appending to the complain vector
						for k:=2;k<2+maxLinesToGate;k++{

							if int(gateTuples[i][k])==totalNumberOfGates {
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

func Mproof(index int,merkletree [][hashFunctionOutputBitSize]byte) ([][hashFunctionOutputBitSize]byte){

	depth:=int(mathlib.Log2(float64(totalNumberOfGates)))+1
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


func Operation(Op int,operationInputs [][]byte,decodedOutputs [][]byte,Receiver_MerkleRootOfEncInp [hashFunctionOutputBitSize]byte)  ([]byte){

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
			merkleTreeOfDecryptedInputVectors:=make([][]byte,noOfInputGates)
			for k:=0;k<noOfInputGates;k++{
						merkleTreeOfDecryptedInputVectors[k]=decodedOutputs[k]
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
