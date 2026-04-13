package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `
	DELETE FROM books
	WHERE id = $1; 
	`
	// that $1 thing works btw!!!

	_, err := conn.Exec(ctx, sqlQuery, id)
	return err

}
