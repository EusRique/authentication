package factory

import (
	"github.com/EusRique/authentication/application/usecase"
	"github.com/EusRique/authentication/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func UserUseCaseFactory(database *gorm.DB) usecase.UserUseCase {
	userRepository := repository.UserRepositoryDb{Db: database}

	userUseCase := usecase.UserUseCase{
		UserRepository: &userRepository,
	}

	return userUseCase
}
