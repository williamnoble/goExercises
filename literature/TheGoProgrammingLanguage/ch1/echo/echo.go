package main

import (
	"fmt"
	"os"
)

func main() {
	argString := ""
	separator := " "
	for _, arg := range os.Args[1:] {
		argString += separator + arg
	}
	fmt.Println(argString)
}
