package model

type User struct {
	Username string
	Bcrypted string
}

type UserStore interface {
	GetUser(username string) User
}
