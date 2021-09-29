package main

import (
	"fmt"
	"net/http"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	books := map[string]Book{
		"001": {Title: "Doraemon", Author: "Fujiko F. Fujio"},
		"002": {Title: "Harry Potter", Author: "J.K. Rowling"},
		"003": {Title: "The Lord of The Ring", Author: "J.R.R. Tolkien"},
	}

	// curl --location --request GET 'http://localhost:8080/getBook?code=001'
	http.HandleFunc("/getBook", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		bookCode := query.Get("code")
		book, exists := books[bookCode]
		if !exists {
			w.WriteHeader(404)
			w.Write([]byte("book not found"))
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Title: %s, Author: %s\n", book.Title, book.Author)))
	})

	// curl --location --request POST 'http://localhost:8080/addBook' \
	// --header 'Content-Type: application/x-www-form-urlencoded' \
	// --data-urlencode 'code=004' \
	// --data-urlencode 'title=Bleach' \
	// --data-urlencode 'author=Tite Kubo'
	http.HandleFunc("/addBook", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			w.Write([]byte("internal server error"))
			return
		}
		newBookCode := r.PostForm.Get("code")
		var newBook Book = Book{}
		newBook.Title = r.PostForm.Get("title")
		newBook.Author = r.PostForm.Get("author")
		books[newBookCode] = newBook
		fmt.Println(books)
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", nil)

}
