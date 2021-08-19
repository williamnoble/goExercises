package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	// it's not necessary to declare a variable e.g p := &Person, return &p
	return &Person{
		name: name,
		age:  age,
	}
}

func main() {
	a := newPerson("William", 102)
	fmt.Printf("%s, %d", a.name, a.age)

}
