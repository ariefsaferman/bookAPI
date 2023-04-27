package handler

import (
	"bookAPI/entity"
	"bookAPI/repository"
	"bookAPI/utils/response"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": repository.DataBooks,
	})
}

func AddBook(ctx *gin.Context) {
	var newBook entity.BookDTO

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad request").Error(), err.Error())
		return
	}

	book := entity.ToBook(newBook)
	book.ID = len(repository.DataBooks) + 1
	repository.DataBooks = append(repository.DataBooks, book)

	response.SendSuccess(ctx, http.StatusOK, book)
}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIdInt, _ := strconv.Atoi(bookID)
	condition := false
	var book entity.Book

	for _, val := range repository.DataBooks {
		if val.ID == bookIdInt {
			book = val
			condition = true
			break
		}

	}

	if !condition {
		response.SendError(ctx, http.StatusNotFound, errors.New("BAD_REQUEST").Error(), "book is not found")
		return
	}

	response.SendSuccess(ctx, http.StatusOK, book)

}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIdInt, _ := strconv.Atoi(bookID)
	condition := false
	var book entity.BookDTO

	if err := ctx.ShouldBindJSON(&book); err != nil {
		response.SendError(ctx, http.StatusNotFound, errors.New("BAD_REQUEST").Error(), "bad request body")
		return
	}

	updatedBook := entity.ToBook(book)
	for i, val := range repository.DataBooks {
		if bookIdInt == val.ID {
			updatedBook.ID = val.ID
			condition = true
			repository.DataBooks[i] = updatedBook
			break
		}
	}

	if !condition {
		response.SendError(ctx, http.StatusNotFound, errors.New("BAD_REQUEST").Error(), "book is not found")
		return
	}

	response.SendSuccess(ctx, http.StatusOK, updatedBook)

}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIdInt, _ := strconv.Atoi(bookID)
	condition := false

	for i, val := range repository.DataBooks {
		if bookIdInt == val.ID {
			repository.DataBooks = append(repository.DataBooks[:i], repository.DataBooks[i+1:]...)
			condition = true
			break
		}
	}

	if !condition {
		response.SendError(ctx, http.StatusNotFound, errors.New("BAD_REQUEST").Error(), "book is not found")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v successfuly deleted", bookIdInt),
	})
}
