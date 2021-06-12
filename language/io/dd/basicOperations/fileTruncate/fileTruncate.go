package main

import (
	"log"
	"os"
)

var (
	err error
)

func main() {
	err = os.Truncate("truncateMe.txt", 10)
	if err != nil {
		log.Fatal(err)
	}
}
