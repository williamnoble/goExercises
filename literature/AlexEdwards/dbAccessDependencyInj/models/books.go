package models

import (
	"database/sql"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks(db *sql.DB) ([]*Book, error) {
	query := "SELECT * FROM literature"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer rows.Close()

	var books []*Book
	for rows.Next() {
		book := Book{}

		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)

		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
