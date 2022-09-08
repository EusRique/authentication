package handlers

import (
	"net/http"

	"github.com/EusRique/authentication/application/factory"
	"github.com/EusRique/authentication/application/model"
	"github.com/EusRique/authentication/application/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserRestService struct {
	UserUseCase usecase.UserUseCase
}

func NewUsers(UserUseCase usecase.UserUseCase) *UserRestService {
	return &UserRestService{
		UserUseCase: UserUseCase,
	}
}

func (u *UserRestService) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "JSON Invalido"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	userUseCase := factory.UserUseCaseFactory(db)
	newUser, err := userUseCase.CreatedUser(user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"results": newUser})
}
