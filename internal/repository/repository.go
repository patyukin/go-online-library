package repository

import (
	"github.com/patyukin/go-online-library/pkg/db"
)

type Repository struct {
	db db.Client
}

func New(db db.Client) *Repository {
	return &Repository{
		db: db,
	}
}
