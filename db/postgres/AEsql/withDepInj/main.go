package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/williamnoble/goExercises/db/postgres/AEsql/withDepInj/models"
)

var db *sql.DB

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.NewDB("postgres://@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db}

	http.HandleFunc("/literature", env.booksIndex)
	http.ListenAndServe(":3000", nil)
}
func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := env.db.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
