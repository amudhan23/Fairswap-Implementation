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
