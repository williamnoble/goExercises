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

	count := make(map[string]int)
	count["one"] = 0
	count["two"] = 0
	for i := 0; i < 80; i++ {
		choice := s[rand.Intn(len(s))]
		if choice == "statementOne" {
			count["one"] += 1
		} else {
			count["two"] += 1
		}
		if err := switcher(choice); err != nil {
			log.Fatal(err.Error())
		}
	}
	fmt.Println("**TABLE**")
	total := count["one"] + count["two"]
	fmt.Printf("1st:\t%d\n2nd:\t%d\nTotal:\t%d",
		count["one"], count["two"], total)
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
