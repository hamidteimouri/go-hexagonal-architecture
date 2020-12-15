package entities

import "time"

type Blog struct {
	ID          int64
	AuthorID    int64
	Title       string
	Description string
	Display     bool
	DeletedAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
