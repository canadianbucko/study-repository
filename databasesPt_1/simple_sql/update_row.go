package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// которая принимает в аргументах модель книги, находит в базе данных в таблице books книгу с ID'шником,
// как у переданной модели, и этой книге в базе данных обновляет все поля на те, что были у переданной модели.
func UpdateRow(ctx context.Context, conn *pgx.Conn, book BookModel) error {
	sqlQuery := `
	UPDATE books
	SET title=$2, author=$3, review=$4, release_date=$5, was_read=$6, added_at=$7, read_at=$8
	WHERE id=$1	
	`
	_, err := conn.Exec(ctx,
		sqlQuery,
		book.ID,
		book.Title,
		book.Author,
		book.Review,
		book.ReleaseDate,
		book.WasRead,
		book.AddedAt,
		book.ReadAt)
	return err

}
