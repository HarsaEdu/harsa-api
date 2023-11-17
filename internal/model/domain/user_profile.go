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
	UserID      uint      `gorm:"type:int;not null" json:"user_id"`
	Class       string    `gorm:"type:varchar(255);" json:"class"`
	ImageUrl    string    `gorm:type:varchar(255)" json:"image_url"`
	FirstName   string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName    string    `gorm:"type:varchar(255)" json:"last_name"`
	DateBirth   time.Time `gorm:"type:date"json:"date_birth"`
	Bio         string    `gorm:"type:varchar(255);" json:"bio"`
	Gender      Gender    `gorm:"type:enum('m','f')" json:"gender"`
	PhoneNumber string    `gorm:"type:varchar(13);not null" json:"phone_number"`
	City        string    `gorm:"type:varchar(255);not null" json:"city"`
	Address     string    `gorm:"type:text" json:"address"`
	Job         string    `gorm:"type:varchar(255)" json:"job"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	User        User      `gorm:"foreignKey:UserID;references:ID"`
}
