package main

import (
	"fmt"
	"strconv"
)

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Author, b.Title)
}

func WriteToTerminal(x fmt.Stringer) {
	fmt.Println(x.String())
	fmt.Println(x)
	// You can either explictely call x.String() or use x as it satisfies stringer interface already
}

type Count int

func (c Count) String() string {
	// Convert an INT to a STRING (Itoa) however we use int(c) as using custom type.
	return strconv.Itoa(int(c))
}

func main() {
	b := Book{"William Noble", "The Art of Go"}
	WriteToTerminal(b)

	c := Count(24)
	WriteToTerminal(c)
}
