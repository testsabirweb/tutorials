package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/testsabirweb/tutorials/go/crud/handlers"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", handlers.GetBook).Methods(http.MethodGet)

	router.HandleFunc("/book", handlers.AddBook).Methods(http.MethodPost)

	router.HandleFunc("/book/{id}", handlers.UpdateBook).Methods(http.MethodPut)

	router.HandleFunc("/book/{id}", handlers.DeleteBook).Methods(http.MethodDelete)

	log.Println("Api is running")

	http.ListenAndServe(":4000", router)

}
