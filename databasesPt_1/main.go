package main

import (
	"context"
	"fmt"
	"someproject/sample_struct"
	"someproject/simple_connection"
	"someproject/simple_creation"
	"someproject/simple_insert"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("ha, looks like it's okay, we're in! we're connected!") // just a text

	if err := simple_creation.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	fmt.Println("hey, some table was created successfully!")

	myBook := sample_struct.NewBook("19124", "balls  stupid", "it's not alright", 1949, false, nil)

	if err := simple_insert.InsertThat(ctx, conn, myBook); err != nil {
		panic(err)
	}
	fmt.Println("looks like we gut")
}
