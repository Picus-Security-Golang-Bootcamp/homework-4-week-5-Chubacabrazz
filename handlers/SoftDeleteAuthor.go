package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chubacabrazz/book-db/book_services/author"
	"github.com/gorilla/mux"
)

//SoftDeleteAuthor deletes given Author_ID informations from DB, then prints it wih json.encoder
func (h handler) SoftDeleteAuthor(w http.ResponseWriter, r *http.Request) {
	var author author.Author
	vars := mux.Vars(r)
	authorID := vars["id"]
	result := h.DB.Where("book_id = ?", authorID).First(&author)
	if result.Error != nil {
		fmt.Println(result.Error)
	} else {
		fmt.Println("Success, book soft-deleted:", authorID)
		h.DB.Where("book_id = ?", authorID).Delete(&author)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(author)

}
