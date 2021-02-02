package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Result ONE"
	}()

	select {
	case r := <-c1:
		fmt.Println(r)
	case <-time.After(1 * time.Second):
		fmt.Println("1st: Timeout Out")
	}

	c2 := make(chan string, 2)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Result TWO"
	}()
	select {
	case r := <-c2:
		fmt.Println(r)
	case <-time.After(3 * time.Second):
		fmt.Println("2nd: Timed Out")
	}
}
