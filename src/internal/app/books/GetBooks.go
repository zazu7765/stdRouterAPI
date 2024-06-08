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
		res := BookResponseList{}
		query, err := queries.GetAllBooks(ctx)
		if err != nil {
			log.Println(err)
			res.Status = http.StatusInternalServerError
		}
		books := make([]BookResponse, len(query))
		for i, q := range query {
			b := CreateBookResponse(q)
			books[i] = b
		}

		res.Results = books
		res.Status = http.StatusOK
		response, err := json.Marshal(res)
		if err != nil {
			log.Println("Error marshalling payload:", err)
			res.Status = http.StatusInternalServerError
			res.Results = nil
		}

		response, err = json.Marshal(res)
		_, err = w.Write(response)
		if err != nil {
			log.Println("Error writing response:", err)
		}
	}
}
