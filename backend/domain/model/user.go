package model

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type UserRepositoryInterface interface {
	CreatedUser(user *User) error
	FindUserByEmail(email string) (*User, error)
}
type User struct {
	ID                   uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name                 string `gorm:"type:varchar(40)" json:"name"`
	Email                string `gorm:"type:varchar(40)" json:"email"`
	Password             string `gorm:"type:varchar(200)" json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`

	CreatedAt time.Time  `gorm:"type:timestamp;autoCreateTime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamp;autoUpdateTime" json:"deleted_at"`
}

func (user *User) IsValid() []string {
	var errStrings []string

	if user.Name == "" {
		errorName := fmt.Errorf("nome é obrigatório")
		errStrings = append(errStrings, errorName.Error())
	}

	if user.Email == "" {
		errorEmail := fmt.Errorf("email é obrigatório")
		errStrings = append(errStrings, errorEmail.Error())
	}

	if user.Password == "" {
		errorPassword := fmt.Errorf("senha é obrigatório")
		errStrings = append(errStrings, errorPassword.Error())
	}

	if user.PasswordConfirmation == "" {
		errorPasswordConfirmation := fmt.Errorf("confirmação de senha é obrigatório")
		errStrings = append(errStrings, errorPasswordConfirmation.Error())
	}

	if user.Password != user.PasswordConfirmation {
		errorComparePassword := fmt.Errorf("as senhas não conferem")
		errStrings = append(errStrings, errorComparePassword.Error())
	}

	if errStrings != nil {
		return errStrings
	}

	return nil
}

func SHA256Encoder(password string) string {
	passwordEncoder := sha256.Sum256([]byte(password))

	return fmt.Sprintf("%x", passwordEncoder)
}

func NewUser(name, email, password, passwordConfirmation string) (*User, []string) {
	passwordEncoder := SHA256Encoder(password)

	user := User{
		Name:                 name,
		Email:                email,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
	}

	err := user.IsValid()
	if err != nil {
		return nil, err
	}

	user.Password = passwordEncoder

	return &user, nil
}
