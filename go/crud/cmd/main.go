package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/testsabirweb/tutorials/go/crud/pkg/db"
	"github.com/testsabirweb/tutorials/go/crud/pkg/handlers"
)

func main() {

	DB := db.Init()
	h := handlers.New(DB)

	router := mux.NewRouter()
	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", h.GetBook).Methods(http.MethodGet)

	router.HandleFunc("/book", h.AddBook).Methods(http.MethodPost)

	router.HandleFunc("/book/{id}", h.UpdateBook).Methods(http.MethodPut)

	router.HandleFunc("/book/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("Api is running")

	http.ListenAndServe(":4000", router)

}
