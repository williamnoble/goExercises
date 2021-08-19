package models

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks() ([]*Book, error) {
	//goland:noinspection SqlResolve
	rows, err := db.Query("SELECT * FROM literature")
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer rows.Close()

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
