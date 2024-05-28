package models

import (
	"fmt"

	"github/fokaz-c/go-book-manager/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	if err := db.AutoMigrate(&Book{}); err != nil {
		fmt.Printf("Failed to migrate Book schema: %v\n", err)
	}
}

func (b *Book) CreateBook() (*Book, error) {
	result := db.Create(b)
	if result.Error != nil {
		return nil, result.Error
	}
	return b, nil
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Delete(&book)
	return book
}
