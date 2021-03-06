package handler

import (
	"fmt"
	"net/http"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Dendi Anugerah",
		"bio":  "Software Engineer",
	})
}

func HelloHanlder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":       "Hello World",
		"description": "Hi iam learning",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func PostBookHandler(c *gin.Context){
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil{

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on fields %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}