

type sendertoreciever struct {
		Id int                                        //id
		Keycommit [32]byte														//key commitment
		EncodedOutputOfGates [TotNumOfEncOutVecToMer][32]byte		//Encoded output of gates
}
