package main

type ToDo struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	Status   int8   `json:"status"`
	Priority int8   `json:"priority"`
}
