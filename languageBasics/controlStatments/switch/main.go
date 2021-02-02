package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	s := []string{
		"statementOne",
		"statementTwo",
	}
	rand.Seed(time.Now().Unix())
	st := s[rand.Intn(len(s))]
	if err := switcher(st); err != nil {
		log.Fatal(err)
	}
}

func switcher(s string) error {
	switch s {
	case "statementOne":
		fmt.Println("Statement One")
		return nil
	case "statementTwo":
		fmt.Println("Statement Two")
		return nil
	default:
		fmt.Println("Input not found")
		return errors.New("err: invalid input")
	}
}
