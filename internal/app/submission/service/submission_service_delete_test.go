package service

import (
	"errors"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestDelete_Success(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	id := 1
	userId := uint(1)
	role := "some_role"

	// Mock the CekUserIDfromSubmission method with no error
	mockSubmissionRepository.On("CekUserIDfromSubmission", uint(id), userId, role).Return(nil)

	// Mock the FindById method with no error
	mockSubmissionRepository.On("FindById", id).Return(&domain.Submissions{ID: uint(id)}, nil)

	// Mock the Delete method with no error
	mockSubmissionRepository.On("Delete", id).Return(nil)

	// Call the function you want to test
	err := submissionService.Delete(id, userId, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockSubmissionRepository.AssertExpectations(t)
}

func TestDelete_CekUserIDfromSubmissionError(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	id := 1
	userId := uint(1)
	role := "some_role"

	// Mock the CekUserIDfromSubmission method with an error
	mockSubmissionRepository.On("CekUserIDfromSubmission", uint(id), userId, role).Return(errors.New("cek user id error"))

	// Call the function you want to test
	err := submissionService.Delete(id, userId, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when cek user id")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertNotCalled(t, "FindById")
	mockSubmissionRepository.AssertNotCalled(t, "Delete")
}

func TestDelete_SubmissionNotFound(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	id := 1
	userId := uint(1)
	role := "some_role"

	// Mock the CekUserIDfromSubmission method with no error
	mockSubmissionRepository.On("CekUserIDfromSubmission", uint(id), userId, role).Return(nil)

	// Mock the FindById method with no error, but submission not found
	mockSubmissionRepository.On("FindById", id).Return(nil, nil)

	// Call the function you want to test
	err := submissionService.Delete(id, userId, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "submission not found")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertNotCalled(t, "Delete")
}

func TestDelete_SubmissionRepositoryDeleteError(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	id := 1
	userId := uint(1)
	role := "some_role"

	// Mock the CekUserIDfromSubmission method with no error
	mockSubmissionRepository.On("CekUserIDfromSubmission", uint(id), userId, role).Return(nil)

	// Mock the FindById method with no error
	mockSubmissionRepository.On("FindById", id).Return(&domain.Submissions{ID: uint(id)}, nil)

	// Mock the Delete method with an error
	mockSubmissionRepository.On("Delete", id).Return(errors.New("delete submission error"))

	// Call the function you want to test
	err := submissionService.Delete(id, userId, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when deleting")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertExpectations(t)
}
