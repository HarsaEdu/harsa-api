package service

import (
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate_Success(t *testing.T) {
	mockRepository := new(mocks.SubmissionAnswerRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()

	// Create a SubmissionAnswerServiceImpl with the mock repositories
	submissionAnswerService := NewSubmissionAnswer(mockRepository, mockValidator, mockCloudinaryUploader, nil)

	request := &web.SubmissionAnswerUpdateRequest{ /* provide valid update request data */ }
	id := 1
	idUser := 1

	// Mock the FindById method with no error
	mockRepository.On("FindById", id).Return(&domain.SubmissionAnswer{ID: uint(id)}, nil)

	// Mock the Uploader method with no error
	mockCloudinaryUploader.On("Uploader", nil, "file", "submission_answer", true).Return("file_url.pdf", nil)

	// Mock the Update method with no error
	mockRepository.On("Update", mock.Anything, id, idUser).Return(nil)

	// Call the function you want to test
	err := submissionAnswerService.Update(nil, request, id, idUser)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

