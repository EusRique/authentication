package routes

import (
	"fmt"

	"github.com/EusRique/authentication/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func configRoutes(api *gin.RouterGroup) {
	fmt.Println("Hello World")
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
