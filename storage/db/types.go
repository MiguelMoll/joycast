package db

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/MiguelMoll/joycast/storage"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Name       string
	Email      string
	OauthToken postgres.Jsonb
}

func dbUser(user *storage.User) (*User, error) {
	if user == nil {
		return nil, nil
	}

	du := User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	bytes, err := json.Marshal(user.OauthToken)
	if err != nil {
		return nil, err
	}

	du.OauthToken = postgres.Jsonb{bytes}

	return &du, nil
}

func storeUser(user *User) (*storage.User, error) {
	if user == nil {
		return nil, nil
	}

	su := storage.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	if len(user.OauthToken.RawMessage) > 0 {
		var token storage.OauthToken
		if err := json.Unmarshal(user.OauthToken.RawMessage, &token); err != nil {
			return nil, err
		}

		su.OauthToken = &token
	}

	return &su, nil
}
