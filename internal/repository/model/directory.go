package model

import "time"

type Directory struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	PromotionID int64     `db:"promotion_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
