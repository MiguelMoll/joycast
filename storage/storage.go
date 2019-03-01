package storage

import (
	"github.com/MiguelMoll/joycast/types"
)

type Store interface {
	CreateUser(user *types.User) error
	GetUser(id uint) (*types.User, error)
	SaveUser(user *types.User) error
}
