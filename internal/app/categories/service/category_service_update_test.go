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

func TestUpdateCategory_Success(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)
	validate := validator.New()

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, validate, nil)

	// Define test data
	request := web.CategoryUpdateRequest{
		Name: "UpdatedCategory",
		// Add other fields as needed
	}
	categoryID := 1

	// Set up mock expectations
	mockCategoryRepo.On("FindById", categoryID).Return(&domain.Category{ID: 1, Name: "ExistingCategory"}, nil)
	mockCategoryRepo.On("FindByName", request.Name).Return(nil, nil) // Assume the name is not already used
	mockCategoryRepo.On("Update", mock.AnythingOfType("*domain.Category"), categoryID).Return(nil)

	// Call the function you want to test
	err := categoryService.Update(request, categoryID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

func TestUpdateCategory_CategoryNotFound(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)
	validate := validator.New()

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, validate, nil)

	// Define test data
	request := web.CategoryUpdateRequest{
		Name: "UpdatedCategory",
		// Add other fields as needed
	}
	categoryID := 1

	// Set up mock expectations for category not found
	mockCategoryRepo.On("FindById", categoryID).Return(nil, fmt.Errorf("simulated error"))

	// Call the function you want to test
	err := categoryService.Update(request, categoryID)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "category not found")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

func TestUpdateCategory_DuplicateCategoryName(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)
	validate := validator.New()

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, validate, nil)

	// Define test data
	request := web.CategoryUpdateRequest{
		Name: "UpdatedCategory",
		// Add other fields as needed
	}
	categoryID := 1

	// Set up mock expectations for duplicate category name
	mockCategoryRepo.On("FindById", categoryID).Return(&domain.Category{ID: 1, Name: "ExistingCategory"}, nil)
	mockCategoryRepo.On("FindByName", request.Name).Return(&domain.Category{ID: 2, Name: "UpdatedCategory"}, nil)

	// Call the function you want to test
	err := categoryService.Update(request, categoryID)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "category name already exist")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

func TestUpdateCategory_ErrorUpdatingCategory(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)
	validate := validator.New()

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, validate, nil)

	// Define test data
	request := web.CategoryUpdateRequest{
		Name: "UpdatedCategory",
		// Add other fields as needed
	}
	categoryID := 1

	// Set up mock expectations for an error during category update
	mockCategoryRepo.On("FindById", categoryID).Return(&domain.Category{ID: 1, Name: "ExistingCategory"}, nil)
	mockCategoryRepo.On("FindByName", request.Name).Return(nil, nil)
	mockCategoryRepo.On("Update", mock.AnythingOfType("*domain.Category"), categoryID).Return(fmt.Errorf("simulated error"))

	// Call the function you want to test
	err := categoryService.Update(request, categoryID)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when updating : simulated error")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

// Add more test cases as needed for different scenarios
