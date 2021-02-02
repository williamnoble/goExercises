package main

import "fmt"

type Helloer interface {
	Hello(string)
}

type Greeting string

func (g Greeting) Hello(name string) {
	fmt.Println(g+",", name)
}

type Invitation struct {
	event string
}

func (in Invitation) Hello(name string) {
	fmt.Printf("Welcome to my %s, %s", in.event, name)
}
func main() {
	var h Helloer
	h = Greeting("Hello")
	h.Hello("Gopher")

	h = Invitation{event: "birthday party"}
	h.Hello("Kitty")
}
