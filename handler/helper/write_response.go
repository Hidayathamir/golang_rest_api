package helper

import "github.com/gin-gonic/gin"

func WriteRespon(c *gin.Context, status int, message string, data any) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}
