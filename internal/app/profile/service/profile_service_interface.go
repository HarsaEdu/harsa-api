package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/profile/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ProfileService interface {
	CreateProfile(ctx echo.Context, profile *web.CreateProfileRequest, userID uint) error
	GetProfileByID(id uint) (*web.GetProfileResponse, error)
	UpdateProfile(ctx echo.Context, profile *web.UpdateProfileRequest, id uint) error
}

type ProfileServiceImpl struct {
	ProfileRepository  repository.ProfileRepository
	Validator          *validator.Validate
	cloudinaryUploader cloudinary.CloudinaryUploader
}

func NewProfileService(pr repository.ProfileRepository, validator *validator.Validate, cloudinary cloudinary.CloudinaryUploader) ProfileService {
	return &ProfileServiceImpl{ProfileRepository: pr, Validator: validator, cloudinaryUploader: cloudinary}
}