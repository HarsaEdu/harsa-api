package service

import (
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCategory(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)
	validate := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CategoryServiceImpl with the mock repositories
	categoryService := NewCategoryService(mockCategoryRepo, validate, mockCloudinaryUploader)

	// Define test data
	request := web.CategoryCreateRequest{
		Name: "TestCateg",
		// Add other fields as needed
	}

	// Set up mock expectations
	mockCategoryRepo.On("FindByName", request.Name).Return(nil, nil) // Assume name is not already used
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "categories", false).Return("image.png", nil)
	mockCloudinaryUploader.On("Uploader", mock.Anything, "icon", "categories", false).Return("icon.png", nil)
	mockCategoryRepo.On("Create", mock.AnythingOfType("*domain.Category")).Return(nil)

	// Call the function you want to test
	err := categoryService.Create(nil, request)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestCreateCategory_NameAlreadyExists(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)
	validate := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CategoryServiceImpl with the mock repositories
	categoryService := NewCategoryService(mockCategoryRepo, validate, mockCloudinaryUploader)

	// Define test data
	request := web.CategoryCreateRequest{
		Name: "TestCateg",
		// Add other fields as needed
	}

	// Set up mock expectations to simulate existing name
	mockCategoryRepo.On("FindByName", request.Name).Return(&domain.Category{}, nil)

	// Call the function you want to test
	err := categoryService.Create(nil, request)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "category name already exist")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestCreateCategory_InvalidImageFormat(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)
	validate := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CategoryServiceImpl with the mock repositories
	categoryService := NewCategoryService(mockCategoryRepo, validate, mockCloudinaryUploader)

	// Define test data
	request := web.CategoryCreateRequest{
		Name: "TestCateg",
		// Add other fields as needed
	}

	// Set up mock expectations for image upload with invalid format
	mockCategoryRepo.On("FindByName", request.Name).Return(nil, nil)
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "categories", false).Return("invalidImage.pdf", nil)

	// Call the function you want to test
	err := categoryService.Create(nil, request)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid file format")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestCreateCategory_InvalidIconFormat(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)
	validate := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CategoryServiceImpl with the mock repositories
	categoryService := NewCategoryService(mockCategoryRepo, validate, mockCloudinaryUploader)

	// Define test data
	request := web.CategoryCreateRequest{
		Name: "TestCateg",
		// Add other fields as needed
	}

	// Set up mock expectations for icon upload with invalid format
	mockCategoryRepo.On("FindByName", request.Name).Return(nil, nil)
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "categories", false).Return("validImage.png", nil)
	mockCloudinaryUploader.On("Uploader", mock.Anything, "icon", "categories", false).Return("invalidIcon.pdf", nil)

	// Call the function you want to test
	err := categoryService.Create(nil, request)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid file format")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}
