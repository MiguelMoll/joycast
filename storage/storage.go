package storage

import (
	"golang.org/x/oauth2"
)

type Store interface {
	SaveToken(user *User, token *oauth2.Token) error
	GetToken(user *User) (*oauth2.Token, error)
}

type User struct {
	ID    string
	Name  string
	Email string
}
