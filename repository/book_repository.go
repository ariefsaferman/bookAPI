package repository

import "bookAPI/entity"

var DataBooks []entity.Book

func Init() {
	DataBooks = []entity.Book{
		{ID: 1, Title: "Book of Happiness", Author: "JK Rowling", Desc: "Magic"},
		{ID: 2, Title: "Book of Sadness", Author: "Rowan Atkins", Desc: "Comedy"},
	}
}
