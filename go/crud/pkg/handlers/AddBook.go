package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/testsabirweb/tutorials/go/crud/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Fatalln(err)
	}
	result := h.DB.Create(&book)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created!!!")
}
