//fn fnKeycommit
//commits(hash's) the key

//input parameters
//key : key to be commited

//returns the commitment of key

func fnKeycommit(key [32]byte) ([32]byte){

			var toBEHashedKey []byte
			var HashedKey [32]byte

			for j:=0;j<32;j++{
				toBEHashedKey=append(toBEHashedKey,key[j])
			}

			HashedAfter:=kec.Keccak256(toBEHashedKey)

			for j:=0;j<32;j++{
					HashedKey[j]=HashedAfter[j]
			}

			return HashedKey
}
