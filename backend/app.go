package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	a.Router.HandleFunc("/todos/{id}", a.UpdateTodoHandler).Methods("PUT")
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
		sendInternalServerError(w, err)
	}
}

func (a *App) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo ToDo

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&todo)

	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	defer r.Body.Close()

	createdTodo, err := a.Repo.CreateToDo(todo)

	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	jsonData, err := json.Marshal(createdTodo)

	if err != nil {
		sendInternalServerError(w, err)
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

func (a *App) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo ToDo

	vars := mux.Vars(r)

	log.Println(vars)
	log.Println(r.URL.Path)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	todo.ID = id

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&todo)

	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	defer r.Body.Close()

	updatedTodo, err := a.Repo.UpdateToDo(todo)

	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	jsonData, err := json.Marshal(updatedTodo)

	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonData)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func sendInternalServerError(w http.ResponseWriter, err error) {
	log.Fatal(err)
	w.WriteHeader(http.StatusInternalServerError)
}
