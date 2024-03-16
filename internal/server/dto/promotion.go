package dto

import "time"

type Promotion struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	UserName    string      `json:"userName"`
	BookName    string      `json:"bookName"`
	AuthorName  string      `json:"authorName"`
	Active      bool        `json:"active"`
	Count       int8        `json:"count"`
	StartAt     time.Time   `json:"startAt"`
	Filters     []Filter    `json:"filters"`
	Directories []Directory `json:"directories"`
}

type Filter struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	UserName    string    `json:"userName"`
	BookName    string    `json:"bookName"`
	AuthorName  string    `json:"authorName"`
	StartAt     time.Time `db:"start_at"`
	NextAfter   int       `db:"next_after"`
	PromotionID int64     `db:"promotion_id"`
}

type Directory struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	PromotionID int64     `db:"promotion_id"`
}
