package service


import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


// TestGetAll_Success tests the case when fetching all subscription plans is successful.
func TestGetAll_Success(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, nil, nil)

	// Mock the SubsPlanRepository GetAllActive method with valid results
	expectedResult := domain.SubsPlan{
		ID: uint(1),
	} // Replace with your actual struct

	subplais:=[]domain.SubsPlan{
		expectedResult,
	}
	mockRepository.On("GetAllActive", mock.Anything, mock.Anything, mock.Anything).Return(subplais, int64(10), nil)

	// Call the function you want to test
	results, pagination, err := subsPlanService.GetAll(0, 10, "searchTerm")

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestGetAll_ErrorSubsPlanNotFound tests the case when no subscription plans are found.
func TestGetAll_ErrorSubsPlanNotFound(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, nil, nil)

	// Mock the SubsPlanRepository GetAllActive method with no plans found
	mockRepository.On("GetAllActive", mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), fmt.Errorf("error"))

	// Call the function you want to test
	results, pagination, err := subsPlanService.GetAll(0, 10, "searchTerm")

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "subs plan not found")
	assert.Nil(t, results)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

// TestGetAll_ErrorInternalServerError tests the case when there's an internal server error.
func TestGetAll_ErrorInternalServerError(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, nil, nil)

	// Mock the SubsPlanRepository GetAllActive method with an internal server error
	mockRepository.On("GetAllActive", mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(1), fmt.Errorf("error"))

	// Call the function you want to test
	results, pagination, err := subsPlanService.GetAll(0, 10, "searchTerm")

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "internal Server Error")
	assert.Nil(t, results)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

