package main

type ToDo struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Status   int8   `json:"done"`
	Priority int8   `json:"priority"`
}
