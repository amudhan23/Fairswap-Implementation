package main

import(
	"fmt"
	kec "github.com/ethereum/go-ethereum/crypto"
	"os"
	"io"
	"log"
	"encoding/json"
	"sync"
	"strconv"
	 // "reflect"

	"time"
  // "context"
  "crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"crypto/rand"
  "github.com/ethereum/go-ethereum/common"
	"math/big"
   import_contracts "./contracts"
)

var buffer_size int
var noOfFileChunks int
var keyLength int

const waitingTime=20  //max time to complete a stage
const hashFunctionOutputBitSize=32 //hash function's output size(32 bytes in this case)

var contractAddress common.Address

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
		channel_SenToRec:=make(chan []byte)				//sender sends the byterarray(of Id,Keycommit,EncryptedOutputOfGates)

		f, err := os.OpenFile("text2	.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
				log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		wrt := io.MultiWriter(os.Stdout, f)
		log.SetOutput(wrt)


	 	if(len(os.Args)<2){
			buffer_size=32
			noOfFileChunks=8
			keyLength=buffer_size
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
			keyLength=buffer_size
		}

		//sender address generation
		authSender,_ := AuthAndAddressGeneration()
		// fmt.Println("\n->Sender\nSender Address :",senderAddress)

		//receiver address generation
		authReceiver,_ := AuthAndAddressGeneration()
		// fmt.Println("\n->Receiver\nReceiver Address :",receiverAddress)

		//blockchain creation
		alloc := make(core.GenesisAlloc)
		alloc[authSender.From] = core.GenesisAccount{Balance: big.NewInt(11000000000000000)}  //1.1 ether
		alloc[authReceiver.From] = core.GenesisAccount{Balance: big.NewInt(250000000000000000)} //2.5 ether
		//client
		client := backends.NewSimulatedBackend(alloc,60000000000)

		//go routines
		// ctx := context.Background()
		var wg sync.WaitGroup
		wg.Add(2)
		//go routines
		go Sender(authSender,client,channel_SenToRec,&wg)
		go Receiver_malicious_HashFileMismatch(authReceiver,client,channel_SenToRec,&wg)

		wg.Wait()

}

