package realm

import "github.com/MiguelMoll/joycast/types"

type UserOption func(u *UserService) error

type UserService struct {
	repo UserRepo
}

func NewUserService(opts ...UserOption) {}

func UserStore(repo UserRepo) UserOption {
	return func(u *UserService) error {
		u.repo = repo
		return nil
	}
}

func (u *UserService) Create(user *types.User) (uint, error) {
	return u.repo.UserCreate(user)
}

func (u *UserService) Delete(id uint) error {
	return u.repo.UserDelete(id)
}

func (u *UserService) Get(id uint) (*types.User, error) {
	return u.repo.UserGet(id)
}

func (u *UserService) Update(user *types.User) error {
	return u.repo.UserUpdate(user)
}
