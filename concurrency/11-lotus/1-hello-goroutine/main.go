package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go hello("boring")
	fmt.Println("I am listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You are boring, i'm leaving")
}
