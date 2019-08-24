package main

import(
	"fmt"
	kec "github.com/ethereum/go-ethereum/crypto"
	"os"
	"io"
	"log"
	"encoding/json"
	"reflect"

	"time"
  "context"
  "crypto/ecdsa"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
// "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/core"
	"crypto/rand"
  "github.com/ethereum/go-ethereum/common"
	"math/big"
   import_contracts "./contracts"
)

const buffer_size=32//byte array size (size of each byte array sent from sender to the receiver)
const noOfFileChunks=8//total number of input gates
const waitingTime=15  //max time to complete a stage
const keyLength=32//key size
const hashFunctionOutputBitSize=32 //hash function's output size(32 bytes in this case)

var contractAddress common.Address

type SenderToReceiverStruct struct {
		EncryptedFileChunks [noOfFileChunks][buffer_size]byte
}

type SenderToContractStruct struct{
	Price big.Int
	HashedFileChunks [][hashFunctionOutputBitSize]byte
	HashedEncryptedFileChunks [][hashFunctionOutputBitSize]byte
}

func main(){

			//channels
		channel_SenToRec:=make(chan []byte)				//sender sends the byterarray(of Id,Keycommit,EncryptedOutputOfGates)
		channel_RecToMainIni:=make(chan string)		//Receiver sends a notificaion once he has finished his entire tasks
		channel_RecToMainRev:=make(chan string)
		channel_SenToMain:=make(chan string)



		var ctx context.Context
		clientNode1,err := ethclient.Dial("/home/amudhan/workspace2/node1/geth.ipc")
		if err != nil {
					log.Fatal(err)
		}
		fmt.Println(reflect.TypeOf(clientNode1))

		accountAddressNode1:=common.HexToAddress("bf41fd7f78de80a85bdce83b154108e4de19f8c0")
		accountNode1orSenderInitialBalance, err := clientNode1.BalanceAt(context.Background(),accountAddressNode1, nil)
		if err != nil{
				log.Fatal(err)
		}
		fmt.Println("-->Main\naccountNode1orSenderInitialBalance : ",accountNode1orSenderInitialBalance,"\n Account address : ",accountAddressNode1)

		var node1Reader io.Reader
		node1Reader, err = os.Open("/home/amudhan/workspace2/node1/keystore/UTC--2019-08-08T18-29-45.601668268Z--bf41fd7f78de80a85bdce83b154108e4de19f8c0")
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


		//var ctx context.Context
		clientNode2,err := ethclient.Dial("/home/amudhan/workspace2/node2/geth.ipc")
		if err != nil {
					log.Fatal(err)
		}
		fmt.Println(reflect.TypeOf(clientNode2))

		accountAddressNode2:=common.HexToAddress("850a23231bb1c329fc472260e6366603bfe9c3fc")
		accountNode2orReceiverBalance, err := clientNode2.BalanceAt(context.Background(),accountAddressNode2, nil)
		if err != nil{
				log.Fatal(err)
		}
		fmt.Println("\naccountNode2BalanceInitial : ",accountNode2orReceiverBalance,"\n AccountNode2 address : ",accountAddressNode2)

		var node2Reader io.Reader
		node2Reader, err = os.Open("/home/amudhan/workspace/node2/keystore/UTC--2019-08-08T18-30-20.635313092Z--850a23231bb1c329fc472260e6366603bfe9c3fc")
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
		go Sender(authSender,clientNode1,channel_SenToRec,channel_SenToMain)
		go Receiver(authReceiver,clientNode2,channel_SenToRec,channel_RecToMainIni,channel_RecToMainRev)


		//channels... Parties communicate to the main function
		ReceicverInittilaizationMessage:=<-channel_RecToMainIni
		fmt.Println("\n->Main fn\n",ReceicverInittilaizationMessage)

		SenderMessageToMain:=<-channel_SenToMain
		fmt.Println("\n->Main fn\n",SenderMessageToMain)

		ReceiverRevealMessage:=<-channel_RecToMainRev
		fmt.Println("\n->Main fn\n",ReceiverRevealMessage)

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

//Sender function
//The functoin executes honest sender's role in fairswap
//Input parameters:
//authsender(*bind.TransactOpts) - the sender in the blockchain
//client()*ethclient.Client) - blockchain(here it is a backend blockchain)
//channel_SenToRec(chan []byte) - to send the byte array to the receiver
//channel_SenToMain(chan string) - to communicate with the main function
func Sender(authSender *bind.TransactOpts,client *ethclient.Client,channel_SenToRec chan []byte,channel_SenToMain chan string){

	//reading from a file "a.txt" and storing bytearray of 32 bytes in FileChunks of size equal to noOfFileChunks
	fmt.Println("\n->Sender\nBegins")
	FileChunks:=FileRead()
	fmt.Println("->Sender\nFileChunks:",FileChunks)

	//objects
	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct


	setPrice(&SenToConObject)
	key:=keyGenerate()
	//maliciouskey:=keyGenerate()
	HashedFileChunks:=HashChunks(FileChunks)
	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedFileChunks=append(SenToConObject.HashedFileChunks,HashedFileChunks[i])
	}

	encryptedFileChunks := EncryptFileChunks(FileChunks,key)
	//maliciousEncryptedFileChunks := MaliciousEncryptedFileChunks(FileChunks,key,maliciouskey)

	//fmt.Println("sender encrypt file chunks",encryptedFileChunks)

	HashedEncryptedFileChunks:=HashChunks(encryptedFileChunks)
	//maliciousHashIndexExchangeFileChunks := createMaliciousHashIndexExchangeFileChunk(HashedEncryptedFileChunks)
	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedEncryptedFileChunks=append(SenToConObject.HashedEncryptedFileChunks,HashedEncryptedFileChunks[i])
	}


	/*Deploy contract*/
	var err error
	//authSender.Value=big.NewInt(1000000000000)
	contractAddress, _,_, err=import_contracts.DeployProtocol2(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.HashedFileChunks,
		SenToConObject.HashedEncryptedFileChunks,
	)
	if err!=nil{
		log.Fatal(err)
	}


	fmt.Println("\n->Sender \n Contract Deployed")

	/*SenToRecObject initialize*/
	SenToRecObject.EncryptedFileChunks=encryptedFileChunks


	//sending the encodedGateOutputs,keycommitment and receiver entry key to the receiver through a byte array(got after applying json marshal to the SenToRecObject)
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}
	channel_SenToRec<-byteArrayFromSenToRec
	fmt.Println("\n->Sender \nByteArray sent to Receiver")

	fmt.Println("\n->Sender\n Initialization Finished")

	time.Sleep(time.Second*6)
	//creating a contract instance
	instance,err:= import_contracts.NewProtocol2(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	//The contract is in initialized stage
	//Once the contract is initialized, the receiver has time equal to waitingTime to accept and pay ether to the contract.
	//the for loop checks for the change in stage.
	//case 1 : If the receiver doesn't accept within the time limit, RevealedBool remains false after the for loop. As a result,
	//				"Receiver have not accepted within the timeout, Key not Revealed" is sent to the main function and the receiver doesn't
	//        proceed with the further protocol.
	//case 2 : If the receiver pays less than the price,then the contract stage goes to finishedAndMaliciousReciever.The sender sends
	//				"Malicious Receiver-Dind't pay enough ether" to the main function  and finishes the protocol execution here. The ether is sent back
	//				to the receiver by the contract.
	//case 3: The receiver accepts within the time. The contract stage goes to accepted and the sender has to reveal the key. RevealedBool becomes
	//				true.


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
		channel_SenToMain<-"Either the reciever is not interested now or Sender might have exchanged the indices of hashedEncryptedFileChunks"
	}else{

			if(!RevealedBool){
				channel_SenToMain<-"Receiver has not accepted within the timeout, Key not Revealed"

			}else{

							//Sender Reveals the key
							var keyToCon []byte
							for i:=0;i<keyLength;i++{
								keyToCon=append(keyToCon,key[i])
							}

							_,err=instance.KeyReveal(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},keyToCon)
							if err!=nil{
								log.Fatal(err)
							}

							/*fmt.Println("\n->Sender\nReceiver have accepted within the timeout, Key Revealed")
							channel_SenToMain<-"I have not revealed the key"*/
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
						fmt.Println("\n->Sender\nIt seems I have wrongly encrypted filechunks")
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
		//	fmt.Println(Phase)
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


//Receiver function
//The functoin executes honest receiver's role in fairswap
//Input parameters:
//authReceiver(*bind.TransactOpts) - the receiver in the blockchain
//client()*ethclient.Client) - blockchain(here it is a backend blockchain)
//channel_SenToRec(chan []byte) - to receiver the byte array from the sender
//channel_RecToMainIni(chan string) - to notify the main function for the completion of initialization phase
//channel_RecToMainRev(chan string) - to notify the main function for the completion of reveal phase
func Receiver(authReceiver *bind.TransactOpts,client *ethclient.Client,channel_SenToRec chan []byte,channel_RecToMainIni chan string,channel_RecToMainRev chan string){


		byteArrayFromSenToRec:=<-channel_SenToRec
		var Receiver_SenToRecObj SenderToReceiverStruct

		err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
		if err != nil {
				log.Fatal(err)
		}

		time.Sleep(time.Second*6)

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
		for i=0;i<noOfFileChunks;i++ {
			value,err:=instance.HashedEncryptedFileChunks(nil,big.NewInt(i))
			if err!=nil{
				log.Fatal(err)
			}
			if(value!=Hash(Receiver_SenToRecObj.EncryptedFileChunks[i])){
				boolHashedEncryptedFileCheck=false
			}
		}
	}

		if(boolHashedEncryptedFileCheck){


				gasLimit := uint64(230000000)


				value:=big.NewInt(5000)

			_	,err=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
				// GasLimit:gasLimit,
			})
			if err != nil {
					log.Fatal(err)
			}
			client.Commit()


			fmt.Println("\n->Receiver\n Finished Accepting and Initialization phase ends")
			//Sending the finished accepted stage message to the main function
			channel_RecToMainIni<-"Accepted stage got over"


			//The contract is in accepted stage
			//The sender
			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
			//case 2 : The sender doesn't reveal the key within time. AcceptedBool remains false after the for loop. Receiver later calls the contract function
			//			   ReceiverGetEther to get the ether.
			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and AcceptedBool is true after the for loop.
			AcceptedBool:=false
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
					AcceptedBool=true
					break;
				}
				time.Sleep(time.Second *1)

			}



			if(!AcceptedBool){

					fmt.Println("\n->Receiver\nSender has not revealed the key...ReceiverGetEther function called")
				//	gasLimit := uint64(210000000000000)

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


					//channel_RecToMainIni<-"Sender hasn't revealed the key within the Timeout"
					channel_RecToMainRev<-"Sender hasn't revealed the key within the Timeout"

			} else{

				//Extracting function being called
				var _contractHashedFileChunks [noOfFileChunks][hashFunctionOutputBitSize]byte
				var i int64
				for i=0;i<noOfFileChunks;i++{
						value,err:=instance.HashedFileChunks(nil,big.NewInt(i))
						if err!=nil{
								log.Fatal(err)
						}
						_contractHashedFileChunks[i]=value
				}

				var _ReceiverKey [keyLength]byte
				for i:=0;i<keyLength;i++{
					_ReceiverKey[i]=contract_Key[i]
				}

				isComplain,index:=Extract(_contractHashedFileChunks,Receiver_SenToRecObj.EncryptedFileChunks,_ReceiverKey)
				//malicious reciever

				//index nil denotes no complain
				if(!isComplain){
					fmt.Println("\n->Receiver\nNo complain")
					channel_RecToMainRev<-"Sender has correctly calculated the encryption"
				} else{

						//complain function being called
						var complaintChunk []byte
						for i:=0;i<buffer_size;i++{
							  complaintChunk=append(complaintChunk,Receiver_SenToRecObj.EncryptedFileChunks[index][i])
						}

						_,err=instance.Complain(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value:nil,
							GasLimit:gasLimit,
						},complaintChunk,big.NewInt(int64(index)))
						if err!=nil{
							log.Fatal(err)
						}


						Phase,err:=instance.Phase(nil)
						if err!=nil{
							log.Fatal(err)
						}
						if(Phase==5){
							fmt.Println("\n->Receiver\nSender Has wrongly calulated the encrypted inputs")
							channel_RecToMainRev<-"sender-Wrong encrypted inputs"

						}
				}
			}
			channel_RecToMainRev<-"Reveal Phase Finished"
		}	else{

				fmt.Println("\n-> Receiver \n HashedEncryptedFileGotFromSender and HashedEncryptedFileFromContract dindn't match")
				channel_RecToMainIni<-"HashedEncryptedFileGotFromSender and HashedEncryptedFileFromContract dindn't match"
				channel_RecToMainRev<-"HashedEncryptedFileGotFromSender and HashedEncryptedFileFromContract dindn't match "
		}
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


