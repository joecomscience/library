package main

import (
	"github.com/gorilla/mux"
	"github.com/joecomscience/library-api/pkg/books"
	"github.com/joecomscience/library-api/pkg/postgres"
	"log"
	"net/http"
)

func init() {
	postgres.Init()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/book", books.GetBooks).Methods("GET")

	log.Println("start server by listen port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
