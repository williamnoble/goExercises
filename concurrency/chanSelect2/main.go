package main

import "fmt"

import "time"

import "math/rand"

func main() {
	fmt.Println("*** Unbuffered double signal ***")
	signalAck()
	fmt.Println("*** Buffered: Close and range **")
	closeRange()
	fmt.Println("*** Unbuffered select w/ RECEIVE ***")
	selectRecv()
	fmt.Println(" *** Unbuffered select w/ SEND ***")
	selectSend()
	fmt.Println(" *** select drop***")
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
	// we need to close the channel or we panic as all goroutines are asleep
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Printf("-----------\n\n")
}

// *** Unbuffered select w/ RECEIVE ***
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
	fmt.Printf("-----------\n\n")
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
	fmt.Printf("-----------\n\n")
}

// SelectDROP
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
	fmt.Printf("-----------\n\n")
}
