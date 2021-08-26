package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() {
		messages <- "I'm a go Chan"
	}()
	//block til we recieve
	msg := <-messages
	fmt.Println(msg)
}
