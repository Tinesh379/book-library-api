package services

import (
	"book-library-api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var newBook dto.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	dto.BooksList = append(dto.BooksList, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dto.BooksList)
}
