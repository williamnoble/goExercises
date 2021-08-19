package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Username string
}

func main() {
	t := template.New("example")
	t, _ = t.Parse("{{.Username}}")
	p := &Person{Username: "William"}
	err := t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println("encountered an error")
	}
}
