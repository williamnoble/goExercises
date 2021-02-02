package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var store BookmarkStore
var id string

func init() {
	session, err := mgo.Dial("mongodb://mongo:Mongo246@localhost:27017")
	// session, err := mgo.Dial("mongodb://localhost:27100")
	if err != nil {
		log.Fatalf("MongoDB Session: %s\n", err)
	}
	collection := session.DB("bookmarkdb").C("bookmarks")
	store = BookmarkStore{
		C: collection,
	}
}

func createUpdate() {
	bookmark := Bookmark{
		Name:        "mgo",
		Description: "Go driver for MongoDb",
		Location:    "https://www.github.com/go-mgo/mgo",
		Priority:    2,
		CreatedOn:   time.Now(),
		Tags:        []string{"go", "mysql", "mongodb"},
	}

	if err := store.Create(&bookmark); err != nil {
		log.Fatalf("Create: %s\n", err)
	}
	id = bookmark.ID.Hex()
	fmt.Printf("New bookmark created with ID: %s\n", id)

	bookmark.Priority = 1
	if err := store.Update(bookmark); err != nil {
		log.Fatalf("Update: %s\n", err)
	}
	fmt.Println("The value after update: ")
	getByID(id)
	bookmark = Bookmark{
		Name:        "gorethink",
		Description: "Go Rethink This, Bro",
		Location:    "http://rethinktheweb.com",
		Priority:    3,
		CreatedOn:   time.Now(),
		Tags:        []string{"go", "nosql", "rethinkdb"},
	}
	if err := store.Create(&bookmark); err != nil {
		log.Fatalf("Create: %s\n", err)
	}
	id = bookmark.ID.Hex()
	fmt.Printf("new bookbook inserted with id: %s\n", id)
}

func getByID(id string) {
	bookmark, err := store.GetByID(id)
	if err != nil {
		log.Fatalf("Get by Id: %s\n", err)
	}
	fmt.Printf("Name: %s, Description: %s, Priority: %d\n", bookmark.Name, bookmark.Description, bookmark.Priority)
}
func main() {
	createUpdate()
}
