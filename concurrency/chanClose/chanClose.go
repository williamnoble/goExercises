package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, ok := <-jobs // value, open?
			if ok {
				fmt.Println("Received Job", j)

			} else { // Block til Close(Jobs)
				fmt.Println("Received All Jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Send Job", j)
	}

	close(jobs)
	fmt.Println("Sent all Jobs")

	<-done
}
