package main

import (
	"context"
	"database/sql"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"stdRouterAPI/database"
	"time"
)

func formatPublishDate(str string) time.Time {
	date, _ := time.Parse(time.DateOnly, str)
	return date
}

//go:embed schema.sql
var schema string

func run() error {
	ctx := context.Background()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, schema); err != nil {
		return err
	}

	queries := database.New(db)

	books, err := queries.GetAllBooks(ctx)
	if err != nil {
		return err
	}
	log.Println(books)

	addedBook, err := queries.CreateBook(ctx, database.CreateBookParams{
		Title:  "Ready Player One",
		Author: "Ernest Cline",
		Publishdate: sql.NullTime{
			Time:  formatPublishDate("2012-06-05"),
			Valid: true,
		},
		Pagecount: sql.NullInt64{
			Int64: 384,
			Valid: true,
		},
		Readstatus: sql.NullInt64{
			Int64: 0,
			Valid: true,
		},
		CollectionID: sql.NullInt64{
			Valid: false,
		},
	})
	if err != nil {
		return err
	}

	log.Println(addedBook)

	retrievedBook, err := queries.GetBookById(ctx, addedBook)
	if err != nil {
		return err
	}

	log.Println(retrievedBook)

	genreList, err := queries.GetAllGenres(ctx)
	if err != nil {
		return err
	}
	log.Println(genreList)

	createdGenre, err := queries.AddGenre(ctx, "Science Fiction")
	if err != nil {
		return err
	}

	err = queries.AssociateBookWithGenre(ctx, database.AssociateBookWithGenreParams{
		BookID:  sql.NullInt64{Int64: retrievedBook.ID, Valid: true},
		GenreID: sql.NullInt64{Int64: createdGenre, Valid: true},
	})
	if err != nil {
		return err
	}

	book, err := queries.GetBookWithGenres(ctx, addedBook)
	if err != nil {
		return err
	}
	log.Println(book)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}