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

	handler := middleware.Logging(
		middleware.Auth(mux),
	)

	mux.HandleFunc("GET /users", h.GetUsers)
	mux.HandleFunc("GET /users/{id}", h.GetUserByID)
	mux.HandleFunc("POST /users", h.CreateUser)
	mux.HandleFunc("PUT /users/{id}", h.UpdateUser)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		return
	}
}
