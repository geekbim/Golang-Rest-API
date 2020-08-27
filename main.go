package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get All Book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Get Create Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Get Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Get Delete Books
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(
		books,
		Book{
			ID:    "1",
			Isbn:  "1337",
			Title: "Pemrograman Golang",
			Author: &Author{
				FirstName: "Abim",
				LastName:  "Gatya",
			},
		},
	)

	books = append(
		books,
		Book{
			ID:    "2",
			Isbn:  "1227",
			Title: "Pemrograman Javascript",
			Author: &Author{
				FirstName: "Dhanu",
				LastName:  "Widiyuta",
			},
		},
	)

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books", deleteBook).Methods("DELETE")

	// Start Server
	log.Fatal(http.ListenAndServe(":8000", r))
}
