package main

import (
	"net/http"

	"HTTP/handlers"
	"HTTP/middleware"
	"HTTP/storage"
)

func main() {

	st := &storage.UserStorage{
		FileName: "data/users.json",
	}

	h := &handlers.UserHandler{
		Storage: st,
	}

	mux := http.NewServeMux()

	mux.Handle(
		"/users",
		middleware.Logging(
			middleware.Auth(
				http.HandlerFunc(h.GetUsers),
			),
		),
	)

	// TODO:
	// POST /users
	// GET /users/{id}
	// PUT /users/{id}

	http.ListenAndServe(":8080", mux)
}
