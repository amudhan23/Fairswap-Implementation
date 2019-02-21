//fn createMerkleTree
//creates a merkle tree and returns the merkleroot


func createMerkleTree(inputvectors [NoOfInputGates][32]byte) ([32]byte){
	var merkletree [2*NoOfInputGates-1][32]byte

	//input gates
	for i:=0;i<NoOfInputGates;i++{
		for j:=0;j<32;j++{
				merkletree[i][j]=inputvectors[i][j]
		}
	}

	//!inputgates
	k:=0
	for i:=0;i<NoOfInputGates-1;i++{

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
			merkletree[i+NoOfInputGates][j]=Hashed[j]
		}
		k=k+2
	}

	return merkletree[2*NoOfInputGates-2]
}
