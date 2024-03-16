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

func (r *Repository) InsertDirectory() db.Client {
	return r.db
}

// PUT users/1/profile/1
// DELETE users/1/profile/1

// car  users
