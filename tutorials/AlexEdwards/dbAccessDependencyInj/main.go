package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/williamnoble/Projects/xTuts/AlexEdwards/dbAccessDependencyInj/models"
)

type Env struct {
	db *sql.DB
}

func main() {
	db, err := models.NewDB("postgres://postgres:@DatabaseUser246@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}

	http.HandleFunc("/books", env.booksIndex)
	_ = http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := models.AllBooks(env.db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		_, _ = fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
