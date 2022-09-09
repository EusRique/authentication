package usecase

import (
	"github.com/EusRique/authentication/domain/model"
)

type UserUseCase struct {
	UserRepository model.UserRepositoryInterface
}

func (u *UserUseCase) CreatedUser(name, email, password, passwordConfirmation string) (*model.User, []string, error) {
	newUser, errRequiredField := model.NewUser(name, email, password, passwordConfirmation)
	if errRequiredField != nil {
		return nil, errRequiredField, nil
	}

	err := u.UserRepository.CreatedUser(newUser)
	if err != nil {
		return nil, nil, err
	}

	return newUser, nil, nil
}
