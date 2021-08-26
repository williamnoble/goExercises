package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	newFile *os.File
	err     error
)

func main() {
	/*
		1. Return the path for the executable (main)
		2. Return the Directory housing the executable from 1
		3. Create a file at 2
		4. Do nothing of note.
	*/

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	newFile, err := os.Create("testFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	fmt.Println(newFile)

}
