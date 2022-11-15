package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to file handling")

	writeToFile("./myLogFile.txt")

	readFromFile("./myLogFile.txt")
}

// example for writing the text into file
func writeToFile(filepath string) {

	content := "This text need to go into the file"

	file, err := os.Create(filepath)

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("written content length is : ", length)
}

// Example for reading the text from the file
func readFromFile(filepath string) {

	byteData, err := ioutil.ReadFile(filepath)

	checkNilError(err)

	//when we are reading from file we will get the data in byte array so we neeed to convert from byte array to string below is one of the simple method how to do this
	fmt.Println(string(byteData))
}

func checkNilError(err error) {

	if err != nil {
		panic(err)
	}
}
