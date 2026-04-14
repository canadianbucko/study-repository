package simple_sql

import "time"

type BookModel struct {
	ID          int
	Title       string
	Author      string
	Review      string
	ReleaseDate int
	WasRead     bool
	AddedAt     time.Time
	ReadAt      *time.Time
}
