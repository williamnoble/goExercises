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
}

type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

func main() {
	b := Book{"William Noble", "The Art of Go"}
	WriteToTerminal(b)

	c := Count(24)
	WriteToTerminal(c)
}
