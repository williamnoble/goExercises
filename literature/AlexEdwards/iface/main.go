package main

import (
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	Title  string
	Author string
}

type w interface {
	String() string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

func WriteLog(item w) {
	log.Println(item.String())
}

func main() {
	book := Book{"Alice in Wonderland", "Lewis Carol"}
	WriteLog(book)

	count := Count(394)
	WriteLog(count)
}
