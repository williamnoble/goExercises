package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool) // true | false
	go worker(done)

	fmt.Println("Blocking til <- done")
	<-done
	fmt.Println("Happy to FInish")

}
