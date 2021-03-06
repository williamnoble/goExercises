package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	fmt.Println("Starting Server")
	db, err := sql.Open("postgres", "postgres://@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully pinged")

	//goland:noinspection SqlResolve
	rows, err := db.Query("SELECT * FROM literature")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	books := make([]*Book, 0)
	for rows.Next() {
		book := new(Book)
		err = rows.Scan(&book.isbn, &book.title, &book.author, &book.price)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, book := range books {
		fmt.Printf("%s, %s, %s, £%.2f\n", book.isbn, book.author, book.title, book.price)
	}
}
