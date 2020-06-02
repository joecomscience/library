package api

import (
	"fmt"
	"github.com/gorilla/context"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request)  {
	c := context.Get(r, "postgres")
	fmt.Printf("context: %v\n", c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("book"))
}