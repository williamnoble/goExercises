package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	rawString := "Raw String Literal"
	byteArray := []byte("Hello Byte-World")

	// Concert between raw string literal and byte array
	byteArrayToString := string(byteArray)
	stringToBytes := []byte(rawString)

	// Default, create a new buffer
	bytesToBuffer := bytes.NewBuffer(byteArray)
	fmt.Printf("byteArray: %s byteArrayToString: %s byteArrayToString: %s\n", byteArray, byteArrayToString, string(byteArray))

	// Initialise a buffer
	bufferOne := new(bytes.Buffer)
	bufferOne.Write(byteArray)
	bufferTwo := bytes.NewBufferString(rawString)
	bufferThree := bytes.NewBuffer(byteArray)
	fmt.Printf("Buffers: 1.%v\t 2.%v\t 3.%v\n", bufferOne, bufferTwo, bufferThree)

	// Read bytes and convert to string
	read, err := ioutil.ReadAll(bufferThree)
	if err != nil {
		print(err.Error())
	}
	output := string(read)

	fmt.Println("New Buffer", bufferThree, "to String:", output)
	fmt.Println()

	//
	rawLiteral := "Life is like a box of chocolates"
	reader := bytes.NewReader([]byte(rawLiteral))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Print(scanner.Text(), " ")
	}

	_ = rawString
	_ = byteArrayToString
	_ = stringToBytes
	_ = bytesToBuffer
}
