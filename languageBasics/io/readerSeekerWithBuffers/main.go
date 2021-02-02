package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	s := strings.NewReader("a\r\nb")
	fmt.Printf("a\r\nb")
	fmt.Printf("\n=>Example 1 - r.ReadLine()\n")
	r := bufio.NewReader(s)
	for {
		token, _, err := r.ReadLine()
		if len(token) > 0 {
			fmt.Printf("Token (ReadLine): %q\n", token)
		}
		if err != nil {
			break
		}
	}
	fmt.Printf("\n=>Example 2 - r.ReadByte(\\n)\n")
	s.Seek(0, io.SeekStart)
	r.Reset(s)
	for {
		token, err := r.ReadBytes('\n')
		fmt.Printf("Token (ReadBytes): %q\n", token)
		if err != nil {
			break
		}
	}
	fmt.Printf("\n=>Example 3 - r.NewScanner(default")
	s.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		fmt.Printf("Token (Scanner): %q\n", scanner.Text())
	}
}
