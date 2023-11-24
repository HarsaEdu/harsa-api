package web

import "time"

type UpdateProfileRequest struct {
	ImageUrl    string    `gorm:"type:varchar(255)" json:"image_url" form:"image"`
	FirstName   string    `gorm:"type:varchar(50)" json:"first_name" form:"first_name" validate:"required"`
	LastName    string    `gorm:"type:varchar(50)" json:"last_name" form:"last_name" validate:"required"`
	DateBirth   time.Time `gorm:"type:date"json:"date_birth" form:"date_birth" validate:"required"`
	Bio         string    `gorm:"type:varchar(255)" json:"bio" form:"bio" validate:"required"`
	Gender      string    `gorm:"type:enum('f','m')" json:"gender" form:"gender" validate:"required"`
	PhoneNumber string    `gorm:"type:varchar(20)" json:"phone_number" form:"phone_number" validate:"required"`
	City        string    `gorm:"type:varchar(20)" json:"city" form:"city" validate:"required"`
	Address     string    `gorm:"type:varchar(255)" json:"address" form:"address" validate:"required"`
	Job         string    `gorm:"type:varchar(20)" json:"job" form:"job" validate:"required"`
}

type CreateProfileRequest struct {
	FirstName   string    `gorm:"type:varchar(50)" json:"first_name" form:"first_name" validate:"required"`
	LastName    string    `gorm:"type:varchar(50)" json:"last_name" form:"last_name"`
	DateBirth   time.Time `gorm:"type:date"json:"date_birth" form:"date_birth" validate:"required"`
	PhoneNumber string    `gorm:"type:varchar(20)" json:"phone_number" form:"phone_number" validate:"required"`
}
