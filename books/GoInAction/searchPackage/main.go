package main

import (
	"log"
	"os"

	"github.com/williamnoble/goExercises/books/GoInAction/searchPackage/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
