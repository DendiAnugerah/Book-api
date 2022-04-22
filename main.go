package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/book-api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}
	db.AutoMigrate(&book.Book{})
	//CRUD

	// --- Create ---
	// book := book.Book{}
	// book.Title = "Man Tiger"
	// book.Price = 90000
	// book.Discount = 10
	// book.Rating = 5
	// book.Description = "Ini adalah buku yang sangat bagus"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	log.Println("Error creating book record")
	// }

	// --- Declare --- 
	// var book book.Book

	// --- Read ---
	// err = db.Debug().First(&book).Error
	if err != nil {
		log.Println("Error creating book record")
	}
	// --- Update --- 
	// book.Title = "Man Tiger (Evolved)"
	// err = db.Save(&book).Error

	// --- Delete ---
	// err = db.delete


	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id/:title", handler.BookHandler)
	v1.GET("/hello", handler.HelloHanlder)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBookHandler)

	router.Run(":8888")
} 

