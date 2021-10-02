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
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
		return
	} else {
		log.Println("ERequest error", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&model.Book{})
	}

}

//Get Single Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	param := mux.Vars(r) //Get params
	book, err := dao.GetBook(param["id"])

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
		return
	} else {
		log.Println("ERequest error", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(book)
	}

}

//Create Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	var book model.Book
	err2 := json.NewDecoder(r.Body).Decode(&book)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERequest error", err2)
		json.NewEncoder(w).Encode(&model.Book{})
		return
	}
	book, err := dao.CreateBook(book)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERequest error", err)
		json.NewEncoder(w).Encode(&model.Book{})
	}
}

//Update Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	param := mux.Vars(r) //Get params
	var book model.Book
	err2 := json.NewDecoder(r.Body).Decode(&book)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERequest error", err2)
		json.NewEncoder(w).Encode(&model.Book{})
		return
	}
	if param["id"] != book.ID {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERequest error", err2)
		json.NewEncoder(w).Encode(&model.Book{})
		return
	}
	book, err := dao.UpdateBook(param["id"], book)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Request error", err)
		json.NewEncoder(w).Encode(&model.Book{})
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
		w.WriteHeader(http.StatusBadRequest)
	}
}
