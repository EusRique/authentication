package middleware

import (
	"net/http"

	"github.com/EusRique/authentication/infrastructure/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Token Invalid (1)"})
			c.Abort()
			return
		}

		token := header[len(Bearer_schema):]

		if !auth.NewJwtService().ValidateToken(token) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Token Invalid (2)"})
			c.Abort()
			return
		}

		c.Next()
	}
}
