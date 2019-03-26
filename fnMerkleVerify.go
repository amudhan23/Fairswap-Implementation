func Mverify(complaint [depth][32]byte,_index int,ciphertextRoot [32]byte) (bool){
	_value:=complaint[0]

	var Hashed []byte
	var i uint
	for i=0; i<depth-1; i++{
		if ((_index & (1<<i))>>i == 1)	{

			var toBeHashed []byte
			for j:=0;j<32;j++{
				toBeHashed=append(toBeHashed,complaint[i+1][j])
			}
			for j:=0;j<32;j++{
				toBeHashed=append(toBeHashed, _value[j])
			}

			Hashed = kec.Keccak256(toBeHashed)

			for j:=0;j<32;j++{
				_value[j]=Hashed[j]

			}

		}else{

			var toBeHashed []byte
			for j:=0;j<32;j++{
				toBeHashed=append(toBeHashed,_value[j])
			}
			for j:=0;j<32;j++{
				toBeHashed=append(toBeHashed,complaint[i+1][j])
			}

			Hashed = kec.Keccak256(toBeHashed)

			for j:=0;j<32;j++{
				_value[j]=Hashed[j]
			}

		}

	}
	return (_value==ciphertextRoot)
}
