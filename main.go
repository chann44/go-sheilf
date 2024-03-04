package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Writer *Writer `json:"writter"`
}

type Writer struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}

var books []Book

func main() {
	r := mux.NewRouter()
	books = append(books, Book{Title: "HELL", ID: "12", Writer: &Writer{
		ID:   "2",
		Name: "VIKASH",
	}})
	r.HandleFunc("/", GetAllBooks).Methods("GET")
	r.HandleFunc("/delete/{id}", DeleteBook).Methods("GET")
	r.HandleFunc("/book/{id}", GetBook).Methods("GET")
	r.HandleFunc("/c", CreateBook).Methods("POST")

	http.ListenAndServe(":8000", r)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode("REMOVEd the book")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
		json.NewEncoder(w).Encode("Book Not found")

	}

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}
