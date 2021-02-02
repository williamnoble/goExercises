package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	integerList := []int{8, 42, 84, 39, 43, 24}
	go printCount(c) // start Channel
	for _, i := range integerList {
		c <- i
	}

	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of Main")
}

func printCount(c chan int) {
	n := 0
	for n >= 0 {
		n = <-c
		fmt.Print(n, " ")
	}
}
