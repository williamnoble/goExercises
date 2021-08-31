package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	rand.Seed(time.Now().UnixNano())
	go boring("boring", c)
	for i := 0; i < rand.Intn(10); i++ {
		fmt.Printf("You Say: %q\n", <-c)
	}

	fmt.Println("")
}
