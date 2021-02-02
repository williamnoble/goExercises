package main

import "fmt"

func main() {
	a, b := Names()
	fmt.Println(a, b)
}

func Names() (string, string) {
	return "Foo", "Bar"
}
