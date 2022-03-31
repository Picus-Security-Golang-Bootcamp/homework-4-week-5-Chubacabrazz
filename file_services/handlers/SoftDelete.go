package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chubacabrazz/book-db/book_services/book"
	"github.com/gorilla/mux"
)

func (h handler) SoftDelete(w http.ResponseWriter, r *http.Request) {
	var book book.Book
	vars := mux.Vars(r)
	bookID := vars["id"]
	result := h.DB.Where("book_id = ?", bookID).First(&book)
	if result.Error != nil {
		fmt.Println(result.Error)
	} else {
		fmt.Println("Success, book soft-deleted:", bookID)
		h.DB.Where("book_id = ?", bookID).Delete(&book)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(book)

}
