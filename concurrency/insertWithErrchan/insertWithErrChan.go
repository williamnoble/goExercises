// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/jT4-vZBpMm
// This sample program demonstrates how to use a buffered
// channel to receive results from other goroutines in a guaranteed way.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const numInserts = 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("main begin: Inserts Started")
	// Buffered channel to receive information about any possible insert.
	ch := make(chan error, numInserts) // 10
	// Number of responses we need to handle.
	var waitResponses int
	for i := 0; i < numInserts; i++ {
		if isNecessary() { // random: Int(0,2) -> true or false?
			waitResponses++
			go func(id int) {
				ch <- insertDoc(id)
			}(i)
		}
	}
	// Process the insert results as they complete.
	for {
		// Wait for a response from a goroutine.
		err := <-ch
		// Display the result.
		if err != nil {
			log.Println("Received error:", err)
		} else {
			log.Println("Received nil error")
		}
		// Decrement the wait count and determine if we are done.
		waitResponses--
		if waitResponses == 0 {
			break
		}
	}
	log.Println("Inserts Complete")
}

func insertDoc(id int) error {
	log.Println("Insert document: ", id)

	// Randomize if the insert fails or not.
	if rand.Intn(2) == 0 {
		return fmt.Errorf("Document ID: %d", id)
	}

	return nil
}

func isNecessary() bool {
	// Randomize if this insert is necessary or not.
	if rand.Intn(2) == 0 {
		return true
	}

	return false
}
