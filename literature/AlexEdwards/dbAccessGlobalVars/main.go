package main

import (
	"fmt"
	models2 "github.com/williamnoble/goExercises/literature/AlexEdwards/dbAccessGlobalVars/models"
	"net/http"
)

func main() {
	models2.InitDB("postgres://postgres:@/bookstore?sslmode=disable")

	http.HandleFunc("/literature", booksIndex)
	_ = http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := models2.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		_, _ = fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}

}
