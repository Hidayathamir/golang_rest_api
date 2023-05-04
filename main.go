package main

import (
	"fmt"

	"github.com/Hidayathamir/golang_rest_api/database"
	"github.com/Hidayathamir/golang_rest_api/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		logger.GetLog().Error(err)
		panic(errors.Wrap(err, "main.main"))
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println(db)
		c.JSON(200, gin.H{
			"message": "paaaangg",
		})
	})
	r.Run()
}
