package main

import (
	"log"
	"net/http"

	"github.com/Chubacabrazz/book-db/book_services/author"
	book "github.com/Chubacabrazz/book-db/book_services/book"
	postgres "github.com/Chubacabrazz/book-db/db"
	"github.com/Chubacabrazz/book-db/file_services/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	//Set environment variables : database infos.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")
	h := handlers.NewConn(db)

	//Set router and RestAPI
	router := mux.NewRouter()

	//router.Handlefunc for handling http requests.
	router.HandleFunc("/books", h.ListBooks).Methods(http.MethodGet)
	router.HandleFunc("/authors", h.ListAuthors).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/search/{word}", h.SearchWord).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.SoftDelete).Methods(http.MethodDelete)
	router.HandleFunc("/authors/{id}", h.SoftDeleteAuthor).Methods(http.MethodDelete)
	log.Println("API is running!!")
	http.ListenAndServe(":8000", router) //for setting port and keeping server up.

	// Repositories
	bookRepo := book.NewBookRepository(db)
	bookRepo.Migrations()
	//bookRepo.InsertData()
	//bookRepo.Buy(20, 33)
	//bookRepo.FindByAuthor(<authorname>)
	//bookRepo.GetByID()
	//bookRepo.SearchWord("lord")
	//bookRepo.SoftDeletebyID(33)
	//bookRepo.List()

	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migration()
	//authorRepo.InsertData()
	//authorRepo.List()
	//authorRepo.FindByWord(<name>)

}
