package main

import "fmt"

func main() {
	AnonymousCallingFn(PrintHelloWorld)
	AnotherAnonymousFunction("Delta")
}

func AnonymousCallingFn(f func()) {
	fmt.Println("Fn1 Called")
	f()
}

func PrintHelloWorld() {
	fmt.Println("Hello World")
}

func AnotherAnonymousFunction(x string) {
	func() {
		fmt.Println(x)
	}()
}
