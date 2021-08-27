package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//User Exposes a User to the GORM
type User struct {
	gorm.Model
	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "All Users endpoint hit")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to DB")
	}

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)
	//goland:noinspection GoUnhandledErrorResult
	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")
	}

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	fmt.Println(name)
	fmt.Println(email)

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Delete User endpoint hit")
}

//
//func updateUser(w http.ResponseWriter, _ *http.Request) {
//	fmt.Fprintf(w, "Update User endpoint hit")
//}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func initialMigrations() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")
	}
	//goland:noinspection GoUnhandledErrorResult
	db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("Go ORM Tutorial")
	initialMigrations()
	handleRequests()
}
