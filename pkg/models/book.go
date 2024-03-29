package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satyalohit/Bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name         string `gorm:""json:"name"`
	Author       string `json:"author"`
	Publications string `json:"publications"`
}

func init() {
	config.Connect()
	db = config.Getdb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book

	db := db.Where("Id=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Find("Id=?", Id).Delete(book)
	return book
}
