package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type Movie struct {
	Title    string
	Director string
	Year     int
}

func ReadCSV(b io.Reader) ([]Movie, error) {
	r := csv.NewReader(b)
	r.Comma = ';'
	r.Comment = '-'
	var movies []Movie
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := Movie{record[0], record[1], int(year)}
		movies = append(movies, m)
	}
	return movies, nil
}

func AddMoviesFromTest() error {
	// String --> Buffer --> CSV
	in := `
	- first our headers
	movie title;director;year released
	- data
	Guardians of the Galaxy Vol. 2;James Gunn;2017
	Star Wars: Episode Z;Rian Johnson;2017
	Revenge of the Turtle;James Elder;2014
	`

	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}
	fmt.Printf("%#vn", m)
	return nil
}
