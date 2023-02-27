package usecase

import (
	"errors"

	"github.com/EusRique/authentication/domain/model"
	"github.com/EusRique/authentication/infrastructure/auth"
)

type LoginUseCase struct {
	UserRepository model.UserRepositoryInterface
}

func (l *LoginUseCase) Login(email, password string) (string, error, []string) {
	errRequiredField := model.LoginIn(email, password)
	if errRequiredField != nil {
		return "", nil, errRequiredField
	}

	// TODO Talvez buscar o usuário com o email e senha fornecidos
	user, err := l.UserRepository.FindUserByEmail(email)
	if err != nil {
		return "", err, nil
	}

	if user.Password != model.SHA256Encoder(password) {
		return "", errors.New("Credenciais inválidas"), nil
	}

	token, err := auth.NewJwtService().GenerateToken(user)
	if err != nil {
		return "", err, nil
	}

	return token, nil, nil
}
