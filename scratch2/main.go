package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var s []string
	s = append(s, "statementOne", "statementTwo")
	rand.Seed(time.Now().Unix())

	sel := make(map[string]int)
	sel["one"] = 0
	sel["two"] = 0
	for i := 0; i < 80; i ++ {
		choice := s[rand.Intn(len(s))]
		if choice == "statementOne" {
			sel["one"] += 1
		} else {
			sel["two"] += 1
		}
		if err := switcher(choice); err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Printf("\n1st Statement: \t%d\n2nd Statment: \t%d\nTotal: \t\t%d",
		sel["one"], sel["two"], sel["one"]+sel["two"])
}

func switcher(s string) error {
	switch s {
	case "statementOne":
		fmt.Printf("1")
		return nil
	case "statementTwo":
		fmt.Printf("2")
		return nil
	default:
		fmt.Println("")
		return errors.New("err: invalid input")
	}
}