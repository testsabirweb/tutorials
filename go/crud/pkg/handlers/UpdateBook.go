package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/testsabirweb/tutorials/go/crud/pkg/models"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
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

	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	book.Author = updatedBook.Author
	book.Title = updatedBook.Title
	book.Desc = updatedBook.Desc

	if result := h.DB.Save(&book); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
