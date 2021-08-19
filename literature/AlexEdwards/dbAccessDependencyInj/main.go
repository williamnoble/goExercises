package main

import (
	"database/sql"
	"fmt"
	models2 "github.com/williamnoble/goExercises/literature/AlexEdwards/dbAccessDependencyInj/models"
	"log"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func main() {
	db, err := models2.NewDB("postgres://postgres@/bookstore?sslmode=disable")
	if err != nil {
		log.Panic("panic!", err)
	}
	env := &Env{db: db}

	http.HandleFunc("/literature", env.booksIndex)
	err = http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		fmt.Println("Fatal Err :(")
	}
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connected..")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := models2.AllBooks(env.db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		_, _ = fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