func Hash(_toBeHashed [buffer_size]byte) ([32]byte){

	var toBeHashed []byte
	var Hashed [hashFunctionOutputBitSize]byte

	for j:=0;j<buffer_size;j++{
		toBeHashed=append(toBeHashed,_toBeHashed[j])
	}

	HashedAfter:=kec.Keccak256(toBeHashed)

	for j:=0;j<hashFunctionOutputBitSize;j++{
			Hashed[j]=HashedAfter[j]
	}

	return Hashed

}

func keyGenerate() ([keyLength]byte){

		keyGen := make([]byte, keyLength)
		var key [keyLength]byte

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


func Encrypt(plainText [buffer_size]byte,key [keyLength]byte) ([buffer_size]byte) {

			var encryptedtext [buffer_size]byte
			for i:=0;i<buffer_size;i++{
				encryptedtext[i]=key[i]^plainText[i]  //xor operation
			}
			return encryptedtext
}

func setPrice(object *SenderToContractStruct){
	object.Price=*big.NewInt(80)     //800 wei
}


func Extract(_hashedFileChunks [noOfFileChunks][hashFunctionOutputBitSize]byte, _encryptedFileChunks [noOfFileChunks][buffer_size]byte, key [keyLength]byte)  (bool,int){

		for i:=0;i<noOfFileChunks;i++{
			if(_hashedFileChunks[i]!=Hash(Decrypt(_encryptedFileChunks[i],key))){
				return true,i
			}
		}
		return false,0
}

func Decrypt(encryptedtext [buffer_size]byte,key [keyLength]byte) ([buffer_size]byte){


			var plainText [buffer_size]byte
			for i:=0;i<buffer_size;i++{
				plainText[i]=key[i]^encryptedtext[i]
			}

			return plainText
}

func HashChunks(_tobeHashedArray [noOfFileChunks][buffer_size]byte) ([noOfFileChunks][hashFunctionOutputBitSize]byte){

		var hashedArray [noOfFileChunks][hashFunctionOutputBitSize]byte
		for i:=0;i<noOfFileChunks;i++{
			hashedArray[i]=Hash(_tobeHashedArray[i])
		}
		return hashedArray
}


func EncryptFileChunks(_fileChunks [noOfFileChunks][buffer_size]byte,key [keyLength]byte) ([noOfFileChunks][buffer_size]byte){

	var encryptedFileChunks [noOfFileChunks][buffer_size]byte
	for i:=0;i<noOfFileChunks;i++{
		encryptedFileChunks[i]=Encrypt(_fileChunks[i],key)
	}
	return encryptedFileChunks
}

//func createMaliciousEncryptedFileChunk(encryptedFileChunks [noOfFileChunks][buffer_size]byte)

func FileRead() ([noOfFileChunks][buffer_size]byte){
	var currentByte int64 = 0
	fileBuffer := make([]byte, buffer_size)
	file, err := os.Open("a.txt")
	if err != nil {
			log.Fatal(err)
	}
	var FileChunks [noOfFileChunks][buffer_size]byte
	count:=0
	for {
			n,err := file.ReadAt(fileBuffer, currentByte)
			currentByte += buffer_size
			fileBuffer=fileBuffer[:n]
			if count==noOfFileChunks{
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


func createMaliciousHashIndexExchangeFileChunk(HashedEncryptedFileChunks [noOfFileChunks][hashFunctionOutputBitSize]byte) ([noOfFileChunks][hashFunctionOutputBitSize]byte){
	var maliciousHashIndexExchangeFileChunks [noOfFileChunks][hashFunctionOutputBitSize]byte
	for i:=0;i<noOfFileChunks;i++{
			maliciousHashIndexExchangeFileChunks[i]=HashedEncryptedFileChunks[i]
	}
	maliciousHashIndexExchangeFileChunks[0]=HashedEncryptedFileChunks[2]
	maliciousHashIndexExchangeFileChunks[2]=HashedEncryptedFileChunks[0]

	return maliciousHashIndexExchangeFileChunks
}


func MaliciousEncryptedFileChunks(_fileChunks [noOfFileChunks][buffer_size]byte,key [keyLength]byte,maliciouskey [keyLength]byte) ([noOfFileChunks][buffer_size]byte){

	var maliciousEncryptedFileChunks [noOfFileChunks][buffer_size]byte
	for i:=0;i<noOfFileChunks;i++{
		maliciousEncryptedFileChunks[i]=Encrypt(_fileChunks[i],key)
	}
	maliciousEncryptedFileChunks[5]=Encrypt(_fileChunks[5],maliciouskey)
	return maliciousEncryptedFileChunks
}

/*func RecMaliciousEncryptedFileChunks(_fileChunks [noOfFileChunks][buffer_size]byte,key [keyLength]byte,maliciouskey [keyLength]byte) ([noOfFileChunks][buffer_size]byte){

	var maliciousEncryptedFileChunks [noOfFileChunks][buffer_size]byte
	for i:=0;i<noOfFileChunks;i++{
		maliciousEncryptedFileChunks[i]=Encrypt(_fileChunks[i],key)
	}
	maliciousEncryptedFileChunks[5]=Encrypt(_fileChunks[5],maliciouskey)
	return maliciousEncryptedFileChunks
}*/

func Sender_malicious_keyNotReveal(authSender *bind.TransactOpts,client *ethclient.Client,channel_SenToRec chan []byte,channel_SenToMain chan string){

	//reading from a file "a.txt" and storing bytearray of 32 bytes in FileChunks of size equal to noOfFileChunks
	fmt.Println("\n->Sender\nBegins")
	FileChunks:=FileRead()
	fmt.Println("->Sender\nFileChunks:",FileChunks)

	//objects
	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct


	setPrice(&SenToConObject)
	key:=keyGenerate()
	//maliciouskey:=keyGenerate()
	HashedFileChunks:=HashChunks(FileChunks)
	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedFileChunks=append(SenToConObject.HashedFileChunks,HashedFileChunks[i])
	}

	encryptedFileChunks := EncryptFileChunks(FileChunks,key)
	//maliciousEncryptedFileChunks := MaliciousEncryptedFileChunks(FileChunks,key,maliciouskey)

	//fmt.Println("sender encrypt file chunks",encryptedFileChunks)

	HashedEncryptedFileChunks:=HashChunks(encryptedFileChunks)
	//maliciousHashIndexExchangeFileChunks := createMaliciousHashIndexExchangeFileChunk(HashedEncryptedFileChunks)
	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedEncryptedFileChunks=append(SenToConObject.HashedEncryptedFileChunks,HashedEncryptedFileChunks[i])
	}


	/*Deploy contract*/
	var err error
	//authSender.Value=big.NewInt(1000000000000)
	contractAddress, _,_, err=import_contracts.DeployProtocol2(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.HashedFileChunks,
		SenToConObject.HashedEncryptedFileChunks,
	)
	if err!=nil{
		log.Fatal(err)
	}


	fmt.Println("\n->Sender \n Contract Deployed")

	/*SenToRecObject initialize*/
	SenToRecObject.EncryptedFileChunks=encryptedFileChunks


	//sending the encodedGateOutputs,keycommitment and receiver entry key to the receiver through a byte array(got after applying json marshal to the SenToRecObject)
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}
	channel_SenToRec<-byteArrayFromSenToRec
	fmt.Println("\n->Sender \nByteArray sent to Receiver")

	fmt.Println("\n->Sender\n Initialization Finished")

	//creating a contract instance
	instance,err:= import_contracts.NewProtocol2(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	//The contract is in initialized stage
	//Once the contract is initialized, the receiver has time equal to waitingTime to accept and pay ether to the contract.
	//the for loop checks for the change in stage.
	//case 1 : If the receiver doesn't accept within the time limit, RevealedBool remains false after the for loop. As a result,
	//				"Receiver have not accepted within the timeout, Key not Revealed" is sent to the main function and the receiver doesn't
	//        proceed with the further protocol.
	//case 2 : If the receiver pays less than the price,then the contract stage goes to finishedAndMaliciousReciever.The sender sends
	//				"Malicious Receiver-Dind't pay enough ether" to the main function  and finishes the protocol execution here. The ether is sent back
	//				to the receiver by the contract.
	//case 3: The receiver accepts within the time. The contract stage goes to accepted and the sender has to reveal the key. RevealedBool becomes
	//				true.
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
		channel_SenToMain<-"Either the reciever is not interested now or Sender might have exchanged the indices of hashedEncryptedFileChunks"
	}else{

			if(!RevealedBool){
				channel_SenToMain<-"Receiver has not accepted within the timeout, Key not Revealed"

			}else{

							//Sender Reveals the key

							fmt.Println("\n->Sender\nReceiver have accepted within the timeout, Key Revealed")
							channel_SenToMain<-"I have not revealed the key"
					}

	}
}


func Sender_malicious_FileChunkWrongEnc(authSender *bind.TransactOpts,client *ethclient.Client,channel_SenToRec chan []byte,channel_SenToMain chan string){

	//reading from a file "a.txt" and storing bytearray of 32 bytes in FileChunks of size equal to noOfFileChunks
	fmt.Println("\n->Sender\nBegins")
	FileChunks:=FileRead()
	fmt.Println("->Sender\nFileChunks:",FileChunks)

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

	//encryptedFileChunks := EncryptFileChunks(FileChunks,key)
	maliciousEncryptedFileChunks := MaliciousEncryptedFileChunks(FileChunks,key,maliciouskey)

	//fmt.Println("sender encrypt file chunks",encryptedFileChunks)

	HashedEncryptedFileChunks:=HashChunks(maliciousEncryptedFileChunks)
	//maliciousHashIndexExchangeFileChunks := createMaliciousHashIndexExchangeFileChunk(HashedEncryptedFileChunks)
	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedEncryptedFileChunks=append(SenToConObject.HashedEncryptedFileChunks,HashedEncryptedFileChunks[i])
	}


	/*Deploy contract*/
	var err error
	//authSender.Value=big.NewInt(1000000000000)
	contractAddress, _,_, err=import_contracts.DeployProtocol2(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.HashedFileChunks,
		SenToConObject.HashedEncryptedFileChunks,
	)
	if err!=nil{
		log.Fatal(err)
	}


	fmt.Println("\n->Sender \n Contract Deployed")

	/*SenToRecObject initialize*/
	SenToRecObject.EncryptedFileChunks=maliciousEncryptedFileChunks


	//sending the encodedGateOutputs,keycommitment and receiver entry key to the receiver through a byte array(got after applying json marshal to the SenToRecObject)
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}
	channel_SenToRec<-byteArrayFromSenToRec
	fmt.Println("\n->Sender \nByteArray sent to Receiver")

	fmt.Println("\n->Sender\n Initialization Finished")

	//creating a contract instance
	instance,err:= import_contracts.NewProtocol2(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	//The contract is in initialized stage
	//Once the contract is initialized, the receiver has time equal to waitingTime to accept and pay ether to the contract.
	//the for loop checks for the change in stage.
	//case 1 : If the receiver doesn't accept within the time limit, RevealedBool remains false after the for loop. As a result,
	//				"Receiver have not accepted within the timeout, Key not Revealed" is sent to the main function and the receiver doesn't
	//        proceed with the further protocol.
	//case 2 : If the receiver pays less than the price,then the contract stage goes to finishedAndMaliciousReciever.The sender sends
	//				"Malicious Receiver-Dind't pay enough ether" to the main function  and finishes the protocol execution here. The ether is sent back
	//				to the receiver by the contract.
	//case 3: The receiver accepts within the time. The contract stage goes to accepted and the sender has to reveal the key. RevealedBool becomes
	//				true.
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
		channel_SenToMain<-"Either the reciever is not interested now or Sender might have exchanged the indices of hashedEncryptedFileChunks"
	}else{

			if(!RevealedBool){
				channel_SenToMain<-"Receiver has not accepted within the timeout, Key not Revealed"

			}else{

							//Sender Reveals the key
							var keyToCon []byte
							for i:=0;i<keyLength;i++{
								keyToCon=append(keyToCon,key[i])
							}

							_,err=instance.KeyReveal(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},keyToCon)
							if err!=nil{
								log.Fatal(err)
							}

							/*fmt.Println("\n->Sender\nReceiver have accepted within the timeout, Key Revealed")
							channel_SenToMain<-"I have not revealed the key"*/
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
						fmt.Println("\n->Sender\nIt seems I have wrongly encrypted filechunks")
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
		//	fmt.Println(Phase)
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


func Sender_malicious_InterchangeHashEncFileChunks(authSender *bind.TransactOpts,client *ethclient.Client,channel_SenToRec chan []byte,channel_SenToMain chan string){

	//reading from a file "a.txt" and storing bytearray of 32 bytes in FileChunks of size equal to noOfFileChunks
	fmt.Println("\n->Sender\nBegins")
	FileChunks:=FileRead()
	fmt.Println("->Sender\nFileChunks:",FileChunks)

	//objects
	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct


	setPrice(&SenToConObject)
	key:=keyGenerate()
	//maliciouskey:=keyGenerate()
	HashedFileChunks:=HashChunks(FileChunks)
	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedFileChunks=append(SenToConObject.HashedFileChunks,HashedFileChunks[i])
	}

	encryptedFileChunks := EncryptFileChunks(FileChunks,key)
	//maliciousEncryptedFileChunks := MaliciousEncryptedFileChunks(FileChunks,key,maliciouskey)

	//fmt.Println("sender encrypt file chunks",encryptedFileChunks)

	HashedEncryptedFileChunks:=HashChunks(encryptedFileChunks)
	maliciousHashIndexExchangeFileChunks := createMaliciousHashIndexExchangeFileChunk(HashedEncryptedFileChunks)
	for i:=0;i<noOfFileChunks;i++{
		SenToConObject.HashedEncryptedFileChunks=append(SenToConObject.HashedEncryptedFileChunks,maliciousHashIndexExchangeFileChunks[i])
	}


	/*Deploy contract*/
	var err error
	//authSender.Value=big.NewInt(1000000000000)
	contractAddress, _,_, err=import_contracts.DeployProtocol2(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.HashedFileChunks,
		SenToConObject.HashedEncryptedFileChunks,
	)
	if err!=nil{
		log.Fatal(err)
	}


	fmt.Println("\n->Sender \n Contract Deployed")

	/*SenToRecObject initialize*/
	SenToRecObject.EncryptedFileChunks=encryptedFileChunks


	//sending the encodedGateOutputs,keycommitment and receiver entry key to the receiver through a byte array(got after applying json marshal to the SenToRecObject)
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}
	channel_SenToRec<-byteArrayFromSenToRec
	fmt.Println("\n->Sender \nByteArray sent to Receiver")

	fmt.Println("\n->Sender\n Initialization Finished")

	//creating a contract instance
	instance,err:= import_contracts.NewProtocol2(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	//The contract is in initialized stage
	//Once the contract is initialized, the receiver has time equal to waitingTime to accept and pay ether to the contract.
	//the for loop checks for the change in stage.
	//case 1 : If the receiver doesn't accept within the time limit, RevealedBool remains false after the for loop. As a result,
	//				"Receiver have not accepted within the timeout, Key not Revealed" is sent to the main function and the receiver doesn't
	//        proceed with the further protocol.
	//case 2 : If the receiver pays less than the price,then the contract stage goes to finishedAndMaliciousReciever.The sender sends
	//				"Malicious Receiver-Dind't pay enough ether" to the main function  and finishes the protocol execution here. The ether is sent back
	//				to the receiver by the contract.
	//case 3: The receiver accepts within the time. The contract stage goes to accepted and the sender has to reveal the key. RevealedBool becomes
	//				true.
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
		channel_SenToMain<-"Either the reciever is not interested now or Sender might have exchanged the indices of hashedEncryptedFileChunks"
	}else{

			if(!RevealedBool){
				channel_SenToMain<-"Receiver has not accepted within the timeout, Key not Revealed"

			}else{

							//Sender Reveals the key
							var keyToCon []byte
							for i:=0;i<keyLength;i++{
								keyToCon=append(keyToCon,key[i])
							}

							_,err=instance.KeyReveal(&bind.TransactOpts{
									From:authSender.From,
									Signer:authSender.Signer,
									Value: nil,
							},keyToCon)
							if err!=nil{
								log.Fatal(err)
							}

							/*fmt.Println("\n->Sender\nReceiver have accepted within the timeout, Key Revealed")
							channel_SenToMain<-"I have not revealed the key"*/
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
						fmt.Println("\n->Sender\nIt seems I have wrongly encrypted filechunks")
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
		//	fmt.Println(Phase)
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

func Receiver_malicious_HashEncFileMismatch(authReceiver *bind.TransactOpts,client *ethclient.Client,channel_SenToRec chan []byte,channel_RecToMainIni chan string,channel_RecToMainRev chan string){


		byteArrayFromSenToRec:=<-channel_SenToRec
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
		for i=0;i<noOfFileChunks;i++ {
			value,err:=instance.HashedEncryptedFileChunks(nil,big.NewInt(i))
			if err!=nil{
				log.Fatal(err)
			}
			if(value!=Hash(Receiver_SenToRecObj.EncryptedFileChunks[i])){
				boolHashedEncryptedFileCheck=false
			}
		}
	}

		if(boolHashedEncryptedFileCheck){


				gasLimit := uint64(230000000)


				value:=big.NewInt(100)

			_	,err=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
				GasLimit:gasLimit,
			})
			if err != nil {
					log.Fatal(err)
			}



			fmt.Println("\n->Receiver\n Finished Accepting and Initialization phase ends")
			//Sending the finished accepted stage message to the main function
			channel_RecToMainIni<-"Accepted stage got over"


			//The contract is in accepted stage
			//The sender
			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
			//case 2 : The sender doesn't reveal the key within time. AcceptedBool remains false after the for loop. Receiver later calls the contract function
			//			   ReceiverGetEther to get the ether.
			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and AcceptedBool is true after the for loop.
			AcceptedBool:=false
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
					AcceptedBool=true
					break;
				}
				time.Sleep(time.Second *1)

			}



			if(!AcceptedBool){

					fmt.Println("\n->Receiver\nSender has not revealed the key...ReceiverGetEther function called")
				//	gasLimit := uint64(210000000000000)

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


					//channel_RecToMainIni<-"Sender hasn't revealed the key within the Timeout"
					channel_RecToMainRev<-"Sender hasn't revealed the key within the Timeout"

			} else{

				//Extracting function being called
				var _contractHashedFileChunks [noOfFileChunks][hashFunctionOutputBitSize]byte
				var i int64
				for i=0;i<noOfFileChunks;i++{
						value,err:=instance.HashedFileChunks(nil,big.NewInt(i))
						if err!=nil{
								log.Fatal(err)
						}
						_contractHashedFileChunks[i]=value
				}

				var _ReceiverKey [keyLength]byte
				for i:=0;i<keyLength;i++{
					_ReceiverKey[i]=contract_Key[i]
				}

				//isComplain,index:=Extract(_contractHashedFileChunks,Receiver_SenToRecObj.EncryptedFileChunks,_ReceiverKey)
				//malicious reciever
				isComplain:=true
				//index nil denotes no complain
				if(!isComplain){
					fmt.Println("\n->Receiver\nNo complain")
					channel_RecToMainRev<-"Sender has correctly calculated the encryption"
				} else{

						//complain function being called
						var complaintChunk []byte
						index:=4
						/*for i:=0;i<buffer_size;i++{
							  complaintChunk=append(complaintChunk,Receiver_SenToRecObj.EncryptedFileChunks[index][i])
						}*/

						_,err=instance.Complain(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value:nil,
							GasLimit:gasLimit,
						},complaintChunk,big.NewInt(int64(index)))
						if err!=nil{
							log.Fatal(err)
						}


						Phase,err:=instance.Phase(nil)
						if err!=nil{
							log.Fatal(err)
						}
						if(Phase==5){
							fmt.Println("\n->Receiver\nSender Has wrongly calulated the encrypted inputs")
							channel_RecToMainRev<-"sender-Wrong encrypted inputs"

						}
				}
			}
			channel_RecToMainRev<-"Reveal Phase Finished"
		}	else{

				fmt.Println("\n-> Receiver \n HashedEncryptedFileGotFromSender and HashedEncryptedFileFromContract dindn't match")
				channel_RecToMainIni<-"HashedEncryptedFileGotFromSender and HashedEncryptedFileFromContract dindn't match"
				channel_RecToMainRev<-"HashedEncryptedFileGotFromSender and HashedEncryptedFileFromContract dindn't match "
		}
}

