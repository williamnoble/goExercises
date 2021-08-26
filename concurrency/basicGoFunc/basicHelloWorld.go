package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mu sync.WaitGroup

func sayHello(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)

	}
}

// Add a wait lock so it finishes
func main() {
	mu.Add(2)
	go sayHello("world")
	sayHello("Hello")
	mu.Wait()
	fmt.Println("end")

	// Alternatively we call time.Sleep, or we could also use Select {, which would block forever!
	//time.Sleep(1 * time.Second)
}
