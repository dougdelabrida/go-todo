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
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/todos", a.RetrieveTodosHandler).Methods("GET")
	a.Router.HandleFunc("/todos", a.CreateTodoHandler).Methods("POST")
}

func (a *App) RetrieveTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := make([]ToDo, 0)

	if res := a.Repo.GetToDoList(); res != nil {
		todos = res
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(todos)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *App) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo ToDo

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&todo)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	createdTodo, err := a.Repo.CreateToDo(todo)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(createdTodo)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(jsonData)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
