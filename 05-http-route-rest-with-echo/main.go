package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookWithCode struct {
	Code string `json:"code"`
	Book
}

func main() {
	books := map[string]Book{
		"001": {Title: "Doraemon", Author: "Fujiko F. Fujio"},
		"002": {Title: "Harry Potter", Author: "J.K. Rowling"},
		"003": {Title: "The Lord of The Ring", Author: "J.R.R. Tolkien"},
	}

	e := echo.New()

	// curl --location --request GET 'http://localhost:8080/books/001'
	e.GET("/books/:code", func(c echo.Context) error {
		bookCode := c.Param("code")
		book, exist := books[bookCode]
		if !exist {
			return c.String(http.StatusNotFound, "book not found")
		}
		return c.JSON(http.StatusOK, book)
	})

	// curl --location --request GET 'http://localhost:8080/books/'
	e.GET("/books", func(c echo.Context) error {
		return c.JSON(http.StatusOK, books)
	})

	// curl --location --request POST 'http://localhost:8080/books/' \
	// --header 'Content-Type: application/json' \
	// --data-raw '{
	//     "code": "004",
	//     "title": "Bleach",
	//     "author": "Tite Kubo"
	// }'
	e.POST("books/", func(c echo.Context) error {
		bookWithCode := BookWithCode{}
		if err := c.Bind(bookWithCode); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		newBookCode := bookWithCode.Code
		book := Book{Title: bookWithCode.Title, Author: bookWithCode.Author}
		books[newBookCode] = book
		fmt.Println(books)
		return c.JSON(http.StatusOK, book)
	})

	// curl --location --request PUT 'http://localhost:8080/books/004' \
	// --header 'Content-Type: application/json' \
	// --data-raw '{
	//     "title": "Naruto",
	//     "author": "Masashi Kishimoto"
	// }'
	e.PUT("books/:code", func(c echo.Context) error {
		bookCode := c.Param("code")
		book := Book{}
		if err := c.Bind(book); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		books[bookCode] = book
		fmt.Println(books)
		return c.JSON(http.StatusOK, book)
	})

	// curl --location --request DELETE 'http://localhost:8080/books/004'
	e.DELETE("books/:code", func(c echo.Context) error {
		bookCode := c.Param("code")
		deletedBook := books[bookCode]
		delete(books, bookCode)
		return c.JSON(http.StatusOK, deletedBook)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
