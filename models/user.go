package models

import (
	"encoding/json"
	"errors"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u *User) Validate() error {
	if u.ID <= 0 {
		return errors.New("user ID should be positive")
	}
	if u.Name == "" {
		return errors.New("user name can not be empty")
	}
	return nil
}

func DecodeUser(r *http.Request) (User, error) {
	user := User{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&user); err != nil {
		return user, err
	}

	if err := user.Validate(); err != nil {
		return user, err
	}

	return user, nil
}
