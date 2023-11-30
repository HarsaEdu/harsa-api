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
type UserEntity struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	RoleName    string `json:"role_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}
type UserDetail struct {
	UserID        uint      `json:"user_id"`
	UserProfileID uint      `json:"user_profile_id"`
	RoleID        uint      `json:"role_id"`
	RoleName      string    `json:"role_name"`
	Email         string    `json:"email"`
	Username      string    `json:"username"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	PhoneNumber   string    `json:"phone_number"`
	DateBirth     time.Time `json:"date_birth"`
	Bio           string    `json:"bio"`
	Gender        Gender    `json:"gender"`
	City          string    `json:"city"`
	Address       string    `json:"address"`
	Job           string    `json:"job"`
}
