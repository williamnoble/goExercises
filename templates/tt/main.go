package main

import "html/template"

type Name struct {
	Username string
}

func main() {
	t := template.New("example")
	t, err := t.Parse("{{.Username}}")
	p := &Person{Username: "William"}
	err := t.Execute(os.Stdout, t)
}
