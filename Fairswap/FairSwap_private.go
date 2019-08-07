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
	"sync"
	//"github.com/ethereum/go-ethereum/core"
	mathlib "math"
	"crypto/rand"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	//"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strconv"
	import_contracts "./contracts"

)

var bufferSize int64
var noOfInputGates int64
var totalNumberOfGates int64


const maxLinesToGate=2
const keyLength=32
const hashFunctionOutputBitSize=32
const gateTupleSize=32
const waitingTime=20

var contractAddress common.Address

type SenderToReceiverStruct struct {
		EncryptedOutputOfGates [totalNumberOfGates][]byte
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

		channel_SenToRec:=make(chan []byte)

		//logging file
		f, err := os.OpenFile("text1.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
		    log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		wrt := io.MultiWriter(os.Stdout, f)
		log.SetOutput(wrt)

		if(len(os.Args)<2){
			bufferSize=32
			noOfInputGates=8
			totalNumberOfGates=2*noOfInputGates
		}else{
			var err error
			bufferSize, err = strconv.ParseInt(os.Args[1], 10, 64)
			if(err!=nil){
				log.Fatal(err)
			}
			noOfInputGates, err = strconv.ParseInt(os.Args[2], 10, 64)
			if(err!=nil){
				log.Fatal(err)
			}
			totalNumberOfGates=2*noOfInputGates
		}
		fmt.Println(bufferSize,"\t",noOfInputGates,"\t",totalNumberOfGates)

		keystoreNode1 := keystore.NewKeyStore("/home/amudhan/workspace1/node1/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
		allAccountsFromNode1:=keystoreNode1.Accounts()
		fmt.Println("accountsFromKeyStore : ",allAccountsFromNode1)

		clientNode1,err := ethclient.Dial("/home/amudhan/workspace1/node1/geth.ipc")
		if err != nil {
					log.Fatal(err)
		}
		fmt.Println(reflect.TypeOf(clientNode1))

		accountAddressNode1:=common.HexToAddress("bfbd4c070743dfa221f2f5d704dd8bd0f045e730")
		accountNode1orSenderInitialBalance, err := clientNode1.BalanceAt(context.Background(),accountAddressNode1, nil)
		if err != nil{
				log.Fatal(err)
		}
		fmt.Println("-->Main\naccountNode1orSenderInitialBalance : ",accountNode1orSenderInitialBalance,"\n Account address : ",accountAddressNode1)

		var node1Reader io.Reader
		node1Reader, err = os.Open("/home/amudhan/workspace1/node1/keystore/UTC--2019-07-25T15-02-05.471118987Z--bfbd4c070743dfa221f2f5d704dd8bd0f045e730")
		fmt.Println(reflect.TypeOf(node1Reader))
		if err != nil{
				log.Fatal(err)
		}

		authSender,err	:= bind.NewTransactor(node1Reader,"qwerty")
		if err != nil{
				log.Fatal(err)
		}
		fmt.Println(reflect.TypeOf(authSender),"auth address : ", authSender.From)


		fmt.Println("\n\n\n\n")

		//2nd node
		keystoreNode2 := keystore.NewKeyStore("/home/amudhan/workspace1/node2/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
		allAccountsFromNode2 :=keystoreNode2.Accounts()
		fmt.Println("allAccountsFromNode2 : ",allAccountsFromNode2)

		//var ctx context.Context
		clientNode2,err := ethclient.Dial("/home/amudhan/workspace1/node2/geth.ipc")
		if err != nil {
					log.Fatal(err)
		}
		fmt.Println(reflect.TypeOf(clientNode2))

		accountAddressNode2:=common.HexToAddress("c03458ea0965d2b2cb7caf2ef66a98a9fb194140")
		accountNode2orReceiverBalance, err := clientNode2.BalanceAt(context.Background(),accountAddressNode2, nil)
		if err != nil{
				log.Fatal(err)
		}
		fmt.Println("\naccountNode2BalanceInitial : ",accountNode2orReceiverBalance,"\n AccountNode2 address : ",accountAddressNode2)

		var node2Reader io.Reader
		node2Reader, err = os.Open("/home/amudhan/workspace1/node2/keystore/UTC--2019-07-25T15-03-03.735001680Z--c03458ea0965d2b2cb7caf2ef66a98a9fb194140")
		fmt.Println(reflect.TypeOf(node2Reader))
		if err != nil{
				log.Fatal(err)
		}

		authReceiver,err	:= bind.NewTransactor(node2Reader,"qwerty")
		if err != nil{
				log.Fatal(err)
		}
		fmt.Println(reflect.TypeOf(authReceiver),"authNode2 address : ", authReceiver.From)

		//go routines
		ctx := context.Background()
		var wg sync.WaitGroup
		wg.Add(2)
		//go routines
		go Sender(ctx,authSender,clientNode1,channel_SenToRec,&wg)
		go Receiver(ctx,authReceiver,clientNode2,channel_SenToRec,&wg)


		wg.Wait()

		//getting the parties account balance at the end of protocol
		senderFinalBalance, err := clientNode1.BalanceAt(context.Background(),accountAddressNode1, nil)
		if err != nil{
				log.Fatal(err)
		}

		senderDifferenceInFinalAndInitialBalance:=new(big.Int).Sub(accountNode1orSenderInitialBalance,senderFinalBalance)
		fmt.Println("\n-->Main\nSenderFinalAccountBalance : ",senderFinalBalance,"\nSenderDifferenceInFinalAndInitialBalance : ",senderDifferenceInFinalAndInitialBalance)

		receiverFinalBalance, err := clientNode2.BalanceAt(context.Background(),accountAddressNode2, nil)
		if err != nil{
				log.Fatal(err)
		}
		receiverDifferenceInFinalAndInitialBalance:=new(big.Int).Sub(accountNode2orReceiverBalance,receiverFinalBalance)
		fmt.Println("\n-->Main\nReceiverFinalAccountBalance : ",receiverFinalBalance,"\nReceiverDifferenceInFinalAndInitialBalance : ",receiverDifferenceInFinalAndInitialBalance ,"\n")
}


func Sender(ctx context.Context,authSender *bind.TransactOpts,client *ethclient.Client,channel_SenToRec chan []byte,wg *sync.WaitGroup){

	sender_round1_computingTime_calculatingVariables_start := time.Now()

	inputVectors:=FileRead()

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
	log.SetPrefix("\nTime ")
	log.Printf("sender_round1_computingTime_calculatingVariables_time : %v \n", sender_round1_computingTime_calculatingVariables_stop.Sub(sender_round1_computingTime_calculatingVariables_start))

	var err error
	var tx *types.Transaction
	sender_round1_commitTime_deployingContract_start := time.Now()
  _,tx,_, err=import_contracts.DeployFairswap(
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

	contractAddress,err=bind.WaitDeployed(ctx,client,tx)
	if err!=nil{
		log.Fatal(err)
	}

	sender_round1_commitTime_deployingContract_stop := time.Now()
	log.SetPrefix("\nTime ")
	log.Printf("sender_round1_commitTime_deployingContract_time : %v \n", sender_round1_commitTime_deployingContract_stop.Sub(sender_round1_commitTime_deployingContract_start))
	log.SetPrefix("\nGas Cost ")
	log.Println("Sender Round 1 : ",tx.Gas())

	instance,err:= import_contracts.NewFairswap(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	/*header, err := client.HeaderByNumber(context.Background(), nil)
 	if err != nil {
		 log.Fatal(err)
 	}

	block, err := client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		 	log.Fatal(err)
	}

	fmt.Println(block.Time().Uint64())*/

	tx,err=instance.Nowy(&bind.TransactOpts{
		From:authSender.From,
		Signer:authSender.Signer,
		Value: nil,
	})
	if err!=nil{
		log.Fatal(err)
	}

	_,err=bind.WaitMined(ctx,client,tx)
	if err!=nil{
		log.Fatal(err)
	}

	now,err:=instance.Now(nil)
	if err!=nil{
				log.Fatal(err)
	}
	fmt.Println("\n  before sender get ether  now : ",now)


	sender_round1_communicationTime_sendingEncryptedGateOutputs_start := time.Now()


	SenToRecObject.EncryptedOutputOfGates=encodedGateOutputs

	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		//log.Println("error:", err)
	}

	channel_SenToRec<-byteArrayFromSenToRec

	sender_round1_communicationTime_sendingEncryptedGateOutputs_stop := time.Now()
	log.SetPrefix("\nTime ")
	log.Printf("sender_round1_communicationTime_sendingEncryptedGateOutputs_time : %v \n", sender_round1_communicationTime_sendingEncryptedGateOutputs_stop.Sub(sender_round1_communicationTime_sendingEncryptedGateOutputs_start))


	sender_round2_waitingTime_forReceiverToAccept_start := time.Now()

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
		time.Sleep(time.Second *1)

	}

	sender_round2_waitingTime_forReceiverToAccept_stop := time.Now()
	log.SetPrefix("\nTime ")
	log.Printf("sender_round2_waitingTime_forReceiverToAccept_time : %v \n", sender_round2_waitingTime_forReceiverToAccept_stop.Sub(sender_round2_waitingTime_forReceiverToAccept_start))


	if(!RevealedBool){
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

							_,err=bind.WaitMined(ctx,client,tx)
							if err!=nil{
								log.Fatal(err)
							}

							sender_round3_commitTime_revealKey_stop := time.Now()
							log.SetPrefix("\nTime ")
							log.Printf("sender_round3_commitTime_revealKey_time : %v \n", sender_round3_commitTime_revealKey_stop.Sub(sender_round3_commitTime_revealKey_start))
							log.SetPrefix("\nGas Cost ")
							log.Println("Sender Round 3 : ",tx.Gas())*/


			//The contract is in keyRevealed stage
			//Once the sender has revealed the key, the receiver has time until waitingTime to either register a complaint or
			//walk away(not do any execution or transaction--got the correct file)
			//case 1 : If receiver gives a wrong complaint then the contract stage goes to finishedAndMaliciousReciever. The sender sends
			//				"Receiver Wrong complaint" to the main function and finishes the protocol.
			//case 2 : If the receiver doesn't respond, then the stage contract stage is still in keyRevealed and the sender gets the ether
			//				 after the waitingTime

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
			log.SetPrefix("\nTime ")
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
			fmt.Println("\n before sender get ether timeout : ",timeout)

			for{
				header, err := client.HeaderByNumber(context.Background(), nil)
			 	if err != nil {
					 log.Fatal(err)
			 	}

				block, err := client.BlockByNumber(context.Background(), header.Number)
				if err != nil {
					 	log.Fatal(err)
				}

				if(timeout.Cmp(big.NewInt(int64(block.Time().Uint64())))==-1){
					break
				}
				time.Sleep(time.Second*1)
			}

			sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start:=time.Now()

			tx,err=instance.SenderGetEtherReceiverNoComplainMessage(&bind.TransactOpts{
					From:authSender.From,
					Signer:authSender.Signer,
					// Value: big.NewInt(10000000000),
					//GasLimit:gasLimit,
				})
			if err!=nil{
						log.Fatal(err)
			}
			_,err=bind.WaitMined(ctx,client,tx)
			if err!=nil{
				log.Fatal(err)
			}

			log.SetPrefix("\nGas Cost ")
			log.Println("Sender Round 4 ",tx.Gas())
			sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop:=time.Now()
			log.SetPrefix("\nTime ")
			log.Printf("sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_time : %v \n", sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop.Sub(sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start))

			//log.Println("\n->Sender\nProtocol completed Successfully")
			//the protocol gets finished and the sender sends the message to the main function
			}
	}
	wg.Done()
}

//channel_RecToMainRev(chan string) - to notify the main function for the completion of reveal phase
func Receiver(ctx context.Context,authReceiver *bind.TransactOpts,client *ethclient.Client,channel_SenToRec chan []byte,wg *sync.WaitGroup){

	receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start := time.Now()
	//log.Println("\n->Receiver\n Begins")
	byteArrayFromSenToRec:=<-channel_SenToRec

	receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop := time.Now()
	log.SetPrefix("\nTime ")
	log.Printf("receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_time : %v \n", receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop.Sub(receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start))


	//log.Println("\n->Receiver\nByteArray got from sender")

	receiver_round1_computingTime_merkleRootCheck_start := time.Now()

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

	receiver_round1_computingTime_merkleRootCheck_stop := time.Now()
	log.SetPrefix("\nTime ")
	log.Printf("receiver_round1_computingTime_merkleRootCheck_time : %v \n", receiver_round1_computingTime_merkleRootCheck_stop.Sub(receiver_round1_computingTime_merkleRootCheck_start))

	//log.Println("\n->Receiver\nboolValueEncryptedGateOutput(comparision between merkleroot of encryptedGateOutputs from contract and calulated encryptedGateOutputs merkleroot): ",boolValueEncryptedGateOutput)
	//log.Println("\n->Receiver\nboolValueEncryptedGateTuples(comparision between merkleroot of encryptedGateTuples from contract and calulated encryptedGateTuples merkleroot): ",boolValueEncryptedGateTuples)

	if(boolValueEncryptedGateOutput && boolValueEncryptedGateTuples){

				//the boolValueEncode was true, so the receiver calls the accept contract function and pays the required price


			value:=big.NewInt(80)
			//gasLimit := uint64(4576000000000)

		 receiver_round2_commitTime_Accept_start := time.Now()

			tx,err:=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
			//	GasLimit:gasLimit,
			})
			if err != nil {
					log.Fatal(err)
			}
			_,err=bind.WaitMined(ctx,client,tx)
			if err!=nil{
				log.Fatal(err)
			}
			receiver_round2_commitTime_Accept_stop := time.Now()
			log.SetPrefix("\nTime ")
			log.Printf("receiver_round2_commitTime_Accept_time : %v \n", receiver_round2_commitTime_Accept_stop.Sub(receiver_round2_commitTime_Accept_start))

			log.SetPrefix("\nGas Cost ")
			log.Println("Receiver Round 2 : ",tx.Gas())

			//
			Phase,err:=instance.Phase(nil)
			if err!=nil{
				log.Fatal(err)
			}
			fmt.Println("\n Phase : ",Phase)




			//log.Println("\n->Receiver\n Accepted")

			//The contract is in accepted stage
			//The sender
			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
			//case 2 : The sender doesn't reveal the key within time. AcceptedBool remains false after the for loop. Receiver later calls the contract function
			//			   ReceiverGetEther to get the ether.
			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and AcceptedBool is true after the for loop.


		receiver_round3_waitingTime_forSenderToRevealKey_start := time.Now()
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
					AcceptedBool=true
					break;
				}
				if(Phase==4){
					AcceptedBool=true
					//wrong key Revealed
					break;
				}
				time.Sleep(time.Second *1)

			}
			receiver_round3_waitingTime_forSenderToRevealKey_stop := time.Now()
			log.SetPrefix("\nTime ")
			log.Printf("receiver_round3_waitingTime_forSenderToRevealKey_time : %v \n", receiver_round3_waitingTime_forSenderToRevealKey_stop.Sub(receiver_round3_waitingTime_forSenderToRevealKey_start))


			if(!AcceptedBool){

					timeout,err:=instance.Timeout(nil)
					if err!=nil{
								log.Fatal(err)
					}
					fmt.Println("\n before sender get ether timeout : ",timeout)

					//log.Println("\n->Receiver\nSender has not revealed the key...ReceiverGetEther function called")
					for{
						header, err := client.HeaderByNumber(context.Background(), nil)
					 	if err != nil {
							 log.Fatal(err)
					 	}

						block, err := client.BlockByNumber(context.Background(), header.Number)
						if err != nil {
							 	log.Fatal(err)
						}

						if(timeout.Cmp(big.NewInt(int64(block.Time().Uint64())))==-1){
							fmt.Println(big.NewInt(int64(block.Time().Uint64())))
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
							//GasLimit:gasLimit,
						})
					if err!=nil{
								log.Fatal(err)
					}
					_,err=bind.WaitMined(ctx,client,tx)
					if err!=nil{
						log.Fatal(err)
					}
					receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop := time.Now()
					log.SetPrefix("\nGas Cost ")
					log.Println("Receiver Round 3 (Sender Not Revealed Key) : ",tx.Cost())
					log.Printf("\nreceiver_round3_commitTime_senderKeyNotRevealedGettingEther_time : %v \n", receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop.Sub(receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start))
					//client.Commit()

			} else{

				//log.Println("\n->Receiver\nSender has revealed the key....Proceeding further")

				receiver_round4_computingTime_extraction_start := time.Now()

				merkleTreeForGateTuples:=createMerkleTreeForGateTuples(gateTuples)
			 	merkleTreeForEncInp:=createMerkleTreeForEncInp(Receiver_SenToRecObj.EncryptedOutputOfGates)

			 	complain,_,gateComplain,encodedVectors_ThisGate :=Extract(Receiver_SenToRecObj.EncryptedOutputOfGates,contract_Key,gateTuples,merkleTreeForEncInp,merkleTreeForGateTuples,Receiver_MerkleRootOfEncInp,Receiver_MerkleRootOfGateTuples)

				receiver_round4_computingTime_extraction_stop := time.Now()
				log.SetPrefix("\nTime ")
				log.Printf("receiver_round4_computingTime_extraction_time : %v \n", receiver_round4_computingTime_extraction_stop.Sub(receiver_round4_computingTime_extraction_start))
				//log.Println("\n->Receiver\nDecoded Ouptuts : ",decodedOutputs)


				if(complain==nil){
						log.SetPrefix("\nGas Cost ")
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

						_,err=bind.WaitMined(ctx,client,tx)
						if err!=nil{
							log.Fatal(err)
						}

						receiver_round4_commitTime_complain_stop := time.Now()
						log.SetPrefix("\nTime ")
						log.Printf("\nreceiver_round4_commitTime_complain_time : %v \n", receiver_round4_commitTime_complain_stop.Sub(receiver_round4_commitTime_complain_start))

						log.SetPrefix("\nGas Cost ")
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

func FileRead() ([noOfInputGates][buffer_size]byte){
	var currentByte int64 = 0
	fileBuffer := make([]byte, buffer_size)
	file, err := os.Open("a.txt")
	if err != nil {
			log.Fatal(err)
	}
	var FileChunks [noOfInputGates][buffer_size]byte
	count:=0
	for {
			n,err := file.ReadAt(fileBuffer, currentByte)
			currentByte += buffer_size
			fileBuffer=fileBuffer[:n]
			if count==noOfInputGates{
					break
			}
			for j:=0;j<n;j++{
				FileChunks[count][j]=fileBuffer[j]
			}
			if err == io.EOF {
					log.Fatal(err)
			}
			count++
	}
	file.Close()
	return FileChunks
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

/*
func maliciousSender_WrongEncryption(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup){

	//log.Println("\n->Sender\nBegins")
	inputVectors:=FileRead()

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

	var err error
	var tx *types.Transaction
  contractAddress,tx,_, err=import_contracts.DeployFairswap(
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
	log.SetPrefix("Sender Round1 ")
	log.Println(tx.Gas())

	SenToRecObject.EncryptedOutputOfGates=encodedGateOutputs
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		//log.Println("error:", err)
	}

	channel_SenToRec<-byteArrayFromSenToRec
	//log.Println("\n->Sender \nByteArray sent to Receiver")

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
			//log.Println("\n->Sender\nReceiver has not paid enough ether")
			break;
		}
		time.Sleep(time.Second *1)

	}

	Phase,err:=instance.Phase(nil)
	if err!=nil{
			log.Fatal(err)
	}
	if(Phase==1){
		//log.Println("\n->Sender\nReceiver hasnt accepted within time")
	}else{

			if(!RevealedBool){
			}else{

						  tx,err=instance.RevealKey(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},key)
							if err!=nil{
								log.Fatal(err)
							}
							//log.Println("\n->Sender\nKeyRevealed")
							client.Commit()
							log.SetPrefix("Sender Round3 ")
							log.Println(tx.Gas())
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
						//log.Println("\n->Sender\nReceiver has made a wrong complaint")
						break;
					}
					if(Phase==5){
						//log.Println("\n->Sender\nIt seems I have made a mistake(Wrong Encryption)")
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

			//log.Println("\n->Sender\nProtocol completed Successfully")
			//the protocol gets finished and the sender sends the message to the main function
			}
	}
	wg.Done()
}

/*
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
		//log.Println("error:", err)
	}

	channel_SenToRec<-byteArrayFromSenToRec
	//log.Println("\n->Sender \nByteArray sent to Receiver")
	//log.Println("\n->Sender\n Initialization Finished")

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
			//log.Println("\n->Sender\nReceiver has not paid enough ether")
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
		//log.Println("\n->Sender\nReceiver hasnt accepted within time")
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
							//log.Println("\n->Sender :\n I have revealed the wrong key")
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

	//log.Println("\n->Receiver\nboolValueEncryptedGateOutput(comparision between merkleroot of encryptedGateOutputs from contract and calulated encryptedGateOutputs merkleroot): ",boolValueEncryptedGateOutput)
	//log.Println("\n->Receiver\nboolValueEncryptedGateTuples(comparision between merkleroot of encryptedGateTuples from contract and calulated encryptedGateTuples merkleroot): ",boolValueEncryptedGateTuples)

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

			//log.Println("\n->Receiver\n Finished Accepting")
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
					//log.Println("\n->Receiver\ncontract_Key",contract_Key)
					AcceptedBool=true
					break;
				}

				if(Phase==5){
					//log.Println("\n->Receiver\nSender Has revealed the wrong key")
					channel_RecToMainRev<-"Wrong key revealed"

				}
				time.Sleep(time.Second *1)

			}


			if(!AcceptedBool){

					//log.Println("\n->Receiver\nSender has not revealed the key...ReceiverGetEther function called")
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

				//log.Println("\n->Receiver\nSender has revealed the key....Proceeding further")

				merkleTreeForGateTuples:=createMerkleTreeForGateTuples(gateTuples)
			 	merkleTreeForEncInp:=createMerkleTreeForEncInp(Receiver_SenToRecObj.EncryptedOutputOfGates)

				maliciousGateTuples:=maliciousInitializeCircuitTuples_asAnArray()

			 	complain,decodedOutputs,gateComplain,encodedVectors_ThisGate :=Extract(Receiver_SenToRecObj.EncryptedOutputOfGates,contract_Key,maliciousGateTuples,merkleTreeForEncInp,merkleTreeForGateTuples,Receiver_MerkleRootOfEncInp,Receiver_MerkleRootOfGateTuples)

				//log.Println("\n->Receiver\nDecoded Ouptuts : ",decodedOutputs)

				if(complain==nil){
					//log.Println("\n->Receiver\nNo complain")
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
							//log.Println("\n->Receiver\nSender Has wrongly calulated the encrypted inputs")
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

*/
