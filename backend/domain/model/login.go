package model

import (
	"fmt"

	"github.com/badoux/checkmail"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (login *Login) IsValid() []string {
	var errStrings []string
	var err error

	if login.Email == "" {
		errorEmail := fmt.Errorf("o email é obrigatório")
		errStrings = append(errStrings, errorEmail.Error())
	}

	if login.Password == "" {
		errorPassword := fmt.Errorf("a senha obrigatório")
		errStrings = append(errStrings, errorPassword.Error())
	}

	if login.Email != "" {
		if err = checkmail.ValidateFormat(login.Email); err != nil {
			errorValidateEmail := fmt.Errorf("email com formato invalido")
			errStrings = append(errStrings, errorValidateEmail.Error())
		}
	}

	if errStrings != nil {
		return errStrings
	}

	return nil
}

func LoginIn(email, password string) []string {
	login := Login{
		Email:    email,
		Password: password,
	}

	err := login.IsValid()
	if err != nil {
		return err
	}

	return nil
}
