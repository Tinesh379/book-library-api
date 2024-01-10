package controllers

import (
	"book-library-api/constants"
	"book-library-api/services"

	"github.com/gin-gonic/gin"
)

var api string = constants.API_PREFIX
func LibraryBooks() {
	router := gin.Default()
	router.GET(api, services.GetBooks)
	router.POST(api, services.CreateBook)
	router.GET(api+"/:id", services.GetBookById)
	router.Run("localhost:9090")
}
