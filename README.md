# Todo App Next.js + Go

The app connects through a rest API built using gorilla mux and persit data in a mongodb instance.

## Run Locally

Before following these steps, make sure you have Node.js 14.6.0 or newer, Go properly set up and Docker installed.

Clone the project

```bash
  https://github.com/dougdelabrida/go-todo.git
```

Go to the project directory

```bash
  cd go-todo
```

Using docker

```bash
  docker-compose up
```

Or, you can the frontend and backend individually.

First, you'll need to install the dependencies

```bash
  cd frontend && npm install
```

Then, start the dev server

```bash
  npm run dev
```

Similarly for the backend. MONGODB_URI: "your uri"

```bash
  go build && ./todo
```

## Running Tests

To run tests for the frontend, run the following command in the `frontend` directory

```bash
  npm run test
```

To run tests for the backend, run the following command in the `backend` directory. **Make sure you have a standalone instance of mongodb running at localhost:27017**

```bash
  go test
```

## Roadmap

- Add production deployment settings

- Sort todos by priority
