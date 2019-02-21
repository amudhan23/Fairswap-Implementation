

//reading from a file "a.txt" and storing bytearray of 32 bytes in inputvectors of size equal to NoOfInputGates


var currentByte int64 = 0
fileBuffer := make([]byte, BUFFER_SIZE)
file, err := os.Open("a.txt") // For read access.
if err != nil {

		log.Fatal(err)
}
var inputvectors [NoOfInputGates][32]byte
count:=0
for {

		n,err2 := file.ReadAt(fileBuffer, currentByte)
		currentByte += BUFFER_SIZE
		fileBuffer=fileBuffer[:n]
		if count==NoOfInputGates{
				break
		}

		for j:=0;j<n;j++{
			inputvectors[count][j]=fileBuffer[j]
		}

		if err2 == io.EOF {
				break

		}
		count++
}
file.Close()
