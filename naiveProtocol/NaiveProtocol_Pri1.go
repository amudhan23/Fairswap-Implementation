package main

import(
	"fmt"
	kec "github.com/ethereum/go-ethereum/crypto"
	"os"
	"io"
	"log"
	"encoding/json"
	"sync"
	"path/filepath"
	// "reflect"
	"context"
	"strconv"
	"time"
  "crypto/ecdsa"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"crypto/rand"
  "github.com/ethereum/go-ethereum/common"
	"math/big"
   import_contracts "./contracts"
)

//the following parameters are got as command line arguments
var buffer_size int
var noOfFileChunks int
var keyLength int

const waitingTime=20  //max time to complete a stage
const hashFunctionOutputBitSize=32 //hash function's output size(32 bytes in this case)

var contractAddress []common.Address

type SenderToReceiverStruct struct {
		EncryptedFileChunks [][]byte
}

type SenderToContractStruct struct{
	Price big.Int
	HashedFileChunks [][hashFunctionOutputBitSize]byte
	HashedEncryptedFileChunks [][hashFunctionOutputBitSize]byte
}

func main(){

		//channels
					//sender sends the byterarray(channel for sending EncryptedFileChunks from sender to receover)

		//log file as well as std out

		//getting command line arguments for initialization
		var noOfPairsOfBuyersAndSellers int
		var keyStoreDirectoryforBuyers string
		var keyStoreDirectoryforSellers string
		var gethNode1ipc string
		var gethNode2ipc string
		var logFilePath string

	 	if(len(os.Args)<8){
			os.Exit(2)
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
			keyStoreDirectoryforBuyers=os.Args[4]
			keyStoreDirectoryforSellers=os.Args[5]
			gethNode1ipc=os.Args[6]
			gethNode2ipc=os.Args[7]
			logFilePath=os.Args[8]
			keyLength=buffer_size
		}

		f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
				log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		wrt := io.MultiWriter(os.Stdout, f)
		log.SetOutput(wrt)

		fmt.Println("gethNode1ipc : ",gethNode1ipc,"gethNode2ipc : ",gethNode2ipc)


		clientNode1,err := ethclient.Dial(string(gethNode1ipc))
		if err != nil {
					log.Fatal(err)
		}
		clientNode2,err := ethclient.Dial(string(gethNode2ipc))
		if err != nil {
					log.Fatal(err)
		}


		contractAddress=make([]common.Address,noOfPairsOfBuyersAndSellers)

		authSenders:=make([]*bind.TransactOpts,noOfPairsOfBuyersAndSellers)
		authReceivers:=make([]*bind.TransactOpts,noOfPairsOfBuyersAndSellers)

		var files []string

    rootNode1 :=keyStoreDirectoryforBuyers
		err = filepath.Walk(rootNode1, func(path string, info os.FileInfo, err error) error {
				files = append(files, path)
				fmt.Println("path : ",path)
				return nil
		})
		if err != nil {
				panic(err)
		}

		fmt.Println("\n\n buyer len : ",len(files),"\nfiles : ",files)

		fmt.Println("\n\n\nkeyStoreDirectoryforSellers : ",keyStoreDirectoryforSellers)
		fmt.Println("\n\n\nkeyStoreDirectoryforBuyers : ",keyStoreDirectoryforBuyers)

		var files2 []string
		rootNode2:= keyStoreDirectoryforSellers
    err = filepath.Walk(rootNode2, func(path string, info os.FileInfo, err error) error {
			files2 = append(files2, path)
				fmt.Println("path : ",path)
        return nil
    })
    if err != nil {
        panic(err)
    }

		if(len(files)-1<noOfPairsOfBuyersAndSellers || len(files2)-1<noOfPairsOfBuyersAndSellers){
			log.Println("Not enough accounts")
			os.Exit(3)
		}


		var node1Reader io.Reader

		for i:=0;i<noOfPairsOfBuyersAndSellers;i++{

			node1Reader, err = os.Open(files[i+1])
			if err != nil{
					log.Fatal(err)
			}

			authSenders[i],err	= bind.NewTransactor(node1Reader,"qwerty")
			if err != nil{
					log.Fatal(err)
			}

		}

		var node2Reader io.Reader

		for i:=0;i<noOfPairsOfBuyersAndSellers;i++{

			node2Reader, err = os.Open(files2[i+1])
			if err != nil{
					log.Fatal(err)
			}

			authReceivers[i],err	= bind.NewTransactor(node2Reader,"qwerty")
			if err != nil{
					log.Fatal(err)
			}

		}

		// fmt.Println()

		//go routines

		var wg sync.WaitGroup
		wg.Add(2*noOfPairsOfBuyersAndSellers)
		for i:=0;i<noOfPairsOfBuyersAndSellers;i++{
			pairIndex:=i
			ctx := context.Background()
			channel_SenToRec:=make(chan []byte)
			go Sender(ctx,authSenders[pairIndex],clientNode1,channel_SenToRec,&wg,pairIndex)
			go Receiver(ctx,authReceivers[pairIndex],clientNode2,channel_SenToRec,&wg,pairIndex)
		}

		wg.Wait()  //Wait blocks until the WaitGroup counter is zero

}



