package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	newFile, err := os.Create(filepath.FromSlash("../test/test.txt"))
	if err != nil {
		log.Fatal(err)
	}
	newFile.Write([]byte("Simple Byte Array\n"))
	newFile.WriteString("I'm a test String\n")
	_, err = newFile.WriteAt([]byte("TTT"), 2)
	//err = newFile.Truncate(12)
	log.Println(newFile)
	err = newFile.Close()
	if err != nil {
		panic(err)
	}
}
