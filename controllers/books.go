package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Books struct {
	BookID      int    `json:"book_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}

var BooksData = []Books{}

func AddBook(c *gin.Context) {
	var newBook Books

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = len(BooksData) + 1
	BooksData = append(BooksData, newBook)

	c.JSON(http.StatusCreated, gin.H{
		"data": newBook,
	})
}

func GetAllBooks(c *gin.Context) {
	var Books = BooksData
	// var AllDatas []Books
	// condition := false

	// if len(BooksData) > 0 {
	// 	condition = true
	// 	AllDatas = append(AllDatas, BooksData...)
	// }

	// if !condition {
	// 	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
	// 		"error_status":  "Data not found",
	// 		"error_message": "Data is empty",
	// 	})
	// }

	c.JSON(http.StatusOK, gin.H{
		// "books": AllDatas,
		"data": Books,
	})
}

func GetBookById(c *gin.Context) {
	bookID := c.Param("bookID")
	convBookID, _ := strconv.Atoi(bookID)
	condition := false
	var bookData Books

	for i, book := range BooksData {
		if convBookID == book.BookID {
			condition = true
			bookData = BooksData[i]
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookData,
	})
}

func UpdateBook(c *gin.Context) {
	bookID := c.Param("bookID")
	convBookID, _ := strconv.Atoi(bookID)
	condition := false
	var updatedBook Books

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BooksData {
		if convBookID == book.BookID {
			condition = true
			BooksData[i] = updatedBook
			BooksData[i].BookID = convBookID
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully updated", bookID),
	})
}

func DeleteBook(c *gin.Context) {
	bookID := c.Param("bookID")
	convBookID, _ := strconv.Atoi(bookID)
	condition := false
	var bookIndex int

	for i, book := range BooksData {
		if convBookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	copy(BooksData[bookIndex:], BooksData[bookIndex+1:])
	BooksData[len(BooksData)-1] = Books{}
	BooksData = BooksData[:len(BooksData)-1]

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully deleted", bookID),
	})
}
