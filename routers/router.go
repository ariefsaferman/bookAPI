package routers

import (
	"bookAPI/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// list end point
	router.GET("/books", handler.GetAllBook)
	router.POST("/book", handler.AddBook)
	router.GET("/book/:bookID", handler.GetBookById)
	router.PUT("/book/:bookID", handler.UpdateBook)
	router.DELETE("/book/:bookID", handler.DeleteBook)

	return router
}
