package main

import (
	"fmt"
	"runtime"
)

func sayHello(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)

	}
}

func main() {
	go sayHello("world")
	sayHello("hello")
}
