package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//1. fmt.Scanln(&input)
	//2. reader.ReadString()
	//3. scanner.Scan()
	//4. scanner.Scan() without a for loop

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("ONE: scanning standard input, q to quit...")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "q" {
			break
		}
		fmt.Println("1: ", text)
	}

	fmt.Println("TWO: Enter some text")
	var input string
	fmt.Scanln(&input)
	fmt.Println("2: ", input)

	// This and the way below it are the correct way to correct input
	func() {
		fmt.Println("Third Times: Enter some text..")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println("Third Times: Enter some text..")
			fmt.Println("3: ", scanner.Text())
			if scanner.Text() == "q" {
				break
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)

	fmt.Println("**** ex ****")
	// https://stackoverflow.com/questions/20895552/how-can-i-read-from-standard-input-in-the-console
	arr := make([]string, 0)
	s2 := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Text: ")
		// Scans a line from Stdin(Console)
		s2.Scan()
		// Holds the string that scanned
		text := s2.Text()
		if len(text) != 0 {
			fmt.Println(text)
			arr = append(arr, text)
		} else {
			break
		}

	}
	// Use collected inputs
	fmt.Println(arr)

}
