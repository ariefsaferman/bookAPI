package entity

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
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
