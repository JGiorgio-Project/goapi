package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/giorg/goapi/dao"
	"github.com/giorg/goapi/model"
	"github.com/gorilla/mux"
)

//Get All Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	books, err := dao.GetBooks()
	if err == nil {
		json.NewEncoder(w).Encode(books)
		return
	} else {
		log.Println("ERequest error", err)
		json.NewEncoder(w).Encode(&model.Book{})
	}

}

//Get Single Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	param := mux.Vars(r) //Get params
	book, err := dao.GetBook(param["id"])

	if err == nil {
		json.NewEncoder(w).Encode(book)
		return
	} else {
		log.Println("ERequest error", err)
		json.NewEncoder(w).Encode(&model.Book{})
	}

}

//Create Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book, err := dao.CreateBook(book)
	if err != nil {
		log.Println("ERequest error", err)
		json.NewEncoder(w).Encode(&model.Book{})
	} else {
		json.NewEncoder(w).Encode(book)
		return
	}
}

//Update Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book, err := dao.UpdateBook(book)
	if err != nil {
		log.Println("Request error", err)
		json.NewEncoder(w).Encode(&model.Book{})
	} else {
		json.NewEncoder(w).Encode(book)
		return
	}
}

//Delete Book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	param := mux.Vars(r) //Get params
	err := dao.DeleteBook(param["id"])

	if err == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	} else {
		log.Println("ERequest error", err)
	}
}
