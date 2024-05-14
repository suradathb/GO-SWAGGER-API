// controllers/book_controller.go
package controllers

import (
	"net/http"

	"go-swagger-api/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, models.Books)
}

// @Summary Get a book
// @Description Get details of a book by ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Router /books/{id} [get]
func GetBook(c *gin.Context) {
	id := c.Param("id")
	for _, book := range models.Books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// @Summary Create a new book
// @Description Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book to create"
// @Success 201 {object} models.Book
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.Books = append(models.Books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// @Summary Update a book
// @Description Update a book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body models.Book true "Book to update"
// @Success 200 {object} models.Book
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range models.Books {
		if book.ID == id {
			if err := c.ShouldBindJSON(&book); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			models.Books[i] = book
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// @Summary Delete a book
// @Description Delete a book by ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 204 {string} string "No Content"
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range models.Books {
		if book.ID == id {
			models.Books = append(models.Books[:i], models.Books[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
