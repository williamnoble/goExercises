package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}

}
func main() {
	url := os.Args[1]
	data, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return

		_, _ = io.Copy(os.Stdout, data.Body)
		if err := data.Body.Close(); err != nil {
			fmt.Println(err)

		}
	}
