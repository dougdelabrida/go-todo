package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var repo Repo
var app App

func TestInitializeRepo(t *testing.T) {
	repo.Initialize("mongodb://localhost", "mongo-testing", context.TODO())

	_, err := repo.DB.Collection("testing-initialize").InsertOne(context.TODO(), bson.D{{"test", "works"}})

	assertResponse("unexpected error when inserting data", err, nil, t)

	var result bson.D
	err = repo.DB.Collection("testing-initialize").FindOne(context.TODO(), bson.D{{"test", "works"}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		t.Errorf("No result")
	}

	repo.DB.Drop(context.TODO())
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

func assertResponse(description string, received any, expected any, t *testing.T) {
	if received != expected {
		t.Errorf("%v: %v, but got %v", description, expected, received)
	}
}

func recorderHandler(method string, url string, body io.Reader, handlerFunc http.HandlerFunc) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFunc)
	handler.ServeHTTP(rr, req)
	return rr, err
}
