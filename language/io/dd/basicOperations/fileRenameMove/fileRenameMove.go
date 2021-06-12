package main

import "os"

import "log"

func main() {
	origin := "test.txt"
	newPath := "mover/testing.txt"
	err := os.Rename(origin, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
