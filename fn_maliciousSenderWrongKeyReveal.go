//maliciousSender_WrongKeyReveal function
//The sender being malicious reveals a wrong key
//Input parameters:
//authsender(*bind.TransactOpts) - the sender in the blockchain
//client()*backends.SimulatedBackend) - blockchain(here it is a backend blockchain)
//channel_SenToRec(chan []byte) - to send the byte array to the receiver
//channel_SenToMain(chan string) - to communicate with the main function
func maliciousSender_WrongKeyReveal(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,channel_SenToMain chan string){

	//reading from a file "a.txt" and storing bytearray of 32 bytes in inputVectors of size equal to noOfInputGates
	fmt.Println("\n->Sender\nBegins")
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

	fmt.Println("->Sender\nInputVectors:",inputVectors)

	//objects
	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct
	var circuitObjects [totalNumberOfGates]circuit //array of circuit object

	//sets the circuit
	setCircuitTuples(&circuitObjects)

	/*sender fns call*/
	key:=keyGenerate()
	keyCommit:=fnKeycommit(key)
	MRinp:=createMerkleTree(inputVectors)                          //MRinp (merkle root for input vectors)
	//MRcir:=createMerkleTreeForCir(circuitObjects)									//MRcir (merkle root for ciruit)
	encodedGateOutputs,_:=Encode(inputVectors,key,circuitObjects,MRinp)
	MRencout:=createMerkleTreeForEncInp(encodedGateOutputs)					//MRencout(merkle root for encoded outputs)
	recEntryKey:=keyGenerate()

	fmt.Println("\n->Sender \n Key:",key,"\n\nkeyCommit : ",keyCommit,"\n\nMRinp (merkle root for input vectors): ",MRinp,"\n\n MRencout(merkle root for encoded outputs) : ",MRencout)

	/*sentoconobject initialize*/
	idInitializingForContract(&SenToConObject)
	setPrice(&SenToConObject)
	SenToConObject.Keycommit=keyCommit
	//SenToConObject.MerkleRootOfCircuit=MRcir
	SenToConObject.MerkleRootOfEncInp=MRencout
	SenToConObject.ConReceiverEntryKey=recEntryKey

	//Putting gate operations in an array
	CircuitGateOperationArray:=[]*big.Int{}
	for i:=0;i<totalNumberOfGates;i++{
		op:=big.NewInt(int64(circuitObjects[i].Operation))
		CircuitGateOperationArray=append(CircuitGateOperationArray,op)
	}

	/*Deploy contract*/
	contractAddress, _,_, err=import_contracts.DeployFairswap(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.Keycommit,
		SenToConObject.MerkleRootOfEncInp,
		SenToConObject.ConReceiverEntryKey,
		CircuitGateOperationArray,

	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()

	fmt.Println("\n->Sender \n Contract Deployed")

	/*SenToRecObject initialize*/
	idInitializingForRec(&SenToRecObject)
	SenToRecObject.Keycommit=keyCommit
	SenToRecObject.EncodedOutputOfGates=encodedGateOutputs
	SenToRecObject.ReceiverEntryKey=recEntryKey


	//sending the encodedGateOutputs,keycommitment and receiver entry key to the receiver through a byte array(got after applying json marshal to the SenToRecObject)
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

	if(!RevealedBool){
		channel_SenToMain<-"Receiver have not accepted within the timeout, Key not Revealed"

	}else{

					/*malicious act*/
					//Sender Reveals the  wrong key
					notContractKey:=keyGenerate()
					_,err=instance.RevealKey(&bind.TransactOpts{
							From:authSender.From,
							Signer:authSender.Signer,
							Value: nil,
					},notContractKey)
					if err!=nil{
						log.Fatal(err)
					}
					client.Commit()
					fmt.Println("\n->Sender\nReceiver have accepted within the timeout, Key Revealed")
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
			time.Sleep(time.Second *1)
	}

	gasLimit := uint64(2100000000)

	//the sender getting the ether
	_,err=instance.SenderGetEther(&bind.TransactOpts{
			From:authSender.From,
			Signer:authSender.Signer,
			Value: nil,
			GasLimit:gasLimit,
		})
	if err!=nil{
				log.Fatal(err)
	}
	client.Commit()

	//the protocol gets finished and the sender sends the message to the main function
	channel_SenToMain<-"Fairswap completed.Sender Got the money"
}
