package routers

import (
	"github.com/gin-gonic/gin"
	"dts-fga-swswg/go-microservices/tugas-6/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// book routes
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}