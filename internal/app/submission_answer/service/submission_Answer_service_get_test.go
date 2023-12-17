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

func TestGet_Success(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()
	mockSubmission := new(mocks.SubmissionRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader, mockSubmission)

	offset := 0
	limit := 10
	submissionID := 1
	search := "some search"

	mockAnswers := []domain.SubmissionsAnswerDetail{
		{ID: 1, FirstName: "SAda", Status: "submitted"},
		{ID: 2, FirstName: "sADSA", Status: "submitted"},
	}

	mockTotal := int64(len(mockAnswers))

	mockPagination := &web.Pagination{
		Offset:  offset,
		Limit:   limit,
		Total:   mockTotal,

	}

	// Mock the Get method with no error
	mockRepository.On("Get", offset, limit, search, uint(submissionID)).Return(mockAnswers, mockTotal, nil)

	// Call the function you want to test
	result, pagination, err := submissionAnswerService.Get(offset, limit, submissionID, search)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)
	assert.Equal(t, mockPagination, pagination)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
}

func TestGet_NotFound(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()
	mockSubmission := new(mocks.SubmissionRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader, mockSubmission)

	offset := 0
	limit := 10
	submissionID := 1
	search := "some search"

	// Mock the Get method with no error, but no answers found
	mockRepository.On("Get", offset, limit, search, uint(submissionID)).Return(nil, int64(0), fmt.Errorf("get answers error"))

	// Call the function you want to test
	result, pagination, err := submissionAnswerService.Get(offset, limit, submissionID, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Contains(t, err.Error(), "not found")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
}

func TestGet_Error(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()
	mockSubmission := new(mocks.SubmissionRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader, mockSubmission)

	offset := 0
	limit := 10
	submissionID := 1
	search := "some search"

	// Mock the Get method with an error
	mockRepository.On("Get", offset, limit, search, uint(submissionID)).Return(nil, int64(1), fmt.Errorf("get answers error"))

	// Call the function you want to test
	result, pagination, err := submissionAnswerService.Get(offset, limit, submissionID, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Contains(t, err.Error(), "get answers error")

	// Assert that no further calls were made
	mockRepository.AssertExpectations(t)
}
