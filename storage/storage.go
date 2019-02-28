package storage

import (
	"time"
)

type Store interface {
	CreateUser(user *User) error
	GetUser(id string) (*User, error)
}

type User struct {
	ID         uint
	Name       string
	Email      string
	OauthToken *OauthToken
}

type OauthToken struct {
	// AccessToken is the token that authorizes and authenticates
	// the requests.
	AccessToken string `json:"access_token"`

	// TokenType is the type of token.
	// The Type method returns either this or "Bearer", the default.
	TokenType string `json:"token_type,omitempty"`

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string `json:"refresh_token,omitempty"`

	// Expiry is the optional expiration time of the access token.
	//
	// If zero, TokenSource implementations will reuse the same
	// token forever and RefreshToken or equivalent
	// mechanisms for that TokenSource will not be used.
	Expires time.Time `json:"expires,omitempty"`
}
