package main

import "fmt"

var name string

func init() {
	fmt.Println("Initiazing... ")
	name = "William"
}

func main() {
	fmt.Println("I am the main function")
	fmt.Println("Hello " + name)
}
