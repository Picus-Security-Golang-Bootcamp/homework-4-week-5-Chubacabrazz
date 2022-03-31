package author

import (
	"fmt"

	"github.com/Chubacabrazz/book-db/file_services/csv_utils"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

//FindByWord lists the authors with the given word case-insensitively
func (a *AuthorRepository) FindByWord(name string) {
	var authors []Author
	a.db.Where("name ILIKE ? ", "%"+name+"%").Find(&authors)

	for _, author := range authors {
		fmt.Println(author.ToString())
	}
}

//Func List prints all authors from db
func (a *AuthorRepository) List() {
	var authors []Author
	a.db.Find(&authors)

	for _, author := range authors {
		fmt.Println(author.ToString())
	}
}

//DeleteByID does a soft delete to an author with given ID
func (a *AuthorRepository) DeleteById(id int) error {
	var author Author
	result := a.db.First(&author, id)
	if result.Error != nil {
		return result.Error
	} else {
		fmt.Println("Success, soft deleted:", id)
	}
	result = a.db.Delete(&Author{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Func Migration forms a book table in database
func (c *AuthorRepository) Migration() {
	c.db.AutoMigrate(&Author{})
}

// Func InsertData starts a concurrent csv reading operation and write them to database.
func (c *AuthorRepository) InsertData() {
	authors := []Author{}
	for _, book := range csv_utils.BookList {
		newItem := Author{
			Author_ID:   book.Author_ID,
			Author_Name: book.Author}

		authors = append(authors, newItem)
	}

	for _, eachAuthor := range authors {
		c.db.Unscoped().Where(Author{Author_ID: eachAuthor.Author_ID}).Attrs(Author{Author_Name: eachAuthor.Author_Name}).FirstOrCreate(&eachAuthor)
	}

}
