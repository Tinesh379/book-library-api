package services

import (
	"book-library-api/dao"
	"book-library-api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dao.BooksList)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := dao.FindBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not found with this Id"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook dto.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	dao.BooksList = append(dao.BooksList, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func CheckoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing 'id' query parameter"})
		return
	}
	book, err := dao.FindBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

func ReturnBook(c *gin.Context) {

	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing 'id' query parameter"})
		return
	}
	book, err := dao.FindBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}
