package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Name string
}

var friendOne = Friend{
	Name: "Melanie",
}
var friendTwo = Friend{Name: "Jessica"}

type Person struct {
	Username string
	Emails   []string
	Friends  []*Friend
}

var t = template.New("index")
var tmpl, _ = t.Parse(`hello {{.Username}}!
            {{range .Emails}}
                an email {{.}}
            {{end}}
            {{with .Friends}}
            {{range .}}
                friend: {{.Name}}
            {{end}}
            {{end}}
            `)

func main() {

	p := Person{
		Username: "Frank",
		Emails:   []string{"frank@example.com", "frank2@example.com"},
		Friends:  []*Friend{&friendOne, &friendTwo},
	}

	_ = tmpl.Execute(os.Stdout, p)
}