//Sender function
//The functoin executes honest sender's role in fairswap
func Sender(ctx context.Context,authSender *bind.TransactOpts,client *ethclient.Client,channel_SenToRec chan []byte,wg *sync.WaitGroup,pairIndex int){

	/*Computing all the variables in sender to receiever structure and sender to contract structure*/
	/**//*computing variables time starts*/
	sender_round1_computingTime_calculatingVariables_start := time.Now()
		//reading from a file "a.txt" and storing bytearray of 32 bytes in FileChunks of size equal to noOfFileChunks
	FileChunks:=FileRead()

	//objects
	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct

	setPrice(&SenToConObject)
	key:=keyGenerate()
	HashedFileChunks:=HashChunks(FileChunks)
	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedFileChunks=append(SenToConObject.HashedFileChunks,HashedFileChunks[i])
	}

 	encryptedFileChunks := EncryptFileChunks(FileChunks,key)
	HashedEncryptedFileChunks:=HashChunks(encryptedFileChunks)

	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedEncryptedFileChunks=append(SenToConObject.HashedEncryptedFileChunks,HashedEncryptedFileChunks[i])
	}

	sender_round1_computingTime_calculatingVariables_stop:=time.Now()
	/**//*computing variables time ends*/
	log.SetPrefix("\nNaive Protocol - Time ")
	// log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - [buffer_size] %d",buffer_size,"[noOfFileChunks] %d",noOfFileChunks)
	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round1_computingTime_calculatingVariables_time : ", int64(sender_round1_computingTime_calculatingVariables_stop.Sub(sender_round1_computingTime_calculatingVariables_start)/time.Nanosecond))





	/*Deploy contract*/
	var err error
	var tx *types.Transaction
	/**//*commit time for contract deployment starts*/
	sender_round1_commitTime_deployingContract_start := time.Now()
	_,tx,_, err=import_contracts.DeployProtocol2(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.HashedFileChunks,
		SenToConObject.HashedEncryptedFileChunks,
	)
	if err!=nil{
		log.Fatal(err)
	}

	contractAddress[pairIndex],err=bind.WaitDeployed(ctx,client,tx)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("\n\nContract Address : ",contractAddress)
	sender_round1_commitTime_deployingContract_stop := time.Now()
	/**//*commit time for contract deployment ends*/
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round1_commitTime_deployingContract_time : ", int64(sender_round1_commitTime_deployingContract_stop.Sub(sender_round1_commitTime_deployingContract_start)/time.Nanosecond))
	log.SetPrefix("\nNaive Protocol - Gas Cost ")
	log.Println("Sender Round 1 : ",tx.Gas())




	/*send encrypted file chunks to receiver as a byte array*/
	/**//*sender sending byte array to receiver time starts*/
	sender_round1_communicationTime_sendingEncryptedGateOutputs_start := time.Now()
	SenToRecObject.EncryptedFileChunks=encryptedFileChunks

	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}
	channel_SenToRec<-byteArrayFromSenToRec

	sender_round1_communicationTime_sendingEncryptedGateOutputs_stop := time.Now()
	/**//*sender sending byte array to receiver time ends*/
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round1_communicationTime_sendingEncryptedGateOutputs_time : ", int64(sender_round1_communicationTime_sendingEncryptedGateOutputs_stop.Sub(sender_round1_communicationTime_sendingEncryptedGateOutputs_start)/time.Nanosecond))




	/*Sender waiting until the receiver accepts*/
	/**//*Sender wait time for receiver to accept starts*/
	sender_round2_waitingTime_forReceiverToAccept_start := time.Now()

	//creating a contract instance
	instance,err:= import_contracts.NewProtocol2(contractAddress[pairIndex],client)
	if err != nil {
			log.Fatal(err)
	}

	//waiting for the receiver to accept
	//Sender calls the contract state variable Phase after every 1 second to know the status of the acceptance.
	//If true, receiverAcceptedStatus is set as true, loop breaks and proceeded further
	//If receiverAcceptedStatus is false after the timeout period, the sender has not accepted and the protocol stops.

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
	/**//*Sender wait time for receiver to accept ends*/
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round2_waitingTime_forReceiverToAccept_time : ", int64(sender_round2_waitingTime_forReceiverToAccept_stop.Sub(sender_round2_waitingTime_forReceiverToAccept_start)/time.Nanosecond))





	if(!receiverAcceptedStatus){
		//Receiver has not accepted within time
			}else{


							/*Sender Reveals the key after the receiver has accepted within the time limit*/
							/**//*Key reveal commit time starts*/
							sender_round3_commitTime_revealKey_start := time.Now()
							var keyToCon []byte
							for i:=0;i<keyLength;i++{
								keyToCon=append(keyToCon,key[i])
							}

							//key reveal functon called
							tx,err=instance.KeyReveal(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},keyToCon)
							if err!=nil{
								log.Fatal(err)
							}

							_,err=bind.WaitMined(ctx,client,tx)
							if err!=nil{
								log.Fatal(err)
							}

							sender_round3_commitTime_revealKey_stop := time.Now()
							/**//*Key reveal commit time ends*/
							log.SetPrefix("\nNaive Protocol - Time ")
							log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round3_commitTime_revealKey_time : ", int64(sender_round3_commitTime_revealKey_stop.Sub(sender_round3_commitTime_revealKey_start)/time.Nanosecond))
							log.SetPrefix("\nNaive Protocol - Gas Cost ")
							log.Println("Sender Round 3 : ",tx.Gas())



							//Waiting for receiver to give complain or a confirmation that he has received the correct file.
							//If Receiver gives a confirmation, then the phase goes to finished in the contract.The loop
							//breaks and the protocol from the sender side is finished.
							//If receiver doesnt give confirmation or gives a complain, then the phase is till in the 'keyRevealed' after the deadline,
							//then the receiver then calls the SenderGetEther function to get the ether.
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
									log.SetPrefix("\nNaive Protocol - Time ")
									log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_time : ", int64(sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_stop.Sub(sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_start)/time.Nanosecond))

									//Phase
									Phase,err:=instance.Phase(nil)
									if err!=nil{
										log.Fatal(err)
									}

									if(Phase==3){

										 //sender getting the ether since the receiver has not finished the protocol
										 //The following loop is to check whether the current blockchain time has exceeded the deadline time set for receiver to produce a
										 //complain or give a confirmation.
										 timeout,err:=instance.Timeout(nil)
										 if err!=nil{
													 log.Fatal(err)
										 }

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


										//senderGetEther function is being after the receiver walked away from the protocol after key is revealed.
										/**//*Sender commit time for Get ether function call to be mined starts*/
										sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start:=time.Now()
										tx,err=instance.SenderGetEther(&bind.TransactOpts{
												From:authSender.From,
												Signer:authSender.Signer,
											})
										if err!=nil{
													log.Fatal(err)
										}

										_,err=bind.WaitMined(ctx,client,tx)
										if err!=nil{
											log.Fatal(err)
										}

										log.SetPrefix("\nNaive Protocol - Gas Cost ")
										log.Println("Sender Round 4 ",tx.Gas())
										sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop:=time.Now()
										/**//*Sender commit time for Get ether function call to be mined ends*/
										log.SetPrefix("\nNaive Protocol - Time ")
										log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_time : ", int64(sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop.Sub(sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start)/time.Nanosecond))
							}
				}
	wg.Done()
}


