package main

import (
	"fmt"
	"time"
)

var users = map[int]string{
	1: "Delila",
	2: "Rosemary",
}

func UserFromId(id int) <-chan string {
	c := make(chan string, 1)
	go func() {
		name, ok := users[id]
		time.Sleep(100 * time.Millisecond)
		if !ok {
			close(c)
			return
		}
		c <- name
	}()
	return c
}

func main() {
	s := time.Now()
	userOne := UserFromId(1)
	userTwo := UserFromId(2)
	fmt.Println(<-userOne, <-userTwo)

	fmt.Println(time.Since(s))
}
