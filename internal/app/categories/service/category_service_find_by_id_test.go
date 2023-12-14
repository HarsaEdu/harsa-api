package service

import (
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFindById_CategoryFound(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, nil, nil)

	// Define test data
	categoryID := 1
	expectedCategory := &domain.Category{ID: uint(categoryID), Name: "TestCategory"}

	// Set up mock expectations
	mockCategoryRepo.On("FindById", categoryID).Return(expectedCategory, nil)

	// Call the function you want to test
	result, err := categoryService.FindById(categoryID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedCategory, result)

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

func TestFindById_CategoryNotFound(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, nil, nil)

	// Define test data
	categoryID := 1

	// Set up mock expectations for a category not found scenario
	mockCategoryRepo.On("FindById", categoryID).Return(nil, nil)

	// Call the function you want to test
	result, err := categoryService.FindById(categoryID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "category not found")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

