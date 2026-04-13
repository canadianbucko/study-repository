package sample_struct

import "time"

type Book struct {
	Title       string
	Author      string
	Review      string
	ReleaseDate int
	WasRead     bool
	AddedAt     time.Time
	ReadAt      *time.Time
}

func NewBook(title string, author string, review string, releaseDate int, wasRead bool, readAt *time.Time) Book {
	return Book{
		Title:       title,
		Author:      author,
		Review:      review,
		ReleaseDate: releaseDate,
		WasRead:     wasRead,
		AddedAt:     time.Now(),
		ReadAt:      readAt,
	}
}

// CREATE TABLE IF NOT EXISTS books(
// 	id SERIAL PRIMARY KEY,
// 	title VARCHAR(100) NOT NULL,
// 	author VARCHAR(100) NOT NULL,
// 	review VARCHAR(1000),
// 	release_date INTEGER NOT NULL,
// 	was_read BOOLEAN,
// 	added_at TIMESTAMP NOT NULL,
// 	read_at TIMESTAMP

// );
// `
