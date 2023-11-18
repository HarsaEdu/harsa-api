package web

import "time"

type ProfileRequest struct {
	ImageUrl    string    `gorm:"type:varchar(255)" json:"image_url" form:"image"`
	FirstName   string    `gorm:"type:varchar(50)" json:"first_name" form:"first_name"`
	LastName    string    `gorm:"type:varchar(50)" json:"last_name" form:"last_name"`
	DateBirth   time.Time `gorm:"type:date"json:"date_birth" form:"date_birth"`
	Bio         string    `gorm:"type:varchar(255)" json:"bio" form:"bio"`
	Gender      string    `gorm:"type:enum('f','m')" json:"gender" form:"gender"`
	PhoneNumber string    `gorm:"type:varchar(20)" json:"phone_number" form:"phone_number"`
	City        string    `gorm:"type:varchar(20)" json:"city" form:"city"`
	Address     string    `gorm:"type:varchar(255)" json:"address" form:"address"`
	Job         string    `gorm:"type:varchar(20)" json:"job" form:"job"`
}
