package handler

import (
	"bookAPI/entity"
	"bookAPI/repository"
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
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := entity.ToBook(newBook)
	book.ID = len(repository.DataBooks) + 1
	repository.DataBooks = append(repository.DataBooks, book)

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
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
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v is not found", bookIdInt),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})

}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIdInt, _ := strconv.Atoi(bookID)
	condition := false
	var book entity.BookDTO

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
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
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v is not found", bookIdInt),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": updatedBook,
	})

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
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v is not found", bookIdInt),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v successfuly deleted", bookIdInt),
	})
}
