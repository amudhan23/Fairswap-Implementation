

type sendertocontract struct{
	Id int
	Price big.Int
	Keycommit [32]byte
	merkleRootOfCircuit [32]byte
	merkleRootOfEncInp [32]byte
}
