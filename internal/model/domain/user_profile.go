package domain

import (
	"time"
)

type Gender string

const (
	Male   Gender = "m"
	Female Gender = "f"
)

type UserProfile struct {
	ID          uint      `gorm:"type:int;primarykey" json:"id"`
	UserID      uint      `gorm:"type:int;not null" json:"user_id" form:"user_id"`
	ImageUrl    string    `gorm:"type:varchar(255);default:'https://placewaifu.com/image/200/200'" json:"image_url" form:"image"`
	FirstName   string    `gorm:"type:varchar(255);not null" json:"first_name" form:"first_name"`
	LastName    string    `gorm:"type:varchar(255)" json:"last_name" form:"last_name"`
	DateBirth   time.Time `gorm:"type:date" json:"date_birth" form:"date_birth"`
	Bio         string    `gorm:"type:varchar(255);" json:"bio" form:"bio"`
	Gender      Gender    `gorm:"type:enum('m','f')" json:"gender" form:"gender"`
	PhoneNumber string    `gorm:"type:varchar(13);not null" json:"phone_number" form:"phone_number"`
	City        string    `gorm:"type:varchar(255);not null" json:"city" form:"city"`
	Address     string    `gorm:"type:text" json:"address" form:"address"`
	Job         string    `gorm:"type:varchar(255)" json:"job" form:"job"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ProfileDetail struct {
	UserID        uint      `json:"user_id"`
	UserProfileID uint      `json:"user_profile_id"`
	ImageUrl      string    `json:"image_url" form:"image"`
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
