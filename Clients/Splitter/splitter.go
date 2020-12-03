package splitter

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Printf("Ingrese nombre de archivo: ")
// 	libro, _ := reader.ReadString('\n')
// 	libro = strings.TrimSuffix(libro, "\n")
// 	libro = strings.TrimSuffix(libro, "\r")

// 	// splitter(libro)

// 	combiner(libro, 3)

// }

// Combiner ... asda
func Combiner(name string, totalPartsNum uint64, libroarray [][]byte) {

	// just for fun, let's recombine back the chunked files in a new file

	newFileName := "a" + name
	_, err := os.Create(newFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//set the newFileName file to APPEND MODE!!
	// open files r and w

	file, err := os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// IMPORTANT! do not defer a file.Close when opening a file for APPEND mode!
	// defer file.Close()

	// just information on which part of the new file we are appending
	var writePosition int64 = 0

	for j := uint64(0); j < totalPartsNum; j++ {

		// calculate the bytes size of each chunk
		// we are not going to rely on previous data and constant
		chunkBufferBytes := libroarray[j]
		var chunkSize int64 = int64(len(chunkBufferBytes))
		// chunkBufferBytes := make([]byte, chunkSize)

		fmt.Println("Agregando en posicion : [", writePosition, "] bytes")
		writePosition = writePosition + chunkSize

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
		// write/save buffer to disk
		//ioutil.WriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

		n, err := file.Write(chunkBufferBytes)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		file.Sync() //flush to disk

		// free up the buffer for next cycle
		// should not be a problem if the chunk size is small, but
		// can be resource hogging if the chunk size is huge.
		// also a good practice to clean up your own plate after eating

		chunkBufferBytes = nil // reset or empty our buffer

		fmt.Println("Se han escrito ", n, " bytes")

		fmt.Println("Juntando parte [", j, "] en : ", newFileName)
	}

	// now, we close the newFileName
	file.Close()
}

// Splitter ...
func Splitter(name string) uint64 {
	fileToBeChunked := name // change here!

	file, err := os.Open(fileToBeChunked)

	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
		return 0
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = (250 * (1 << 10)) // 250 KB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Dividiendo el archivo en %d partes.\n", totalPartsNum)

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		// write to disk
		fileName := name + "_" + strconv.FormatUint(i, 10)
		_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

		fmt.Println("Dividido a : ", fileName)
	}
	return totalPartsNum

}
