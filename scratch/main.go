package main

import (
	"bytes"
	"fmt"
	"os"
)

type Book struct {
	author, title string
	year          int
}

func main() {
	dst, err := os.OpenFile("./book.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer dst.Close()

	bookList := []Book{
		{
			author: "William Noble",
			title:  "The Art of Typing",
			year:   2020,
		},
		{
			author: "Billbo Baggins",
			title:  "The Art of Typography",
			year:   2019,
		},
		{
			author: "Jeffrey Rose",
			title:  "Roses are blue and so are you",
			year:   1880,
		},
	}

	b := new(bytes.Buffer)
	for _, book := range bookList {
		// fmt.Fprintf(b, "%s - %s - %d \n", book.title, book.author, book.year)
		fmt.Fprintf(b, "%s - %s - %d", book.title, book.author, book.year)
		b.WriteRune('\n')
	}
	if _, err := b.WriteTo(dst); err != nil {
		fmt.Println("Error", err)
		return
	}
}
