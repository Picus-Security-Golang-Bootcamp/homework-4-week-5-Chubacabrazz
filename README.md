# Library App in GOLANG

### Build with Go. Used Gorm , PostgreSQL , Godotenv , Gorilla Mux Packages. Also used Postman.

This app reads a csv list with workerpool and writes them to db then you can access them with RestAPI. It works with http requests.

For checking APIs, I used Postman. My domain for this example is "localhost:8000"

#### •Func List : Lists all books from DB. ||| Usage: Generate a GET request at "domain/books" 

![listpostman](https://user-images.githubusercontent.com/77194087/161068024-3e84809e-6dfe-4577-bc52-4bc373c4a128.png)

#### •Func GetBook : Lists a book of given Book_ID from DB. ||| Usage: Generate a GET request at "domain/books/id" 

![getbookpostman](https://user-images.githubusercontent.com/77194087/161070098-579238dd-f898-47c5-b110-9221b667540a.png)

#### •Func SearchWord : Classic insensitive search with Book_Name from DB. ||| Usage: Generate a GET request at "domain/books/search/word" 

![searchpostman](https://user-images.githubusercontent.com/77194087/161072159-5edf9346-8af2-4ae7-b2bf-dde438c68298.png)

#### •Func SoftDelete : Soft deletes a book from DB. ||| Usage: Generate a DELETE request at "domain/books/id" 

![deletepostman](https://user-images.githubusercontent.com/77194087/161073311-e4af1221-325d-4be8-bd80-1a8e24532c9b.png)

