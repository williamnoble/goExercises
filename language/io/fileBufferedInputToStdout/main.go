package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path")
		return
	}

	f, _ := os.Open(os.Args[1])
	defer f.Close()

	var b = make([]byte, 4)
	var err error

	for n := 0; err == nil; {
		n, err = f.Read(b)
		if err == nil {
			// Append most recent bytes to the end of the string
			fmt.Print(string(b[:n]))
		}
	}
	if err != nil && err != io.EOF {
		fmt.Println("\n\nError", err)
	}
}
