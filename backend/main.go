package main

import (
	"context"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	repo := new(Repo)
	repo.Initialize(os.Getenv("MONGODB_URI"), "todo-app", ctx)

	app := App{}
	app.Initialize(ctx, repo)

	app.Run(":4000")
}
