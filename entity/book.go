package entity

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookDTO struct {
	Title       string
	Author      string
	Description string
}

func ToBook(dto BookDTO) Book {
	return Book{
		Title:       dto.Title,
		Author:      dto.Author,
		Description: dto.Description,
	}
}
