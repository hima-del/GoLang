//https://www.youtube.com/watch?v=SonwZ6MF5BE
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

//function to get all books

func getBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {
	//init router
	r := mux.NewRouter()
	r.HandleFunc("/getbook", getBook)
}
