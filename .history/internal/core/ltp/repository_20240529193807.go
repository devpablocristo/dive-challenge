package ltp

import (
	"context"
)

type Repository struct {
	db *InMemDB
}

func NewRepository() RepositoryPort {
	db := make(InMemDB)
	return &Repository{
		db: &db,
	}
}

func (r *Repository) GetLTP(ctx context.Context, ID string) error {
	return nil
}
