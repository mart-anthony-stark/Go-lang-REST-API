package main

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"	
)

// Book struct model
type Book struct{
	ID		string `json:"id"`
	Isbn	string `json:"isbn"`
	Title	string `json:"title"`
	Author	*Author `json:"author"`
}
// Author struct
type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Initial book
var books []Book

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main(){
	// Init router
	r := mux.NewRouter()
	
	// Hardcoded data - @todo: add database
	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handlers enpoint
	r.HandleFunc("/books", getBooks).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}