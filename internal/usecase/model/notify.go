package model

import "time"

type Notify struct {
	ID          int64     `db:"id"`
	NotifyTime  time.Time `db:"notify_time"`
	PromotionID int64     `db:"promotion_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}
