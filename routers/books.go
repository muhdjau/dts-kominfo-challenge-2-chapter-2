package routers

import (
	"challenge-chapter-2-sesi-2/controllers"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()

	r.POST("/books", controllers.AddBook)
	r.GET("/books", controllers.GetAllBooks)
	r.GET("/books/:bookID", controllers.GetBookById)
	r.PUT("/books/:bookID", controllers.UpdateBook)
	r.DELETE("/books/:bookID", controllers.DeleteBook)

	return r
}
