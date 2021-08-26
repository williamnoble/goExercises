package main

import "fmt"

func main() {
	jobs := make(chan string, 5) // buffer of 5 jobs, to which we send 15 jobs!
	done := make(chan bool)      // signal

	go func() {
		for {
			job, open := <-jobs
			if open {
				fmt.Println("Job Received okay!, JOB: ", job)
			} else {
				fmt.Println("All Jobs Received")
				done <- true // event signal
				return
			}
		}
	}()

	for x := 1; x <= 15; x++ {
		job := fmt.Sprintf("Job ID #%d", x)
		jobs <- job
	}

	close(jobs)
	// Note, this line will _probably_ print before all jobs have printed to stdout.
	// The goroutine will not panic as we have a go routine which checks job, open -> we need a wg!
	fmt.Println("All jobs have been sent, now we block on \"done\" channel")

	<-done
	fmt.Println("Done blocked as it was meant to and now main finishes")
}
