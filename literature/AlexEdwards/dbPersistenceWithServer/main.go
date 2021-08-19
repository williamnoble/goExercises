package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

// curl -i -X POST -d $bookCreate "http://localhost:3000/books/create"

func main() {
	http.HandleFunc("/literature", booksIndex)
	http.HandleFunc("/literature/show", booksShow)
	http.HandleFunc("/literature/create", booksCreate)
	_ = http.ListenAndServe(":3000", nil)
}

//goland:noinspection SqlResolve
func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM literature")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer rows.Close()
	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		_, _ = fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

//goland:noinspection SqlResolve
func booksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), 405)
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), 400)
		return
	}

	isbn2 := r.URL.Query().Get("isbn")
	fmt.Println(isbn)
	fmt.Println(isbn2)

	row := db.QueryRow("SELECT * FROM literature WHERE isbn = $1", isbn)

	bk := new(Book)
	err := row.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	_, _ = fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
}

func booksCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), 405)
		return
	}

	isbn := r.FormValue("isbn")
	title := r.FormValue("title")
	author := r.FormValue("author")

	if isbn == "" || title == "" || author == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), 400)
		return
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 32)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), 400)
		return
	}

	//goland:noinspection SqlResolve
	result, err := db.Exec("INSERT INTO literature VALUES($1, $2, $3, $4)", isbn, title, author, price)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}

	fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", isbn, rowsAffected)
}
