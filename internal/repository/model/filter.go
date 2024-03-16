package model

import "time"

type Filter struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	UserName    string    `db:"user_name"`
	BookName    string    `db:"book_name"`
	AuthorName  string    `db:"author_name"`
	StartAt     time.Time `db:"start_at"`
	NextAfter   int       `db:"next_after"`
	PromotionID int64     `db:"promotion_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}
