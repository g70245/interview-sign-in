package model

import (
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var Users []User

const (
	Username = "test"
	Salt     = "o)t6A@2S'(6!("
	Password = "abcd1234"
)

type UserStoreMock struct{}

func init() {
	bcrypted, err := bcrypt.GenerateFromPassword([]byte(Salt+Password), viper.GetInt("bcrypt.cost"))
	if err == nil {
		Users = append(Users, User{Username, Salt, string(bcrypted)})
	}
}

func (m UserStoreMock) GetUser(username string) User {
	var user User

	for _, u := range Users {
		if u.Username == username {
			user = u
		}
	}

	return user
}
