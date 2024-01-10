package main

import (
	"book-library-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/books", services.GetBooks)
	router.POST("/api/books", services.GetBooks)
	router.Run("localhost:9090")
}
