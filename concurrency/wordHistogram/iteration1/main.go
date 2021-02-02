package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink, and no it's not water .....",
		"The dark bird of prey is a bird which lands on a small inlet ...",
	}

	histogram := make(map[string]int)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for _, vline := range data {
			words := strings.Split(vline, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				if word == "..." || word == "...." || word == "....." {
					break
				} else {
					histogram[word]++
				}

			}
		}
	}()

	<-done
	for k, v := range histogram {
		fmt.Printf("%s\t (%d)\n", k, v)
	}

}
