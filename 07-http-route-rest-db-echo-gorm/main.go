package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
}

func main() {
	connectionString := "root:toor@tcp(localhost:3306)/alta?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Book{})

	// https://github.com/go-gorm/gorm/issues/3145#issuecomment-658230163

	e := echo.New()

	// curl --location --request GET 'http://localhost:8080/books/1'
	e.GET("/books/:code", func(c echo.Context) error {
		bookId, err := strconv.Atoi(c.Param("code"))
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusBadRequest, "invalid id")
		}
		var book Book
		if err := db.First(&book, bookId).Error; err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		if book.ID == 0 {
			return c.String(http.StatusNotFound, "book not found")
		}
		return c.JSON(http.StatusOK, book)
	})

	// curl --location --request GET 'http://localhost:8080/books/'
	e.GET("/books/", func(c echo.Context) error {
		var books []Book
		if err := db.Find(&books).Error; err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		return c.JSON(http.StatusOK, books)
	})

	// curl --location --request POST 'http://localhost:8080/books/' \
	// --header 'Content-Type: application/json' \
	// --data-raw '{
	//     "id": 4,
	//     "title": "Bleach",
	//     "author": "Tite Kubo"
	// }'
	e.POST("books/", func(c echo.Context) error {
		book := Book{}
		// fmt.Printf("Book sebelum bind %#v\n", book)
		if err := c.Bind(&book); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		// fmt.Printf("Book setelah bind %#v\n", book)
		fmt.Printf("Before insert: %#v\n", book)
		if err := db.Save(&book).Error; err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		return c.JSON(http.StatusOK, book)
	})

	// curl --location --request PUT 'http://localhost:8080/books/4' \
	// --header 'Content-Type: application/json' \
	// --data-raw '{
	//     "title": "Naruto",
	//     "author": "Masashi Kishimoto"
	// }'
	e.PUT("books/:code", func(c echo.Context) error {
		bookId, err := strconv.Atoi(c.Param("code"))
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusBadRequest, "invalid id")
		}
		fmt.Println("Isi bookId ", bookId)
		var book Book
		fmt.Printf("Isi book sebelum select %#v\n", book)
		if err := db.First(&book, bookId).Error; err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		if book.ID == 0 {
			return c.String(http.StatusNotFound, "book not found")
		}
		fmt.Printf("Isi book setelah select %#v\n", book)
		if err := c.Bind(&book); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		fmt.Printf("Isi book setelah bind %#v\n", book)
		fmt.Printf("Before update: %#v\n", book)
		if err := db.Save(&book).Error; err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		return c.JSON(http.StatusOK, book)
	})

	// curl --location --request DELETE 'http://localhost:8080/books/4'
	e.DELETE("books/:code", func(c echo.Context) error {
		bookId, err := strconv.Atoi(c.Param("code"))
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusBadRequest, "invalid id")
		}
		var book Book
		if err := db.First(&book, bookId).Error; err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		if book.ID == 0 {
			return c.String(http.StatusNotFound, "book not found")
		}
		if err := db.Delete(&book).Error; err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		return c.JSON(http.StatusOK, book)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
