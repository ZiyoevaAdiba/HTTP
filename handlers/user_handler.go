package handlers

import (
	"HTTP/models"
	"HTTP/storage"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Storage *storage.UserStorage
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, _ *http.Request) {
	// getting all users from the storage
	allUsers, err := h.Storage.GetAll()
	if err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	// encoding all users to JSON and writing the response
	if err := json.NewEncoder(w).Encode(allUsers); err != nil {
		fmt.Println("error writing response", err)
		return
	}
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	user, err := h.Storage.GetByID(id)
	if err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}

		if errors.Is(err, models.ErrUserNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
			return
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		fmt.Println("error writing response", err)
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// decode the user from the body
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := h.Storage.Create(user); err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}

		if errors.Is(err, models.ErrUserAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
			return
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := h.Storage.Update(id, user); err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}
		if errors.Is(err, models.ErrUserNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
			return
		}
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}
}
