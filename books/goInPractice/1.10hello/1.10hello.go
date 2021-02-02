package main

import "fmt"

func getName() string {
	return "world"
}

func main() {
	name := getName()
	fmt.Println("Hello ", name)
}
