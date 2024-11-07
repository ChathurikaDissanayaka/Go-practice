// handlers/book_handler.go

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"crud-api/models"
)

var books []models.Book
var nextID = 1

// GetBooks handles GET requests and returns all books in JSON format
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook handles GET requests for a single book based on ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// CreateBook handles POST requests to add a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	book.ID = nextID
	nextID++

	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook handles PUT requests to update an existing book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if book.ID == id {
			var updatedBook models.Book
			json.NewDecoder(r.Body).Decode(&updatedBook)

			books[i].Title = updatedBook.Title
			books[i].Author = updatedBook.Author
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// DeleteBook handles DELETE requests to remove a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			fmt.Fprintf(w, "Book with ID %d was deleted", id)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
