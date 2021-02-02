package main

import "fmt"

import "time"

func main() {
	fmt.Printf("\n=> Basics of send and receive")
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
	<-ch
	fmt.Println("event recieved")

}
