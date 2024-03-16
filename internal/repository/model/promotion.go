package model

import "time"

type Promotion struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Active    bool      `db:"active"`
	Count     int8      `db:"count"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
