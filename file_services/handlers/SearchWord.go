package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chubacabrazz/book-db/book_services/book"
	"github.com/gorilla/mux"
)

func (h handler) SearchWord(w http.ResponseWriter, r *http.Request) {
	var book []book.Book
	vars := mux.Vars(r)
	search := vars["word"]
	result := h.DB.Where("book_name ILIKE ? ", "%"+search+"%").Find(&book)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(book)

}
