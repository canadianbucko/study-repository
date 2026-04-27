package sql

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	connectionString := os.Getenv("CONN_STRING") //include .env    export  - makefile не забудь  чтобы он импортировал
	//trueConnString := "postgres://postgres:1234@localhost:54332/mydb?sslmode=disable"
	conn, err := pgx.Connect(ctx, connectionString)

	return conn, err
}
