package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindById_Success(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	// Mock the SubsPlanRepository FindById method with a valid result
	expectedResult := &domain.SubsPlan{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedResult, nil)

	// Call the function you want to test
	result, err := subsPlanService.FindById(1) // Replace 1 with an actual ID

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestFindById_ErrorSubsPlanNotFound tests the case when the subscription plan is not found.
func TestFindById_ErrorSubsPlanNotFound(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	// Mock the SubsPlanRepository FindById method with no plan found
	mockRepository.On("FindById", mock.Anything).Return(nil, fmt.Errorf("eror"))

	// Call the function you want to test
	result, err := subsPlanService.FindById(1) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "subs plan not found")
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestFindById_ErrorInternalServerError tests the case when there's an internal server error.
func TestFindById_ErrorInternalServerError(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	expectedResult := &domain.SubsPlan{
		ID: uint(1),
	}
	// Mock the SubsPlanRepository FindById method with an internal server error
	mockRepository.On("FindById", mock.Anything).Return(expectedResult, fmt.Errorf("eror"))

	// Call the function you want to test
	result, err := subsPlanService.FindById(1) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "internal Server Error")
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)

}

func TestFindById_ErrorNotFound(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	// Mock the SubsPlanRepository FindById method with an internal server error
	mockRepository.On("FindById", mock.Anything).Return(nil, nil)

	// Call the function you want to test
	result, err := subsPlanService.FindById(1) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "subs plan not found")
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)

}

// ... (You can continue adding more test cases as needed)