//Sender function
//The functoin executes honest sender's role in fairswap
func Sender(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup){

	//reading from a file "a.txt" and storing bytearray of 32 bytes in FileChunks of size equal to noOfFileChunks
	sender_round1_computingTime_calculatingVariables_start := time.Now()
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
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round1_computingTime_calculatingVariables_time : %v \n", sender_round1_computingTime_calculatingVariables_stop.Sub(sender_round1_computingTime_calculatingVariables_start))


	/*Deploy contract*/
	var err error
	var tx *types.Transaction
	sender_round1_commitTime_deployingContract_start := time.Now()
	contractAddress,tx,_, err=import_contracts.DeployProtocol2(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.HashedFileChunks,
		SenToConObject.HashedEncryptedFileChunks,
	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()
	sender_round1_commitTime_deployingContract_stop := time.Now()
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round1_commitTime_deployingContract_time : %v \n", sender_round1_commitTime_deployingContract_stop.Sub(sender_round1_commitTime_deployingContract_start))
	log.SetPrefix("\nNaive Protocol - Gas Cost ")
	log.Println("Sender Round 1 : ",tx.Gas())


	sender_round1_communicationTime_sendingEncryptedGateOutputs_start := time.Now()

	SenToRecObject.EncryptedFileChunks=encryptedFileChunks

	//sending the encoded file chunks to the receiver through the channel
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}
	channel_SenToRec<-byteArrayFromSenToRec

	sender_round1_communicationTime_sendingEncryptedGateOutputs_stop := time.Now()
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round1_communicationTime_sendingEncryptedGateOutputs_time : %v \n", sender_round1_communicationTime_sendingEncryptedGateOutputs_stop.Sub(sender_round1_communicationTime_sendingEncryptedGateOutputs_start))

	sender_round2_waitingTime_forReceiverToAccept_start := time.Now()

	//creating a contract instance
	instance,err:= import_contracts.NewProtocol2(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}


	//waiting for the receiver to accept
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
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round2_waitingTime_forReceiverToAccept_time : %v \n", sender_round2_waitingTime_forReceiverToAccept_stop.Sub(sender_round2_waitingTime_forReceiverToAccept_start))


	if(!receiverAcceptedStatus){
		//Receiver has not accepted within time
			}else{

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
							client.Commit()

							sender_round3_commitTime_revealKey_stop := time.Now()
							log.SetPrefix("\nNaive Protocol - Time ")
							log.Printf("sender_round3_commitTime_revealKey_time : %v \n", sender_round3_commitTime_revealKey_stop.Sub(sender_round3_commitTime_revealKey_start))
							log.SetPrefix("\nNaive Protocol - Gas Cost ")
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
									log.SetPrefix("\nNaive Protocol - Time ")
									log.Printf("sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_time : %v \n", sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_stop.Sub(sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_start))

									//Phase
									Phase,err:=instance.Phase(nil)
									if err!=nil{
										log.Fatal(err)
									}

									if(Phase==3){
									 	//sender getting the ether since the receiver has not finished the protocol

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

										tx,err=instance.SenderGetEther(&bind.TransactOpts{
												From:authSender.From,
												Signer:authSender.Signer,
											})
										if err!=nil{
													log.Fatal(err)
										}
										client.Commit()

										log.SetPrefix("\nNaive Protocol - Gas Cost ")
										log.Println("Sender Round 4 ",tx.Gas())
										sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop:=time.Now()
										log.SetPrefix("\nNaive Protocol - Time ")
										log.Printf("sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_time : %v \n", sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop.Sub(sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start))
							}
				}
	wg.Done()
}


//Receiver function

