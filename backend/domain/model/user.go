package model

import "time"

type User struct {
	ID       uint   `gorm:"type:uint;primary_key" json:"id"`
	Name     string `gorm:"type:varchar(40)" json:"name"`
	Email    string `gorm:"type:varchar(40)" json:"email"`
	Password string `gorm:"type:varchar(40)" json:"password"`

	CreatedAt time.Time  `gorm:"type:timestamp;autoCreateTime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamp;autoUpdateTime" json:"deleted_at"`
}
