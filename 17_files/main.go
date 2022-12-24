package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Learning how to handle files")

	content := "This is sample text that I want to add to a file - VinayChunchu"

	file, err := os.Create("./textOutput.txt")

	// if err != nil {
	// 	// We can use panic keyword in case of error. Panic will shutdown execution and shows error
	// 	// This is common way to handle errors
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length output of file writing : ", length)
	defer file.Close()

	readFile("./textOutput.txt")

}

func readFile(filePath string) {
	incomingBytes, err := ioutil.ReadFile(filePath)
	checkNilError(err)

	fmt.Println("Content of the file: ", incomingBytes)

	// ways to handle byte data
	fmt.Println("Content of the file: ", string(incomingBytes))
}

func checkNilError(err error) {
	if err != nil {
		// We can use panic keyword in case of error. Panic will shutdown execution and shows error
		// This is common way to handle errors
		panic(err)
	}
}
