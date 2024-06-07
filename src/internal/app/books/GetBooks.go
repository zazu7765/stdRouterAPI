package books

import (
	"context"
	"encoding/json"
	"github.com/zazu7765/stdRouterAPI/src/internal/database"
	"log"
	"net/http"
)

func GetBooks(queries *database.Queries, ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GetAllBooks Request")
		books, err := queries.GetAllBooks(ctx)
		if err != nil {
			log.Println(err)
		}
		// TODO: create struct to hide unnecessary fields and send request through there instead
		payload, err := json.Marshal(books)
		if err != nil {
			log.Println("Error marshalling payload:", err)
		}
		_, err = w.Write(payload)
		if err != nil {
			log.Println("Error writing response:", err)
		}
	}
}
