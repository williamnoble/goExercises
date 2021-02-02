package main

import "fmt"

import "time"

import "math/rand"

func main() {
	signalAck()
	closeRange()
	selectRecv()
	selectSend()
	selectDrop()
}

// *** Unbuffered double signal ***
func signalAck() {
	ch := make(chan string)

	go func() {
		fmt.Println(<-ch)
		ch <- "ok, job done"
	}()

	ch <- "do this first"

	// blocked on recieve ("ok, job done")
	fmt.Println(<-ch)
}

// *** Buffered: Close and range **
func closeRange() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

// *** Unbuffered select w/ RECIEVE ***
func selectRecv() {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "SelectRecv: work"
	}()

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timed out")
	}
}

// *** Unbuffered select w/ SEND ***
func selectSend() {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		fmt.Println(<-ch)
	}()

	select {
	case ch <- "selectSend: work":
		fmt.Println("selectSend: Send Work")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}
}

func selectDrop() {
	ch := make(chan int, 5)

	go func() {
		for v := range ch {
			fmt.Println("recv", v)
		}
	}()

	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
			fmt.Println("sent work:", i)
		default:
			fmt.Println("drop", i)
		}
	}
}
