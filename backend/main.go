package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", RetrieveTodosHandler)
	err := http.ListenAndServe(":4000", r)

	if err != nil {
		log.Fatal(err)
	}
}

func RetrieveTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := []ToDo{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(todos)

	if err != nil {
		log.Fatal(err)
	}
}
