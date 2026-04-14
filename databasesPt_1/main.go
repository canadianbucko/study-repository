package main

import (
	"context"
	"fmt"
	"someproject/sample_struct"
	"someproject/simple_sql"
	"time"

	"github.com/k0kubun/pp"
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

	myBook := sample_struct.NewBook("113234", "someth3ing idk", "it's 2alright", 19429, false, nil)

	if err := simple_sql.InsertThat(ctx, conn, myBook); err != nil {
		panic(err)
	}
	fmt.Println("looks like we gut")

	if err := simple_sql.DeleteRow(ctx, conn, 2); err != nil {
		panic(err)
	}

	fmt.Println("works?")

	books, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}
	pp.Print(books)

	// Author      string
	// Review      string
	// ReleaseDate int
	// WasRead     bool
	// AddedAt     time.Time
	// ReadAt

	for _, book := range books { // вот мы идем в базу данных и получаем все данные из нее и тут блять
		// идем и сука по каждой из этой херни смотрим где же блять book.ID == 4.
		// и если нашли то вэтом экземпляре меняем все
		// и записываем обратно
		if book.ID == 4 { //that's just stupid but it'll do
			book.Title = "retard"
			book.Author = "retard"
			book.ReleaseDate = 0
			book.WasRead = false
			book.AddedAt = time.Now()
			now := time.Now()
			book.ReadAt = &now

			if err := simple_sql.UpdateRow(ctx, conn, book); err != nil {
				panic(err)
			}
			break
		}
		pp.Print("fuck off ")
	}

	sliceIds := []int{1, 2, 3}
	if err := simple_sql.DeleteByID(ctx, conn, sliceIds); err != nil {
		panic(err)
	}

}
