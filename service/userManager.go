package service

import (
	"app/helper/status"
	"app/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserManager struct {
	UserStore *model.UserStore
}

func (m *UserManager) Signin(username, password string) error {
	user := (*m.UserStore).GetUser(username)

	if user == (model.User{}) {
		return errors.New(status.UserNotFound)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Bcrypted), []byte(user.Salt+password)); err != nil {
		return errors.New(status.UserPasswordIncorrect)
	}

	return nil
}
