package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var repo Repo
var app App

func TestInitializeRepo(t *testing.T) {
	_, _ = repo.Initialize("mongodb://localhost", "mongo-testing", context.TODO())

	_, err := repo.DB.Collection("testing-initialize").InsertOne(context.TODO(), bson.D{{"test", "works"}})

	assertResponse("unexpected error when inserting data", err, nil, t)

	var result bson.D
	err = repo.DB.Collection("testing-initialize").FindOne(context.TODO(), bson.D{{"test", "works"}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		t.Errorf("No result")
	}

	destroyDatabase(repo.DB, t)
}

func TestListTodosEmpty(t *testing.T) {
	repo.Initialize("mongodb://localhost", "mongo-testing", context.TODO())
	app.Initialize(context.TODO(), &repo)

	rr, err := recorderHandler("GET", "/todos", nil, app.RetrieveTodosHandler)

	if err != nil {
		t.Fatal(err)
	}

	expected := "[]"

	assertResponse("unexpected status code", rr.Code, http.StatusOK, t)
	assertResponse("unexpected Content-Type", rr.Header().Get("Content-Type"), "application/json", t)
	assertResponse("handler returned unexpected body", strings.TrimSpace(rr.Body.String()), expected, t)
}

func TestCreateTodo(t *testing.T) {
	_, _ = repo.Initialize("mongodb://localhost", "mongo-testing", context.TODO())
	app.Initialize(context.TODO(), &repo)

	var requestBody = []byte(`{"text":"First task","status":1,"priority":1}`)

	rr, _ := recorderHandler("POST", "/todos", bytes.NewBuffer(requestBody), app.CreateTodoHandler)

	assertResponse("unexpected status code", rr.Code, http.StatusCreated, t)

	var m map[string]interface{}
	_ = json.Unmarshal(rr.Body.Bytes(), &m)

	assertResponse("text should be the same: ", m["text"], "First task", t)
	assertResponse("status should be the same: ", m["status"], float64(1), t)
	assertResponse("priority should be the same: ", m["priority"], float64(1), t)

	destroyDatabase(repo.DB, t)
}

func TestListTodos(t *testing.T) {
	_, _ = repo.Initialize("mongodb://localhost", "mongo-testing", context.TODO())
	app.Initialize(context.TODO(), &repo)

	var requestBody = []byte(`{"text":"First ToDo","status":1,"priority":1}`)

	_, _ = recorderHandler("POST", "/todos", bytes.NewBuffer(requestBody), app.CreateTodoHandler)

	rr, err := recorderHandler("GET", "/todos", nil, app.RetrieveTodosHandler)

	if err != nil {
		t.Fatal(err)
	}

	if err != nil {
		t.Error("json Marshal error")
	}

	assertResponse("unexpected status code", rr.Code, http.StatusOK, t)
	assertResponse("unexpected Content-Type", rr.Header().Get("Content-Type"), "application/json", t)
	if strings.TrimSpace(rr.Body.String()) == "[]" {
		t.Fatal("empty list of ToDos")
	}

	destroyDatabase(repo.DB, t)
}

func TestUpdateTodo(t *testing.T) {
	// Create Todo
	requestBody := []byte(`{"text":"Not updated","status":1,"priority":1}`)

	_, _ = recorderHandler("POST", "/todos", bytes.NewBuffer(requestBody), app.CreateTodoHandler)

	// Check if it was created

	rr, _ := recorderHandler("GET", "/todos", nil, app.RetrieveTodosHandler)

	createdTodo := getTodo(rr.Body.Bytes(), 0)
	assertResponse("text should be the same: ", createdTodo.Text, "Not updated", t)

	// Update Todo

	requestBody = []byte(`{"text":"Updated","status":1,"priority":1}`)

	req, _ := http.NewRequest("PUT", "/todos/{id}", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": createdTodo.ID.Hex()})

	rr = httptest.NewRecorder()
	handler := http.HandlerFunc(app.UpdateTodoHandler)
	handler.ServeHTTP(rr, req)

	// Check if it was updated

	rr, _ = recorderHandler("GET", "/todos", nil, app.RetrieveTodosHandler)

	updatedTodo := getTodo(rr.Body.Bytes(), 0)
	assertResponse("text should should be: ", updatedTodo.Text, "Updated", t)

	destroyDatabase(repo.DB, t)
}

func TestDeleteTodo(t *testing.T) {
	// Create Todo
	requestBody := []byte(`{"text":"ToDo to delete","status":1,"priority":1}`)

	_, _ = recorderHandler("POST", "/todos", bytes.NewBuffer(requestBody), app.CreateTodoHandler)

	// Check if it was created

	rr, _ := recorderHandler("GET", "/todos", nil, app.RetrieveTodosHandler)

	createdTodo := getTodo(rr.Body.Bytes(), 0)
	assertResponse("text should be the same: ", createdTodo.Text, "ToDo to delete", t)

	// Update Todo

	req, _ := http.NewRequest("DELETE", "/todos/{id}", nil)
	req = mux.SetURLVars(req, map[string]string{"id": createdTodo.ID.Hex()})

	rr = httptest.NewRecorder()
	handler := http.HandlerFunc(app.DeleteTodoHandler)
	handler.ServeHTTP(rr, req)

	// Check if it was updated

	rr, _ = recorderHandler("GET", "/todos", nil, app.RetrieveTodosHandler)

	firstTodo := getTodo(rr.Body.Bytes(), 0)
	assertResponse("text should should be: ", firstTodo, (*ToDo)(nil), t)

	destroyDatabase(repo.DB, t)
}

func destroyDatabase(db *mongo.Database, t *testing.T) {
	err := db.Drop(context.TODO())

	if err != nil {
		t.Fatal("Failed to destroy database")
	}
}

func assertResponse(description string, received any, expected any, t *testing.T) {
	if received != expected {
		t.Error(reflect.TypeOf(received), reflect.TypeOf(expected))
		t.Errorf("%v: %v, but got %v", description, expected, received)
	}
}

func recorderHandler(method string, url string, body io.Reader, handlerFunc http.HandlerFunc) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, body)

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(rr, req)
	return rr, err
}

func getTodo(bytes []byte, pos int64) *ToDo {
	var todos []ToDo
	_ = json.Unmarshal(bytes, &todos)

	if len(todos) == 0 {
		return nil
	}

	return &todos[pos]
}
