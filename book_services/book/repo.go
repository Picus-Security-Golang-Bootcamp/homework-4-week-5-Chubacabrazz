package book

import (
	"fmt"
	"strconv"

	"github.com/Chubacabrazz/book-db/file_services/csv_utils"
	"gorm.io/gorm"
)

var csvfile string = "books.csv"

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) List() {
	var books []Book
	b.db.Find(&books)

	for _, thebook := range books {
		fmt.Println(thebook.ToString())
	}
}

//Func SoftDeletebyID applies a soft delete to a book
func (c *BookRepository) SoftDeletebyID(id int) error {
	var book Book
	bookID := strconv.Itoa(id)
	result := c.db.Where("book_id = ?", bookID).First(&book)
	if result.Error != nil {
		return result.Error
	} else {
		fmt.Println("Success, book soft-deleted:", id)
		c.db.Where("book_id = ?", bookID).Delete(&Book{})
	}
	//result = c.db.Delete(book, bookID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Func Buy does purchase for a book with given ID and quantity
func (c *BookRepository) Buy(quantity, id int) error {
	var book Book
	bookID := strconv.Itoa(id)
	result := c.db.Where("book_id = ?", bookID).First(&book)
	stock := book.Book_Stock
	if result.Error != nil {
		return result.Error
	} else if stock < quantity {
		fmt.Printf("we don't have that much. we have: %d", stock)
		return nil
	} else {
		fmt.Printf("Shopping successfull. You bought: %s", book.Book_Name)
	}

	result = c.db.Model(&book).Where("id = ? AND book_stock >= ?", bookID, quantity).
		Update("book_stock", gorm.Expr("book_stock - ?", quantity))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Func FindByAuthor finds books with Author Name
func (c *BookRepository) FindByAuthor(Author string) []Book {
	var books []Book
	c.db.Where(`"Author" = ?`, Author).Find(&books)
	for _, thebook := range books {
		fmt.Println(thebook.ToString())
	}
	return books
}

//Func SearchWord lists the books with the given word. (case insensitive)
func (b *BookRepository) SearchWord(name string) {
	var books []Book
	b.db.Where("book_name ILIKE ? ", "%"+name+"%").Find(&books)

	for _, thebook := range books {
		fmt.Println(thebook.ToString())
	}
}

//Func GetByID prints book details of given ID
func (c *BookRepository) GetByID(id int) (*Book, error) {
	var book Book
	bookID := strconv.Itoa(id)
	result := c.db.Where("book_id = ?", bookID).First(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(book.ToString())

	return &book, nil
}

//Func GetByID hard deletes book details of given ID
func (c *BookRepository) DeleteById(id int) error {
	result := c.db.Delete(&Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *BookRepository) Migrations() {
	c.db.AutoMigrate(&Book{})
}

// Func InsertData starts a concurrent csv reading operation and write them to database.
func (c *BookRepository) InsertData() {
	csv_utils.ReadBooksWithWorkerPool(csvfile)
	books := []Book{}
	for _, book := range csv_utils.BookList {
		newItem := Book{
			Book_ID:    book.Book_ID,
			Book_Name:  book.Book_Name,
			Book_Price: book.Book_Price,
			Book_Page:  book.Book_Page,
			Book_Stock: book.Book_Stock,
			Book_Scode: book.Book_Scode,
			Book_ISBN:  book.Book_ISBN,
			Author:     book.Author}
		books = append(books, newItem)
	}
	for _, eachBook := range books {
		c.db.Unscoped().Where(Book{Book_Name: eachBook.Book_Name}).FirstOrCreate(&eachBook)
	}

}
