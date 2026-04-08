package main

import (
	"context"
	"someproject/simple_connection"
)

func main() {
	ctx := context.Background()
	simple_connection.CheckConnection(ctx)

}
