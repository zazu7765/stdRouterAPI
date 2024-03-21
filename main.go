package main

import (
	"context"
	"database/sql"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zazu7765/stdRouterAPI/database"
	"log"
	"time"
)

// The heresy of sinners will fall to the power of time.DateOnly
func formatPublishDate(str string) time.Time {
	date, _ := time.Parse(time.DateOnly, str)
	return date
}

// Embeded SQL file as a string to upload to database
//
//go:embed schema.sql
var schema string

func run() error {
	ctx := context.Background()
	// Temporary memory database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	// Set tables through schema file
	if _, err := db.ExecContext(ctx, schema); err != nil {
		return err
	}

	// Create queries instance
	queries := database.New(db)

	// Get all books (should return empty)
	books, err := queries.GetAllBooks(ctx)
	if err != nil {
		return err
	}
	log.Println(books)

	// Add book to database (Associate Genres later) (Will Probably wrapper function this)
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

	// Get book by ID from database
	retrievedBook, err := queries.GetBookById(ctx, addedBook)
	if err != nil {
		return err
	}

	log.Println(retrievedBook)

	// Get All Genres from database (should be empty)
	genreList, err := queries.GetAllGenres(ctx)
	if err != nil {
		return err
	}
	log.Println(genreList)

	// Add Genre option to database
	createdGenre, err := queries.AddGenre(ctx, "Science Fiction")
	if err != nil {
		return err
	}

	// Associate Genre to Book and vice versa in join table
	err = queries.AssociateBookWithGenre(ctx, database.AssociateBookWithGenreParams{
		BookID:  sql.NullInt64{Int64: retrievedBook.ID, Valid: true},
		GenreID: sql.NullInt64{Int64: createdGenre, Valid: true},
	})
	if err != nil {
		return err
	}

	// Return book with genre, not in a struct yet just as a row which means I will probably wrap my own struct
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
