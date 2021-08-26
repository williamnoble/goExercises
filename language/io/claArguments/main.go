package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	fmt.Println(calcArgs(args...))
}

// spread a variable number of args. We can pass zero or more strings. Note the output of os.Args is []string
// thus we need to spread in our call to calcArgs(args...) vs calcArgs(args).
func calcArgs(args ...string) int {
	sum := 0

	for i := range args {
		strconv.Atoi(args[i])
		sum += i
	}

	return sum
}
