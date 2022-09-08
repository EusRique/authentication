package routes

import (
	rest "github.com/EusRique/authentication/application/rest/handlers"
	"github.com/EusRique/authentication/application/usecase"
	"github.com/EusRique/authentication/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func configRoutes(api *gin.RouterGroup) {
	user := rest.NewUsers(usecase.UserUseCase{})

	api.GET("/", rest.Alive)

	api.POST("/new-user", user.CreateUser)
}

func Start(port string) {
	r := gin.New()
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	defaultMiddlewares := []gin.HandlerFunc{middleware.Database()}

	version := "v1"
	api := r.Group(version+"/api", defaultMiddlewares...)

	configRoutes(api)
	r.Run(":" + port)
}
