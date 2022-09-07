package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Database() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("[middleware] Database()")

	}
}
