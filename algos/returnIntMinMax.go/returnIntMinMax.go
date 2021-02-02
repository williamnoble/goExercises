package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x, y, z := 3, 4, 5
	maxXY := max(x, y)
	maxYZ := max(y, z)

	fmt.Printf("The Maximum of x(%d) and y(%d) is %d \n", x, y, maxXY)
	fmt.Printf("The Maximum of Y(%d) and Z(%d) is %d \n", y, z, maxYZ)
}
