package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteByID(ctx context.Context, conn *pgx.Conn, id []int) error {
	sqlQuery := `
	DELETE FROM books
	WHERE ID=ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, id)
	return err
}
