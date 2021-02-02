package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello My Little Froglet!\n")
	_ = ioutil.WriteFile("file1", data, 0644)
	fileOneByteArray, _ := ioutil.ReadFile("file1")
	fmt.Println(string(fileOneByteArray)) // []byte ==> string "Hello..."

	fileTwo, _ := os.Create("file2")
	defer fileTwo.Close()
	fileTwoN, _ := fileTwo.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", fileTwoN)

	fileTwo, _ = os.Open("file2")
	defer fileTwo.Close()
	b := bytes.Buffer{}
	by, _ = fileTwo.Read(b.Bytes())
	fmt.Printf("Read %d bytes from file \n", by)
	fmt.Println(string(by))
}
