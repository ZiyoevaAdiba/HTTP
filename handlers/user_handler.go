package handlers

import (
	"net/http"

	"HTTP/storage"
)

type UserHandler struct {
	Storage *storage.UserStorage
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO
}
