//fn createMerkleTreeForCir
//the function creates a merkle tree for the circuit with each gate's byte array(got after doing json.Marshal to each ciruit gate) as input
//input parameters
//cicuitobjects : array of ciruit gate object

//returns merkleroot

//Unlike the merkletreeconstruction for encryptedGateOutputs,where the input vectors were not hashed(in the sense they were stored as it is), here the inputvectors
//to the circuit merkletreeconstruction is hashed and stored to maintain the uniformity in number of bytes in each gate array




func createMerkleTreeForCir(circuitobjects [TotalNumberOfgates]circuit) ([32]byte){

	var circuitByteArray [TotalNumberOfgates][]byte  //contains the json.marshal byte array of each gate

	for i:=0;i<TotalNumberOfgates;i++ {
		b, err := json.Marshal(circuitobjects[i])
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
