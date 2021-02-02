package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Println("Why Hello/Hi there my strange friend.. said", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Printf("La la la la.. hit the chorus %q\n", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s, Call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func (h Human) String() string {
	return "Name: " + h.name + ", Age: " + strconv.Itoa(h.age)
}
func main() {
	mike := Student{Human{"Mike", 25, "01234 333444"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "01234 453444"}, "Harvard", 100.00}
	sam := Employee{Human{"Sam", 36, "02834 334534"}, "GoLang Inc", 1000}
	tom := Employee{Human{"Tom", 26, "01234 948493"}, "Things Ltd", 9000}

	var i Men
	h := Human{"William", 32, "Junior Developer"}
	fmt.Println("Human String(): ", h)
	i = mike
	fmt.Println("This is Mike, he's a student")
	i.SayHi()
	i.Sing("Ohhh November Rain!")

	i = sam
	fmt.Println("Hey! I'm Sam, i'm a loyal Employee")
	i.SayHi()
	i.Sing("Ohh there's no slave labour here... cough")

	fmt.Println("A slice of men interfaces...")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, tom
	for _, value := range x {
		value.SayHi()
	}
}
