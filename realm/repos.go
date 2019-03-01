package realm

import "github.com/MiguelMoll/joycast/types"

type UserRepository interface {
	UserCreate(user *types.User) (uint, error)
	UserDelete(id uint) error
	UserGet(id uint) (*types.User, error)
	UserUpdate(user *types.User) error
}
