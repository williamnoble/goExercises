package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

type Event struct{}
type Item struct {
	kind string
}

var items = map[string]Item{
	"a": Item{"gopher"},
	"b": Item{"rabbit"},
}

func doSlowThing() { time.Sleep(10 * time.Millisecond) }

// Fetch immediately returns, then fetches the item and
// invokes f in a goroutine when the item is available.
// If the item does not exist,
// Fetch invokes f on the zero Item.
func Fetch(name string, f func(Item)) {
	go func() {
		item := items[name]
		doSlowThing()
		f(item)
	}()
}

func main() {
	start := time.Now()

	n := int32(0)
	Fetch("a", func(i Item) {
		fmt.Println(i)
		if atomic.AddInt32(&n, 1) == 2 {
			fmt.Println(time.Since(start))
			os.Exit(0)
		}
	})
	Fetch("b", func(i Item) {
		fmt.Println(i)
		if atomic.AddInt32(&n, 1) == 2 {
			fmt.Println(time.Since(start))
			os.Exit(0)
		}
	})

	select {}
}
