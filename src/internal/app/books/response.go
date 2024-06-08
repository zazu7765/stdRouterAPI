package books

import (
	"github.com/zazu7765/stdRouterAPI/src/internal/database"
	"strconv"
)

type BookResponseList struct {
	Status  int            `json:"status"`
	Results []BookResponse `json:"results"`
}

type BookResponse struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publishDate"`
	Status      string `json:"status"`
	ISBN        string `json:"isbn"`
	Collection  string `json:"collection"`
}

func CreateBookResponse(book database.Book) BookResponse {
	pdate := ""
	status := ""
	collection := ""

	if book.Publishdate.Valid == true {
		pdate = strconv.Itoa(book.Publishdate.Time.Year())
	}

	if book.Readstatus.Valid == true {
		switch book.Readstatus.Int64 {
		case 0:
			status = "true"
		case 1:
			status = "false"
		}
	}

	if book.CollectionID.Valid == true {
		collection = strconv.FormatInt(book.CollectionID.Int64, 10)
	}

	return BookResponse{
		book.Title,
		book.Author,
		pdate,
		status,
		book.Isbn,
		collection,
	}
}
