package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tmpl := template.New("index")
	tmpl, _ = tmpl.Parse("{{.Name}}")

	p := struct {
		Name string
	}{
		Name: "William",
	}

	err := tmpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatal(err)
	}
}
