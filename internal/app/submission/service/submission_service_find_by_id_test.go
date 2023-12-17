package service

import (
	"errors"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestFindById_Success(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repository
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	id := 1

	// Mock the FindById method with no error
	mockSubmissionRepository.On("FindById", id).Return(&domain.Submissions{ID: uint(id)}, nil)

	// Call the function you want to test
	result, err := submissionService.FindById(id)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockSubmissionRepository.AssertExpectations(t)
}

func TestFindById_SubmissionNotFound(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repository
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	id := 1

	// Mock the FindById method with no error, but submission not found
	mockSubmissionRepository.On("FindById", id).Return(nil, nil)

	// Call the function you want to test
	result, err := submissionService.FindById(id)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "submission not found")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertExpectations(t)
}

func TestFindById_SubmissionRepositoryError(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repository
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	id := 1

	// Mock the FindById method with an error
	mockSubmissionRepository.On("FindById", id).Return(nil, errors.New("find by id error"))

	// Call the function you want to test
	result, err := submissionService.FindById(id)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "submission not found")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertExpectations(t)
}
