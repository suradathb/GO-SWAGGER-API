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


# Go Swagger API Example

This repository contains an example of a Go API with Swagger documentation using the Gin framework.

## Step-by-Step Guide

### Step 1: Create the Data Model

Create a file named `book.go` in the `models` directory.

```go
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
