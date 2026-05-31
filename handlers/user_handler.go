package handlers

import (
	"go.uber.org/zap"
	"httpProject/models"
	"httpProject/pkg/Logger"
	"httpProject/storage"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type UserHandler struct {
	storage *storage.UserStorage
}

func New(storage *storage.UserStorage) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, _ *http.Request) {
	// getting all users from the storage
	allUsers, err := h.storage.GetAll()
	if err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	// encoding all users to JSON and writing the response
	if err := json.NewEncoder(w).Encode(allUsers); err != nil {
		Logger.L.Error("Could not write response",
			zap.String("error msg", err.Error()))

		fmt.Println("error writing response", err)
		return
	}
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		Logger.L.Error("Could not parse id from path")

		errDTO := models.ErrorDTO{Message: err.Error()}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	user, err := h.storage.GetByID(id)
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
		Logger.L.Error("Could not write response",
			zap.String("error msg", err.Error()))

		fmt.Println("error writing response", err)
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := models.DecodeUser(r)
	if err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := h.storage.Create(user); err != nil {
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
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	user, err := models.DecodeUser(r)
	if err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := h.storage.Update(id, user); err != nil {
		errDTO := models.ErrorDTO{Message: err.Error()}
		if errors.Is(err, models.ErrUserNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
			return
		}
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}
}
