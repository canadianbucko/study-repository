package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

func CheckConnection(ctx context.Context) {
	connectionString := "postgres://postgres:11757@localhost:5432/postgres"

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		panic(err)
	}
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("ha, looks like it's okay, we're in! we're connected!") // just a text
}
