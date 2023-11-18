package web

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

type GetProfileResponse struct {
	Class       string        `gorm:"type:varchar(50)" json:"class"`
	ImageUrl    string        `gorm:"type:varchar(255)" json:"image_url"`
	FirstName   string        `gorm:"type:varchar(50)" json:"first_name"`
	LastName    string        `gorm:"type:varchar(50)" json:"last_name"`
	DateBirth   time.Time     `gorm:"type:int" json:"date_birth"`
	Bio         string        `gorm:"type:varchar(255)" json:"bio"`
	Gender      domain.Gender `gorm:"type:enum('f','m')" json:"gender"`
	PhoneNumber string        `gorm:"type:varchar(20)" json:"phone_number"`
	City        string        `gorm:"type:varchar(20)" json:"city" form:"city"`
	Address     string        `gorm:"type:varchar(255)" json:"address"`
	Job         string        `gorm:"type:varchar(20)" json:"job"`
}
