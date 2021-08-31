package main

import (
	"fmt"
	"sync"
)

// This is an example of Fanning Out. Essentially multiple functions read from the same channel until it is closed.

func convertListToChanSend(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(done <-chan struct{}, input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range input {
			select {
			//out <- n * n
			case out <- n * n:
			case <-done:
				return
			}
		}
		close(out)
	}()
	return out
}

// Our Square function is returning a single <- Chan Int. We can have multiple goRoutines
// reading off this channel. We range over the main input "cs" and each value "c" is passed
// to one of our goRoutines which then does work this data.
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			// Whilst we still have a value from range c we pass this value to out otherwise we are done
			case out <- n:
			case <-done:
				return
			}
		}
		wg.Done()
	}
	// How many goroutines do we have? Do work with TWO input (<- chan int)
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Wait til both goRoutines have completed then close the out channel
	go func() {
		wg.Wait() // Ensure we have actually completed, lest send to closed chan panic!
		close(out)
	}()
	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)

	input := convertListToChanSend(2, 3)

	c1 := square(done, input)
	c2 := square(done, input)

	output := merge(done, c1, c2)
	fmt.Println(<-output)

}
