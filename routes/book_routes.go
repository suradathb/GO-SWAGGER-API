// routes/book_routes.go
package routes

import (
	"go-swagger-api/controllers"

	"github.com/gin-gonic/gin"
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
