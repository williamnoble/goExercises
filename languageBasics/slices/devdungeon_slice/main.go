package main

import "fmt"

func main() {

	numRows := 10

	// Initialize a ten length slice of empty slices
	grid := make([][]int, numRows)

	// Verify it is a slice of ten empty slices
	fmt.Println(grid)

	// Initialize those 10 empty slices
	for i := 0; i < numRows; i++ {
		grid[i] = make([]int, 4)
	}

	// grid is a 2d slice of ints with dimensions 10x4
	fmt.Println(grid)
}
