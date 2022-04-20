package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rouling"},
	{ID: "2", Title: "The Lord of the king", Author: "J. R. R. Tolkien"},
}

// GET all books
func booksPage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, books)
}

// POST a new book
func PostBook(ctx *gin.Context) {
	var book Book
	// Receive request body here
	// myString := "Hi"
	// fmt.Println(*&myString)  // prints "Hi"
	err := ctx.ShouldBind(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	books = append(books, book)
	ctx.JSON(http.StatusCreated, books)

}

// DELETE
func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	ctx.Status(http.StatusNoContent)
}

// GET single book
func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, book := range books {
		if book.ID == id {
			book:= books[i]
			ctx.JSON(http.StatusOK,book)
			break
		}
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": "item not found",
	})
}

func main() {
	fmt.Println("hello gin")
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	// GET all books route
	r.GET("/books", booksPage)
	r.POST("/book", PostBook)
	// DELETE book by ID
	r.DELETE("/books/:id", Delete)

	// GET single book
	r.GET("/books/:id",GetBook)

	r.Run()
}
