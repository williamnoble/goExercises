package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("1: Recieved Message", msg)
	default:
		fmt.Println("1: No Message Recieved")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("2: Sent message", msg)
	default:
		fmt.Println("2: No Message Sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("3: Recieved Messsage", msg)
	case sig := <-signals:
		fmt.Println("3: Recieved Signal", sig)
	default:
		fmt.Println("3: No Activity")
	}
}
