package simple_sql

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	connectionString := os.Getenv("CONN_STRING")

	conn, err := pgx.Connect(ctx, connectionString)

	return conn, err
}
