package api

import "net/http"

func GetBooks(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("book"))
}