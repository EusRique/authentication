package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/EusRique/authentication/domain/model"
	"github.com/golang-jwt/jwt/v4"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJwtService() *jwtService {
	return &jwtService{
		secretKey: os.Getenv("JWT_SECRET_KEY"),
		issure:    "auth-api",
	}
}

type Claim struct {
	Sum   uint   `json:"sum"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(user *model.User) (string, error) {
	claim := &Claim{
		uint(user.ID),
		user.Name,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenRefresh, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenRefresh, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}
