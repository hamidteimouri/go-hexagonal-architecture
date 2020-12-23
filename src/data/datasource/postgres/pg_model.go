package postgres

import "time"

type PgBlogModel struct {
	ID          int64 `gorm:"primaryKey"`
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
