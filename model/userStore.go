package model

type User struct {
	Username string
	Salt     string
	Bcrypted string
}

type UserStore interface {
	GetUser(username string) User
}
