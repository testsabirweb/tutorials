package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/testsabirweb/tutorials/go/crud/mocks"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}