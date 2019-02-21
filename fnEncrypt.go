//Encrypt function
//The function is used to encrypt(XOR) the given vector.

//Input Parameters
//key : The key with which the XORing is done.
//plaintext : the input vector

//Returns the encrypt output

//Operation performed : plaintext (XOR) key


func Enc(key [32]byte,plaintext [32]byte) ([32]byte) {

			var keyPlusIndex []byte //appends the key(input parameter) and index "i"(range : 0 to Total Number of gates-1 )

			//In this specific case we are solving,the index("0") always remains the same.The funcion can be easily extended to accomadate varying index
			k:=0
			for i:=0;i<len(key);i++{
				keyPlusIndex=append(keyPlusIndex,key[i])
			}
			keyPlusIndex=append(keyPlusIndex,byte(k))
			key0 := kec.Keccak256(keyPlusIndex)


			var key032 [32]byte
			for i:=0;i<len(key0);i++{
				key032[i]=key0[i]
			}

			var encryptedtext [32]byte
			for i:=0;i<32;i++{
				encryptedtext[i]=key032[i]^plaintext[i]  //xor operation
			}
			return encryptedtext
}
