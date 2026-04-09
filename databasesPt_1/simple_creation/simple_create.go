package simple_creation

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books(
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		author VARCHAR(100) NOT NULL,
		review VARCHAR(1000),
		release_date INTEGER NOT NULL,
		was_read BOOLEAN,
		added_at TIMESTAMP NOT NULL,
		read_at TIMESTAMP
		
	);
	`
	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
