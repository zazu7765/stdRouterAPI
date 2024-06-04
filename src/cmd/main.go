package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/zazu7765/stdRouterAPI/src/internal/database"
	_ "modernc.org/sqlite"
)

// The heresy of sinners will fall to the power of time.DateOnly
func formatPublishDate(str string) (time.Time, error) {
	str2 := strings.Split(str, ".")
	if len(str) < 1 {
		fmt.Println("Str has 2 parts")
	}

	intstr, err := strconv.ParseUint(str2[0], 10, 16)
	if err != nil {
		return time.Now(), err
	}

	// fmt.Println(intstr)
	date := time.Date(int(intstr), 1, 1, 0, 0, 0, 0, time.UTC)
	// fmt.Println(date)
	return date, nil
}

// Embeded SQL file as a string to upload to database
//
//go:embed sql/schema.sql
var schema string

func populateDB(ctx context.Context, q *database.Queries, f string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	reader.Read()

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
		title := record[0]
		authors := record[1]
		publish_date := record[2]
		ISBN := record[3]
		genre := record[4]
		date, err := formatPublishDate(publish_date)
		if err != nil {
			fmt.Println(err)
			continue
		}
		pDate := sql.NullTime{
			Time:  date,
			Valid: true,
		}
		bookfields := database.CreateBookParams{
			Title:       title,
			Author:      authors,
			Publishdate: pDate,
			Isbn:        ISBN,
			Readstatus: sql.NullInt64{
				Int64: 0,
				Valid: true,
			},
			CollectionID: sql.NullInt64{
				Valid: false,
			},
		}
		_, err = q.CreateBook(ctx, bookfields)
		if err != nil {
			return err
		}

		genreSplit := strings.Split(genre, ",")
		for _, g := range genreSplit {
			gid, err := q.AddGenre(ctx, g)
			if err != nil {
				log.Println(err)
				log.Println("Adding Genre: ", gid)
			}
			// err = q.AssociateBookWithGenre(ctx, database.AssociateBookWithGenreParams{
			// 	BookID:  sql.NullInt64{Int64: b, Valid: true},
			// 	GenreID: sql.NullInt64{Int64: g, Valid: true},
			// })
			// if err != nil {
			// 	return err
			// }
		}
	}
	return nil
}

func run() error {
	filePath := filepath.Join("src", "configs", "books.csv")
	ctx := context.Background()
	// Temporary memory database for testing
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return err
	}

	// Set tables through schema file
	if _, err := db.ExecContext(ctx, schema); err != nil {
		return err
	}

	// Create queries instance
	queries := database.New(db)

	err = populateDB(ctx, queries, filePath)
	if err != nil {
		return err
	}
	// Get all books (should return empty)
	books, err := queries.GetAllBooks(ctx)
	if err != nil {
		return err
	}
	log.Println(books)

	// Add book to database (Associate Genres later) (Will Probably wrapper function this)
	timeBook, _ := formatPublishDate("2012-06-05")
	addedBook, err := queries.CreateBook(ctx, database.CreateBookParams{
		Title:  "Ready Player One",
		Author: "Ernest Cline",
		Publishdate: sql.NullTime{
			Time:  timeBook,
			Valid: true,
		},
		Isbn: "0307887448",
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
	// routes := []server.Route{}
	// route1 := server.Route{
	// 	Name:    "HelloWorld",
	// 	Method:  "GET",
	// 	Pattern: "/hello",
	// 	Handler: func(w http.ResponseWriter, r *http.Request) {
	// 		log.Println("Hello World Request")
	// 		w.Write([]byte("Hello World!"))
	// 	},
	// }
	// routes = append(routes, route1)
	// log.Println("Starting REST Server on :8080")
	// log.Fatal(http.ListenAndServe(":8080", server.NewRouter(routes)))
}
