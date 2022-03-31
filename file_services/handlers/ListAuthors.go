package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chubacabrazz/book-db/book_services/author"
)

func (h handler) ListAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []author.Author
	if result := h.DB.Find(&authors); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(authors)

}
