package main

import "fmt"

func main() {
	// declare empty array


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

}
