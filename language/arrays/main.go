package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//TODO refactor into multiple functions, add sum func
	var x [3]int
	x[0] = 0
	x[1] = 10
	x[2] = 20
	total := 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Printf("Arr: %v has total: %d\n", x, total)

	total = 0

	// declare array literal
	y := [4]int{46, 67, 86, 32}
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Printf("Arr: %v has total: %d\n", y, total)

	// use elipsis to declare array
	var z = [...]int{1, 2, 3}
	_ = z

	structArray := [...]struct {
		name string
		age  int
	}{
		{"Isabel", 34},
		{"Jessica", 21},
	}

	var stringArray = [2]string{0: "Hello", 1: "Goodbyte ;)"}
	_, _ = stringArray, structArray

	rand.Seed(time.Now().UnixNano())

	var b [6]int
	for i := 0; i < len(b); i++ {

		b[i] = rand.Intn(1000)
	}
	sum := sumArray(b)
	fmt.Printf("sum: %d\n", sum)
}

func sumArray(a [6]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}
