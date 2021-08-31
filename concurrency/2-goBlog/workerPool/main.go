package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker #", id, "started and finished job: ", job)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("->Worker #", id, "finished job", job)
		results <- job * 2
	}
}

func main() {
	const numJobs = 5
	jobs, results := make(chan int, numJobs), make(chan int, numJobs)

	// 1. Create 3 Workers with IDs: 1,2,3.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 2. Create x Jobs for each worker, closing the Job chan once work is complete
	for job := 1; job <= numJobs; job++ {
		jobs <- job
	}
	close(jobs)

	// 3. Read from the results channel, looping depending on num of jobs created
	for a := 1; a <= numJobs; a++ {
		<-results
	}

}
