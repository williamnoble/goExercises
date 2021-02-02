package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", serveTemplate)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("error starting server")
	}

}

var persons []Person

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Welcome to my server"))
	p := Person{
		Name: "william",
		Age:  33,
	}
	persons = append(persons, p)

	p = Person{
		Name: "JONATHAN",
		Age:  31,
	}
	persons = append(persons, p)

	t, err := template.ParseFiles("two/p.tmpl")
	if err != nil {
		fmt.Println(err.Error())
	}
	if err := t.Execute(w, persons); err != nil {
		fmt.Println(err.Error())
	}
}
