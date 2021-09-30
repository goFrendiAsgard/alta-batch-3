package main

import (
	"bufio"
	"fmt"
	"os"
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
	for {
		fmt.Println("1. Get book\n2. Add book\n3. Exit")
		userInput := Scanln()
		if userInput == "3" {
			break
		}
		switch userInput {
		case "1":
			fmt.Println("Book Code: ")
			bookCode := Scanln()
			book, exists := books[bookCode]
			if exists {
				fmt.Printf("Title: %s, Author: %s\n", book.Title, book.Author)
			} else {
				fmt.Println("book not found")
			}
		case "2":
			fmt.Println("New Book Code: ")
			newBookCode := Scanln()
			var newBook Book = Book{}
			fmt.Println("New Book Title: ")
			newBook.Title = Scanln()
			fmt.Println("New Book Author: ")
			newBook.Author = Scanln()
			books[newBookCode] = newBook
		}
	}
}

func Scanln() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		return line
	}
	return ""
}
