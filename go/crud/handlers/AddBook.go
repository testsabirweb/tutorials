package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/testsabirweb/tutorials/go/crud/mocks"
	"github.com/testsabirweb/tutorials/go/crud/models"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Fatalln(err)
	}
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created!!!")
}
