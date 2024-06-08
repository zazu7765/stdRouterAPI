package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/csv"
	"fmt"
	books "github.com/zazu7765/stdRouterAPI/src/internal/app/books"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/zazu7765/stdRouterAPI/src/internal/database"
	"github.com/zazu7765/stdRouterAPI/src/internal/server"
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

	_, err = reader.Read()
	if err != nil {
		return err
	}

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	log.Println("Loading Records from books.csv...")
	for _, record := range records {
		title := record[0]
		authors := record[1]
		publishDate := record[2]
		ISBN := record[3]
		genre := record[4]
		date, err := formatPublishDate(publishDate)
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
		b, err := q.CreateBook(ctx, bookfields)
		if err != nil {
			return err
		}

		genreSplit := strings.Split(genre, ",")
		for _, g := range genreSplit {
			_, _ = q.AddGenre(ctx, g)
			gid, err := q.GetGenreByName(ctx, g)
			if err != nil {
				log.Println(err)
			}
			// fmt.Println("Genre ID: ", gid)
			err = q.AssociateBookWithGenre(ctx, database.AssociateBookWithGenreParams{
				BookID:  sql.NullInt64{Int64: b, Valid: true},
				GenreID: sql.NullInt64{Int64: gid.ID, Valid: true},
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func run() (*http.ServeMux, error) {
	filePath := filepath.Join("src", "configs", "books.csv")
	ctx := context.Background()
	// Temporary memory database for testing
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}

	// Set tables through schema file
	if _, err := db.ExecContext(ctx, schema); err != nil {
		return nil, err
	}

	// Create queries instance
	queries := database.New(db)

	err = populateDB(ctx, queries, filePath)
	if err != nil {
		return nil, err
	}
	routes := []server.Route{
		{
			Name:    "Ping",
			Method:  "GET",
			Pattern: "/ping",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				log.Println("Ping Request")
				_, err := w.Write([]byte("Pong"))
				if err != nil {
					log.Println("Error writing response:", err)
				}
			},
		},
		{
			Name:    "GetAllBooks",
			Method:  "GET",
			Pattern: "/GetAllBooks",
			Handler: books.GetBooks(queries, ctx),
		},
	}
	log.Println("Starting REST Server on :8080")
	return server.NewRouter(routes), nil
}

func main() {
	s, err := run()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", s))
}
