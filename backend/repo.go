package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	DB  *mongo.Database
	ctx context.Context
}

func (r *Repo) Initialize(uri string, databaseName string, ctx context.Context) {
	r.ctx = ctx

	client, err := mongo.Connect(r.ctx, options.Client().ApplyURI(uri))

	r.DB = client.Database(databaseName)

	if err != nil {
		panic("Connection with database failed")
	}
}

func (r *Repo) GetToDoList() {
	collection := r.DB.Collection("todos")
	cur, err := collection.Find(r.ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(r.ctx) {
		var result bson.D

		err := cur.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		print(result)
		// do something with result....
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
