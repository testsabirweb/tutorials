package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/testsabirweb/tutorials/go/crud/mocks"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatalln(err)
	}

	for index, book := range mocks.Books {
		if book.Id == id {
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
