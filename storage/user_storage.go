package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"HTTP/models"
)

type UserStorage struct {
	Mu       sync.Mutex
	FileName string
}

func New(filename string) *UserStorage {
	return &UserStorage{
		FileName: filename,
	}
}

func (s *UserStorage) GetAll() ([]models.User, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	return readUsersFromFile(s.FileName)
}

func (s *UserStorage) GetByID(id int) (*models.User, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	users, err := readUsersFromFile(s.FileName)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, models.ErrUserNotFound
}

func (s *UserStorage) Create(user models.User) error {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	users, err := readUsersFromFile(s.FileName)
	if err != nil {
		return err
	}

	for _, existingUser := range users {
		if existingUser.ID == user.ID {
			return models.ErrUserAlreadyExists
		}
	}
	users = append(users, user)

	return writeToFile(s.FileName, users)
}

func (s *UserStorage) Update(id int, user models.User) error {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	users, err := readUsersFromFile(s.FileName)
	if err != nil {
		return err
	}

	for i, userRange := range users {
		if userRange.ID == id {
			users[i] = user
			return writeToFile(s.FileName, users)
		}
	}

	return errors.New("user not found")
}

func readUsersFromFile(name string) ([]models.User, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func writeToFile(name string, users []models.User) error {
	data, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(name, data, 0644)
}
