package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCategory_CategoryFound(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, nil, nil)

	// Define test data
	categoryID := 1

	// Set up mock expectations
	mockCategoryRepo.On("FindById", categoryID).Return(&domain.Category{ID: uint(categoryID), Name: "TestCategory"}, nil)
	mockCategoryRepo.On("Delete", categoryID).Return(nil)

	// Call the function you want to test
	err := categoryService.Delete(categoryID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

func TestDeleteCategory_CategoryNotFound(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, nil, nil)

	// Define test data
	categoryID := 1

	// Set up mock expectations for a category not found scenario
	mockCategoryRepo.On("FindById", categoryID).Return(nil, nil)

	// Call the function you want to test
	err := categoryService.Delete(categoryID)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "category not found")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

func TestDeleteCategory_ErrorDeleting(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, nil, nil)

	// Define test data
	categoryID := 1

	// Set up mock expectations for an error during deletion
	mockCategoryRepo.On("FindById", categoryID).Return(&domain.Category{ID: uint(categoryID), Name: "TestCategory"}, nil)
	mockCategoryRepo.On("Delete", categoryID).Return(fmt.Errorf("simulated deletion error"))

	// Call the function you want to test
	err := categoryService.Delete(categoryID)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when deleting : simulated deletion error")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

