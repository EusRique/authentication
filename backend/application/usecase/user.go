package usecase

import (
	"github.com/EusRique/authentication/domain/model"
)

type UserUseCase struct {
	UserRepository model.UserRepositoryInterface
}

func (u *UserUseCase) CreatedUser(name, email, password string) (*model.User, error) {
	newUser, err := model.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}

	err = u.UserRepository.CreatedUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
