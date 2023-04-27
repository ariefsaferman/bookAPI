package response

import "github.com/gin-gonic/gin"

func SendError(c *gin.Context, statusCode int, errCode string, msg string) {
	c.JSON(statusCode, gin.H{
		"code":    errCode,
		"message": msg,
	})
}

func SendSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}
