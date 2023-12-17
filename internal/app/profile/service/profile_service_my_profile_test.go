package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestGetMyProfile(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))
	request := web.UserGetByIDRequest{
		ID: 1,
	}

	profile := domain.ProfileDetailMobile{
		UserID:        1,
		UserProfileID: 1,
	}

	mockRepo.On("FindByUserID", request.ID).Return(&profile, nil)
	res, err := profileService.MyProfile(&request)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	mockRepo.AssertExpectations(t)
}

func TestGetMyProfileValidationError(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))
	request := web.UserGetByIDRequest{}
	res, err := profileService.MyProfile(&request)

	assert.Error(t, err)
	assert.Nil(t, res)
	mockRepo.AssertExpectations(t)
}

func TestGetMyProfileNotFound(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))
	request := web.UserGetByIDRequest{
		ID: 1,
	}

	mockRepo.On("FindByUserID", request.ID).Return(nil, fmt.Errorf("not found"))
	res, err := profileService.MyProfile(&request)

	assert.Error(t, err)
	assert.Nil(t, res)
	mockRepo.AssertExpectations(t)
}
