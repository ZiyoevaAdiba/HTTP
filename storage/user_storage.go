package storage

import (
	"sync"

	"HTTP/models"
)

type UserStorage struct {
	Mu       sync.Mutex
	FileName string
}

func (s *UserStorage) GetAll() ([]models.User, error) {
	// TODO
	return nil, nil
}

func (s *UserStorage) GetByID(id int) (*models.User, error) {
	// TODO
	return nil, nil
}

func (s *UserStorage) Create(user models.User) error {
	// TODO
	return nil
}

func (s *UserStorage) Update(id int, user models.User) error {
	// TODO
	return nil
}
