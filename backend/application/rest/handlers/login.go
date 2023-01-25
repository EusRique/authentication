package handlers

import (
	"net/http"

	"github.com/EusRique/authentication/application/factory"
	"github.com/EusRique/authentication/application/model"
	"github.com/EusRique/authentication/application/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type LoginRestService struct {
	UserUseCase usecase.UserUseCase
}

func LoginIn(UserUseCase usecase.UserUseCase) *LoginRestService {
	return &LoginRestService{
		UserUseCase: UserUseCase,
	}
}

func (l *LoginRestService) Login(c *gin.Context) {
	var login model.Login

	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "JSON Invalido"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	user := factory.LoginUseCaseFactory(db)
	token, err, errRequiredField := user.Login(login.Email, login.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if errRequiredField != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errRequiredField})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}
