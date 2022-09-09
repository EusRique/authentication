package usecase

import (
	"errors"

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

	user, err := u.UserRepository.FindUserByEmail(email)
	if err != nil {
		return nil, nil, err
	}

	if user.ID == 0 {
		err = u.UserRepository.CreatedUser(newUser)
		if err != nil {
			return nil, nil, err
		}

		return newUser, nil, nil
	}

	return newUser, nil, errors.New("já existe um usuário com esse email cadastrado")
}
