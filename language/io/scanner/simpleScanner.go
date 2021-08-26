package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// Let's do a scanner
	stringLiteral := "Life is like a box of chocolates, you never know what you're going to get."
	b := []byte(stringLiteral)

	// create a byteReader
	byteReader := bytes.NewReader(b)
	_ = byteReader
	stringReader := strings.NewReader(stringLiteral)

	// pass a byteReader to a Scanner
	s := bufio.NewScanner(stringReader)

	// define how the scanner splits (default is bufio.ScanLines)
	s.Split(bufio.ScanWords)

	// call scanner.Split
	for s.Scan() {
		fmt.Printf("%q ", s.Text())
	}
}
