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
