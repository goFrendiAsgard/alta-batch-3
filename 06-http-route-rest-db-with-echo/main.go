package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type BookData struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Book struct {
	Id int `json:"id"`
	BookData
}

func main() {
	db, err := sql.Open("mysql", "root:toor@/alta")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	e := echo.New()

	// curl --location --request GET 'http://localhost:8080/books/001'
	e.GET("/books/:code", func(c echo.Context) error {
		bookId := c.Param("code")
		results, err := db.Query("SELECT id, title, author FROM books WHERE id=?", bookId)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		defer results.Close()
		exist := results.Next()
		if !exist {
			return c.String(http.StatusNotFound, "book not found")
		}
		var book Book
		if err := results.Scan(&book.Id, &book.Title, &book.Author); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		return c.JSON(http.StatusOK, book)
	})

	// curl --location --request GET 'http://localhost:8080/books/'
	e.GET("/books/", func(c echo.Context) error {
		books := []Book{}
		results, err := db.Query("SELECT id, title, author FROM books")
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		defer results.Close()
		for results.Next() {
			var book Book
			if err := results.Scan(&book.Id, &book.Title, &book.Author); err != nil {
				fmt.Println(err)
				return c.String(http.StatusInternalServerError, "internal server error")
			}
			books = append(books, book)
		}
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
		bookData := BookData{}
		if err := c.Bind(&bookData); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		insert, err := db.Query("INSERT INTO books(title, author) VALUES ( ?, ?)", bookData.Title, bookData.Author)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		defer insert.Close()
		return c.JSON(http.StatusOK, bookData)
	})

	// curl --location --request PUT 'http://localhost:8080/books/004' \
	// --header 'Content-Type: application/json' \
	// --data-raw '{
	//     "title": "Naruto",
	//     "author": "Masashi Kishimoto"
	// }'
	e.PUT("books/:code", func(c echo.Context) error {
		bookId := c.Param("code")
		bookData := BookData{}
		if err := c.Bind(&bookData); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		update, err := db.Query("UPDATE books SET title=?, author=? WHERE id=?", bookData.Title, bookData.Author, bookId)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		defer update.Close()
		return c.JSON(http.StatusOK, bookData)
	})

	// curl --location --request DELETE 'http://localhost:8080/books/004'
	e.DELETE("books/:code", func(c echo.Context) error {
		bookId := c.Param("code")
		delete, err := db.Query("DELETE FROM books WHERE id=?", bookId)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "internal server error")
		}
		defer delete.Close()
		return c.JSON(http.StatusOK, bookId)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
