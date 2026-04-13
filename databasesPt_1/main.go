package main

import (
	"context"
	"fmt"
	"someproject/sample_struct"
	"someproject/simple_sql"
)

func main() {
	ctx := context.Background()
	conn, err := simple_sql.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("ha, looks like it's okay, we're in! we're connected!") // just a text

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	fmt.Println("hey, some table was created successfully!")

	myBook := sample_struct.NewBook("19124", "balls  stupid", "it's not alright", 1949, false, nil)

	if err := simple_sql.InsertThat(ctx, conn, myBook); err != nil {
		panic(err)
	}
	fmt.Println("looks like we gut")

	if err := simple_sql.DeleteRow(ctx, conn, 2); err != nil {
		panic(err)
	}

	fmt.Println("works?")
}
