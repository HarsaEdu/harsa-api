package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestSetStatus_Success(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()

	// Create a SubsPlanServiceImpl with the mock repository
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, nil)

	id := uint(1)
	request := &web.SubsPlanUpdateStatus{
		IsActive: true,
	}

	// Mock the FindById method with no error
	mockRepository.On("FindById", int(id)).Return(&domain.SubsPlan{ID: id}, nil)

	// Mock the UpdateStatus method with no error
	mockRepository.On("UpdateStatus", request.IsActive, id).Return(nil)

	// Call the function you want to test
	err := subsPlanService.SetStatus(request, id)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

func TestSetStatus_SubPlanNotFound(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()

	// Create a SubsPlanServiceImpl with the mock repository
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, nil)

	id := uint(1)
	request := &web.SubsPlanUpdateStatus{
		IsActive: true,
	}

	// Mock the FindById method with no error, but subs plan not found
	mockRepository.On("FindById", int(id)).Return(nil, fmt.Errorf("find by id error"))

	// Call the function you want to test
	err := subsPlanService.SetStatus(request, id)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "subs plan not found")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
}

func TestSetStatus_UpdateStatusError(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()

	// Create a SubsPlanServiceImpl with the mock repository
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, nil)

	id := uint(1)
	request := &web.SubsPlanUpdateStatus{
		IsActive: true,
	}

	// Mock the FindById method with no error
	mockRepository.On("FindById", int(id)).Return(&domain.SubsPlan{ID: id}, nil)

	// Mock the UpdateStatus method with an error
	mockRepository.On("UpdateStatus", request.IsActive, id).Return(fmt.Errorf("update status error"))

	// Call the function you want to test
	err := subsPlanService.SetStatus(request, id)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "update status error")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
}
