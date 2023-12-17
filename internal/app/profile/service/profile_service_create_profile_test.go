package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/cloudinary"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateProfile(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	req := httptest.NewRequest(http.MethodPost, "/path", strings.NewReader(""))

	e := echo.New()
	ctx := e.NewContext(req, httptest.NewRecorder())

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))
	dateBirth := time.Unix(1609459200, 0)

	request := web.CreateProfileRequest{
		FirstName:   "harsa",
		LastName:    "edu",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "08080808080",
		ImageUrl:    "test.url",
	}

	mockRepo.On("IsExists", uint(1)).Return(false)
	mockRepo.On("CreateProfile", mock.AnythingOfType("*domain.UserProfile")).Return(nil)
	err := profileService.CreateProfile(ctx, &request, uint(1))

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreateProfileValidationError(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	req := httptest.NewRequest(http.MethodPost, "/path", strings.NewReader(""))

	e := echo.New()
	ctx := e.NewContext(req, httptest.NewRecorder())

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))
	dateBirth := time.Unix(1609459200, 0)

	request := web.CreateProfileRequest{
		FirstName: "harsa",
		LastName:  "edu",
		DateBirth: dateBirth,
		Gender:    "f",
		ImageUrl:  "test.url",
	}

	err := profileService.CreateProfile(ctx, &request, uint(1))

	require.Error(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreateProfileExists(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	req := httptest.NewRequest(http.MethodPost, "/path", strings.NewReader(""))

	e := echo.New()
	ctx := e.NewContext(req, httptest.NewRecorder())

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))
	dateBirth := time.Unix(1609459200, 0)

	request := web.CreateProfileRequest{
		FirstName:   "harsa",
		LastName:    "edu",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "08080808080",
		ImageUrl:    "test.url",
	}

	mockRepo.On("IsExists", uint(1)).Return(true)
	err := profileService.CreateProfile(ctx, &request, uint(1))

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreateProfileResultError(t *testing.T) {
	config, _ := configs.LoadConfig()
	mockRepo := new(mocks.ProfileRepository)
	validate := validator.New()

	req := httptest.NewRequest(http.MethodPost, "/path", strings.NewReader(""))

	e := echo.New()
	ctx := e.NewContext(req, httptest.NewRecorder())

	profileService := NewProfileService(mockRepo, validate, cloudinary.NewClodinaryUploader(&config.Cloudinary))
	dateBirth := time.Unix(1609459200, 0)

	request := web.CreateProfileRequest{
		FirstName:   "harsa",
		LastName:    "edu",
		DateBirth:   dateBirth,
		Gender:      "f",
		PhoneNumber: "08080808080",
		ImageUrl:    "test.url",
	}

	mockRepo.On("IsExists", uint(0)).Return(false)
	mockRepo.On("CreateProfile", mock.AnythingOfType("*domain.UserProfile")).Return(fmt.Errorf("error"))
	err := profileService.CreateProfile(ctx, &request, uint(0))

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}
