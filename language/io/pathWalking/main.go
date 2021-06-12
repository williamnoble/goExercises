package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var c struct {
	files int
	dirs  int
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please specify a path")
	}
	root, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println("Cannot get absolute path", err)
	}
	fmt.Println("Listing files in", root)

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			c.dirs++
		} else {
			c.files++
		}
		fmt.Println("-", path)
		return nil
	})
	fmt.Printf("Total: %d files in %d directoriess", c.files, c.dirs)
}
