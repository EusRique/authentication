package model

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type UserRepositoryInterface interface {
	CreatedUser(user *User) error
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

func SHA256Encoder(password string) string {
	passwordEncoder := sha256.Sum256([]byte(password))

	return fmt.Sprintf("%x", passwordEncoder)
}

func NewUser(name, email, password string) (*User, error) {
	passwordEncoder := SHA256Encoder(password)

	user := User{
		Name:     name,
		Email:    email,
		Password: passwordEncoder,
	}

	return &user, nil
}
