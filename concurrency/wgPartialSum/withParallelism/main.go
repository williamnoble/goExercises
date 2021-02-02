package main

import (
	"fmt"
	"sync"
)

const MAX = 1000
const workers = 2

func main() {
	values := make(chan int)
	result := make(chan int, workers)
	var wg sync.WaitGroup

	go func() {
		for i := 1; i < MAX; i++ {
			if (i%3) == 0 || (i%5) == 0 {
				values <- i
			}
		}
		close(values)
	}()

	work := func() {
		defer wg.Done()
		r := 0
		for i := range values {
			r += i
		}
		result <- r
	}

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go work()
	}

	wg.Wait()
	close(result)
	total := 0
	for pr := range result {
		total += pr
	}

	fmt.Println("Total:", total)
}
