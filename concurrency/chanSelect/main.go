package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "first call"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "second call"
	}()

	// this is ugly. We block forever until we have two loops
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// This is much better, we block until we give both time to complete with the final case.
	// If we remove the for { then we select one channel then finish.
	// With the for we block forever until ..
	//for {
	//	select {
	//	case msg1 := <-c1:
	//		fmt.Println("received", msg1)
	//	case msg2 := <-c2:
	//		fmt.Println("received", msg2)
	//	case <- time.After(4 * time.Second):
	//		return
	//	}
	//}
}
