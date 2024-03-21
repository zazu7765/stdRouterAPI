// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package database

import (
	"context"
	"database/sql"
)

const addBookToCollection = `-- name: AddBookToCollection :exec
UPDATE Books
SET collection_id = ?
WHERE id = ?
`

type AddBookToCollectionParams struct {
	CollectionID sql.NullInt64
	ID           int64
}

// Add a book to a collection
func (q *Queries) AddBookToCollection(ctx context.Context, arg AddBookToCollectionParams) error {
	_, err := q.db.ExecContext(ctx, addBookToCollection, arg.CollectionID, arg.ID)
	return err
}

const addGenre = `-- name: AddGenre :one
INSERT OR IGNORE INTO Genres (name)
VALUES (?)
RETURNING ID
`

// Add a genre if it doesn't exist
func (q *Queries) AddGenre(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRowContext(ctx, addGenre, name)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const associateBookWithGenre = `-- name: AssociateBookWithGenre :exec
INSERT INTO BookGenres (book_id, genre_id)
VALUES (?, ?)
`

type AssociateBookWithGenreParams struct {
	BookID  sql.NullInt64
	GenreID sql.NullInt64
}

// Associate a book with a genre
func (q *Queries) AssociateBookWithGenre(ctx context.Context, arg AssociateBookWithGenreParams) error {
	_, err := q.db.ExecContext(ctx, associateBookWithGenre, arg.BookID, arg.GenreID)
	return err
}

const createBook = `-- name: CreateBook :one
INSERT INTO Books (title, author, publishDate, pageCount, readStatus, collection_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id
`

type CreateBookParams struct {
	Title        string
	Author       string
	Publishdate  sql.NullTime
	Pagecount    sql.NullInt64
	Readstatus   sql.NullInt64
	CollectionID sql.NullInt64
}

// Create a new book
func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createBook,
		arg.Title,
		arg.Author,
		arg.Publishdate,
		arg.Pagecount,
		arg.Readstatus,
		arg.CollectionID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createCollection = `-- name: CreateCollection :one
INSERT INTO Collections (title)
VALUES (?)
RETURNING id
`

// Create a new collection
func (q *Queries) CreateCollection(ctx context.Context, title string) (int64, error) {
	row := q.db.QueryRowContext(ctx, createCollection, title)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE
FROM Books
WHERE id = ?
`

// Delete a book
func (q *Queries) DeleteBook(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBook, id)
	return err
}

const deleteCollection = `-- name: DeleteCollection :exec
DELETE
FROM Collections
WHERE id = ?
`

// Delete a collection
func (q *Queries) DeleteCollection(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCollection, id)
	return err
}

const getAllBooks = `-- name: GetAllBooks :many
SELECT id, title, author, publishdate, pagecount, readstatus, collection_id
FROM Books
`

// List all books
func (q *Queries) GetAllBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, getAllBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.Publishdate,
			&i.Pagecount,
			&i.Readstatus,
			&i.CollectionID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllCollections = `-- name: GetAllCollections :many
SELECT id, title
FROM Collections
`

// List all collections
func (q *Queries) GetAllCollections(ctx context.Context) ([]Collection, error) {
	rows, err := q.db.QueryContext(ctx, getAllCollections)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(&i.ID, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllGenres = `-- name: GetAllGenres :many
SELECT id, name
FROM Genres
`

// List all genres
func (q *Queries) GetAllGenres(ctx context.Context) ([]Genre, error) {
	rows, err := q.db.QueryContext(ctx, getAllGenres)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Genre
	for rows.Next() {
		var i Genre
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBookById = `-- name: GetBookById :one
SELECT id, title, author, publishdate, pagecount, readstatus, collection_id
FROM Books
WHERE id = ?
`

// Retrieve a specific book by ID
func (q *Queries) GetBookById(ctx context.Context, id int64) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBookById, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Publishdate,
		&i.Pagecount,
		&i.Readstatus,
		&i.CollectionID,
	)
	return i, err
}

const getBookWithGenres = `-- name: GetBookWithGenres :one
SELECT b.id, b.title, b.author, b.publishdate, b.pagecount, b.readstatus, b.collection_id, g.id, g.name
FROM Books b
         JOIN BookGenres bg ON b.id = bg.book_id
         JOIN Genres g ON bg.genre_id = g.id
WHERE b.id = ?
`

type GetBookWithGenresRow struct {
	ID           int64
	Title        string
	Author       string
	Publishdate  sql.NullTime
	Pagecount    sql.NullInt64
	Readstatus   sql.NullInt64
	CollectionID sql.NullInt64
	ID_2         int64
	Name         string
}

// Retrieve a book and its associated genres
func (q *Queries) GetBookWithGenres(ctx context.Context, id int64) (GetBookWithGenresRow, error) {
	row := q.db.QueryRowContext(ctx, getBookWithGenres, id)
	var i GetBookWithGenresRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Publishdate,
		&i.Pagecount,
		&i.Readstatus,
		&i.CollectionID,
		&i.ID_2,
		&i.Name,
	)
	return i, err
}

const getCollectionById = `-- name: GetCollectionById :one
SELECT id, title
FROM Collections
WHERE id = ?
`

// Retrieve a specific collection by ID
func (q *Queries) GetCollectionById(ctx context.Context, id int64) (Collection, error) {
	row := q.db.QueryRowContext(ctx, getCollectionById, id)
	var i Collection
	err := row.Scan(&i.ID, &i.Title)
	return i, err
}

const markBookRead = `-- name: MarkBookRead :exec
UPDATE Books
SET readStatus = 1 -- Assuming 1 represents "read" status, adjust if needed
WHERE id = ?
`

// Mark a book as read
func (q *Queries) MarkBookRead(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, markBookRead, id)
	return err
}

const updateBook = `-- name: UpdateBook :exec
UPDATE Books
SET title         = ?,
    author        = ?,
    publishDate   = ?,
    pageCount     = ?,
    readStatus    = ?,
    collection_id = ?
WHERE id = ?
`

type UpdateBookParams struct {
	Title        string
	Author       string
	Publishdate  sql.NullTime
	Pagecount    sql.NullInt64
	Readstatus   sql.NullInt64
	CollectionID sql.NullInt64
	ID           int64
}

// Update an existing book
func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.ExecContext(ctx, updateBook,
		arg.Title,
		arg.Author,
		arg.Publishdate,
		arg.Pagecount,
		arg.Readstatus,
		arg.CollectionID,
		arg.ID,
	)
	return err
}

const updateCollection = `-- name: UpdateCollection :exec
UPDATE Collections
SET title = ?
WHERE id = ?
`

type UpdateCollectionParams struct {
	Title string
	ID    int64
}

// Update an existing collection
func (q *Queries) UpdateCollection(ctx context.Context, arg UpdateCollectionParams) error {
	_, err := q.db.ExecContext(ctx, updateCollection, arg.Title, arg.ID)
	return err
}
