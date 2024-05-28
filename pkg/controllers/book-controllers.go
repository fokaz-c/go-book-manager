package controllers

import (
	"encoding/json"
	"fmt"
	"github/fokaz-c/go-book-manager/pkg/models"
	"github/fokaz-c/go-book-manager/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error parsing")
	}
	bookDetails, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	newBook := &models.Book{}

	if err := utils.ParseBody(r, newBook); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	createdBook, err := newBook.CreateBook()
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(createdBook)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error Deleting")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	booksID := vars["bookID"]
	ID, err := strconv.ParseInt(booksID, 0, 0)
	if err != nil {
		fmt.Println("Error Updating")
	}
	booksDetail, db := models.GetBookByID(ID)
	if updateBook.Name != "" {
		booksDetail.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		booksDetail.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		booksDetail.Publication = updateBook.Publication
	}

	db.Save(&booksDetail)
	res, _ := json.Marshal(booksDetail)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
