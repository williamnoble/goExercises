package main

import (
	"fmt"
	"strings"
)

var users = map[int]string{
	1: "Delila",
	2: "Rosemary",
	3: "Samantha",
	4: "Charles",
	5: "Will",
	6: "William",
}

func Search(subStr string) <-chan string {
	queue := make(chan string)
	go func() {
		defer close(queue)
		for _, user := range users {
			if assert := strings.Contains(user, subStr); !assert {
				continue
			}
			queue <- user
		}
	}()
	return queue
}

func main() {
	// for with range as we may expect more than one result. Will `will` match "Will" & "William"
	for u := range Search("Will") {
		fmt.Println(u)
	}

}
