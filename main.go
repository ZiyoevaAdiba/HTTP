package main

import (
	"net/http"

	"HTTP/handlers"
	"HTTP/middleware"
	"HTTP/storage"
)

func main() {

	userStorage := storage.New("data/users.json")

	userHandler := handlers.New(userStorage)

	mux := http.NewServeMux()

	handler := middleware.Logging(
		middleware.Auth(mux),
	)

	mux.HandleFunc("GET /users", userHandler.GetUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUserByID)
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		return
	}
}
