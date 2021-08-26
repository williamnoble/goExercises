package main

import "fmt"

//   fROM THE CHAN   <- CHAN <- TO THE CHAN

func transfer(channelA chan<- string, msg string) {
	channelA <- msg
}

func pong(outputChannel chan<- string, channelA <-chan string) {
	msg := <-channelA
	outputChannel <- msg
}

func main() {
	channelA := make(chan string, 1)
	outputChannel := make(chan string, 1)

	// send a msg to channel A. chan <- [We are sending TO the chan].
	transfer(channelA, "Dear Channel A")

	// Get the msg from channelA and then send to output channel.
	// outputChannel chan <- [We are sending to the Chan and Receiving fROM chan A]
	pong(outputChannel, channelA)
	fmt.Println(<-outputChannel)

}
