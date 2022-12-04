package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDo struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Text     string             `json:"text"`
	Status   int8               `json:"status"`
	Priority int8               `json:"priority"`
}
