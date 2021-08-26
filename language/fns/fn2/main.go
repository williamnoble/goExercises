package main

import "fmt"

func main() {
	a := func() {
		AnonymousFn(printHelloWorld)
	}
	a()

	AnonymousFuncWithInput("Zebras are real")
}

func AnonymousFn(f func()) {
	fmt.Println("Calling f")
	f()
	fmt.Println("Called f")
}

func printHelloWorld() {
	fmt.Println("Hello World")
}

func AnonymousFuncWithInput(s string) {
	func() {
		fmt.Println(s)
	}()
}
