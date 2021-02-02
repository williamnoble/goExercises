package main

import "fmt"

type Employee interface {
	Language() string
	//Age() int
	//Random() (string, error)
}

type Engineer struct {
	Name string
}

func (e *Engineer) Language() string {
	return e.Name + " programs in Go"
}

func main() {
	// This will throw an error
	var programmers []Employee
	william := &Engineer{Name: "William"}
	jonathan := &Engineer{Name: "Jonathan"}
	ben := &Engineer{Name: "Benjamin"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, william, jonathan, ben)
	for i := 0; i < len(programmers); i++ {
		//fmt.Println(programmers[i])
		fmt.Printf("%s\n", programmers[i].Language())
	}
}
