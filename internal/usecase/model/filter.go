package model

import "time"

// [1,2,3,4,5,5]

type Filter struct {
	ID             int64      `db:"id"`
	MinAge         *int       `db:"min_age"`
	MaxAge         *[]int     `db:"max_age"`
	RegisterDate   *time.Time `db:"register_date"`
	LastActivity   *time.Time `db:"last_activity"`
	NotifyDatetime *time.Time `db:"notify_datetime"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
}
