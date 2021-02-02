package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specific a path")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer f.Close()
	r := bufio.NewReader(f)

	var rowCount int // Keep track of the number of Rows
	for err == nil { // DON'T STOP TIL WE GET ERR.EOF
		var b []byte
		for moar := true; err == nil && moar; {
			b, moar, err = r.ReadLine()
			if err == nil {
				fmt.Print(string(b))
			}
		}
		// Each time moar is false a line is completely read
		if err == nil {
			fmt.Println()
			rowCount++
		}
	}
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	fmt.Println("\nRow Count:", rowCount)
}
