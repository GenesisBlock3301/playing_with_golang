package main

import (
	"github.com/gin-gonic/gin"
	"bookweb/controllers"
	"bookweb/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/books",controllers.FindBooks)
		v1.POST("/books",controllers.CreateBook)
		v1.GET("/books/:id",controllers.FindBook)
		v1.PUT("/books/:id",controllers.UpdateBook)
		v1.DELETE("/books/:id",controllers.DeleteBook)
	}
	// r.GET("/books",controllers.FindBooks)
	r.Run()
}
