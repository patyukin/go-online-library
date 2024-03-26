package model

import (
	"database/sql"
	"gopkg.in/guregu/null.v3"
	"time"
)

type Promotion struct {
	ID          int64        `db:"id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	Comment     null.String  `db:"comment"`
	Status      bool         `db:"status"`
	Type        string       `db:"type"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type PromotionFilter struct {
	PromotionID int64 `db:"promotion_id"`
	FilterID    int64 `db:"filter_id"`
}

type PromotionDirectory struct {
	PromotionID int64 `db:"promotion_id"`
	DirectoryID int64 `db:"directory_id"`
}
