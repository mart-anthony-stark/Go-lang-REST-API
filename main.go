package main

import(
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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

}

func main(){
	// Init router
	r := mux.NewRouter()

	// Route handlers enpoint
	r.HandleFuc("/books", getBooks).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}