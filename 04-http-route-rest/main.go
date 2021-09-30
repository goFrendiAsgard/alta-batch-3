package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		jsonDecoder := json.NewDecoder(r.Body)
		bookCode := strings.TrimPrefix(r.URL.Path, "/books/")
		fmt.Println(r.Method)
		switch r.Method {

		// curl --location --request GET 'http://localhost:8080/books/'
		// curl --location --request GET 'http://localhost:8080/books/001'
		case "GET":
			if bookCode == "" {
				jsonBooks, err := json.Marshal(books)
				if err != nil {
					WriteInternalServerError(w, err)
					return
				}
				w.WriteHeader(200)
				w.Write(jsonBooks)
				return
			}
			book, exists := books[bookCode]
			if !exists {
				w.WriteHeader(404)
				w.Write([]byte("book not found"))
				return
			}
			jsonBook, err := json.Marshal(book)
			if err != nil {
				WriteInternalServerError(w, err)
				return
			}
			w.WriteHeader(200)
			w.Write(jsonBook)
			return

		// curl --location --request POST 'http://localhost:8080/books/' \
		// --header 'Content-Type: application/json' \
		// --data-raw '{
		//     "code": "004",
		//     "title": "Bleach",
		//     "author": "Tite Kubo"
		// }'
		case "POST":
			bookWithCode := BookWithCode{}
			if err := jsonDecoder.Decode(&bookWithCode); err != nil {
				WriteInternalServerError(w, err)
				return
			}
			newBookCode := bookWithCode.Code
			book := Book{Title: bookWithCode.Title, Author: bookWithCode.Author}
			books[newBookCode] = book
			fmt.Println(books)
			jsonBook, err := json.Marshal(book)
			if err != nil {
				WriteInternalServerError(w, err)
				return
			}
			w.WriteHeader(201)
			w.Write(jsonBook)
			return

		// curl --location --request PUT 'http://localhost:8080/books/004' \
		// --header 'Content-Type: application/json' \
		// --data-raw '{
		//     "title": "Naruto",
		//     "author": "Masashi Kishimoto"
		// }'
		case "PUT":
			book := Book{}
			if err := jsonDecoder.Decode(&book); err != nil {
				WriteInternalServerError(w, err)
				return
			}
			books[bookCode] = book
			fmt.Println(books)
			jsonBook, err := json.Marshal(book)
			if err != nil {
				WriteInternalServerError(w, err)
				return
			}
			w.WriteHeader(200)
			w.Write(jsonBook)
			return

		// curl --location --request DELETE 'http://localhost:8080/books/004'
		case "DELETE":
			deletedBook := books[bookCode]
			delete(books, bookCode)
			jsonBook, err := json.Marshal(deletedBook)
			if err != nil {
				WriteInternalServerError(w, err)
				return
			}
			w.WriteHeader(200)
			w.Write(jsonBook)
			return
		}

		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func WriteInternalServerError(w http.ResponseWriter, err error) {
	fmt.Println(err)
	w.WriteHeader(500)
	w.Write([]byte("internal server error"))
}
