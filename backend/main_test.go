package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListTodosEmpty(t *testing.T) {
	rr, err := recorderHandler("GET", "/todos", nil, RetrieveTodosHandler)

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
