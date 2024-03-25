package model

import (
	"database/sql"
	"time"
)

type Directory struct {
	ID        int64        `db:"id"`
	Name      []int        `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

//Row.Scan(&directory.ID, &directory.Name, &directory.CreatedAt, &directory.UpdatedAt, &directory.DeletedAt)
