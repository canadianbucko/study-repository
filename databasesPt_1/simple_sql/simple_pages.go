package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/k0kubun/pp"
)

func SelectPage(ctx context.Context, conn *pgx.Conn, limit int, offset int) ([]BookModel, error) {
	sqlQuery := `
	SELECT id, title, author, review, release_date, was_read, added_at, read_at
	FROM books
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	booksSlice := make([]BookModel, 0)

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

		booksSlice = append(booksSlice, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return booksSlice, nil
}

func ListPage(ctx context.Context, conn *pgx.Conn, pageSize int) error {
	page := 1
	offset := 0

	for {
		books, err := SelectPage(ctx, conn, pageSize, offset)
		if err != nil {
			return err
		}

		if len(books) == 0 {
			pp.Println("hey, that's a last page, thank you")
			break
		}

		pp.Print("printing page №", page, books)

		page++

		offset += pageSize

	}
	return nil
}
