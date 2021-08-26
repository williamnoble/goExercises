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

func serveTemplate(w http.ResponseWriter, _ *http.Request) {
	var persons []Person

	alice := Person{"Alice", 21}
	bob := Person{"Bob", 67}
	persons = append(persons, alice, bob)

	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(t.Name())
	w.Header().Set("Content-Type", "text/html")
	// t.Execure passes the io.Writer and data interface{}: persons
	if err := t.Execute(w, persons); err != nil {
		fmt.Println(err.Error())
	}
}
