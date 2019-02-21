//fn CheckcreateMerkleTreeForEncInp
//constructs merkletree and compares between the merkleroot got from construction and merkleroot passed as function parameter
//input parameters
//encryptedGateOutputs : Vectors of encrypted output of the circuit gates
//merklerootforEncInput : merkle root of Encrypted Input
//Returns bool value,true if merklerootforcircuit matches with the merkleroot got from construction,else false
func CheckcreateMerkleTreeForEncInp(encryptedGateOutputs [TotNumOfEncOutVecToMer][32]byte, merklerootforEncInput [32]byte) (bool){
	var merkletree [2*TotNumOfEncOutVecToMer-1][32]byte


	//inputgates
	for i:=0;i<TotNumOfEncOutVecToMer;i++{
		for j:=0;j<32;j++{
				merkletree[i][j]=encryptedGateOutputs[i][j]
		}
	}


	//!inputgates
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
		Hashed=kec.Keccak256(toBeHashed)
		for j:=0;j<32;j++{
			merkletree[i+TotNumOfEncOutVecToMer][j]=Hashed[j]
		}
		k=k+2
	}


	for j:=0;j<32;j++{
		if(merkletree[2*TotNumOfEncOutVecToMer-2][j]!=merklerootforEncInput[j]){
			return false
		}
	}
	return true
}