func Receiver(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup){

		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start := time.Now()

		//getting the encoded filechunks array from sender
		byteArrayFromSenToRec:=<-channel_SenToRec

		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop := time.Now()
		log.SetPrefix("\nNaive Protocol - Time ")
		log.Printf("receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_time : %v \n", receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop.Sub(receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start))

		receiver_round1_computingTime_merkleRootCheck_start := time.Now()

		var Receiver_SenToRecObj SenderToReceiverStruct

		err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
		if err != nil {
				log.Fatal(err)
		}

		instance,err := import_contracts.NewProtocol2(contractAddress,client)
		if err != nil {
				log.Fatal(err)
		}

	 _contractNoOfFileChunks,err := instance.NoOfFileChunks(nil)
	 if err!=nil{
		 log.Fatal(err)
	 }

	 boolHashedEncryptedFileCheck:=true

	 //checking whether encoded file chunks got from the sender matches with those got from sender
	 if(_contractNoOfFileChunks.Cmp(big.NewInt(int64(len(Receiver_SenToRecObj.EncryptedFileChunks))))==0){
		var i int64
		for i=0;i<int64(noOfFileChunks);i++ {
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
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("receiver_round1_computingTime_merkleRootCheck_time : %v \n", receiver_round1_computingTime_merkleRootCheck_stop.Sub(receiver_round1_computingTime_merkleRootCheck_start))


	if(boolHashedEncryptedFileCheck){

				receiver_round2_commitTime_Accept_start := time.Now()

				value:=big.NewInt(100)

				tx,err:=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
				// GasLimit:gasLimit,
			})
			if err != nil {
					log.Fatal(err)
			}
			client.Commit()

			receiver_round2_commitTime_Accept_stop := time.Now()
			log.SetPrefix("\nNaive Protocol - Time ")
			log.Printf("receiver_round2_commitTime_Accept_time : %v \n", receiver_round2_commitTime_Accept_stop.Sub(receiver_round2_commitTime_Accept_start))
			log.SetPrefix("\nNaive Protocol - Gas Cost ")
			log.Println("Receiver Round 2 : ",tx.Gas())


			//waiting for the sender to reveal key
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
				if(Phase==4){
					senderKeyRevealedStatus=true
					break;
				}
				time.Sleep(time.Second *1)
			}
			receiver_round3_waitingTime_forSenderToRevealKey_stop := time.Now()
			log.SetPrefix("\nNaive Protocol - Time ")
			log.Printf("receiver_round3_waitingTime_forSenderToRevealKey_time : %v \n", receiver_round3_waitingTime_forSenderToRevealKey_stop.Sub(receiver_round3_waitingTime_forSenderToRevealKey_start))



			if(!senderKeyRevealedStatus){
						//sender has not revealed the wrong key
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

					 	receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start := time.Now()
						tx,err=instance.ReceiverGetEther(&bind.TransactOpts{
								From:authReceiver.From,
								Signer:authReceiver.Signer,
								Value: nil,
								// GasLimit:gasLimit,
							})
						if err!=nil{
									log.Fatal(err)
						}
						client.Commit()
						receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop := time.Now()
						log.SetPrefix("\nNaive Protocol - Gas Cost ")
						log.Println("Receiver Round 3 (Sender Not Revealed Key) : ",tx.Cost())
						log.SetPrefix("\nNaive Protocol - Time ")
						log.Printf("receiver_round3_commitTime_senderKeyNotRevealedGettingEther_time : %v \n", receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop.Sub(receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start))


			} else{

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
						//malicious reciever
						receiver_round4_computingTime_extraction_stop := time.Now()
						log.SetPrefix("\nNaive Protocol - Time ")
						log.Printf("receiver_round4_computingTime_extraction_time : %v \n", receiver_round4_computingTime_extraction_stop.Sub(receiver_round4_computingTime_extraction_start))

						//index nil denotes no complain
						if(!isComplain){
							log.SetPrefix("\nNaive Protocol - Gas Cost ")
								log.Println("Receiver Round 4 : ",0)
						} else{

								//complain function being called
								receiver_round4_commitTime_complain_start := time.Now()
								var complaintChunk []byte
								for i:=0;i<buffer_size;i++{
									  complaintChunk=append(complaintChunk,Receiver_SenToRecObj.EncryptedFileChunks[index][i])
								}

								tx,err=instance.Complain(&bind.TransactOpts{
									From:authReceiver.From,
									Signer:authReceiver.Signer,
									Value:nil,
									// GasLimit:gasLimit,
								},complaintChunk,big.NewInt(int64(index)))
								if err!=nil{
									log.Fatal(err)
								}
								client.Commit()

								receiver_round4_commitTime_complain_stop := time.Now()
								log.SetPrefix("\nNaive Protocol - Time ")
								log.Printf("receiver_round4_commitTime_complain_time : %v \n", receiver_round4_commitTime_complain_stop.Sub(receiver_round4_commitTime_complain_start))
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
        // handle error here
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
	object.Price=*big.NewInt(80)     //800 wei
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

		// var hashedArray [][hashFunctionOutputBitSize]byte
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

//func createMaliciousEncryptedFileChunk(encryptedFileChunks [noOfFileChunks][buffer_size]byte)




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

func Sender_malicious_keyNotReveal(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup){

	//reading from a file "a.txt" and storing bytearray of 32 bytes in FileChunks of size equal to noOfFileChunks
	sender_round1_computingTime_calculatingVariables_start := time.Now()
	FileChunks:=FileRead()

	//objects
	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct


	setPrice(&SenToConObject)
	key:=keyGenerate()
	// maliciouskey:=keyGenerate()
	HashedFileChunks:=HashChunks(FileChunks)
	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedFileChunks=append(SenToConObject.HashedFileChunks,HashedFileChunks[i])
	}

 	encryptedFileChunks := EncryptFileChunks(FileChunks,key)
  // maliciousEncryptedFileChunks := MaliciousEncryptedFileChunks(FileChunks,key,maliciouskey)

	//fmt.Println("sender encrypt file chunks",encryptedFileChunks)

	HashedEncryptedFileChunks:=HashChunks(encryptedFileChunks)
	// maliciousHashIndexExchangeFileChunks := createMaliciousHashIndexExchangeFileChunk(HashedEncryptedFileChunks)
	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedEncryptedFileChunks=append(SenToConObject.HashedEncryptedFileChunks,HashedEncryptedFileChunks[i])
	}

	sender_round1_computingTime_calculatingVariables_stop:=time.Now()
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round1_computingTime_calculatingVariables_time : %v \n", sender_round1_computingTime_calculatingVariables_stop.Sub(sender_round1_computingTime_calculatingVariables_start))


	/*Deploy contract*/
	var err error
	var tx *types.Transaction
	sender_round1_commitTime_deployingContract_start := time.Now()
	contractAddress,tx,_, err=import_contracts.DeployProtocol2(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.HashedFileChunks,
		SenToConObject.HashedEncryptedFileChunks,
	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()
	sender_round1_commitTime_deployingContract_stop := time.Now()
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round1_commitTime_deployingContract_time : %v \n", sender_round1_commitTime_deployingContract_stop.Sub(sender_round1_commitTime_deployingContract_start))
	log.SetPrefix("\nNaive Protocol - Gas Cost ")
	log.Println("Sender Round 1 : ",tx.Gas())


	sender_round1_communicationTime_sendingEncryptedGateOutputs_start := time.Now()
	/*SenToRecObject initialize*/
	SenToRecObject.EncryptedFileChunks=encryptedFileChunks

	//sending the encodedGateOutputs,keycommitment and receiver entry key to the receiver through a byte array(got after applying json marshal to the SenToRecObject)
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}
	channel_SenToRec<-byteArrayFromSenToRec

	sender_round1_communicationTime_sendingEncryptedGateOutputs_stop := time.Now()
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round1_communicationTime_sendingEncryptedGateOutputs_time : %v \n", sender_round1_communicationTime_sendingEncryptedGateOutputs_stop.Sub(sender_round1_communicationTime_sendingEncryptedGateOutputs_start))
	//creating a contract instance

	sender_round2_waitingTime_forReceiverToAccept_start := time.Now()

	instance,err:= import_contracts.NewProtocol2(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}


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
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round2_waitingTime_forReceiverToAccept_time : %v \n", sender_round2_waitingTime_forReceiverToAccept_stop.Sub(sender_round2_waitingTime_forReceiverToAccept_start))


	if(!receiverAcceptedStatus){
		//Receiver has not accepted within time
			}else{

			}
	wg.Done()
}

func Sender_malicious_FileChunkWrongEnc(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup){

	//reading from a file "a.txt" and storing bytearray of 32 bytes in FileChunks of size equal to noOfFileChunks
	sender_round1_computingTime_calculatingVariables_start := time.Now()
	FileChunks:=FileRead()

	//objects
	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct


	setPrice(&SenToConObject)
	key:=keyGenerate()
	maliciouskey:=keyGenerate()
	HashedFileChunks:=HashChunks(FileChunks)
	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedFileChunks=append(SenToConObject.HashedFileChunks,HashedFileChunks[i])
	}

 	// encryptedFileChunks := EncryptFileChunks(FileChunks,key)
  maliciousEncryptedFileChunks := MaliciousEncryptedFileChunks(FileChunks,key,maliciouskey)

	//fmt.Println("sender encrypt file chunks",encryptedFileChunks)

	HashedEncryptedFileChunks:=HashChunks(maliciousEncryptedFileChunks)

	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedEncryptedFileChunks=append(SenToConObject.HashedEncryptedFileChunks,HashedEncryptedFileChunks[i])
	}

	sender_round1_computingTime_calculatingVariables_stop:=time.Now()
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round1_computingTime_calculatingVariables_time : %v \n", sender_round1_computingTime_calculatingVariables_stop.Sub(sender_round1_computingTime_calculatingVariables_start))


	/*Deploy contract*/
	var err error
	var tx *types.Transaction
	sender_round1_commitTime_deployingContract_start := time.Now()
	contractAddress,tx,_, err=import_contracts.DeployProtocol2(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.HashedFileChunks,
		SenToConObject.HashedEncryptedFileChunks,
	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()
	sender_round1_commitTime_deployingContract_stop := time.Now()
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round1_commitTime_deployingContract_time : %v \n", sender_round1_commitTime_deployingContract_stop.Sub(sender_round1_commitTime_deployingContract_start))
	log.SetPrefix("\nNaive Protocol - Gas Cost ")
	log.Println("Sender Round 1 : ",tx.Gas())


	sender_round1_communicationTime_sendingEncryptedGateOutputs_start := time.Now()
	/*SenToRecObject initialize*/
	SenToRecObject.EncryptedFileChunks=maliciousEncryptedFileChunks

	//sending the encodedGateOutputs,keycommitment and receiver entry key to the receiver through a byte array(got after applying json marshal to the SenToRecObject)
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}
	channel_SenToRec<-byteArrayFromSenToRec

	sender_round1_communicationTime_sendingEncryptedGateOutputs_stop := time.Now()
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round1_communicationTime_sendingEncryptedGateOutputs_time : %v \n", sender_round1_communicationTime_sendingEncryptedGateOutputs_stop.Sub(sender_round1_communicationTime_sendingEncryptedGateOutputs_start))
	//creating a contract instance

	sender_round2_waitingTime_forReceiverToAccept_start := time.Now()

	instance,err:= import_contracts.NewProtocol2(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}


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
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("sender_round2_waitingTime_forReceiverToAccept_time : %v \n", sender_round2_waitingTime_forReceiverToAccept_stop.Sub(sender_round2_waitingTime_forReceiverToAccept_start))


	if(!receiverAcceptedStatus){
		//Receiver has not accepted within time
			}else{

							sender_round3_commitTime_revealKey_start := time.Now()
							var keyToCon []byte
							for i:=0;i<keyLength;i++{
								keyToCon=append(keyToCon,key[i])
							}

							tx,err=instance.KeyReveal(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},keyToCon)
							if err!=nil{
								log.Fatal(err)
							}
							client.Commit()

							sender_round3_commitTime_revealKey_stop := time.Now()
							log.SetPrefix("\nNaive Protocol - Time ")
							log.Printf("sender_round3_commitTime_revealKey_time : %v \n", sender_round3_commitTime_revealKey_stop.Sub(sender_round3_commitTime_revealKey_start))
							//client.Commit()
							log.SetPrefix("\nNaive Protocol - Gas Cost ")
							log.Println("Sender Round 3 : ",tx.Gas())
							/*fmt.Println("\n->Sender\nReceiver have accepted within the timeout, Key Revealed")
							channel_SenToMain<-"I have not revealed the key"*/

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
			log.Printf("sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_time : %v \n", sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_stop.Sub(sender_round4_waitingTime_forReceiverToApproveOrProduceComplain_start))

			// gasLimit := uint64(2100000000)

			//the sender getting the ether
			Phase,err:=instance.Phase(nil)
			if err!=nil{
				log.Fatal(err)
			}
		//	fmt.Println(Phase)

			if(Phase==3){
			 //time.Sleep(time.Second *60)
			 // count:=0

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

			tx,err=instance.SenderGetEther(&bind.TransactOpts{
					From:authSender.From,
					Signer:authSender.Signer,
					// Value: big.NewInt(10000000000),
					// GasLimit:gasLimit,
				})
			if err!=nil{
						log.Fatal(err)
			}
			client.Commit()

			log.SetPrefix("\nNaive Protocol - Gas Cost ")
			log.Println("Sender Round 4 ",tx.Gas())
			sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop:=time.Now()
			log.SetPrefix("\nNaive Protocol - Time ")
			log.Printf("sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_time : %v \n", sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_stop.Sub(sender_round4_commitTime_getEtherOnSuccesfulCompletionOfProtocol_start))

			}
		}
	wg.Done()
}


