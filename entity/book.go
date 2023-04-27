package entity

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

type BookDTO struct {
	Title  string
	Author string
	Desc   string
}

func ToBook(dto BookDTO) Book {
	return Book{
		Title:  dto.Title,
		Author: dto.Author,
		Desc:   dto.Desc,
	}
}
