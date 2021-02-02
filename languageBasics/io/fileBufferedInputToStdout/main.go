package main

import (
	"fmt"
	"io"
	"os"
)

// This example uses a byte array as a buffer, obviously because we're over-writing the contents
// on each loop we need to append to the end of the string on each pass
// Interesting for loop (try to memorise)
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error", err)
		return
	}

	defer f.Close()

	var b = make([]byte, 16)

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
