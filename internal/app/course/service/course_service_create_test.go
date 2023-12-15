package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CourseServiceImpl with the mock repositories and dependencies
	courseService := NewCourseService(mockCourseRepo, mockValidator, mockCloudinaryUploader)


	request := &web.CourseCreateRequest{
		UserId: uint(1),
		Title: "sadsa",
		CategoryID: 1,
	}
	instructorID := uint(1)

	// Set up mock expectations for Uploader (CloudinaryUploader)
	mockCloudinaryUploader.On("Uploader", mock.Anything, "file", "courses", false).Return("mockedImageUrl", nil)

	// Set up mock expectations for Create (CourseRepository)
	mockCourseRepo.On("Create", mock.Anything).Return(nil)

	// Call the function you want to test
	err := courseService.Create(nil, request, instructorID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockCloudinaryUploader.AssertExpectations(t)
	mockCourseRepo.AssertExpectations(t)
}

func TestCreate_ValidatorError(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CourseServiceImpl with the mock repositories and dependencies
	courseService := NewCourseService(mockCourseRepo, mockValidator, mockCloudinaryUploader)


	request := &web.CourseCreateRequest{
		
	}
	instructorID := uint(1)

	
	// Set up mock expectations for Create (CourseRepository)
	mockCourseRepo.On("Create", mock.Anything).Return(nil)

	// Call the function you want to test
	err := courseService.Create(nil, request, instructorID)

	// Assert the result
	require.Error(t, err)
	assert.NotNil(t, err)
}

func TestCreate_ImageUploaderError(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CourseServiceImpl with the mock repositories and dependencies
	courseService := NewCourseService(mockCourseRepo, mockValidator, mockCloudinaryUploader)

	request := &web.CourseCreateRequest{
		UserId:     uint(1),
		Title:      "sadsa",
		CategoryID: 1,
	}
	instructorID := uint(1)


	// Set up mock expectations for Uploader (CloudinaryUploader)
	mockCloudinaryUploader.On("Uploader", mock.Anything, "file", "courses", false).Return("", fmt.Errorf("image upload error"))

	// Call the function you want to test
	err := courseService.Create(nil, request, instructorID)

	// Assert the result
	require.Error(t, err)
	assert.NotNil(t, err)
	assert.Equal(t, "error when uploading image : image upload error", err.Error())

	mockCloudinaryUploader.AssertExpectations(t)
	mockCourseRepo.AssertExpectations(t)
}