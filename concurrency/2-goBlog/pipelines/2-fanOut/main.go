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

func square(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range input {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Our Square function is returning a single <- Chan Int. We can have multiple goRoutines
// reading off this channel. We range over the main input "cs" and each value "c" is passed
// to one of our goRoutines which then does work this data.
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
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
	//queue := convertListToChanSend(2, 3)
	//squaredQueue := 1-square(queue)
	//fmt.Println(<-squaredQueue)
	//fmt.Println(<-squaredQueue)

	input := convertListToChanSend(2, 3)
	c1 := square(input)
	c2 := square(input)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}

}
