-- name: CreateBook :one
-- Create a new book
INSERT INTO Books (title, author, publishDate, pageCount, readStatus, collection_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id;

-- name: GetBookById :one
-- Retrieve a specific book by ID
SELECT *
FROM Books
WHERE id = ?;

-- name: GetAllBooks :many
-- List all books
SELECT *
FROM Books;

-- name: UpdateBook :exec
-- Update an existing book
UPDATE Books
SET title         = ?,
    author        = ?,
    publishDate   = ?,
    pageCount     = ?,
    readStatus    = ?,
    collection_id = ?
WHERE id = ?;

-- name: MarkBookRead :exec
-- Mark a book as read
UPDATE Books
SET readStatus = 1 -- Assuming 1 represents "read" status, adjust if needed
WHERE id = ?;

-- name: AddBookToCollection :exec
-- Add a book to a collection
UPDATE Books
SET collection_id = ?
WHERE id = ?;

-- name: DeleteBook :exec
-- Delete a book
DELETE
FROM Books
WHERE id = ?;

-- name: CreateCollection :one
-- Create a new collection
INSERT INTO Collections (title)
VALUES (?)
RETURNING id;

-- name: GetAllCollections :many
-- List all collections
SELECT *
FROM Collections;

-- name: GetCollectionById :one
-- Retrieve a specific collection by ID
SELECT *
FROM Collections
WHERE id = ?;

-- name: UpdateCollection :exec
-- Update an existing collection
UPDATE Collections
SET title = ?
WHERE id = ?;

-- name: DeleteCollection :exec
-- Delete a collection
DELETE
FROM Collections
WHERE id = ?;

-- name: GetAllGenres :many
-- List all genres
SELECT *
FROM Genres;

-- name: AddGenre :one
-- Add a genre if it doesn't exist
INSERT OR IGNORE INTO Genres (name)
VALUES (?)
RETURNING ID;

-- name: AssociateBookWithGenre :exec
-- Associate a book with a genre
INSERT INTO BookGenres (book_id, genre_id)
VALUES (?, ?);

-- name: GetBookWithGenres :one
-- Retrieve a book and its associated genres
SELECT b.*, g.*
FROM Books b
         JOIN BookGenres bg ON b.id = bg.book_id
         JOIN Genres g ON bg.genre_id = g.id
WHERE b.id = ?;