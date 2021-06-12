package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

type Location struct {
	Street  string
	ZipCode string
}

type User struct {
	Username  string
	Locations map[string]Location
}

type UserPage struct {
	Title string
	Users []User
}

func main() {
	message, err := ioutil.ReadFile("tpl.tmpl")
	if err != nil {
		panic(err)
	}

	t, err := template.New("UsersPage").Parse(string(message))
	if err != nil {
		panic(err)
	}

	p := UserPage{
		Title: "A Working Example of Templates",
		Users: []User{
			{§
				Username: "William",
				Locations: map[string]Location{
					"Home": {
						Street:  "GoLand",
						ZipCode: "2018.3",
					},
				},
			},
		},
	}

	err = t.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
