package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/profile/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ProfileService interface {
	CreateProfile(ctx echo.Context, profile *web.ProfileRequest, userID uint) error
	GetProfileByID(id uint) (*domain.UserProfile, error)
	UpdateProfile(ctx echo.Context, profile *domain.UserProfile, id uint) error
}

type ProfileServiceImpl struct {
	ProfileRepository  repository.ProfileRepository
	Validator          *validator.Validate
	cloudinaryUploader cloudinary.CloudinaryUpdloader
}

func NewProfileService(pr repository.ProfileRepository, validator *validator.Validate, cloudinary cloudinary.CloudinaryUpdloader) ProfileService {
	return &ProfileServiceImpl{ProfileRepository: pr, Validator: validator, cloudinaryUploader: cloudinary}
}
