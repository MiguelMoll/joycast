package realm

import "github.com/MiguelMoll/joycast/types"

type UserRepo interface {
	Create(user *types.User) (uint, error)
	Delete(id uint) error
	Find(id uint) (*types.User, error)
	Update(user *types.User) error
}
