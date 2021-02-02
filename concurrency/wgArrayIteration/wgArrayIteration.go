package main

import (
	"fmt"
	"sync"
)

func printer(msg string, wg *sync.WaitGroup) {
	wg.Done()
	fmt.Printf("Hello %s\n", msg)
}

func main() {
	greeks := []string{"plato", "socrates", "aristotle", "archimedes", "pythagoros", "democritus"}

	var wg sync.WaitGroup

	fmt.Println("Init")

	for _, philospoher := range greeks {
		wg.Add(1)
		go printer(philospoher, &wg)
	}

	wg.Wait()
	fmt.Println("End")
}
