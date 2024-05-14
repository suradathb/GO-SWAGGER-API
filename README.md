# GO-SWAGGER-API

-- Documentation Structure Project
go-swagger-api/
├── main.go
├── controllers/
│   └── book_controller.go
├── models/
│   └── book.go
├── routes/
│   └── book_routes.go
└── docs/
    └── ... (Swagger generated docs)


Step-by-Step Guide
Step 1: Create the Data Model
Create a file named book.go in the models directory.

go
Copy code
// models/book.go
package models

type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Year   string `json:"year"`
}

var Books = []Book{
    {ID: "1", Title: "Go Programming", Author: "John Doe", Year: "2021"},
    {ID: "2", Title: "Advanced Go", Author: "Jane Doe", Year: "2022"},
}
Step 2: Create Controller Functions
Create a file named book_controller.go in the controllers directory.

go
Copy code
// controllers/book_controller.go
package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "go-swagger-api/models"
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
Step 3: Set Up Routes
Create a file named book_routes.go in the routes directory.

go
Copy code
// routes/book_routes.go
package routes

import (
    "github.com/gin-gonic/gin"
    "go-swagger-api/controllers"
)

func RegisterBookRoutes(r *gin.Engine) {
    api := r.Group("/api/v1")
    {
        api.GET("/books", controllers.GetBooks)
        api.GET("/books/:id", controllers.GetBook)
        api.POST("/books", controllers.CreateBook)
        api.PUT("/books/:id", controllers.UpdateBook)
        api.DELETE("/books/:id", controllers.DeleteBook)
    }
}
Step 4: Set Up Main File
Update the main.go file to initialize the router and register the routes.

go
Copy code
// main.go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files" // swagger embed files
    "github.com/swaggo/gin-swagger" // gin-swagger middleware
    _ "go-swagger-api/docs" // swagger documentation
    "go-swagger-api/routes"
)

// @title Go Swagger API Example
// @version 1.0
// @description This is a sample server for a Go Swagger API example.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
    r := gin.Default()

    // Swagger route
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Register routes
    routes.RegisterBookRoutes(r)

    // Start server
    r.Run(":8080")
}
Generate Swagger Docs
If you haven't already, generate the Swagger documentation by running:

sh
Copy code
swag init
This will generate the necessary Swagger documentation in the docs directory.

Running the API
Run your API server:

sh
Copy code
go run main.go
Access the Swagger UI:
Open your browser and go to http://localhost:8080/swagger/index.html to view your API documentation.