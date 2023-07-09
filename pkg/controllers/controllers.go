package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// "context"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	"github.com/Shashank-Panda/crud-mongoDB/pkg/models"
	"github.com/Shashank-Panda/crud-mongoDB/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	library := models.GetAllBooks()
	res, _ := json.Marshal(library)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	//since CreateBook is a struct of type Book defined in models
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	models.DeleteBook(Id)
	// res, _:=json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/jason")
	w.WriteHeader(http.StatusOK)
	// w.Write("Book deleted")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var UpdateBook = &models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	res, _ := json.Marshal(models.UpdateBook(ID, *UpdateBook))
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// BookDetails, db := models.GetBookById(ID)
	// if UpdateBook.Title != ""{
	// 	BookDetails.Title = UpdateBook.Title
	// }
	// if UpdateBook.Author != ""{
	// 	BookDetails.Author = UpdateBook.Author
	// }
	// if UpdateBook.Isbn != ""{
	// 	BookDetails.Isbn = UpdateBook.Isbn
	// }
	// db.Save()
}
