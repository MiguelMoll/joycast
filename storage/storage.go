package storage

import (
	"golang.org/x/oauth2"
)

type Store interface {
	CreateUser(user *User) error
	GetUser(id uint) (*User, error)
	SaveUser(user *User) error
}

type User struct {
	ID         uint
	Name       string
	Email      string
	OauthToken *oauth2.Token
}
