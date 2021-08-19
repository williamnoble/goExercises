package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name string
}

func main() {
	t := template.New("test")
	t, err := t.Parse("{{.Name}}")
	p := Person{Name: "William"}
	out := t.Execute(os.Stdout, p)
	if out != nil {
		panic(err)
	}
}
