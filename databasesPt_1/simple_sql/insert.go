package simple_sql

import (
	"context"
	"someproject/sample_struct"

	"github.com/jackc/pgx/v5"
)

func InsertThat(ctx context.Context, conn *pgx.Conn, book sample_struct.Book) error {
	sqlQuery := `
	INSERT INTO books (title, author, review, release_date, was_read, added_at, read_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	_, err := conn.Exec(ctx, sqlQuery, book.Title, book.Author, book.Review, book.ReleaseDate, book.WasRead, book.AddedAt, book.ReadAt)
	return err
}

// sqlQuery := `
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
