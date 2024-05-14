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
	{ID: "3", Title: "Blockchain Go", Author: "Suradath Bangnikrai", Year: "2022"},
}
