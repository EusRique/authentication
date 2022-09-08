package repository

import (
	"errors"

	"github.com/EusRique/authentication/domain/model"
	"github.com/jinzhu/gorm"
)

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (u *UserRepositoryDb) CreatedUser(user *model.User) error {
	err := u.Db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryDb) FindUserByEmail(email string) (*model.User, error) {
	var user model.User

	resultQuery := u.Db.First(&user, "email = ?", email).Error

	if errors.Is(resultQuery, gorm.ErrRecordNotFound) {
		return &user, nil
	}

	if resultQuery != nil {
		return &user, resultQuery
	}

	return &user, nil
}
