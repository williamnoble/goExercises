package db

import (
	"context"
	"database/sql"

	"github.com/williamnoble/goExercises/scratch/meo/schema"
)

type PostgresRespository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRespository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRespository{
		db,
	}, nil

}

func (r *PostgresRespository) Close() {
	r.db.Close()
}

func (r *PostgresRespository) InsertMeow(ctx context.Context, meow schema.Meow) error {
	stmt := "INSERT INTO meows(id, body, created_at) VALUES($1, $2, $3)"
	_, err := r.db.Exec(stmt, meow.ID, meow.Body, meow.CreatedAt)
	return err
}

func (r *PostgresRespository) ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error) {
	stmt := "SELECT * FROM meows ORDER BY id DESC OFFSET $1 LIMIT $2"
	rows, err := r.db.Query(stmt, skip, take)
	if err != nil{
		return nil, error
	}

	defer rows.Close()
	meows := []schema.Meow{}
	for rows.Next {
	if err = rows.Scan(&meow.ID, &meow.Body, &meow.CreatedAt); err == nil {
		meows = append(meows, meow)
	  }
	}
	if err = rows.Err(); err != nil {
	  return nil, err
	}
  
	return meows, nil
  }
}