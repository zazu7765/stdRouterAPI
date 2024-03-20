# Simple API Example in Golang
Models a bookshelf:
- Books 
  - Title
  - Author
  - Genres
  - Publish Date
  - Page Count
  - Read?
  - Collection
- Collections
  - Title
  - Books

Queries will be generated using `sqlc`, while the API itself will be implemented using `net/http` 