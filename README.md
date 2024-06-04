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


## Setup (Mac/Linux only)
1. Make all scripts executable (`chmod +x`)
2. run `run_all.sh`
  - If you want to only rebuild, run `build.sh`
  - To run after building, run `./bin/stdRouterApi`
  - If you accidentally deleted something or uninstalled a dependency, run `bootstrap.sh`

## Setup (Windows)
Make sure that the bin directory exists under this project folder!
```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
sqlc generate -f src/configs/sqlc.yaml
go mod tidy
go build -o bin/stdRouterApi ./src/cmd/main.go
```
