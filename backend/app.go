package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	*Repo
	Router *mux.Router
}

func (a *App) Initialize(ctx context.Context, repo *Repo) {
	a.Repo = repo

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	err := http.ListenAndServe(addr, a.Router)

	if err != nil {
		log.Fatal(err)
	}

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/todos", a.RetrieveTodosHandler).Methods("GET")
}

func (a *App) RetrieveTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := []ToDo{}

	a.Repo.GetToDoList()

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
