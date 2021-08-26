package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func (w *Worker) process(c <-chan int) int {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		select {
		case c <- rand.Int():
		default:
			fmt.Println("dropped")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
