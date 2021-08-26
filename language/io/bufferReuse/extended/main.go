package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer(make([]byte, 4)) // return *Buffer

	texts := []string{
		"into the wind",
		"unto the fray",
		"beneath the depths, I find a dragon",
	}

	for index := range texts {
		b.Reset() // clear the buffer
		b.WriteString(texts[index])
		fmt.Printf("len: %d, cap: %d, data: %s.\n", b.Len(), b.Cap(), b)
	}
}
