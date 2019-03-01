package realm

import "github.com/MiguelMoll/joycast/types"

type UserOption func(u *userService) error

type userService struct {
	repo UserRepo
}

func NewUserService(opts ...UserOption) {}

func UserStore(repo UserRepo) UserOption {
	return func(u *userService) error {
		u.repo = repo
		return nil
	}
}

func (u *userService) Create(user *types.User) (uint, error) {
	return u.repo.Create(user)
}

func (u *userService) Delete(id uint) error {
	return u.repo.Delete(id)
}

func (u *userService) Find(id uint) (*types.User, error) {
	return u.repo.Find(id)
}

func (u *userService) Update(user *types.User) error {
	return u.repo.Update(user)
}
