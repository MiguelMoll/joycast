package types

import "golang.org/x/oauth2"

type User struct {
	ID         uint
	Name       string
	Email      string
	OauthToken *oauth2.Token
}
