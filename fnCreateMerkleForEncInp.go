//Function  createMerkleTreeForEncInp
//constructs the merkle tree for the enoded inputs and returns the merkleroot

//Input Parameters :
//encryptedGateOutputs : Vectors of encrypted output of the circuit gates

//Returns merkleroot


func createMerkleTreeForEncInp(encryptedGateOutputs [TotNumOfEncOutVecToMer][32]byte) ([32]byte){
	var merkletree [2*TotNumOfEncOutVecToMer-1][32]byte

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<TotNumOfEncOutVecToMer;i++{
		for j:=0;j<32;j++{
				merkletree[i][j]=encryptedGateOutputs[i][j]
		}
	}

	//!inputgates
	//merkle tree construction
	k:=0
	for i:=0;i<TotNumOfEncOutVecToMer-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<32;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<32;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input

		for j:=0;j<32;j++{
			merkletree[i+TotNumOfEncOutVecToMer][j]=Hashed[j]
		}
		k=k+2
	}

	return merkletree[2*TotNumOfEncOutVecToMer-2]  //returns the merkleroot
}
