package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "todos"
const timeout = 10 * time.Second

type Repo struct {
	DB *mongo.Database
}

func (r *Repo) Initialize(uri string, databaseName string, ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	r.DB = client.Database(databaseName)

	if err != nil {
		panic("Connection with database failed")
	}

	return client, err
}

func (r *Repo) GetToDoList() []ToDo {
	var toDos []ToDo

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	collection := r.DB.Collection(collectionName)
	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	if err = cur.All(ctx, &toDos); err != nil {
		panic(err)
	}

	return toDos
}

func (r *Repo) CreateToDo(todo ToDo) (ToDo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	result, err := r.DB.Collection(collectionName).InsertOne(ctx, todo)

	if err != nil {
		panic(err)
	}

	return ToDo{
		ID:       result.InsertedID.(primitive.ObjectID),
		Text:     todo.Text,
		Status:   todo.Status,
		Priority: todo.Priority,
	}, err
}