func Receiver_malicious_HashEncFileMismatch(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup){

		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start := time.Now()

		byteArrayFromSenToRec:=<-channel_SenToRec

		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop := time.Now()
		log.SetPrefix("\nNaive Protocol - Time ")
		log.Printf("receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_time : %v \n", receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop.Sub(receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start))

		receiver_round1_computingTime_merkleRootCheck_start := time.Now()

		var Receiver_SenToRecObj SenderToReceiverStruct

		err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
		if err != nil {
				log.Fatal(err)
		}

		instance,err := import_contracts.NewProtocol2(contractAddress,client)
		if err != nil {
				log.Fatal(err)
		}

	 _contractNoOfFileChunks,err := instance.NoOfFileChunks(nil)
	 if err!=nil{
		 log.Fatal(err)
	 }

	 boolHashedEncryptedFileCheck:=true

	 //if(big.NewInt(int64(len(Receiver_SenToRecObj.EncryptedFileChunks)))==_contractNoOfFileChunks)
	 if(_contractNoOfFileChunks.Cmp(big.NewInt(int64(len(Receiver_SenToRecObj.EncryptedFileChunks))))==0){
		var i int64
		for i=0;i<int64(noOfFileChunks);i++ {
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
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("receiver_round1_computingTime_merkleRootCheck_time : %v \n", receiver_round1_computingTime_merkleRootCheck_stop.Sub(receiver_round1_computingTime_merkleRootCheck_start))

	if(boolHashedEncryptedFileCheck){

				receiver_round2_commitTime_Accept_start := time.Now()

				value:=big.NewInt(100)

				tx,err:=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
				// GasLimit:gasLimit,
			})
			if err != nil {
					log.Fatal(err)
			}
			client.Commit()

			receiver_round2_commitTime_Accept_stop := time.Now()
			log.SetPrefix("\nNaive Protocol - Time ")
			log.Printf("receiver_round2_commitTime_Accept_time : %v \n", receiver_round2_commitTime_Accept_stop.Sub(receiver_round2_commitTime_Accept_start))
			log.SetPrefix("\nNaive Protocol - Gas Cost ")
			log.Println("Receiver Round 2 : ",tx.Gas())
			//Sending the finished accepted stage message to the main function


			//The contract is in accepted stage
			//The sender
			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
			//case 2 : The sender doesn't reveal the key within time. senderKeyRevealedStatus remains false after the for loop. Receiver later calls the contract function
			//			   ReceiverGetEther to get the ether.
			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and senderKeyRevealedStatus is true after the for loop.
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
				if(Phase==4){
					senderKeyRevealedStatus=true
					break;
				}
				time.Sleep(time.Second *1)
			}
			receiver_round3_waitingTime_forSenderToRevealKey_stop := time.Now()
			log.SetPrefix("\nNaive Protocol - Time ")
			log.Printf("receiver_round3_waitingTime_forSenderToRevealKey_time : %v \n", receiver_round3_waitingTime_forSenderToRevealKey_stop.Sub(receiver_round3_waitingTime_forSenderToRevealKey_start))



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

			 	receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start := time.Now()
				tx,err=instance.ReceiverGetEther(&bind.TransactOpts{
						From:authReceiver.From,
						Signer:authReceiver.Signer,
						Value: nil,
						// GasLimit:gasLimit,
					})
				if err!=nil{
							log.Fatal(err)
				}
				client.Commit()
				receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop := time.Now()
				log.SetPrefix("\nNaive Protocol - Gas Cost ")
				log.Println("Receiver Round 3 (Sender Not Revealed Key) : ",tx.Cost())
				log.SetPrefix("\nNaive Protocol - Time ")
				log.Printf("receiver_round3_commitTime_senderKeyNotRevealedGettingEther_time : %v \n", receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop.Sub(receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start))


			} else{

				//Extracting function being called

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

				isComplain=true
				//malicious reciever
				receiver_round4_computingTime_extraction_stop := time.Now()
				log.SetPrefix("\nNaive Protocol - Time ")
				log.Printf("receiver_round4_computingTime_extraction_time : %v \n", receiver_round4_computingTime_extraction_stop.Sub(receiver_round4_computingTime_extraction_start))

				//index nil denotes no complain
				if(!isComplain){
					log.SetPrefix("\nNaive Protocol - Gas Cost ")
						log.Println("Receiver Round 4 : ",0)
				} else{

						//complain function being called
						receiver_round4_commitTime_complain_start := time.Now()
						var complaintChunk []byte
						index=4

						tx,err=instance.Complain(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value:nil,
							// GasLimit:gasLimit,
						},complaintChunk,big.NewInt(int64(index)))
						if err!=nil{
							log.Fatal(err)
						}
						client.Commit()

						receiver_round4_commitTime_complain_stop := time.Now()
						log.SetPrefix("\nNaive Protocol - Time ")
						log.Printf("receiver_round4_commitTime_complain_time : %v \n", receiver_round4_commitTime_complain_stop.Sub(receiver_round4_commitTime_complain_start))
						log.SetPrefix("\nNaive Protocol - Gas Cost ")
						log.Println("Receiver Round 4 : ",tx.Cost())

				}
			}
		}	else{
				// log.Println("\n->Receiver\n Hash encrypted file from the contract and the hash of encrypted file from the contract doesnt match")
		}
	wg.Done()
}


