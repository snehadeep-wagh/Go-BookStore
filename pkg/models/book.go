package models

import (
	"github.com/jinzhu/gorm"
	"github.com/snehadeep-wagh/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	// gorm.Model
	Id int64 `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectToDB()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func CreateBook(b *Book) *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	return book, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
