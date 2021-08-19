package models

import "database/sql"

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func (db *DB) AllBooks() ([]*Book, error) {
	//goland:noinspection SqlResolve
	rows, err := db.Query("SELECT * FROM literature")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}
