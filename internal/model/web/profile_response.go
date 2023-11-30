package web

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

type GetProfileResponse struct {
	ID          uint          `json:"profile_id"`
	ImageUrl    string        `json:"image_url"`
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	DateBirth   time.Time     `json:"date_birth"`
	Bio         string        `json:"bio"`
	Gender      domain.Gender `json:"gender"`
	PhoneNumber string        `json:"phone_number"`
	City        string        `json:"city" form:"city"`
	Address     string        `json:"address"`
	Job         string        `json:"job"`
}
