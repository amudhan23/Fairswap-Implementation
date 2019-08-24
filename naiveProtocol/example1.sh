#/bin/sh

echo "Enter the byte size of each file chunk "
echo "Enter the number of file chunks "
echo "Enter the number of pairs that should run fairswap protocol parallely "
echo "Enter the number of pairs that should run naive protocol parallely "
echo "Enter the keystore path directory for buyers "
echo "Enter the keystore pathe directory for sellers "
echo "Enter the geth rpc of node 1 "
echo "Enter the geth rpc of node 2 "
echo "Log file Path "
read filechunksize numberOfFileChunks numberOfPairsFairswap numberOfPairsNaive keyStoreDirectoryforBuyers keyStoreDirectoryforSellers gethNode1ipc gethNode2ipc logFilePath
echo $filechunksize $numberOfFileChunks $numberOfPairsFairswap $numberOfPairsNaive $keyStoreDirectoryforBuyers $keyStoreDirectoryforSellers $gethNode1ipc $gethNode2ipc $logFilePath

if [ $numberOfPairsNaive -gt 0 ]
then
 	go run NaiveProtocol_Pri1.go $filechunksize $numberOfFileChunks $numberOfPairsNaive $keyStoreDirectoryforBuyers $keyStoreDirectoryforSellers $gethNode1ipc $gethNode2ipc $logFilePath
fi

if [ $numberOfPairsFairswap -gt 0 ]
then
	go run Fairswap_Private.go $filechunksize $numberOfFileChunks $numberOfPairsFairswap $keyStoreDirectoryforBuyers $keyStoreDirectoryforSellers $gethNode1ipc $gethNode2ipc $logFilePath
fi
