package main

import (
	"fmt"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(100 * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func main() {
	will := boring("Will")
	//jacob := boring("Jacob")

	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-will)
	//	fmt.Println(<-jacob)
	//}

	for msg := range will {
		fmt.Println(msg)
	}

	fmt.Println("end of fn")
}
