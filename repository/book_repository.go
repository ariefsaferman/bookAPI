package repository

import (
	"bookAPI/db"
	"bookAPI/entity"
	"errors"
	"log"

	"gorm.io/gorm"
)

func GetAllBook(books []*entity.Book) ([]*entity.Book, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}

	err = db.Find(&books).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("book is not found")
		return nil, errors.New("error geting data")
	}

	return books, nil
}

func AddBook(book *entity.Book) (*entity.Book, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}

	err = db.Create(&book).Error
	if err != nil {
		log.Println("error insert a book")
		return nil, errors.New("error insert a book")
	}

	return book, nil
}

func GetBookById(id uint) (*entity.Book, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}

	var book *entity.Book
	err = db.Where("id = ?", id).First(&book).Error
	if err != nil {
		log.Println("error find a book with id ", id)
		return nil, errors.New("error find a book")
	}

	return book, nil
}

func UpdateBookById(id uint, updatedBook *entity.Book) (*entity.Book, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}

	book, err := GetBookById(id)
	if err != nil {
		log.Println("error find a book with id ", id)
		return nil, errors.New("error find a book")
	}

	book.Author = updatedBook.Author
	book.Title = updatedBook.Title
	book.Description = updatedBook.Description

	err = db.Save(&book).Error
	if err != nil {
		log.Println("error updating book")
		return nil, errors.New("error update a book")
	}

	return book, nil
}

func DeleteBook(id uint) error {
	var book *entity.Book

	db, err := db.ConnectDB()
	if err != nil {
		return err
	}

	err = db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		log.Println("error updating book")
		return err
	}

	return nil
}
