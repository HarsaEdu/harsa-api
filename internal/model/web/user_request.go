package web

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

type UserCreateRequest struct {
	FirstName   string        `json:"first_name" validate:"required"`
	LastName    string        `json:"last_name"`
	DateBirth   string        `json:"date_birth" validate:"required"`
	Bio         string        `json:"bio" validate:"required"`
	Gender      domain.Gender `json:"gender" validate:"required"`
	PhoneNumber string        `json:"phone_number" validate:"required"`
	City        string        `json:"city" validate:"required"`
	Address     string        `json:"address" validate:"required"`
	Email       string        `json:"email" validate:"required,email"`
	Username    string        `json:"username" validate:"required"`
	Password    string        `json:"password" validate:"required,min=8,max=255"`
	RoleID      uint          `json:"role_id" validate:"required"`
	Job         string        `json:"job" validate:"required"`
}

type UserProfileUpdateRequest struct {
	ID          uint          `json:"id" validate:"required"`
	FirstName   string        `json:"first_name" validate:"required"`
	LastName    string        `json:"last_name"`
	DateBirth   string        `json:"date_birth" validate:"required"`
	Bio         string        `json:"bio" validate:"required"`
	Gender      domain.Gender `json:"gender" validate:"required"`
	PhoneNumber string        `json:"phone_number" validate:"required"`
	City        string        `json:"city" validate:"required"`
	Address     string        `json:"address" validate:"required"`
	Job         string        `json:"job" validate:"required"`
}

type UserUpdateRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"omitempty,min=8,max=255"`
	RoleID   uint   `json:"role_id" validate:"required"`
}