func Receiver_malicious_HashFileMismatch(authReceiver *bind.TransactOpts,client *ethclient.Client,channel_SenToRec chan []byte,channel_RecToMainIni chan string,channel_RecToMainRev chan string){


		byteArrayFromSenToRec:=<-channel_SenToRec
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
		for i=0;i<noOfFileChunks;i++ {
			value,err:=instance.HashedEncryptedFileChunks(nil,big.NewInt(i))
			if err!=nil{
				log.Fatal(err)
			}
			if(value!=Hash(Receiver_SenToRecObj.EncryptedFileChunks[i])){
				boolHashedEncryptedFileCheck=false
			}
		}
	}

		if(boolHashedEncryptedFileCheck){


				gasLimit := uint64(230000000)


				value:=big.NewInt(100)

			_	,err=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
				GasLimit:gasLimit,
			})
			if err != nil {
					log.Fatal(err)
			}



			fmt.Println("\n->Receiver\n Finished Accepting and Initialization phase ends")
			//Sending the finished accepted stage message to the main function
			channel_RecToMainIni<-"Accepted stage got over"


			//The contract is in accepted stage
			//The sender
			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
			//case 2 : The sender doesn't reveal the key within time. AcceptedBool remains false after the for loop. Receiver later calls the contract function
			//			   ReceiverGetEther to get the ether.
			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and AcceptedBool is true after the for loop.
			AcceptedBool:=false
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
					AcceptedBool=true
					break;
				}
				time.Sleep(time.Second *1)

			}



			if(!AcceptedBool){

					fmt.Println("\n->Receiver\nSender has not revealed the key...ReceiverGetEther function called")
				//	gasLimit := uint64(210000000000000)

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


					//channel_RecToMainIni<-"Sender hasn't revealed the key within the Timeout"
					channel_RecToMainRev<-"Sender hasn't revealed the key within the Timeout"

			} else{

				//Extracting function being called
				var _contractHashedFileChunks [noOfFileChunks][hashFunctionOutputBitSize]byte
				var i int64
				for i=0;i<noOfFileChunks;i++{
						value,err:=instance.HashedFileChunks(nil,big.NewInt(i))
						if err!=nil{
								log.Fatal(err)
						}
						_contractHashedFileChunks[i]=value
				}

				var _ReceiverKey [keyLength]byte
				for i:=0;i<keyLength;i++{
					_ReceiverKey[i]=contract_Key[i]
				}

				//isComplain,index:=Extract(_contractHashedFileChunks,Receiver_SenToRecObj.EncryptedFileChunks,_ReceiverKey)
				//malicious reciever
				isComplain:=true
				//index nil denotes no complain
				if(!isComplain){
					fmt.Println("\n->Receiver\nNo complain")
					channel_RecToMainRev<-"Sender has correctly calculated the encryption"
				} else{

						//complain function being called
						var complaintChunk []byte
						index:=4
						for i:=0;i<buffer_size;i++{
							  complaintChunk=append(complaintChunk,Receiver_SenToRecObj.EncryptedFileChunks[index][i])
						}

						_,err=instance.Complain(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value:nil,
							GasLimit:gasLimit,
						},complaintChunk,big.NewInt(int64(index)))
						if err!=nil{
							log.Fatal(err)
						}


						Phase,err:=instance.Phase(nil)
						if err!=nil{
							log.Fatal(err)
						}
						if(Phase==5){
							fmt.Println("\n->Receiver\nSender Has wrongly calulated the encrypted inputs")
							channel_RecToMainRev<-"sender-Wrong encrypted inputs"

						}
				}
			}
			channel_RecToMainRev<-"Reveal Phase Finished"
		}	else{

				fmt.Println("\n-> Receiver \n HashedEncryptedFileGotFromSender and HashedEncryptedFileFromContract dindn't match")
				channel_RecToMainIni<-"HashedEncryptedFileGotFromSender and HashedEncryptedFileFromContract dindn't match"
				channel_RecToMainRev<-"HashedEncryptedFileGotFromSender and HashedEncryptedFileFromContract dindn't match "
		}
}
