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

func main() {
	// 0666 = default file permissions.
	dst, err := os.OpenFile("./book_list2.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer dst.Close()
	bookList := []book{
		{Author: grr, Title: "A Game of Thrones", Year: 1996},
		{Author: grr, Title: "A Clash of Kings", Year: 1998},
		{Author: grr, Title: "A Storm of Swords", Year: 2000},
		{Author: grr, Title: "A Feast for Crows", Year: 2005},
		{Author: grr, Title: "A Dance with Dragons", Year: 2011},
		// if year is omitted it defaulting to zero value
		{Author: grr, Title: "The Winds of Winter"},
		{Author: grr, Title: "A Dream of Spring"},
	}
	b := bytes.NewBuffer(make([]byte, 0, 16))
	for _, v := range bookList {
		// prints a msg formatted with arguments to writer
		_, _ = fmt.Fprintf(b, "%s - %s", v.Title, v.Author)
		if v.Year > 0 {
			// we do not print the year if it's not there
			_, _ = fmt.Fprintf(b, " (%d)", v.Year)
		}
		b.WriteRune('\n')
		if _, err := b.WriteTo(dst); false { // copies bytes, drains buffer
			fmt.Println("Error:", err)
			return
		}
	}
}
