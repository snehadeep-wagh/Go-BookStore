package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/snehadeep-wagh/go-bookstore/pkg/models"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var books []models.Book
	books = models.GetAllBooks()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var book models.Book
	Params := mux.Vars(r)
	bookId := Params["bookId"]
	Id, _ := strconv.ParseInt(bookId, 0, 0)
	book, _ = models.GetBookById(Id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var book models.Book

	// decode the body
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(book.Author)
	// fmt.Println(book.Name)
	// fmt.Println(book.Publication)
	models.CreateBook(&book)

	json.NewEncoder(w).Encode(book)
	defer r.Body.Close()
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("content-type", "application/json")
	id, _ := strconv.ParseInt(params["bookId"], 0, 0)
	val := models.DeleteBook(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(val)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["bookId"], 0, 0)
	models.DeleteBook(id)
	var newBook models.Book
	json.NewDecoder(r.Body).Decode(&newBook)
	newBook.Id = id
	models.CreateBook(&newBook)
	json.NewEncoder(w).Encode(newBook)
}
