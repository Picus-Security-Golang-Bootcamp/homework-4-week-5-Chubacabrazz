package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Chubacabrazz/book-db/book_services/book"
	"github.com/gorilla/mux"
)

//BuyBook finds given ID in DB, if stock is enough; buys it, then updates new stock on DB.
func (h handler) BuyBook(w http.ResponseWriter, r *http.Request) {
	var book book.Book
	vars := mux.Vars(r)
	bookID := vars["id"]
	quantity, _ := strconv.Atoi(vars["quantity"])
	result := h.DB.Where("book_id = ?", bookID).First(&book)
	stock := book.Book_Stock
	if result.Error != nil {
		fmt.Println(result.Error)
	} else if stock < quantity {
		fmt.Fprintf(w, "we don't have that much. we have: %d books in our stock", stock)
	} else {
		fmt.Fprintf(w, "Shopping successfull. You bought: %s , %d Remain in stock", book.Book_Name, stock-quantity)
	}
	newStock := stock - quantity

	result = h.DB.Model(&book).Where("book_id = ? AND book_stock >= ?", bookID, quantity).
		Update("book_stock", newStock)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
