package main

import (
	"github.com/gorilla/mux"
	"github.com/joecomscience/library/pkg/api"
	"github.com/joecomscience/library/pkg/middleware"
	"github.com/joecomscience/library/pkg/postgres"
	"log"
	"net/http"
)


func main() {
	db := postgres.Init()
	m := middleware.ContextInjector{DB: db}

	r := mux.NewRouter()
	r.Use(m.Injector)
	r.HandleFunc("/book", api.GetBooks).Methods("GET")

	log.Println("start server by listen port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
