package service

import (
	"errors"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestGetAll_Success(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repository
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	moduleId := 1
	search := "some search"
	mockSubmissions := []domain.Submissions{{ID: 1, Title: "Submission 1"}, {ID: 2, Title: "Submission 2"}}
	expectedResult := []web.SubmissionsResponseWeb{{Id: 1, SubmissionTitle: "Submission 1"}, {Id: 2, SubmissionTitle: "Submission 2"}}
	mockTotal := int64(len(mockSubmissions))

	// Mock the GetAllWeb method with no error
	mockSubmissionRepository.On("GetAllWeb", moduleId, search).Return(expectedResult, mockTotal, nil)

	// Call the function you want to test
	result, err := submissionService.GetAll(moduleId, search)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)

	// Assert that the expected calls were made
	mockSubmissionRepository.AssertExpectations(t)
}

func TestGetAll_SubmissionNotFound(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repository
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	moduleId := 1
	search := "some search"

	// Mock the GetAllWeb method with no error, but submissions not found
	mockSubmissionRepository.On("GetAllWeb", moduleId, search).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, err := submissionService.GetAll(moduleId, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "submission not found")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertExpectations(t)
}

func TestGetAll_SubmissionRepositoryError(t *testing.T) {
	mockSubmissionRepository := new(mocks.SubmissionRepository)
	mockValidator := validator.New()

	// Create a SubmissionServiceImpl with the mock repository
	submissionService := NewSubmissionService(mockSubmissionRepository, *mockValidator)

	// Define test data
	moduleId := 1
	search := "some search"

	// Mock the GetAllWeb method with an error
	mockSubmissionRepository.On("GetAllWeb", moduleId, search).Return(nil, int64(0), errors.New("get all web error"))

	// Call the function you want to test
	result, err := submissionService.GetAll(moduleId, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "internal Server Error")

	// Assert that no further calls were made
	mockSubmissionRepository.AssertExpectations(t)
}
