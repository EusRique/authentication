package model

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/badoux/checkmail"
)

type UserRepositoryInterface interface {
	CreatedUser(user *User) error
	FindUserByEmail(email string) (*User, error)
}
type User struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(40)" json:"name"`
	Email    string `gorm:"type:varchar(40)" json:"email"`
	Password string `gorm:"type:varchar(200)" json:"password"`

	CreatedAt time.Time  `gorm:"type:timestamp;autoCreateTime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamp;autoUpdateTime" json:"deleted_at"`
}

func (user *User) IsValid(passwordConfirmation string) []string {
	var errStrings []string
	var err error

	if user.Name == "" {
		errorName := fmt.Errorf("Nome é obrigatório")
		errStrings = append(errStrings, errorName.Error())
	}

	if user.Email == "" {
		errorEmail := fmt.Errorf("Email é obrigatório")
		errStrings = append(errStrings, errorEmail.Error())
	}

	if user.Email != "" {
		if err = checkmail.ValidateFormat(user.Email); err != nil {
			errorValidateEmail := fmt.Errorf("Email com formato invalido")
			errStrings = append(errStrings, errorValidateEmail.Error())
		}
	}
	if user.Password == "" {
		errorPassword := fmt.Errorf("Senha é obrigatório")
		errStrings = append(errStrings, errorPassword.Error())
	}

	if passwordConfirmation == "" {
		errorPasswordConfirmation := fmt.Errorf("Confirmação de senha é obrigatório")
		errStrings = append(errStrings, errorPasswordConfirmation.Error())
	}

	if user.Password != passwordConfirmation {
		errorComparePassword := fmt.Errorf("As senhas não conferem")
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
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.IsValid(passwordConfirmation)
	if err != nil {
		return nil, err
	}

	user.Password = passwordEncoder

	return &user, nil
}
