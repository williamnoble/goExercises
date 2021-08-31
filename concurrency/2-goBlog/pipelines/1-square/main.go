package main

import "fmt"

// Construct a pipeline.
// 1. Gen convert a []int -> chan int which emits integers in a list
// We return a channel which is a send Channel i.e. <- chan
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

// We are squaring the numbers which come in and returning them to a channel
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

func main() {
	queue := convertListToChanSend(2, 3)
	squaredQueue := square(queue)
	fmt.Println(<-squaredQueue)
	fmt.Println(<-squaredQueue)
}
