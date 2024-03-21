CREATE TABLE Books
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    title         TEXT NOT NULL,
    author        TEXT NOT NULL,
    publishDate   DATETIME,
    pageCount     INTEGER,
    readStatus    INTEGER,
    collection_id INTEGER,
    FOREIGN KEY (collection_id) REFERENCES Collections (id)
);

CREATE TABLE Collections
(
    id    INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL
);

CREATE TABLE Genres
(
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE BookGenres
(
    book_id  INTEGER,
    genre_id INTEGER,
    FOREIGN KEY (book_id) REFERENCES Books (id),
    FOREIGN KEY (genre_id) REFERENCES Genres (id),
    PRIMARY KEY (book_id, genre_id)
)