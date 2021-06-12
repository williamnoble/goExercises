package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("--------?---------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", 1)
		/*
			-1 = Replace infinitely
			0 = No replacement
			1 = Replace once
			2 = Repliace twice...

		*/
		if strings.Compare("hi", text) == 0 {
			fmt.Println(`A Generic "Hi" without user input`)
		}

		inputLength := len(text)
		if strings.Contains(text, "hello") {
			if len(text) == 5 {
				fmt.Println("A different form of Hello")
			} else {
				fmt.Println("Why Hello There " + text[6:inputLength])
			}
		}

	}
}
