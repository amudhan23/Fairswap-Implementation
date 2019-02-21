




//address(sender,receiver) generation
//sender private key generation
privatekeysender,err := crypto.GenerateKey()
authsender := bind.NewKeyedTransactor(privatekeysender)

//reciever private key generation
privatekeyreceiver,err:=crypto.GenerateKey()
authreceiver :=bind.NewKeyedTransactor(privatekeyreceiver)

//genesisalloc
alloc := make(core.GenesisAlloc)
alloc[authsender.From] = core.GenesisAccount{Balance: big.NewInt(1100000000000000000)}  //110ether
alloc[authreceiver.From] = core.GenesisAccount{Balance: big.NewInt(2500000000000000000)} //250eether

//client
client := backends.NewSimulatedBackend(alloc,6000000)

//sender address generation
publickeysender := privatekeysender.Public()
publickeyECDSAsender, ok := publickeysender.(*ecdsa.PublicKey)
if !ok {
		log.Fatal("error casting public key to ECDSA")
}
senderaddress := crypto.PubkeyToAddress(*publickeyECDSAsender)
sendernonce, err := client.PendingNonceAt(context.Background(), senderaddress)
if err != nil {
		log.Fatal(err)
}
authsender.Nonce = big.NewInt(int64(sendernonce))
authsender.Value = big.NewInt(0)

//receiveraddress generation
publickeyreceiver := privatekeyreceiver.Public()
publickeyECDSAreceiver, ok := publickeyreceiver.(*ecdsa.PublicKey)
if !ok {
		log.Fatal("error casting public key to ECDSA")
}
receiveraddress := crypto.PubkeyToAddress(*publickeyECDSAreceiver)
receivernonce, err := client.PendingNonceAt(context.Background(), receiveraddress)
if err != nil {
		log.Fatal(err)
}
authreceiver.Nonce = big.NewInt(int64(receivernonce))
authreceiver.Value = big.NewInt(0)
