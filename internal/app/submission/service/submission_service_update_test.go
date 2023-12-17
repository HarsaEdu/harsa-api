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

func TestUpdate_Success(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	submissionId := 1
	userID := uint(1)
	role := "some_role"
	request := &web.SubmissionUpdateRequest{Title: "Updated Title"}
	mockSubmission := &domain.Submissions{ID: uint(submissionId), Title: "Original Title"}

	// Mock the CekUserIDfromSubmission method with no error
	mockSubmissionRepository.On("CekUserIDfromSubmission", uint(submissionId), userID, role).Return(nil)

	// Mock the FindById method with no error
	mockSubmissionRepository.On("FindById", submissionId).Return(mockSubmission, nil)

	// Mock the Update method with no error
	mockSubmissionRepository.On("Update", mock.Anything, submissionId).Return(nil)

	// Call the function you want to test
	err := submissionService.Update(nil, request, submissionId, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockSubmissionRepository.AssertExpectations(t)
}

func TestUpdate_CekUserIDfromSubmissionError(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	submissionId := 1
	userID := uint(1)
	role := "some_role"
	request := &web.SubmissionUpdateRequest{Title: "Updated Title"}

	// Mock the CekUserIDfromSubmission method with an error
	mockSubmissionRepository.On("CekUserIDfromSubmission", uint(submissionId), userID, role).Return(fmt.Errorf("cek user id error"))

	// Call the function you want to test
	err := submissionService.Update(nil, request, submissionId, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when cek user id")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertNotCalled(t, "FindById")
	mockSubmissionRepository.AssertNotCalled(t, "Update")
}

func TestUpdate_SubmissionNotFound(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	submissionId := 1
	userID := uint(1)
	role := "some_role"
	request := &web.SubmissionUpdateRequest{Title: "Updated Title"}

	// Mock the CekUserIDfromSubmission method with no error
	mockSubmissionRepository.On("CekUserIDfromSubmission", uint(submissionId), userID, role).Return(nil)

	// Mock the FindById method with no error, but submission not found
	mockSubmissionRepository.On("FindById", submissionId).Return(nil, nil)

	// Call the function you want to test
	err := submissionService.Update(nil, request, submissionId, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "submission not found")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertNotCalled(t, "Update")
}


func TestUpdate_UpdateError(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	submissionId := 1
	userID := uint(1)
	role := "some_role"
	request := &web.SubmissionUpdateRequest{Title: "Updated Title"}
	mockSubmission := &domain.Submissions{ID: uint(submissionId), Title: "Original Title"}

	// Mock the CekUserIDfromSubmission method with no error
	mockSubmissionRepository.On("CekUserIDfromSubmission", uint(submissionId), userID, role).Return(nil)

	// Mock the FindById method with no error
	mockSubmissionRepository.On("FindById", submissionId).Return(mockSubmission, nil)

	// Mock the Update method with an error
	mockSubmissionRepository.On("Update", mock.Anything, submissionId).Return(fmt.Errorf("update submission error"))

	// Call the function you want to test
	err := submissionService.Update(nil, request, submissionId, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when updating submission")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertExpectations(t)
}
