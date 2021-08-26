package main

import (
	"bytes"
	"fmt"
	"os"
)

const grr = "G.R.R. Martin"

type book struct {
	Author, Title string
	Year          int
}

var bookList = []book{
	{Author: grr, Title: "A Game of Thrones", Year: 1996},
	{Author: grr, Title: "A Clash of Kings", Year: 1998},
	{Author: grr, Title: "A Storm of Swords", Year: 2000},
	{Author: grr, Title: "A Feast for Crows", Year: 2005},
	{Author: grr, Title: "A Dance with Dragons", Year: 2011},
	// if year is omitted it defaulting to zero value
	{Author: grr, Title: "The Winds of Winter"},
	{Author: grr, Title: "A Dream of Spring"},
}

//goland:noinspection ALL
func main() {
	f, _ := os.OpenFile("./books_delete.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()

	// bytes.NewBuffer returns a Pointer to a Buffer
	b := bytes.NewBuffer([]byte{})
	for _, v := range bookList {
		fmt.Fprintf(b, "%s %s", v.Author, v.Title)
		// Write a new line to the buffer
		b.WriteRune('\n')
		// Write the Buffer (b Buffer) WriteTo to an io.Reader (f) from os.OpenFile. Note pointer receiver
		if _, err := b.WriteTo(f); err != nil {
			fmt.Println("error")
			return
		}
	}

}
