package main

import (
	"log"
	"os"

	"github.com/williamnoble/Projects/books/GoInAction/searchPackage/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
