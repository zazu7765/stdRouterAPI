-- name: create_book :one
-- Create a new book
INSERT INTO books (title, author, publishDate, pageCount, readStatus, collection_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id;

-- name: get_book_by_id :one
-- Retrieve a specific book by ID
SELECT * FROM books WHERE id = ?;

-- name: update_book :exec
-- Update an existing book
UPDATE books
SET title = ?, author = ?, publishDate = ?, pageCount = ?, readStatus = ?, collection_id = ?
WHERE id = ?;

-- name: add_book_to_collection :exec
-- Add a book to a collection
UPDATE books
SET collection_id = ?
WHERE id = ?;

-- name: delete_book :exec
-- Delete a book
DELETE FROM books WHERE id = ?;

-- name: create_collection :one
-- Create a new collection
INSERT INTO collections (title)
VALUES (?)
RETURNING id;

-- name: get_collection_by_id :one
-- Retrieve a specific collection by ID
SELECT * FROM collections WHERE id = ?;

-- name: update_collection :exec
-- Update an existing collection
UPDATE collections
SET title = ?
WHERE id = ?;

-- name: delete_collection :exec
-- Delete a collection
DELETE FROM collections WHERE id = ?;

-- name: add_genre :exec
-- Add a genre if it doesn't exist
INSERT OR IGNORE INTO Genres (name) VALUES (?);