package book

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Book_ID    string `gorm:"primaryKey"`
	Book_Name  string
	Book_Price string
	Book_Page  string
	Book_Stock int
	Book_Scode string
	Book_ISBN  string
	Author     string `gorm:"type:varchar(100);column:Author"`
}

func (Book) TableName() string {
	return "Book"
}

func (c *Book) ToString() string {
	return fmt.Sprintf("Book ID:%s,  Name:%s,  Price:%s,  Page:%s ,  Author:%s", c.Book_ID, c.Book_Name, c.Book_Price, c.Book_Page, c.Author)
}

func (c *Book) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Book (%s) deleting...", c.Book_ID)
	return nil
}
