package main

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	repo := new(Repo)

	client, err := repo.Initialize(os.Getenv("MONGODB_URI"), "todo-app", ctx)

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := App{}
	app.Initialize(ctx, repo)

	app.Run(":4000")
}
