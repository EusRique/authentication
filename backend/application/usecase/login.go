package usecase

import (
	"errors"

	"github.com/EusRique/authentication/domain/model"
	"github.com/EusRique/authentication/infrastructure/auth"
)

type LoginUseCase struct {
	UserRepository model.UserRepositoryInterface
}

func (l *LoginUseCase) Login(email, password string) (interface{}, error) {
	// TODO Criar model Login para validação de campos vazios no coração da aplicação

	user, err := l.UserRepository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != model.SHA256Encoder(password) {
		return nil, errors.New("Invalid credentials")
	}

	token, err := auth.NewJwtService().GenerateToken(uint(user.ID))
	if err != nil {
		return nil, err
	}

	return token, nil
}
