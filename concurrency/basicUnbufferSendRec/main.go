package main

import "fmt"

import "time"

func main() {
	fmt.Printf("\n=> Basics send")
	basicSendRecv()

	fmt.Printf("\n=> Close a channel to signal an event\n")
	signalClose()
}

func basicSendRecv() {
	ch := make(chan string)
	go func() {
		ch <- "hello"
	}()
	fmt.Println(<-ch)

}

func signalClose() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("signal event")
		close(ch)
	}()

	// docs: After the last value has been received from a closed channel c,
	// any receive from c will succeed without blocking,
	// returning the zero value for the channel element
	x := <-ch
	fmt.Println(x)
	fmt.Println("event received")

}
