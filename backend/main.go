package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	repo := Repo{}
	repo.Initialize(os.Getenv("MONGODB_URI"), "todo-app", ctx)

	repo.GetToDoList()

	r := mux.NewRouter()
	r.HandleFunc("/todos", RetrieveTodosHandler)
	err := http.ListenAndServe(":4000", r)

	if err != nil {
		log.Fatal(err)
	}
}

func RetrieveTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := []ToDo{}

	err := json.NewEncoder(w).Encode(todos)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
