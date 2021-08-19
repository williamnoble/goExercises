package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer(make([]byte, 2))
	entry := []string{
		`a`,        // 1
		`bb`,       // 2
		`ccc`,      // 3
		`dddd`,     // 4
		`eeeeeeee`, // 8
	}

	for i := range entry {
		b.Reset()
		b.WriteString(entry[i])
		fmt.Printf("Entry: %s\tLength: %d\tCapacity: %d\n", b, b.Len(), b.Cap())
	}
}
