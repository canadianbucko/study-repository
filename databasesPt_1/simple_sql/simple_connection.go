package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	connectionString := "postgres://postgres:1234@localhost:5432/mydb"

	conn, err := pgx.Connect(ctx, connectionString)

	return conn, err
}