//Receiver function
//The functoin executes honest receiver's role in fairswap
func Receiver(ctx context.Context,	authReceiver *bind.TransactOpts,client *ethclient.Client,channel_SenToRec chan []byte,wg *sync.WaitGroup,pairIndex int){


		//Initially the receiver waits for sender to send encrypted file chunks
		/**//*Waiting time for receiver to get the encrypted file chunks starts*/
		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start := time.Now()

		//getting the encoded filechunks array from sender
		byteArrayFromSenToRec:=<-channel_SenToRec

		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop := time.Now()
		/**//*Waiting time for receiver to get the encrypted file chunks ends*/
		log.SetPrefix("\nNaive Protocol - Time ")
		log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_time : %d \n", int64(receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop.Sub(receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start)/time.Nanosecond))



		//the receiver does the following things before accepting
		//1. Unmarshals the encrypted file byte arrray got from sender
		//2. Gets the hashes of encrypted file chunks put in the contract by the sender
		//3. Compares the Hashes of encrypted file got from the contract with those of hashes calculated from the encrypted file chunks got from the sender.
		/**//*the time for receiver to do all these 3 procedure starts*/
		receiver_round1_computingTime_merkleRootCheck_start := time.Now()

		var Receiver_SenToRecObj SenderToReceiverStruct

		//1.Unmarshalling
		err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
		if err != nil {
				log.Fatal(err)
		}

		fmt.Println("\n\nReceiver contract address : ",contractAddress[pairIndex])

		instance,err := import_contracts.NewProtocol2(contractAddress[pairIndex],client)
		if err != nil {
				log.Fatal(err)
		}

		fmt.Println("hi")
				//get the number of filechunks from the contract
	 _contractNoOfFileChunks,err := instance.NoOfFileChunks(nil)
	 if err!=nil{
		 log.Fatal(err)
	 }

	 fmt.Println("Hello")

	 boolHashedEncryptedFileCheck:=true
	 //3.Comparision
	 //Conditions checked :
	 //a) whether the number of hashed encrypted files chunks given to the contract matches with the number of encrypted file chunks got from sender
	 //b) Whether the hash of encrypted file chunk got from sender matches with the hash of encrypted file chunk given to the contract
	 //if both these conditions satisfy then the protocol is proceeded further else discarded by the receiver
	 if(_contractNoOfFileChunks.Cmp(big.NewInt(int64(len(Receiver_SenToRecObj.EncryptedFileChunks))))==0){
		var i int64
		for i=0;i<int64(noOfFileChunks);i++ {
			//2.Get Hashes of encrypted file from contract
			value,err:=instance.HashedEncryptedFileChunks(nil,big.NewInt(i))
			if err!=nil{
				log.Fatal(err)
			}
			if(value!=Hash(Receiver_SenToRecObj.EncryptedFileChunks[i])){
				boolHashedEncryptedFileCheck=false
			}
		}
	}else{
		boolHashedEncryptedFileCheck=false
	}

	receiver_round1_computingTime_merkleRootCheck_stop := time.Now()
	/**//*the time for receiver to do all these 3 procedure ends*/
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round1_computingTime_merkleRootCheck_time : ", int64(receiver_round1_computingTime_merkleRootCheck_stop.Sub(receiver_round1_computingTime_merkleRootCheck_start)/time.Nanosecond))





	//Proceeds further if the comparison is positive
	if(boolHashedEncryptedFileCheck){


			fmt.Println("Inside")

			//the receiver accepts
			/**//*Commit time for receiver to accept starts*/
			receiver_round2_commitTime_Accept_start := time.Now()

			value:=big.NewInt(100)

				tx,err:=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
			})
			if err != nil {
					log.Fatal(err)
			}
			_,err=bind.WaitMined(ctx,client,tx)
			if err!=nil{
				log.Fatal(err)
			}

			receiver_round2_commitTime_Accept_stop := time.Now()
			/**//*Commit time for receiver to accept ends*/
			log.SetPrefix("\nNaive Protocol - Time ")
			log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round2_commitTime_Accept_time : ",int64(receiver_round2_commitTime_Accept_stop.Sub(receiver_round2_commitTime_Accept_start)/time.Nanosecond))
			log.SetPrefix("\nNaive Protocol - Gas Cost ")
			log.Println("Receiver Round 2 : ",tx.Gas())




			//waiting for the sender to reveal the key
			//If the sender reveals the key within the deadline the loop breaks and proceeded further for the extraction.
			//If the sender doesn't reveal the key within the deadline, receiver will call the ReceiverGetEther function to get the ether.
			/**//* the sender waiting time for sender to reveal the key starts*/
			receiver_round3_waitingTime_forSenderToRevealKey_start := time.Now()
			senderKeyRevealedStatus:=false
			var contract_Key []byte
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
				time.Sleep(time.Second *1)
			}
			receiver_round3_waitingTime_forSenderToRevealKey_stop := time.Now()
			/**//* the sender waiting time for sender to reveal the key ends*/
			log.SetPrefix("\nNaive Protocol - Time ")
			log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round3_waitingTime_forSenderToRevealKey_time : ", int64(receiver_round3_waitingTime_forSenderToRevealKey_stop.Sub(receiver_round3_waitingTime_forSenderToRevealKey_start)/time.Nanosecond))



			if(!senderKeyRevealedStatus){



					 //sender has not revealed the wrong key
					 timeout,err:=instance.Timeout(nil)
		 			 if err!=nil{
		 						 log.Fatal(err)
		 			 }

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

					 // receiverGetEther function being called
					 /**//*the commit time for receiverGetEther function call starts*/
					 	receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start := time.Now()
						tx,err=instance.ReceiverGetEther(&bind.TransactOpts{
								From:authReceiver.From,
								Signer:authReceiver.Signer,
								Value: nil,
							})
						if err!=nil{
									log.Fatal(err)
						}

						_,err=bind.WaitMined(ctx,client,tx)
						if err!=nil{
							log.Fatal(err)
						}
						receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop := time.Now()
						/**//*the commit time for receiverGetEther function call ends*/
						log.SetPrefix("\nNaive Protocol - Gas Cost ")
						log.Println("Receiver Round 3 (Sender Not Revealed Key) : ",tx.Cost())
						log.SetPrefix("\nNaive Protocol - Time ")
						log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round3_commitTime_senderKeyNotRevealedGettingEther_time : ", int64(receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop.Sub(receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start)/time.Nanosecond))




			} else{

						//Extraction
						//1. The receiver gets the hashes of file chunks from the contract
						//2. Extraction function is called : The function basically decrypts the encrypted file chunks got from the sender hashes it and then compares
						// with the hash of file chunks got from the contract. If there is some discrepancy then the index of that gate is returned by the function.
						//Then a complain is given.
						//If the comparision is positive, then the receiver can either give a confirmation that he received the correct file or else walk away.
						//In our implemenation doesn't give this confirmation
						/**//*to calculate the time required for receiver to perfoem the extraction function starts*/
						receiver_round4_computingTime_extraction_start := time.Now()

						 _contractHashedFileChunks:=make([][hashFunctionOutputBitSize]byte,noOfFileChunks)
						var i int64
						for i=0;i<int64(noOfFileChunks);i++{
								value,err:=instance.HashedFileChunks(nil,big.NewInt(i))
								if err!=nil{
										log.Fatal(err)
								}
								_contractHashedFileChunks[i]=value
						}

						_ReceiverKey:=make([]byte,keyLength)
						for i:=0;i<keyLength;i++{
							_ReceiverKey[i]=contract_Key[i]
						}

						isComplain,index:=Extract(_contractHashedFileChunks,Receiver_SenToRecObj.EncryptedFileChunks,_ReceiverKey)
						receiver_round4_computingTime_extraction_stop := time.Now()
						/**//*to calculate the time required for receiver to perfoem the extraction function ends*/
						log.SetPrefix("\nNaive Protocol - Time ")
						log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round4_computingTime_extraction_time : ", int64(receiver_round4_computingTime_extraction_stop.Sub(receiver_round4_computingTime_extraction_start)/time.Nanosecond))

						//index nil denotes no complain
						if(!isComplain){
							log.SetPrefix("\nNaive Protocol - Gas Cost ")
								log.Println("Receiver Round 4 : ",0)
						} else{

								//complain function being called
								/**//*commit time for receiver to produce the complaint starts*/
								receiver_round4_commitTime_complain_start := time.Now()
								var complaintChunk []byte
								for i:=0;i<buffer_size;i++{
									  complaintChunk=append(complaintChunk,Receiver_SenToRecObj.EncryptedFileChunks[index][i])
								}

								tx,err=instance.Complain(&bind.TransactOpts{
									From:authReceiver.From,
									Signer:authReceiver.Signer,
								},complaintChunk,big.NewInt(int64(index)))
								if err!=nil{
									log.Fatal(err)
								}

								_,err=bind.WaitMined(ctx,client,tx)
								if err!=nil{
									log.Fatal(err)
								}

								receiver_round4_commitTime_complain_stop := time.Now()
								/**//*commit time for receiver to produce the complaint ends*/
								log.SetPrefix("\nNaive Protocol - Time ")
								log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round4_commitTime_complain_time : ", int64(receiver_round4_commitTime_complain_stop.Sub(receiver_round4_commitTime_complain_start)/time.Nanosecond))
								log.SetPrefix("\nNaive Protocol - Gas Cost ")
								log.Println("Receiver Round 4 : ",tx.Cost())

				}
			}
		}	else{
				// log.Println("\n->Receiver\n Hash encrypted file from the contract and the hash of encrypted file from the contract doesnt match")
		}
	wg.Done()
}


