package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCategories_Success(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "test"

	// Set up mock expectations
	mockCategoryRepo.On("GetAll", offset, limit, search).Return([]domain.Category{{ID: 1, Name: "TestCategory"}}, int64(1), nil)

	// Call the function you want to test
	result, pagination, err := categoryService.GetAll(offset, limit, search)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)
	assert.Equal(t, 1, len(result))

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

func TestGetAllCategories_NoCategoriesFound(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "test"

	// Set up mock expectations for no categories found
	mockCategoryRepo.On("GetAll", offset, limit, search).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := categoryService.GetAll(offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.EqualError(t, err, "category not found")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

func TestGetAllCategories_ErrorGettingCategories(t *testing.T) {
	// Create a mock CategoryRepository
	mockCategoryRepo := new(mocks.CategoryRepository)

	// Create a CategoryServiceImpl with the mock repository
	categoryService := NewCategoryService(mockCategoryRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "test"

	// Set up mock expectations for an error during category retrieval
	mockCategoryRepo.On("GetAll", offset, limit, search).Return(nil, int64(1), fmt.Errorf("simulated retrieval error"))

	// Call the function you want to test
	result, pagination, err := categoryService.GetAll(offset, limit, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.EqualError(t, err, "internal Server Error")

	// Assert that the expected calls were made
	mockCategoryRepo.AssertExpectations(t)
}

// Add more test cases as needed for different scenarios


