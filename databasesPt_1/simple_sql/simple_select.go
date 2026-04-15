package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]BookModel, error) {
	sqlQuery := `
	SELECT id, title, author, review, release_date, was_read, added_at, read_at
	FROM books
	ORDER BY id ASC
	LIMIT 10;
	`

	rows, err := conn.Query(ctx, sqlQuery)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	books := make([]BookModel, 0)

	for rows.Next() {
		var book BookModel

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Review,
			&book.ReleaseDate,
			&book.WasRead,
			&book.AddedAt,
			&book.ReadAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil

}
