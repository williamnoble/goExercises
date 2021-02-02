package main

import (
	"os"
	"strconv"
)

func main() {
	a := os.Args
	sum := 0
	for i := 0; i < len(a); i++ {
		v, _ := strconv.Atoi(a[i])
		sum += v
	}
	println("sum: ", sum)
}
