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

func TestFindById_Success(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()
	mockSubmission := new(mocks.SubmissionRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repositories

	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader, mockSubmission)

	mockSubmissionAnswer := &domain.SubmissionAnswer{
		ID:           1,
		UserID:       1,
		SubmissionID: 1,
		SubmittedUrl: "file_url.pdf",
		Status:       "submitted",
	}

	mockResponse := &web.SubmissionAnswerResponseMobile{
		ID:         1,
		Submission: "file_url.pdf",
		Status:     "submitted",
	}

	// Mock the FindById method with no error
	mockRepository.On("FindById", 1).Return(mockSubmissionAnswer, nil)

	// Call the function you want to test
	result, err := submissionAnswerService.FindById(1)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockResponse, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

func TestFindById_SubmissionAnswerNotFound(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()
	mockSubmission := new(mocks.SubmissionRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repositories

	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader, mockSubmission)

	// Mock the FindById method with no error, but submission answer not found
	mockRepository.On("FindById", 1).Return(nil, fmt.Errorf("find by id error"))

	// Call the function you want to test
	result, err := submissionAnswerService.FindById(1)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "submission answer not found")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
}

func TestFindById_Error(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repository
	submissionAnswerService := NewSubmissionAnswer(mockRepository, nil, nil, nil)

	// Mock the FindById method with an error
	mockRepository.On("FindById", 1).Return(&domain.SubmissionAnswer{}, fmt.Errorf("find by id error"))

	// Call the function you want to test
	result, err := submissionAnswerService.FindById(1)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "internal Server Error")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
}

func TestFindByIdWeb_Success(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()
	mockSubmission := new(mocks.SubmissionRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader, mockSubmission)

	mockResponse := &web.SubmissionAnswerResponseWebById{
		ID:         1,
		SubmittedUrl: "file_url.pdf",
		Status:     "submitted",
	}

	// Mock the FindByIdWeb method with no error
	mockRepository.On("FindByIdWeb", 1).Return(mockResponse, nil)

	// Call the function you want to test
	result, err := submissionAnswerService.FindByIdWeb(1)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockResponse, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

func TestFindByIdWeb_SubmissionAnswerNotFound(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()
	mockSubmission := new(mocks.SubmissionRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader, mockSubmission)

	// Mock the FindByIdWeb method with no error, but submission answer not found
	mockRepository.On("FindByIdWeb", 1).Return(nil, fmt.Errorf("find by id error"))

	// Call the function you want to test
	result, err := submissionAnswerService.FindByIdWeb(1)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "submission answer not found")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
}

func TestFindByIdWeb_Error(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repository
	submissionAnswerService := NewSubmissionAnswer(mockRepository, nil, nil, nil)

	// Mock the FindByIdWeb method with an error
	mockRepository.On("FindByIdWeb", 1).Return(&web.SubmissionAnswerResponseWebById{}, fmt.Errorf("find by id error"))

	// Call the function you want to test
	result, err := submissionAnswerService.FindByIdWeb(1)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "internal Server Error")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
}