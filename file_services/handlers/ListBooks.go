package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chubacabrazz/book-db/book_services/book"
)

//ListBooks lists all books from DB, then prints them with json.encoder
func (h handler) ListBooks(w http.ResponseWriter, r *http.Request) {
	var books []book.Book
	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(books)

}
