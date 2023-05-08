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
	var books []*entity.Book

	result, err := repository.GetAllBook(books)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad request").Error(), err.Error())
		return
	}

	if len(result) == 0 {
		response.SendSuccess(ctx, http.StatusOK, books)
		return
	}

	response.SendSuccess(ctx, http.StatusOK, result)
}

func AddBook(ctx *gin.Context) {
	var newBook entity.BookDTO

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad request").Error(), err.Error())
		return
	}

	book := entity.ToBook(newBook)
	res, err := repository.AddBook(&book)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad query").Error(), err.Error())
		return
	}

	response.SendSuccess(ctx, http.StatusOK, res)
}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIdInt, _ := strconv.Atoi(bookID)

	res, err := repository.GetBookById(uint(bookIdInt))
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad query").Error(), err.Error())
		return
	}

	response.SendSuccess(ctx, http.StatusOK, res)

}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIdInt, _ := strconv.Atoi(bookID)
	var book entity.BookDTO

	if err := ctx.ShouldBindJSON(&book); err != nil {
		response.SendError(ctx, http.StatusNotFound, errors.New("BAD_REQUEST").Error(), "bad request body")
		return
	}

	updatedBook := entity.ToBook(book)

	res, err := repository.UpdateBookById(uint(bookIdInt), &updatedBook)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad query").Error(), err.Error())
		return
	}

	response.SendSuccess(ctx, http.StatusOK, res)

}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIdInt, _ := strconv.Atoi(bookID)

	err := repository.DeleteBook(uint(bookIdInt))
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, errors.New("bad query").Error(), err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v successfuly deleted", bookIdInt),
	})
}
