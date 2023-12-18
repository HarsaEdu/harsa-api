package service

import (
	"errors"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate_Success(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()
	mockSubmission := new(mocks.SubmissionRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader,mockSubmission)
	mockRepository.On("FindByuserIDAndSubmissionID", mock.Anything, mock.Anything).Return(nil, nil)

	idSubmission := 1
	idUser := 1
	mockSubmissionAnswer := &domain.SubmissionAnswer{
		UserID: uint(1),
		SubmissionID: uint(1),
		SubmittedUrl: "file_url.pdf",
		Status: "submitted",
	}

	request := &web.SubmissionAnswerRequest{}

	// Mock the ConvertSubmissionAnswerRequestToSubmissionAnswerDomain method
	mockRepository.On("Create", *mockSubmissionAnswer).Return(nil)
	mockCloudinaryUploader.On("Uploader", nil, "file", "submission_answer", true).Return("file_url.pdf", nil)

	// Call the function you want to test
	err := submissionAnswerService.Create(nil, request, idSubmission, idUser)

	// Assert the result
	assert.NoError(t, err)
	assert.Nil(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestCreate_InvalidFileFormat(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()
	mockSubmission := new(mocks.SubmissionRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader, mockSubmission)

	idSubmission := 1
	idUser := 1

	request := &web.SubmissionAnswerRequest{}
	mockRepository.On("FindByuserIDAndSubmissionID", mock.Anything, mock.Anything).Return(nil, nil)

	// Mock the Uploader method to return an invalid file format
	mockCloudinaryUploader.On("Uploader", nil, "file", "submission_answer", true).Return("file_url.jpg", nil)

	// Call the function you want to test
	err := submissionAnswerService.Create(nil, request, idSubmission, idUser)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid file format")

	// Assert that no further calls were made
	mockRepository.AssertNotCalled(t, "Create")
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestCreate_ErrorUploadingSubmissionAnswer(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()
	mockSubmission := new(mocks.SubmissionRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader, mockSubmission)

	idSubmission := 1
	idUser := 1

	request := &web.SubmissionAnswerRequest{}
	mockRepository.On("FindByuserIDAndSubmissionID", mock.Anything, mock.Anything).Return(nil, nil)

	// Mock the Uploader method to return an error
	mockCloudinaryUploader.On("Uploader", nil, "file", "submission_answer", true).Return("file_url.pdf", errors.New("upload error"))

	// Call the function you want to test
	err := submissionAnswerService.Create(nil, request, idSubmission, idUser)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when uploading submission answer")

	// Assert that no further calls were made
	mockRepository.AssertNotCalled(t, "Create")
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestCreate_ErrorCreatingSubmissionAnswer(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()
	mockSubmission := new(mocks.SubmissionRepository)

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader, mockSubmission)

	idSubmission := 1
	idUser := 1

	mockSubmissionAnswer := &domain.SubmissionAnswer{
		UserID:        uint(idUser),
		SubmissionID:  uint(idSubmission),
		SubmittedUrl:  "file_url.pdf",
		Status:        domain.StatusSubmissionAnswerSubmitted,
	}

	request := &web.SubmissionAnswerRequest{}
	mockRepository.On("FindByuserIDAndSubmissionID", mock.Anything, mock.Anything).Return(nil, nil)

	// Mock the Uploader method
	mockCloudinaryUploader.On("Uploader", nil, "file", "submission_answer", true).Return("file_url.pdf", nil)

	// Mock the Create method to return an error
	mockRepository.On("Create", *mockSubmissionAnswer).Return(errors.New("create error"))

	// Call the function you want to test
	err := submissionAnswerService.Create(nil, request, idSubmission, idUser)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when upload submission answer")

	// Assert that no further calls were made
	mockCloudinaryUploader.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}

