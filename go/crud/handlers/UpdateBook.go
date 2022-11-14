package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/testsabirweb/tutorials/go/crud/mocks"
	"github.com/testsabirweb/tutorials/go/crud/models"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalln(err)
	}
	var updatedBook models.Book
	err = json.Unmarshal(body, &updatedBook)
	if err != nil {
		log.Fatalln(err)
	}

	for index, book := range mocks.Books {
		if book.Id == id {
			book.Author = updatedBook.Author
			book.Title = updatedBook.Title
			book.Desc = updatedBook.Desc

			mocks.Books[index] = book

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
