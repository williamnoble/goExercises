package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	school string
}

func (s Student) String() string {
	str := strconv.Itoa(s.age)
	return fmt.Sprintf("%s %s at %s", s.name, str, s.school)
}

func main() {
	a := Student{
		Human: Human{
			"Jeffrey",
			23,
		},
		school: "The School of Life",
	}

	fmt.Printf("%s", a)
}
