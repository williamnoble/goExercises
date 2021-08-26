package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name      string
	Age       int
	Quotation string
}

type Book struct {
	Title  string
	Author Author
}

func main() {
	b := Book{
		Title: "Together we stand against hate",
		Author: Author{
			Name:      "Ben Braggs",
			Age:       27,
			Quotation: "Together we Stand. Together we Fall",
		},
	}

	bytes, _ := json.Marshal(b)
	bytes2, _ := json.MarshalIndent(b, "", "\t")
	bytes3 := append(bytes2, '\n') //prittyfy
	fmt.Println(string(bytes))
	fmt.Println("TWO")
	fmt.Println(string(bytes2))
	fmt.Println("THREE")
	fmt.Println(string(bytes3))
}
