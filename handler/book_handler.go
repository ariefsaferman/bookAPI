package handler

import (
	"bookAPI/db"
	"bookAPI/entity"
	"bookAPI/utils/response"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBook(ctx *gin.Context) {
	var books []entity.Book
	sqlStatement := `SELECT * FROM books;`

	rows, err := db.DB.Query(sqlStatement)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad query").Error(), err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book entity.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			log.Println("error: ", err)
			return
		}
		books = append(books, book)
	}

	response.SendSuccess(ctx, http.StatusOK, books)
}

func AddBook(ctx *gin.Context) {
	var newBook entity.BookDTO
	sqlStatement := `INSERT INTO books(title, author, description) 
	VALUES($1, $2, $3) Returning *`

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad request").Error(), err.Error())
		return
	}

	book := entity.ToBook(newBook)
	err := db.DB.QueryRow(sqlStatement, book.Title, book.Author, book.Description).Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad query").Error(), err.Error())
		return
	}

	response.SendSuccess(ctx, http.StatusOK, book)
}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIdInt, _ := strconv.Atoi(bookID)
	var book entity.Book
	sqlStatement := `SELECT * FROM books WHERE id = $1`

	err := db.DB.QueryRow(sqlStatement, bookIdInt).Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad query").Error(), err.Error())
		return
	}

	response.SendSuccess(ctx, http.StatusOK, book)

}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIdInt, _ := strconv.Atoi(bookID)
	var book entity.BookDTO
	sqlStatement := `UPDATE books SET title = $2, author = $3, description = $4 WHERE id = $1`

	if err := ctx.ShouldBindJSON(&book); err != nil {
		response.SendError(ctx, http.StatusNotFound, errors.New("BAD_REQUEST").Error(), "bad request body")
		return
	}

	updatedBook := entity.ToBook(book)
	updatedBook.ID = bookIdInt
	_, err := db.DB.Exec(sqlStatement, bookIdInt, updatedBook.Title, updatedBook.Author, updatedBook.Description)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad query").Error(), err.Error())
		return
	}

	response.SendSuccess(ctx, http.StatusOK, updatedBook)

}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIdInt, _ := strconv.Atoi(bookID)
	sqlStatement := `DELETE FROM books WHERE id = $1;`

	_, err := db.DB.Exec(sqlStatement, bookIdInt)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad query").Error(), err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v successfuly deleted", bookIdInt),
	})
}
