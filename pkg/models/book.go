package models

import (
	"github.com/jinzhu/gorm"
	"github.com/raghadasfour-02/Go-Myshelf/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	ID string `json:"id"`
	Title string `gorm:"" json:"title"`
	Author string `json:"author"`
}

// Connect to DB
func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBookById(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

