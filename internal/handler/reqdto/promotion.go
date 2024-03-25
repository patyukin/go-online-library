package reqdto

import (
	"time"
)

type Promotion struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Comment     *string     `json:"comment"`
	Status      bool        `json:"active"`
	Type        string      `json:"type"`
	Filters     []Filter    `json:"filters"`
	Directories []Directory `json:"directories"`
}

type Filter struct {
	ID             int64      `json:"id"`
	MinAge         *int       `json:"min_age"`
	MaxAge         *int       `json:"max_age"`
	RegisterDate   *time.Time `json:"register_date"`
	LastActivity   *time.Time `db:"last_activity"`
	NotifyDatetime *time.Time `db:"notify_datetime"`
}

type Directory struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type Notify struct {
	ID          int64     `json:"id"`
	PromotionID int64     `json:"promotion_id"`
	NotifyTime  time.Time `json:"notify_time"`
}
