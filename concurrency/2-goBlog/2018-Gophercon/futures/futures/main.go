package main

import (
	"fmt"
	"time"
)

// this is basically async/await
type Event struct{}
type Item struct {
	kind string
}

var items = map[string]Item{
	"a": Item{"gopher"},
	"b": Item{"rabbit"},
}

func doSlowThing() { time.Sleep(10 * time.Millisecond) }

func consume(a, b Item) {
	fmt.Println(a, b)
}

// Fetch immediately returns a channel, then fetches
// the requested item and sends it on the channel.
// If the item does not exist,
// Fetch closes the channel without sending.
func Fetch(name string) <-chan Item {
	c := make(chan Item, 1)
	go func() {
		item, ok := items[name]
		doSlowThing()
		if !ok {
			close(c)
			return
		}
		c <- item
	}()
	return c
}

func main() {
	start := time.Now()
	a := Fetch("a")
	b := Fetch("b")
	consume(<-a, <-b)
	fmt.Println(time.Since(start))
}
