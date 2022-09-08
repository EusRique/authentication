package repository

import (
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
