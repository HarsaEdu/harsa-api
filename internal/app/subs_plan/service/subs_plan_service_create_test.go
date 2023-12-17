package service

import (
	"fmt"
	"testing"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate_Success(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanCreateRequest{
		Title: "asda",
		Price: 1000,
		Duration_days: 30,
	}

	// Mock the Uploader method with no error
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs-plan", false).Return("image_url.jpg", nil)

	// Call the function you want to test
	err := subsPlanService.Create(nil, request)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestCreate_InvalidFileFormat(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanCreateRequest{
		// Set your request fields here
	}

	// Mock the Uploader method with no error, but invalid file format
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs-plan", false).Return("image_url.pdf", nil)

	// Call the function you want to test
	err := subsPlanService.Create(nil, request)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'SubsPlanCreateRequest.Title'")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestCreate_ErrorUploadingImage(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanCreateRequest{
		// Set your request fields here
	}


	// Mock the Uploader method with an error
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs-plan", false).Return("", fmt.Errorf("error uploading image"))

	// Call the function you want to test
	err := subsPlanService.Create(nil, request)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error uploading image")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestCreate_ErrorCreatingSubsPlan(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanCreateRequest{
		// Set your request fields here
	}


	// Mock the Uploader method with no error
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs-plan", false).Return("image_url.jpg", nil)

	// Mock the Create method with an error
	mockRepository.On("Create", mock.Anything).Return(fmt.Errorf("error creating subs plan"))

	// Call the function you want to test
	err := subsPlanService.Create(nil, request)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error creating subs plan")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}
