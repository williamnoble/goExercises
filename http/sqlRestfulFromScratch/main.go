package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func dbConn() (db *sql.DB) {
	dbUser := "root"
	dbPass := "root"
	//	dbPort := 3006
	dbName := "goblog"
	dbDriver := "mysql"
	dbURI := fmt.Sprintf("%s%s:%s@/%s", dbDriver, dbUser, dbPass, dbName)
	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		fmt.Println("Failed to connect")
	}
	return db

}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, _ *http.Request) {
	db := dbConn()
	rows, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	var employees []Employee
	for rows.Next() {
		var id int
		var name, city string
		err = rows.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		employees = append(employees, emp)
	}
	//goland:noinspection GoUnhandledErrorResult
	tmpl.ExecuteTemplate(w, "Index", employees)
	//goland:noinspection GoUnhandledErrorResult
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	_ = tmpl.ExecuteTemplate(w, "Show", emp)
	//goland:noinspection GoUnhandledErrorResult
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	_ = tmpl.ExecuteTemplate(w, "Edit", emp)
	//goland:noinspection GoUnhandledErrorResult
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO employee(name, city) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE employee SET name=?, city=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		//goland:noinspection GoUnhandledErrorResult
		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	//goland:noinspection GoUnhandledErrorResult
	delForm.Exec(emp)
	log.Println("DELETE")
	//goland:noinspection GoUnhandledErrorResult
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	//goland:noinspection GoUnhandledErrorResult
	http.ListenAndServe(":8080", nil)
}
