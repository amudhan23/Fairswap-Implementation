//maliciousReceiver_WrongComplain function
//The functoin in which the receiver creates a wrong complaint
//Input parameters:
//authReceiver(*bind.TransactOpts) - the receiver in the blockchain
//client()*backends.SimulatedBackend) - blockchain(here it is a backend blockchain)
//channel_SenToRec(chan []byte) - to receiver the byte array from the sender
//channel_RecToMainIni(chan string) - to notify the main function for the completion of initialization phase
//channel_RecToMainRev(chan string) - to notify the main function for the completion of reveal phase
func maliciousReceiver_WrongComplain(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,channel_RecToMainIni chan string,channel_RecToMainRev chan string){


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

		Receiver_MerkleRootOfEncInp,err:=instance.EncInputRoot(nil)
		if err!=nil{
			log.Fatal(err)
		}

		Receiver_EncodedOutputOfGates:=Receiver_SenToRecObj.EncodedOutputOfGates

		//comparing the merkle root of encodedGateOutputs got from contract and the calculated merkleroot from the encodedGateOutputs got from sender
		boolValueEncode:=CheckcreateMerkleTreeForEncInp(Receiver_EncodedOutputOfGates,Receiver_MerkleRootOfEncInp)

		fmt.Println("\n->Receiver\boolValueEncode(comparision between merkleroot of encodedGateOutputs from contract and calulated encodedGateOutputs merkleroot): ",boolValueEncode)

		//if boolValueEncode is false, then the receiver doesn't proceed with the protocol
		if(boolValueEncode){

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
			gasLimit := uint64(2100000000)

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


				/*malicious act*/
				//creates a wrong circuit for mainpulating a invalid complaint
				var maliciousCircuitObjects [totalNumberOfGates]circuit  //array of circuit object
				//sets the circuit
				setMaliciousCircuitTuples(&maliciousCircuitObjects)

				fmt.Println("\n->Receiver\nSender has revealed the key....Proceeding further")

			 //Receiver_EncodedOutputOfGates[9][0]=3
				merkletree:=ReceiverMerkleTreeCreate(Receiver_EncodedOutputOfGates)

				//Extracting function being called
				//sending malicious circuit to the function
				complain,decodedOutputs,index:=Extract(Receiver_EncodedOutputOfGates,contract_Key,maliciousCircuitObjects,merkletree,Receiver_MerkleRootOfEncInp)

				//index nil denotes no complain
				if(index==nil){
					fmt.Println("\n->Receiver\nNo complain")
					fmt.Println("\n->Receiver\nDecoded Ouptuts : ",decodedOutputs)

				} else{

					fmt.Println("complain :",complain)
					indexes:=[]*big.Int{}
					for i:=0;i<len(index);i++{
						b:=big.NewInt(int64(index[i]))
						indexes=append(indexes,b)
						}

						//complain function being called
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

						//checking whether the given complaint is valid by verifying the stage changes.
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


//for setting a werong circuit
//changing a gate operation from 2 to 1
func setMaliciousCircuitTuples(circuitObjects *[totalNumberOfGates]circuit){

	//Giving tuple value to input gates
	for i:=0;i<noOfInputGates;i++{
			circuitObjects[i].Index = i
			circuitObjects[i].Operation =1
			circuitObjects[i].InputToTheGate = nil
	}

	//Giving tuple values to !inputgates && !lastgate
	k:=0
	for i:=noOfInputGates;i<totalNumberOfGates-1;i++ {
			circuitObjects[i].Index=i
			circuitObjects[i].Operation = 2
			for j:=0;j<maxLinesToGate;j++{
				circuitObjects[i].InputToTheGate=append(circuitObjects[i].InputToTheGate,k)
				k++
			}
	}

	//set wrong gate operation
	circuitObjects[9].Operation = 1

	//Last gate
	circuitObjects[totalNumberOfGates-1].Index = totalNumberOfGates-1
	circuitObjects[totalNumberOfGates-1].Operation = 3
	circuitObjects[totalNumberOfGates-1].InputToTheGate=append(circuitObjects[totalNumberOfGates-1].InputToTheGate,totalNumberOfGates-2)

}
