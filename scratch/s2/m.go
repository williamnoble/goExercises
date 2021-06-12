package main

import "fmt"

type Employee interface {
	Language() string // Lang fn which returns a string
}

type Developer struct {
	Name string
}

func (d *Developer) Language() string {
	return "the developer " + d.Name + " programs in go"
}

func main() {
	a := &Developer{"Tobie"}
	b := &Developer{"Jessica"}
	c := &Developer{"Bert"}

	Developers := []Employee{a, b, c}
	for i := 0; i < len(Developers); i++ {
		fmt.Printf("%s\n", Developers[i].Language())
	}
}