func FileRead() ([][]byte){
	var currentByte int64 = 0
	fileBuffer := make([]byte, buffer_size)
	FileChunks := make([][]byte, noOfFileChunks)
	file, err := os.Open("a.txt")
	if err != nil {
			log.Fatal(err)
	}
	count:=0
	for {
			n,err := file.ReadAt(fileBuffer, currentByte)
			if err!=nil{
				fmt.Println(err)
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


func Hash(_toBeHashed []byte) ([hashFunctionOutputBitSize]byte){

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

func keyGenerate() ([]byte){

		keyGen := make([]byte, keyLength)
		key:=make([]byte,keyLength)

		//Here
	  _, err := rand.Read(keyGen)
    if err != nil {
        log.Fatal(err)
    }


		for i:=0;i<keyLength;i++{
			key[i]=keyGen[i]
		}
		return key
}



func Encrypt(plainText []byte,key []byte) ([]byte) {

			encryptedtext:=make([]byte,buffer_size)
			for i:=0;i<buffer_size;i++{
				encryptedtext[i]=key[i]^plainText[i]  //xor operation
			}
			return encryptedtext
}

func setPrice(object *SenderToContractStruct){
	object.Price=*big.NewInt(80)     //80 wei
}

func Decrypt(encryptedtext []byte,key []byte) ([]byte){

			plainText:=make([]byte,buffer_size)
			for i:=0;i<buffer_size;i++{
				plainText[i]=key[i]^encryptedtext[i]
			}

			return plainText
}

func Extract(_hashedFileChunks [][hashFunctionOutputBitSize]byte, _encryptedFileChunks [][]byte, key []byte)  (bool,int){

		for i:=0;i<noOfFileChunks;i++{
			if(_hashedFileChunks[i]!=Hash(Decrypt(_encryptedFileChunks[i],key))){
				return true,i
			}
		}
		return false,0
}

func HashChunks(_tobeHashedArray [][]byte) ([][hashFunctionOutputBitSize]byte){

		hashedArray:=make([][hashFunctionOutputBitSize]byte,noOfFileChunks)
		for i:=0;i<noOfFileChunks;i++{
			hashedArray[i]=Hash(_tobeHashedArray[i])
		}
		return hashedArray
}


func EncryptFileChunks(_fileChunks [][]byte,key []byte) ([][]byte){

	encryptedFileChunks:=make([][]byte,noOfFileChunks)
	for i:=0;i<noOfFileChunks;i++{
		encryptedFileChunks[i]=Encrypt(_fileChunks[i],key)
	}
	return encryptedFileChunks
}


 func createMaliciousHashIndexExchangeFileChunk(HashedEncryptedFileChunks [][hashFunctionOutputBitSize]byte) ([][hashFunctionOutputBitSize]byte){

	maliciousHashIndexExchangeFileChunks:=make([][hashFunctionOutputBitSize]byte,noOfFileChunks)
	for i:=0;i<noOfFileChunks;i++{
			maliciousHashIndexExchangeFileChunks[i]=HashedEncryptedFileChunks[i]
	}
	maliciousHashIndexExchangeFileChunks[0]=HashedEncryptedFileChunks[2]
	maliciousHashIndexExchangeFileChunks[2]=HashedEncryptedFileChunks[0]

	return maliciousHashIndexExchangeFileChunks
}


func MaliciousEncryptedFileChunks(_fileChunks [][]byte,key []byte,maliciouskey []byte) ([][]byte){

	maliciousEncryptedFileChunks:=make([][]byte,noOfFileChunks)
	for i:=0;i<noOfFileChunks;i++{
		maliciousEncryptedFileChunks[i]=Encrypt(_fileChunks[i],key)
	}
	maliciousEncryptedFileChunks[5]=Encrypt(_fileChunks[5],maliciouskey)
	return maliciousEncryptedFileChunks
}

//
// func Sender_malicious_keyNotReveal(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup){
//
// 	//reading from a file "a.txt" and storing bytearray of 32 bytes in FileChunks of size equal to noOfFileChunks
// 	sender_round1_computingTime_calculatingVariables_start := time.Now()
// 	FileChunks:=FileRead()
//
// 	//objects
// 	var SenToConObject SenderToContractStruct
// 	var SenToRecObject SenderToReceiverStruct
//
//
// 	setPrice(&SenToConObject)
// 	key:=keyGenerate()
// 	// maliciouskey:=keyGenerate()
// 	HashedFileChunks:=HashChunks(FileChunks)
// 	for i:=0;i<noOfFileChunks;i++{
// 		SenToConObject.HashedFileChunks=append(SenToConObject.HashedFileChunks,HashedFileChunks[i])
// 	}
//
//  	encryptedFileChunks := EncryptFileChunks(FileChunks,key)
//   // maliciousEncryptedFileChunks := MaliciousEncryptedFileChunks(FileChunks,key,maliciouskey)
//
// 	//fmt.Println("sender encrypt file chunks",encryptedFileChunks)
//
// 	HashedEncryptedFileChunks:=HashChunks(encryptedFileChunks)
// 	// maliciousHashIndexExchangeFileChunks := createMaliciousHashIndexExchangeFileChunk(HashedEncryptedFileChunks)
// 	for i:=0;i<noOfFileChunks;i++{
// 		SenToConObject.HashedEncryptedFileChunks=append(SenToConObject.HashedEncryptedFileChunks,HashedEncryptedFileChunks[i])
// 	}
//
// 	sender_round1_computingTime_calculatingVariables_stop:=time.Now()
// 	log.SetPrefix("\nNaive Protocol - Time ")
// 	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round1_computingTime_calculatingVariables_time : ", sender_round1_computingTime_calculatingVariables_stop.Sub(sender_round1_computingTime_calculatingVariables_start))
//
//
// 	/*Deploy contract*/
// 	var err error
// 	var tx *types.Transaction
// 	sender_round1_commitTime_deployingContract_start := time.Now()
// 	contractAddress,tx,_, err=import_contracts.DeployProtocol2(
// 		authSender,
// 		client,
// 		&SenToConObject.Price,
// 		SenToConObject.HashedFileChunks,
// 		SenToConObject.HashedEncryptedFileChunks,
// 	)
// 	if err!=nil{
// 		log.Fatal(err)
// 	}
// 	_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
// 	sender_round1_commitTime_deployingContract_stop := time.Now()
// 	log.SetPrefix("\nNaive Protocol - Time ")
// 	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round1_commitTime_deployingContract_time : ", sender_round1_commitTime_deployingContract_stop.Sub(sender_round1_commitTime_deployingContract_start))
// 	log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 	log.Println("Sender Round 1 : ",tx.Gas())
//
//
// 	sender_round1_communicationTime_sendingEncryptedGateOutputs_start := time.Now()
// 	/*SenToRecObject initialize*/
// 	SenToRecObject.EncryptedFileChunks=encryptedFileChunks
//
// 	//sending the encodedGateOutputs,keycommitment and receiver entry key to the receiver through a byte array(got after applying json marshal to the SenToRecObject)
// 	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	channel_SenToRec<-byteArrayFromSenToRec
//
// 	sender_round1_communicationTime_sendingEncryptedGateOutputs_stop := time.Now()
// 	log.SetPrefix("\nNaive Protocol - Time ")
// 	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round1_communicationTime_sendingEncryptedGateOutputs_time : ", sender_round1_communicationTime_sendingEncryptedGateOutputs_stop.Sub(sender_round1_communicationTime_sendingEncryptedGateOutputs_start))
// 	//creating a contract instance
//
// 	sender_round2_waitingTime_forReceiverToAccept_start := time.Now()
//
// 	instance,err:= import_contracts.NewProtocol2(contractAddress,client)
// 	if err != nil {
// 			log.Fatal(err)
// 	}
//
//
// 	receiverAcceptedStatus:=false
// 	for i:=0;i<waitingTime;i++ {
//
// 		Phase,err:=instance.Phase(nil)
// 		if err!=nil{
// 			log.Fatal(err)
// 		}
// 		if(Phase==2){
// 			receiverAcceptedStatus=true
// 			break;
// 		}
// 		time.Sleep(time.Second *1)
// 	}
//
// 	sender_round2_waitingTime_forReceiverToAccept_stop := time.Now()
// 	log.SetPrefix("\nNaive Protocol - Time ")
// 	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round2_waitingTime_forReceiverToAccept_time : ", sender_round2_waitingTime_forReceiverToAccept_stop.Sub(sender_round2_waitingTime_forReceiverToAccept_start))
//
//
// 	if(!receiverAcceptedStatus){
// 		//Receiver has not accepted within time
// 			}else{
//
// 			}
// 	wg.Done()
// }
//
// func Sender_malicious_FileChunkWrongEnc(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup){
//
// 	//reading from a file "a.txt" and storing bytearray of 32 bytes in FileChunks of size equal to noOfFileChunks
// 	sender_round1_computingTime_calculatingVariables_start := time.Now()
// 	FileChunks:=FileRead()
//
// 	//objects
// 	var SenToConObject SenderToContractStruct
// 	var SenToRecObject SenderToReceiverStruct
//
//
// 	setPrice(&SenToConObject)
// 	key:=keyGenerate()
// 	maliciouskey:=keyGenerate()
// 	HashedFileChunks:=HashChunks(FileChunks)
// 	for i:=0;i<noOfFileChunks;i++{
// 		SenToConObject.HashedFileChunks=append(SenToConObject.HashedFileChunks,HashedFileChunks[i])
// 	}
//
//  	// encryptedFileChunks := EncryptFileChunks(FileChunks,key)
//   maliciousEncryptedFileChunks := MaliciousEncryptedFileChunks(FileChunks,key,maliciouskey)
//
// 	//fmt.Println("sender encrypt file chunks",encryptedFileChunks)
//
// 	HashedEncryptedFileChunks:=HashChunks(maliciousEncryptedFileChunks)
//
// 	for i:=0;i<noOfFileChunks;i++{
// 		SenToConObject.HashedEncryptedFileChunks=append(SenToConObject.HashedEncryptedFileChunks,HashedEncryptedFileChunks[i])
// 	}
//
// 	sender_round1_computingTime_calculatingVariables_stop:=time.Now()
// 	log.SetPrefix("\nNaive Protocol - Time ")
// 	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round1_computingTime_calculatingVariables_time : ", sender_round1_computingTime_calculatingVariables_stop.Sub(sender_round1_computingTime_calculatingVariables_start))
//
//
// 	/*Deploy contract*/
// 	var err error
// 	var tx *types.Transaction
// 	sender_round1_commitTime_deployingContract_start := time.Now()
// 	contractAddress,tx,_, err=import_contracts.DeployProtocol2(
// 		authSender,
// 		client,
// 		&SenToConObject.Price,
// 		SenToConObject.HashedFileChunks,
// 		SenToConObject.HashedEncryptedFileChunks,
// 	)
// 	if err!=nil{
// 		log.Fatal(err)
// 	}
// 	_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
// 	sender_round1_commitTime_deployingContract_stop := time.Now()
// 	log.SetPrefix("\nNaive Protocol - Time ")
// 	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round1_commitTime_deployingContract_time : ", sender_round1_commitTime_deployingContract_stop.Sub(sender_round1_commitTime_deployingContract_start))
// 	log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 	log.Println("Sender Round 1 : ",tx.Gas())
//
//
// 	sender_round1_communicationTime_sendingEncryptedGateOutputs_start := time.Now()
// 	/*SenToRecObject initialize*/
// 	SenToRecObject.EncryptedFileChunks=maliciousEncryptedFileChunks
//
// 	//sending the encodedGateOutputs,keycommitment and receiver entry key to the receiver through a byte array(got after applying json marshal to the SenToRecObject)
// 	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	channel_SenToRec<-byteArrayFromSenToRec
//
// 	sender_round1_communicationTime_sendingEncryptedGateOutputs_stop := time.Now()
// 	log.SetPrefix("\nNaive Protocol - Time ")
// 	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round1_communicationTime_sendingEncryptedGateOutputs_time : ", sender_round1_communicationTime_sendingEncryptedGateOutputs_stop.Sub(sender_round1_communicationTime_sendingEncryptedGateOutputs_start))
// 	//creating a contract instance
//
// 	sender_round2_waitingTime_forReceiverToAccept_start := time.Now()
//
// 	instance,err:= import_contracts.NewProtocol2(contractAddress,client)
// 	if err != nil {
// 			log.Fatal(err)
// 	}
//
//
// 	receiverAcceptedStatus:=false
// 	for i:=0;i<waitingTime;i++ {
//
// 		Phase,err:=instance.Phase(nil)
// 		if err!=nil{
// 			log.Fatal(err)
// 		}
// 		if(Phase==2){
// 			receiverAcceptedStatus=true
// 			break;
// 		}
// 		time.Sleep(time.Second *1)
// 	}
//
// 	sender_round2_waitingTime_forReceiverToAccept_stop := time.Now()
// 	log.SetPrefix("\nNaive Protocol - Time ")
// 	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round2_waitingTime_forReceiverToAccept_time : ", sender_round2_waitingTime_forReceiverToAccept_stop.Sub(sender_round2_waitingTime_forReceiverToAccept_start))
//
//
// 	if(!receiverAcceptedStatus){
// 		//Receiver has not accepted within time
// 			}else{
//
// 							sender_round3_commitTime_revealKey_start := time.Now()
// 							var keyToCon []byte
// 							for i:=0;i<keyLength;i++{
// 								keyToCon=append(keyToCon,key[i])
// 							}
//
// 							tx,err=instance.KeyReveal(&bind.TransactOpts{
// 									From:authSender.From,
// 									Signer:authSender.Signer,
// 									Value: nil,
// 							},keyToCon)
// 							if err!=nil{
// 								log.Fatal(err)
// 							}
// 							_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
//
// 							sender_round3_commitTime_revealKey_stop := time.Now()
// 							log.SetPrefix("\nNaive Protocol - Time ")
// 							log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round3_commitTime_revealKey_time : ", sender_round3_commitTime_revealKey_stop.Sub(sender_round3_commitTime_revealKey_start))
// 							//_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
// 							log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 							log.Println("Sender Round 3 : ",tx.Gas())
// 							/*fmt.Println("\n->Sender\nReceiver have accepted within the timeout, Key Revealed")
// 							channel_SenToMain<-"I have not revealed the key"*/
//
// 			sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_start:=time.Now()
//
// 			for i:=0;i<waitingTime;i++{
//
// 					Phase,err:=instance.Phase(nil)
// 					if err!=nil{
// 						log.Fatal(err)
// 					}
// 					if(Phase==4){
// 						break
// 					}
// 					time.Sleep(time.Second *1)
// 			}
//
// 			sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_stop:=time.Now()
// 			log.SetPrefix("\nNaive Protocol - Time ")
// 			log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_time : ", sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_stop.Sub(sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_start))
//
// 			// gasLimit := uint64(2100000000)
//
// 			//the sender getting the ether
// 			Phase,err:=instance.Phase(nil)
// 			if err!=nil{
// 				log.Fatal(err)
// 			}
// 		//	fmt.Println(Phase)
//
// 			if(Phase==3){
// 			 //time.Sleep(time.Second *60)
// 			 // count:=0
//
// 			 timeout,err:=instance.Timeout(nil)
// 			 if err!=nil{
// 						 log.Fatal(err)
// 			 }
//
// 			 for{
//
// 				  _,err:=instance.Nowy(&bind.TransactOpts{
// 	 					From:authSender.From,
// 	 					Signer:authSender.Signer,
// 	 				})
// 					if err!=nil{
// 								log.Fatal(err)
// 					}
// 					_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
// 					Now,err:=instance.Now(nil)
// 					if err!=nil{
// 								log.Fatal(err)
// 					}
// 					if(Now.Cmp(timeout)==1){
//  					 break
//  				 	}
// 					time.Sleep(time.Second*1)
// 			 }
//
// 			 sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start:=time.Now()
//
// 			tx,err=instance.SenderGetEther(&bind.TransactOpts{
// 					From:authSender.From,
// 					Signer:authSender.Signer,
// 					// Value: big.NewInt(10000000000),
// 					// GasLimit:gasLimit,
// 				})
// 			if err!=nil{
// 						log.Fatal(err)
// 			}
// 			_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
//
// 			log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 			log.Println("Sender Round 4 ",tx.Gas())
// 			sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop:=time.Now()
// 			log.SetPrefix("\nNaive Protocol - Time ")
// 			log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_time : ", sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop.Sub(sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start))
//
// 			}
// 		}
// 	wg.Done()
// }
//
//
// func Receiver_malicious_HashEncFileMismatch(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup){
//
// 		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start := time.Now()
//
// 		byteArrayFromSenToRec:=<-channel_SenToRec
//
// 		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop := time.Now()
// 		log.SetPrefix("\nNaive Protocol - Time ")
// 		log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_time : ", receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop.Sub(receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start))
//
// 		receiver_round1_computingTime_merkleRootCheck_start := time.Now()
//
// 		var Receiver_SenToRecObj SenderToReceiverStruct
//
// 		err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
// 		if err != nil {
// 				log.Fatal(err)
// 		}
//
// 		instance,err := import_contracts.NewProtocol2(contractAddress,client)
// 		if err != nil {
// 				log.Fatal(err)
// 		}
//
// 	 _contractNoOfFileChunks,err := instance.NoOfFileChunks(nil)
// 	 if err!=nil{
// 		 log.Fatal(err)
// 	 }
//
// 	 boolHashedEncryptedFileCheck:=true
//
// 	 //if(big.NewInt(int64(len(Receiver_SenToRecObj.EncryptedFileChunks)))==_contractNoOfFileChunks)
// 	 if(_contractNoOfFileChunks.Cmp(big.NewInt(int64(len(Receiver_SenToRecObj.EncryptedFileChunks))))==0){
// 		var i int64
// 		for i=0;i<int64(noOfFileChunks);i++ {
// 			value,err:=instance.HashedEncryptedFileChunks(nil,big.NewInt(i))
// 			if err!=nil{
// 				log.Fatal(err)
// 			}
// 			if(value!=Hash(Receiver_SenToRecObj.EncryptedFileChunks[i])){
// 				boolHashedEncryptedFileCheck=false
// 			}
// 		}
// 	}else{
// 		boolHashedEncryptedFileCheck=false
// 	}
//
// 	receiver_round1_computingTime_merkleRootCheck_stop := time.Now()
// 	log.SetPrefix("\nNaive Protocol - Time ")
// 	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round1_computingTime_merkleRootCheck_time : ", receiver_round1_computingTime_merkleRootCheck_stop.Sub(receiver_round1_computingTime_merkleRootCheck_start))
//
// 	if(boolHashedEncryptedFileCheck){
//
// 				receiver_round2_commitTime_Accept_start := time.Now()
//
// 				value:=big.NewInt(100)
//
// 				tx,err:=instance.Accept(&bind.TransactOpts{
// 				From:authReceiver.From,
// 				Signer:authReceiver.Signer,
// 				Value:value,
// 				// GasLimit:gasLimit,
// 			})
// 			if err != nil {
// 					log.Fatal(err)
// 			}
// 			_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
//
// 			receiver_round2_commitTime_Accept_stop := time.Now()
// 			log.SetPrefix("\nNaive Protocol - Time ")
// 			log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round2_commitTime_Accept_time : ", receiver_round2_commitTime_Accept_stop.Sub(receiver_round2_commitTime_Accept_start))
// 			log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 			log.Println("Receiver Round 2 : ",tx.Gas())
// 			//Sending the finished accepted stage message to the main function
//
//
// 			//The contract is in accepted stage
// 			//The sender
// 			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
// 			//case 2 : The sender doesn't reveal the key within time. senderKeyRevealedStatus remains false after the for loop. Receiver later calls the contract function
// 			//			   ReceiverGetEther to get the ether.
// 			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and senderKeyRevealedStatus is true after the for loop.
// 			receiver_round3_waitingTime_forSenderToRevealKey_start := time.Now()
// 			senderKeyRevealedStatus:=false
// 			var contract_Key []byte
// 			for i:=0;i<waitingTime;i++ {
//
// 				Phase,err:=instance.Phase(nil)
// 				if err!=nil{
// 					log.Fatal(err)
// 				}
// 				if(Phase==3){
// 					contract_Key,err=instance.Key(nil)
// 					if err!=nil{
// 						log.Fatal(err)
// 					}
// 					senderKeyRevealedStatus=true
// 					break;
// 				}
// 				if(Phase==4){
// 					senderKeyRevealedStatus=true
// 					break;
// 				}
// 				time.Sleep(time.Second *1)
// 			}
// 			receiver_round3_waitingTime_forSenderToRevealKey_stop := time.Now()
// 			log.SetPrefix("\nNaive Protocol - Time ")
// 			log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round3_waitingTime_forSenderToRevealKey_time : ", receiver_round3_waitingTime_forSenderToRevealKey_stop.Sub(receiver_round3_waitingTime_forSenderToRevealKey_start))
//
//
//
// 			if(!senderKeyRevealedStatus){
//
// 			 timeout,err:=instance.Timeout(nil)
//  			 if err!=nil{
//  						 log.Fatal(err)
//  			 }
//
//  			 for{
//
//  				  _,err:=instance.Nowy(&bind.TransactOpts{
//  	 					From:authReceiver.From,
//  	 					Signer:authReceiver.Signer,
//  	 				})
//  					if err!=nil{
//  								log.Fatal(err)
//  					}
//  					_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
//  					Now,err:=instance.Now(nil)
//  					if err!=nil{
//  								log.Fatal(err)
//  					}
//  					if(Now.Cmp(timeout)==1){
//   					 break
//   				 	}
//  					time.Sleep(time.Second*1)
//  			 }
//
// 			 	receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start := time.Now()
// 				tx,err=instance.ReceiverGetEther(&bind.TransactOpts{
// 						From:authReceiver.From,
// 						Signer:authReceiver.Signer,
// 						Value: nil,
// 						// GasLimit:gasLimit,
// 					})
// 				if err!=nil{
// 							log.Fatal(err)
// 				}
// 				_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
// 				receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop := time.Now()
// 				log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 				log.Println("Receiver Round 3 (Sender Not Revealed Key) : ",tx.Cost())
// 				log.SetPrefix("\nNaive Protocol - Time ")
// 				log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round3_commitTime_senderKeyNotRevealedGettingEther_time : ", receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop.Sub(receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start))
//
//
// 			} else{
//
// 				//Extracting function being called
//
// 				receiver_round4_computingTime_extraction_start := time.Now()
//
// 				 _contractHashedFileChunks:=make([][hashFunctionOutputBitSize]byte,noOfFileChunks)
// 				var i int64
// 				for i=0;i<int64(noOfFileChunks);i++{
// 						value,err:=instance.HashedFileChunks(nil,big.NewInt(i))
// 						if err!=nil{
// 								log.Fatal(err)
// 						}
// 						_contractHashedFileChunks[i]=value
// 				}
//
// 				_ReceiverKey:=make([]byte,keyLength)
// 				for i:=0;i<keyLength;i++{
// 					_ReceiverKey[i]=contract_Key[i]
// 				}
//
// 				isComplain,index:=Extract(_contractHashedFileChunks,Receiver_SenToRecObj.EncryptedFileChunks,_ReceiverKey)
//
// 				isComplain=true
// 				//malicious reciever
// 				receiver_round4_computingTime_extraction_stop := time.Now()
// 				log.SetPrefix("\nNaive Protocol - Time ")
// 				log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round4_computingTime_extraction_time : ", receiver_round4_computingTime_extraction_stop.Sub(receiver_round4_computingTime_extraction_start))
//
// 				//index nil denotes no complain
// 				if(!isComplain){
// 					log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 						log.Println("Receiver Round 4 : ",0)
// 				} else{
//
// 						//complain function being called
// 						receiver_round4_commitTime_complain_start := time.Now()
// 						var complaintChunk []byte
// 						index=4
//
// 						tx,err=instance.Complain(&bind.TransactOpts{
// 							From:authReceiver.From,
// 							Signer:authReceiver.Signer,
// 							Value:nil,
// 							// GasLimit:gasLimit,
// 						},complaintChunk,big.NewInt(int64(index)))
// 						if err!=nil{
// 							log.Fatal(err)
// 						}
// 						_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
//
// 						receiver_round4_commitTime_complain_stop := time.Now()
// 						log.SetPrefix("\nNaive Protocol - Time ")
// 						log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round4_commitTime_complain_time : ", receiver_round4_commitTime_complain_stop.Sub(receiver_round4_commitTime_complain_start))
// 						log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 						log.Println("Receiver Round 4 : ",tx.Cost())
//
// 				}
// 			}
// 		}	else{
// 				// log.Println("\n->Receiver\n Hash encrypted file from the contract and the hash of encrypted file from the contract doesnt match")
// 		}
// 	wg.Done()
// }
//
//
// func Receiver_malicious_HashFileMismatch(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup){
//
// 		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start := time.Now()
//
// 		byteArrayFromSenToRec:=<-channel_SenToRec
//
// 		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop := time.Now()
// 		log.SetPrefix("\nNaive Protocol - Time ")
// 		log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_time : ", receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop.Sub(receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start))
//
// 		receiver_round1_computingTime_merkleRootCheck_start := time.Now()
//
// 		var Receiver_SenToRecObj SenderToReceiverStruct
//
// 		err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
// 		if err != nil {
// 				log.Fatal(err)
// 		}
//
// 		instance,err := import_contracts.NewProtocol2(contractAddress,client)
// 		if err != nil {
// 				log.Fatal(err)
// 		}
//
// 	 _contractNoOfFileChunks,err := instance.NoOfFileChunks(nil)
// 	 if err!=nil{
// 		 log.Fatal(err)
// 	 }
//
// 	 boolHashedEncryptedFileCheck:=true
//
// 	 //if(big.NewInt(int64(len(Receiver_SenToRecObj.EncryptedFileChunks)))==_contractNoOfFileChunks)
// 	 if(_contractNoOfFileChunks.Cmp(big.NewInt(int64(len(Receiver_SenToRecObj.EncryptedFileChunks))))==0){
// 		var i int64
// 		for i=0;i<int64(noOfFileChunks);i++ {
// 			value,err:=instance.HashedEncryptedFileChunks(nil,big.NewInt(i))
// 			if err!=nil{
// 				log.Fatal(err)
// 			}
// 			if(value!=Hash(Receiver_SenToRecObj.EncryptedFileChunks[i])){
// 				boolHashedEncryptedFileCheck=false
// 			}
// 		}
// 	}else{
// 		boolHashedEncryptedFileCheck=false
// 	}
//
// 	receiver_round1_computingTime_merkleRootCheck_stop := time.Now()
// 	log.SetPrefix("\nNaive Protocol - Time ")
// 	log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round1_computingTime_merkleRootCheck_time : ", receiver_round1_computingTime_merkleRootCheck_stop.Sub(receiver_round1_computingTime_merkleRootCheck_start))
//
// 	if(boolHashedEncryptedFileCheck){
//
// 				receiver_round2_commitTime_Accept_start := time.Now()
//
// 				value:=big.NewInt(100)
//
// 				tx,err:=instance.Accept(&bind.TransactOpts{
// 				From:authReceiver.From,
// 				Signer:authReceiver.Signer,
// 				Value:value,
// 				// GasLimit:gasLimit,
// 			})
// 			if err != nil {
// 					log.Fatal(err)
// 			}
// 			_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
//
// 			receiver_round2_commitTime_Accept_stop := time.Now()
// 			log.SetPrefix("\nNaive Protocol - Time ")
// 			log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round2_commitTime_Accept_time : ", receiver_round2_commitTime_Accept_stop.Sub(receiver_round2_commitTime_Accept_start))
// 			log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 			log.Println("Receiver Round 2 : ",tx.Gas())
// 			//Sending the finished accepted stage message to the main function
//
//
// 			//The contract is in accepted stage
// 			//The sender
// 			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
// 			//case 2 : The sender doesn't reveal the key within time. senderKeyRevealedStatus remains false after the for loop. Receiver later calls the contract function
// 			//			   ReceiverGetEther to get the ether.
// 			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and senderKeyRevealedStatus is true after the for loop.
// 			receiver_round3_waitingTime_forSenderToRevealKey_start := time.Now()
// 			senderKeyRevealedStatus:=false
// 			var contract_Key []byte
// 			for i:=0;i<waitingTime;i++ {
//
// 				Phase,err:=instance.Phase(nil)
// 				if err!=nil{
// 					log.Fatal(err)
// 				}
// 				if(Phase==3){
// 					contract_Key,err=instance.Key(nil)
// 					if err!=nil{
// 						log.Fatal(err)
// 					}
// 					senderKeyRevealedStatus=true
// 					break;
// 				}
// 				if(Phase==4){
// 					senderKeyRevealedStatus=true
// 					break;
// 				}
// 				time.Sleep(time.Second *1)
// 			}
// 			receiver_round3_waitingTime_forSenderToRevealKey_stop := time.Now()
// 			log.SetPrefix("\nNaive Protocol - Time ")
// 			log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round3_waitingTime_forSenderToRevealKey_time : ", receiver_round3_waitingTime_forSenderToRevealKey_stop.Sub(receiver_round3_waitingTime_forSenderToRevealKey_start))
//
//
//
// 			if(!senderKeyRevealedStatus){
//
//
//
// 			 timeout,err:=instance.Timeout(nil)
//  			 if err!=nil{
//  						 log.Fatal(err)
//  			 }
//
//  			 for{
//
//  				  _,err:=instance.Nowy(&bind.TransactOpts{
//  	 					From:authReceiver.From,
//  	 					Signer:authReceiver.Signer,
//  	 				})
//  					if err!=nil{
//  								log.Fatal(err)
//  					}
//  					_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
//  					Now,err:=instance.Now(nil)
//  					if err!=nil{
//  								log.Fatal(err)
//  					}
//  					if(Now.Cmp(timeout)==1){
//   					 break
//   				 	}
//  					time.Sleep(time.Second*1)
//  			 }
//
// 			 	receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start := time.Now()
// 				tx,err=instance.ReceiverGetEther(&bind.TransactOpts{
// 						From:authReceiver.From,
// 						Signer:authReceiver.Signer,
// 						Value: nil,
// 						// GasLimit:gasLimit,
// 					})
// 				if err!=nil{
// 							log.Fatal(err)
// 				}
// 				_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
// 				receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop := time.Now()
// 				log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 				log.Println("Receiver Round 3 (Sender Not Revealed Key) : ",tx.Cost())
// 				log.SetPrefix("\nNaive Protocol - Time ")
// 				log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round3_commitTime_senderKeyNotRevealedGettingEther_time : ", receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop.Sub(receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start))
//
//
// 			} else{
//
// 				//Extracting function being called
//
// 				receiver_round4_computingTime_extraction_start := time.Now()
//
// 				 _contractHashedFileChunks:=make([][hashFunctionOutputBitSize]byte,noOfFileChunks)
// 				var i int64
// 				for i=0;i<int64(noOfFileChunks);i++{
// 						value,err:=instance.HashedFileChunks(nil,big.NewInt(i))
// 						if err!=nil{
// 								log.Fatal(err)
// 						}
// 						_contractHashedFileChunks[i]=value
// 				}
//
// 				_ReceiverKey:=make([]byte,keyLength)
// 				for i:=0;i<keyLength;i++{
// 					_ReceiverKey[i]=contract_Key[i]
// 				}
//
// 				isComplain,index:=Extract(_contractHashedFileChunks,Receiver_SenToRecObj.EncryptedFileChunks,_ReceiverKey)
//
// 				isComplain=true
// 				//malicious reciever
// 				receiver_round4_computingTime_extraction_stop := time.Now()
// 				log.SetPrefix("\nNaive Protocol - Time ")
// 				log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round4_computingTime_extraction_time : ", receiver_round4_computingTime_extraction_stop.Sub(receiver_round4_computingTime_extraction_start))
//
// 				//index nil denotes no complain
// 				if(!isComplain){
// 					log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 						log.Println("Receiver Round 4 : ",0)
// 				} else{
//
// 						//complain function being called
// 						receiver_round4_commitTime_complain_start := time.Now()
//
// 						var complaintChunk []byte
// 						index=4
// 						for i:=0;i<buffer_size;i++{
// 							  complaintChunk=append(complaintChunk,Receiver_SenToRecObj.EncryptedFileChunks[index][i])
// 						}
//
// 						tx,err=instance.Complain(&bind.TransactOpts{
// 							From:authReceiver.From,
// 							Signer:authReceiver.Signer,
// 							Value:nil,
// 							// GasLimit:gasLimit,
// 						},complaintChunk,big.NewInt(int64(index)))
// 						if err!=nil{
// 							log.Fatal(err)
// 						}
// 						_,err=bind.WaitMined(ctx,client,tx)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
//
// 						receiver_round4_commitTime_complain_stop := time.Now()
// 						log.SetPrefix("\nNaive Protocol - Time ")
// 						log.Println("[buffer_size : ",buffer_size,", noOfFileChunks : ",noOfFileChunks,"] - receiver_round4_commitTime_complain_time : ", receiver_round4_commitTime_complain_stop.Sub(receiver_round4_commitTime_complain_start))
// 						log.SetPrefix("\nNaive Protocol - Gas Cost ")
// 						log.Println("Receiver Round 4 : ",tx.Cost())
//
// 				}
// 			}
// 		}	else{
// 				// log.Println("\n->Receiver\n Hash encrypted file from the contract and the hash of encrypted file from the contract doesnt match")
// 		}
// 	wg.Done()
// }
