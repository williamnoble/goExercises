package main

import "fmt"

func main() {

	c := make(chan string, 2)

	c <- "hello"
	c <- "William"

	close(c)
	// fine to read off a closed channel
	fmt.Println(<-c)
	fmt.Println(<-c)

	// this will panic!
	//fmt.Println(<-c)

	// safer to use a range, short for x, ok := range c, will not panic if there is nothing to read!
	for x := range c {
		fmt.Println(x)
	}
}
