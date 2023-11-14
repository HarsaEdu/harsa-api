package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                uint           `gorm:"type:int;primarykey" json:"id"`
	Email             string         `gorm:"type:varchar(255)" json:"email"`
	Username          string         `gorm:"type:varchar(255)" json:"username"`
	Password          string         `gorm:"type:varchar(255)" json:"password"`
	RoleID            uint           `gorm:"type:int;primarykey" json:"role_id"`
	RegistrationToken string         `gorm:"type:varchar(255)" json:"registration_token"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"delete_at"`
	Role              Role           `gorm:"foreignKey:RoleID;references:ID"`
}
