package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProfile(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	req := httptest.NewRequest(http.MethodPost, "/path", strings.NewReader(""))

	e := echo.New()
	ctx := e.NewContext(req, httptest.NewRecorder())
	dateBirth := time.Unix(1609459200, 0)

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))

	request := web.UpdateProfileRequest{
		ImageUrl:    "url",
		FirstName:   "harsa",
		LastName:    "edu",
		Bio:         "bio",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "080808080",
		City:        "jakarta",
		Address:     "jakarta",
		Job:         "Job",
	}

	profile := domain.UserProfile{
		UserID:      1,
		ImageUrl:    "url",
		FirstName:   "harsa",
		LastName:    "edu",
		Bio:         "bio",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "080808080",
		City:        "jakarta",
		Address:     "jakarta",
		Job:         "Job",
	}

	profileExists := domain.ProfileDetail{
		UserID:      1,
		ImageUrl:    "url",
		FirstName:   "harsa",
		LastName:    "edu",
		Bio:         "bio",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "080808080",
		City:        "jakarta",
		Address:     "jakarta",
		Job:         "Job",
	}

	mockRepo.On("GetProfileByID", uint(1)).Return(&profileExists, nil)
	mockRepo.On("UpdateProfile", &profile).Return(nil)
	err := profileService.UpdateProfile(ctx, &request, uint(1))

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProfileValidationError(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	req := httptest.NewRequest(http.MethodPost, "/path", strings.NewReader(""))

	e := echo.New()
	ctx := e.NewContext(req, httptest.NewRecorder())

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))

	request := web.UpdateProfileRequest{
		ImageUrl: "url",
	}
	err := profileService.UpdateProfile(ctx, &request, uint(1))

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProfileNotFound(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	req := httptest.NewRequest(http.MethodPost, "/path", strings.NewReader(""))

	e := echo.New()
	ctx := e.NewContext(req, httptest.NewRecorder())
	dateBirth := time.Unix(1609459200, 0)

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))

	request := web.UpdateProfileRequest{
		ImageUrl:    "url",
		FirstName:   "harsa",
		LastName:    "edu",
		Bio:         "bio",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "080808080",
		City:        "jakarta",
		Address:     "jakarta",
		Job:         "Job",
	}

	mockRepo.On("GetProfileByID", uint(1)).Return(nil, fmt.Errorf("not found"))
	err := profileService.UpdateProfile(ctx, &request, uint(1))

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProfileResultError(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	req := httptest.NewRequest(http.MethodPost, "/path", strings.NewReader(""))

	e := echo.New()
	ctx := e.NewContext(req, httptest.NewRecorder())
	dateBirth := time.Unix(1609459200, 0)

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))

	request := web.UpdateProfileRequest{
		ImageUrl:    "url",
		FirstName:   "harsa",
		LastName:    "edu",
		Bio:         "bio",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "080808080",
		City:        "jakarta",
		Address:     "jakarta",
		Job:         "Job",
	}

	profile := domain.UserProfile{
		UserID:      1,
		ImageUrl:    "url",
		FirstName:   "harsa",
		LastName:    "edu",
		Bio:         "bio",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "080808080",
		City:        "jakarta",
		Address:     "jakarta",
		Job:         "Job",
	}

	profileExists := domain.ProfileDetail{
		UserID:      1,
		ImageUrl:    "url",
		FirstName:   "harsa",
		LastName:    "edu",
		Bio:         "bio",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "080808080",
		City:        "jakarta",
		Address:     "jakarta",
		Job:         "Job",
	}

	mockRepo.On("GetProfileByID", uint(1)).Return(&profileExists, nil)
	mockRepo.On("UpdateProfile", &profile).Return(fmt.Errorf("error"))
	err := profileService.UpdateProfile(ctx, &request, uint(1))

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProfileNilImageURL(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	req := httptest.NewRequest(http.MethodPost, "/path", strings.NewReader(""))

	e := echo.New()
	ctx := e.NewContext(req, httptest.NewRecorder())
	dateBirth := time.Unix(1609459200, 0)

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))

	request := web.UpdateProfileRequest{
		ImageUrl:    "url",
		FirstName:   "harsa",
		LastName:    "edu",
		Bio:         "bio",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "080808080",
		City:        "jakarta",
		Address:     "jakarta",
		Job:         "Job",
	}

	profile := domain.UserProfile{
		UserID:      1,
		ImageUrl:    "",
		FirstName:   "harsa",
		LastName:    "edu",
		Bio:         "bio",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "080808080",
		City:        "jakarta",
		Address:     "jakarta",
		Job:         "Job",
	}

	profileExists := domain.ProfileDetail{
		UserID:      1,
		ImageUrl:    "",
		FirstName:   "harsa",
		LastName:    "edu",
		Bio:         "bio",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "080808080",
		City:        "jakarta",
		Address:     "jakarta",
		Job:         "Job",
	}

	mockRepo.On("GetProfileByID", uint(1)).Return(&profileExists, nil)
	mockRepo.On("UpdateProfile", &profile).Return(fmt.Errorf("error"))
	err := profileService.UpdateProfile(ctx, &request, uint(1))

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
