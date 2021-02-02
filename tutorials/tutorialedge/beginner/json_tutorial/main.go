package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

func main() {
	// author := Author{Sales: 3, Age: 25, Developer: true}
	// book := Book{Title: "Learning Concurrency in Go", Author: author}
	book := Book{
		Title: "Learning Concurrency in Go and play with structure",
		Author: Author{
			Sales:     26,
			Age:       4,
			Developer: true,
		},
	}

	// byteArray, err := json.Marshal(book)
	byteArray, err := json.MarshalIndent(book, "--> ", "   ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}
