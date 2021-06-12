package main

import (
	"bytes"
	"fmt"
	"os"
)

type book struct {
	Author string
	Title  string
	Year   int
}

func main() {
	dst, err := os.OpenFile("bookList.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func() {
		if err := dst.Close(); err != nil {
			panic(err.Error())
		}
	}()

	bookList :=
		[]book{
			{
				Author: "Freddy Fred",
				Title:  "Water is negative space",
				Year:   2020,
			},
			{
				Author: "Betting Betty",
				Title:  "Cooking for miracle workers",
				Year:   1990,
			},
			{
				Author: "Shady Shack",
				Title:  "Shake Shack, who needs veg",
				Year:   1980,
			},
			{
				Author: "Crimson Carmel",
				Title:  "Yes, I am a dessert",
				Year:   1846,
			},
		}

	b := bytes.NewBuffer(make([]byte, 0, 16))
	for _, v := range bookList {
		if v.Year <= 0 {
			return
		}
		_, _ = fmt.Fprintf(b, "%s - %s - %d", v.Title, v.Author, v.Year)
		// if v.Year > 0 {
		// 	_, _ = fmt.Fprintf(b, " (%d)", v.Year)
		// }
		b.WriteRune('\n')
		if _, err := b.WriteTo(dst); err != nil {
			fmt.Println("error", err)
			return
		}
	}
}
