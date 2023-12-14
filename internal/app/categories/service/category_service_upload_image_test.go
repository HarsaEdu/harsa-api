package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUploadImage_Success(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)
	validate := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, validate, mockCloudinaryUploader)

	// Define test data
	request := &web.CategoryUploadImageRequest{}
	categoryID := 1

	// Set up mock expectations for successful image and icon upload
	mockCategoryRepo.On("FindById", categoryID).Return(&domain.Category{ID: 1}, nil)
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "categories", true).Return("image.png", nil)
	mockCloudinaryUploader.On("Uploader", mock.Anything, "icon", "categories", true).Return("icon.png", nil)
	mockCategoryRepo.On("UpdateImage", "image.png", "icon.png", categoryID).Return(nil)

	// Call the function you want to test
	err := categoryService.UploadImage(nil, request, categoryID, true, true)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestUploadImage_CategoryNotFound(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)
	validate := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, validate, mockCloudinaryUploader)

	// Define test data
	request := &web.CategoryUploadImageRequest{}
	categoryID := 1

	// Set up mock expectations for category not found
	mockCategoryRepo.On("FindById", categoryID).Return(nil, fmt.Errorf("simulated error"))

	// Call the function you want to test
	err := categoryService.UploadImage(nil, request, categoryID, true, true)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "category not found")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

// Add more test cases for different scenarios
