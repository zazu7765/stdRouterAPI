CREATE TABLE Books (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       title TEXT NOT NULL,
                       author TEXT NOT NULL,
                       genre_id INTEGER, -- Store genre as a number
                       publishDate DATETIME,
                       pageCount INTEGER,
                       readStatus INTEGER,
                       collection_id INTEGER,
                       FOREIGN KEY (collection_id) REFERENCES Collections(id),
                       FOREIGN KEY (genre_id) REFERENCES Genres(id) -- Foreign key to Genres table
);

CREATE TABLE Collections (
                             id INTEGER PRIMARY KEY AUTOINCREMENT,
                             title TEXT NOT NULL
);

CREATE TABLE Genres (
                        id INTEGER PRIMARY KEY,
                        name TEXT NOT NULL UNIQUE
);