func Receiver_malicious_HashFileMismatch(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,wg *sync.WaitGroup){

		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start := time.Now()

		byteArrayFromSenToRec:=<-channel_SenToRec

		receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop := time.Now()
		log.SetPrefix("\nNaive Protocol - Time ")
		log.Printf("receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_time : %v \n", receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_stop.Sub(receiver_round1_waitingTime_forReceivingEncryptedGateOutputs_start))

		receiver_round1_computingTime_merkleRootCheck_start := time.Now()

		var Receiver_SenToRecObj SenderToReceiverStruct

		err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
		if err != nil {
				log.Fatal(err)
		}

		instance,err := import_contracts.NewProtocol2(contractAddress,client)
		if err != nil {
				log.Fatal(err)
		}

	 _contractNoOfFileChunks,err := instance.NoOfFileChunks(nil)
	 if err!=nil{
		 log.Fatal(err)
	 }

	 boolHashedEncryptedFileCheck:=true

	 //if(big.NewInt(int64(len(Receiver_SenToRecObj.EncryptedFileChunks)))==_contractNoOfFileChunks)
	 if(_contractNoOfFileChunks.Cmp(big.NewInt(int64(len(Receiver_SenToRecObj.EncryptedFileChunks))))==0){
		var i int64
		for i=0;i<int64(noOfFileChunks);i++ {
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
	log.SetPrefix("\nNaive Protocol - Time ")
	log.Printf("receiver_round1_computingTime_merkleRootCheck_time : %v \n", receiver_round1_computingTime_merkleRootCheck_stop.Sub(receiver_round1_computingTime_merkleRootCheck_start))

	if(boolHashedEncryptedFileCheck){

				receiver_round2_commitTime_Accept_start := time.Now()

				value:=big.NewInt(100)

				tx,err:=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
				// GasLimit:gasLimit,
			})
			if err != nil {
					log.Fatal(err)
			}
			client.Commit()

			receiver_round2_commitTime_Accept_stop := time.Now()
			log.SetPrefix("\nNaive Protocol - Time ")
			log.Printf("receiver_round2_commitTime_Accept_time : %v \n", receiver_round2_commitTime_Accept_stop.Sub(receiver_round2_commitTime_Accept_start))
			log.SetPrefix("\nNaive Protocol - Gas Cost ")
			log.Println("Receiver Round 2 : ",tx.Gas())
			//Sending the finished accepted stage message to the main function


			//The contract is in accepted stage
			//The sender
			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
			//case 2 : The sender doesn't reveal the key within time. senderKeyRevealedStatus remains false after the for loop. Receiver later calls the contract function
			//			   ReceiverGetEther to get the ether.
			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and senderKeyRevealedStatus is true after the for loop.
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
				if(Phase==4){
					senderKeyRevealedStatus=true
					break;
				}
				time.Sleep(time.Second *1)
			}
			receiver_round3_waitingTime_forSenderToRevealKey_stop := time.Now()
			log.SetPrefix("\nNaive Protocol - Time ")
			log.Printf("receiver_round3_waitingTime_forSenderToRevealKey_time : %v \n", receiver_round3_waitingTime_forSenderToRevealKey_stop.Sub(receiver_round3_waitingTime_forSenderToRevealKey_start))



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

			 	receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start := time.Now()
				tx,err=instance.ReceiverGetEther(&bind.TransactOpts{
						From:authReceiver.From,
						Signer:authReceiver.Signer,
						Value: nil,
						// GasLimit:gasLimit,
					})
				if err!=nil{
							log.Fatal(err)
				}
				client.Commit()
				receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop := time.Now()
				log.SetPrefix("\nNaive Protocol - Gas Cost ")
				log.Println("Receiver Round 3 (Sender Not Revealed Key) : ",tx.Cost())
				log.SetPrefix("\nNaive Protocol - Time ")
				log.Printf("receiver_round3_commitTime_senderKeyNotRevealedGettingEther_time : %v \n", receiver_round3_commitTime_senderKeyNotRevealedGettingEther_stop.Sub(receiver_round3_commitTime_senderKeyNotRevealedGettingEther_start))


			} else{

				//Extracting function being called

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

				isComplain=true
				//malicious reciever
				receiver_round4_computingTime_extraction_stop := time.Now()
				log.SetPrefix("\nNaive Protocol - Time ")
				log.Printf("receiver_round4_computingTime_extraction_time : %v \n", receiver_round4_computingTime_extraction_stop.Sub(receiver_round4_computingTime_extraction_start))

				//index nil denotes no complain
				if(!isComplain){
					log.SetPrefix("\nNaive Protocol - Gas Cost ")
						log.Println("Receiver Round 4 : ",0)
				} else{

						//complain function being called
						receiver_round4_commitTime_complain_start := time.Now()

						var complaintChunk []byte
						index=4
						for i:=0;i<buffer_size;i++{
							  complaintChunk=append(complaintChunk,Receiver_SenToRecObj.EncryptedFileChunks[index][i])
						}

						tx,err=instance.Complain(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value:nil,
							// GasLimit:gasLimit,
						},complaintChunk,big.NewInt(int64(index)))
						if err!=nil{
							log.Fatal(err)
						}
						client.Commit()

						receiver_round4_commitTime_complain_stop := time.Now()
						log.SetPrefix("\nNaive Protocol - Time ")
						log.Printf("receiver_round4_commitTime_complain_time : %v \n", receiver_round4_commitTime_complain_stop.Sub(receiver_round4_commitTime_complain_start))
						log.SetPrefix("\nNaive Protocol - Gas Cost ")
						log.Println("Receiver Round 4 : ",tx.Cost())

				}
			}
		}	else{
				// log.Println("\n->Receiver\n Hash encrypted file from the contract and the hash of encrypted file from the contract doesnt match")
		}
	wg.Done()
}
