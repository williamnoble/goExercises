package db

import (
	"context"

	"github.com/williamnoble/goExercises/scratch/meo/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error)
}

var impl Repository

func SetRespository(r Repository) {
	impl = r
}

func Close() {
	impl.Close()
}

func InsertMeow(ctx context.Context, meow schema.Meow) error {
	return impl.InsertMeow(ctx, meow)
}

func ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error) {
	return impl.ListMeows(ctx, skip, take)
}
