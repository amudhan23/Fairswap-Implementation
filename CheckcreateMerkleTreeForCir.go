//fn CheckcreateMerkleTreeForCir
//constructs merkletree and compares between the merkleroot got from construction and merkleroot passed as function parameter
//input parameters
//cicuitobjects : array of ciruit gate object
//merklerootforcircuit : merkleroot of circuit

//Returns bool value,true if merklerootforcircuit matches with the merkleroot got from construction,else false


func CheckcreateMerkleTreeForCir(circuitobjects [TotalNumberOfgates]circuit,merklerootforcircuit [32]byte) (bool){

		var circuitByteArray [TotalNumberOfgates][]byte     //contains the json.marshal byte array of each gate

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

	for j:=0;j<32;j++{
		if(merkletree[2*TotalNumberOfgates-2][j]!=merklerootforcircuit[j]){
			return false
			break
		}
	}

	return true
}
