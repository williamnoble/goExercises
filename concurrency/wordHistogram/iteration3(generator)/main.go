package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink, and no it's not water",
		"The dark bird of prey is a bird which lands on a small inlet",
	}

	histogram := make(map[string]int)

	words := words(data)

	for word := range words {
		histogram[word]++
	}

	for k, v := range histogram {
		fmt.Printf("%s\t (%d) \n", k, v)
	}
}

// receive only channel
func words(data []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				out <- word
			}
		}
	}()
	return out
}
