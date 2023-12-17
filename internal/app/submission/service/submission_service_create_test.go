package service

import (
	"errors"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate_Success(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	request := web.SubmissionRequest{
		ModuleID: 1,
		Title:    "aSDA",
		Content:  "Sada",
	}

	userId := uint(1)
	role := "some_role"

	// Mock the CekUserIDfromModuleID method with no error
	mockSubmissionRepository.On("CekUserIDfromModuleID", request.ModuleID, userId, role).Return(nil)

	// Mock the Create method with no error
	mockSubmissionRepository.On("Create", mock.Anything).Return(nil)

	// Call the function you want to test
	err := submissionService.Create(nil, request, userId, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockSubmissionRepository.AssertExpectations(t)
}

func TestCreate_ValidationError(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	request := web.SubmissionRequest{
		// Populate other necessary fields
	}

	userId := uint(1)
	role := "some_role"

	// Call the function you want to test
	err := submissionService.Create(nil, request, userId, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'SubmissionRequest.Title'")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertNotCalled(t, "CekUserIDfromModuleID")
	mockSubmissionRepository.AssertNotCalled(t, "Create")
}

func TestCreate_CekUserIDfromModuleIDError(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	request := web.SubmissionRequest{
		ModuleID: 1,
		Title: "aSDA",
		Content: "Sada",
	}


	userId := uint(1)
	role := "some_role"

	// Mock the CekUserIDfromModuleID method with an error
	mockSubmissionRepository.On("CekUserIDfromModuleID", request.ModuleID, userId, role).Return(errors.New("cek user id error"))

	// Call the function you want to test
	err := submissionService.Create(nil, request, userId, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when cek user id")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertNotCalled(t, "Create")
}

func TestCreate_SubmissionRepositoryCreateError(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repositories
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	request := web.SubmissionRequest{
		ModuleID: 1,
		Title:    "aSDA",
		Content:  "Sada",
	}

	userId := uint(1)
	role := "some_role"

	// Mock the CekUserIDfromModuleID method with no error
	mockSubmissionRepository.On("CekUserIDfromModuleID", request.ModuleID, userId, role).Return(nil)

	// Mock the Create method with an error
	mockSubmissionRepository.On("Create", mock.Anything).Return(errors.New("create submission error"))

	// Call the function you want to test
	err := submissionService.Create(nil, request, userId, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when creating submission")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertExpectations(t)
}
