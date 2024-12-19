package Library

import (
	"fmt"
)

// Book struct definition
type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

// Library struct maintaining a collection of books
type Library struct {
	books map[string]Book
}

// NewLibrary is a constructor for Library
func NewLibrary() *Library {
	return &Library{
		books: make(map[string]Book),
	}
}

// AddBook adds a new book to the library
func (l *Library) AddBook(book Book) {
	l.books[book.ID] = book
	fmt.Println("Book added:", book.Title)
}

// BorrowBook borrows a book by ID if available
func (l *Library) BorrowBook(id string) {
	b, ok := l.books[id]
	if !ok {
		fmt.Println("Book not found!")
		return
	}
	if b.IsBorrowed {
		fmt.Println("Book is already borrowed.")
		return
	}
	b.IsBorrowed = true
	l.books[id] = b
	fmt.Println("Book borrowed:", b.Title)
}

// ReturnBook returns a borrowed book by ID
func (l *Library) ReturnBook(id string) {
	b, ok := l.books[id]
	if !ok {
		fmt.Println("Book not found!")
		return
	}
	if !b.IsBorrowed {
		fmt.Println("Book is not borrowed.")
		return
	}
	b.IsBorrowed = false
	l.books[id] = b
	fmt.Println("Book returned:", b.Title)
}

// ListBooks prints all available (not borrowed) books
func (l *Library) ListBooks() {
	fmt.Println("Available books:")
	for _, b := range l.books {
		if !b.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", b.ID, b.Title, b.Author)
		}
	}
}
