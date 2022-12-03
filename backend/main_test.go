package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeRouter(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HelloHandler)

	handler.ServeHTTP(rr, req)

	expected := "Hello World"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
