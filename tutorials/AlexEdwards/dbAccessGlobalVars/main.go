package main

import (
	"fmt"
	"net/http"

	"github.com/williamnoble/Projects/xTuts/AlexEdwards/dbAccessGlobalVars/models"
)

func main() {
	models.InitDB("postgres://postgres:@DatabaseUser246@localhost/bookstore?sslmode=disable")

	http.HandleFunc("/books", booksIndex)
	_ = http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := models.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}

}
