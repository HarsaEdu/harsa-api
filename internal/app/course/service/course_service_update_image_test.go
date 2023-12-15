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

func TestUpdateImage_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CourseServiceImpl with the mock repositories and dependencies
	courseService := NewCourseService(mockCourseRepo, mockValidator, mockCloudinaryUploader)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.CourseUpdateImageRequest{
		// Set image update request fields
	}

	// Set up mock expectations for CekIdFromCourse
	mockCourseRepo.On("CekIdFromCourse", userID, courseID, role).Return(&domain.Course{
		// Set existing course details
	}, nil)

	// Set up mock expectations for Uploader (CloudinaryUploader)
	mockCloudinaryUploader.On("Uploader", nil, "file", "courses", true).Return("mockedImageUrl", nil)

	// Set up mock expectations for UpdateImage (CourseRepository)
	mockCourseRepo.On("UpdateImage", mock.Anything).Return(nil)

	// Call the function you want to test
	err := courseService.UpdateImage(nil, courseID, userID, role, request)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestUpdateImage_CekIdFromCourseError(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CourseServiceImpl with the mock repositories and dependencies
	courseService := NewCourseService(mockCourseRepo, mockValidator, mockCloudinaryUploader)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.CourseUpdateImageRequest{
		// Set image update request fields
	}

	// Set up mock expectations for CekIdFromCourse
	mockCourseRepo.On("CekIdFromCourse", userID, courseID, role).Return(nil, fmt.Errorf("error checking course ID"))

	// Call the function you want to test
	err := courseService.UpdateImage(nil, courseID, userID, role, request)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, "error when cek id user in course update image : error checking course ID", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestUpdateImage_UploaderError(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CourseServiceImpl with the mock repositories and dependencies
	courseService := NewCourseService(mockCourseRepo, mockValidator, mockCloudinaryUploader)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.CourseUpdateImageRequest{
		// Set image update request fields
	}

	// Set up mock expectations for CekIdFromCourse
	mockCourseRepo.On("CekIdFromCourse", userID, courseID, role).Return(&domain.Course{
		// Set existing course details
	}, nil)

	// Set up mock expectations for Uploader (CloudinaryUploader) error
	mockCloudinaryUploader.On("Uploader", nil, "file", "courses", true).Return("", fmt.Errorf("errorchecking image"))

	err := courseService.UpdateImage(nil, courseID, userID, role, request)

// Assert the result
assert.Error(t, err)
assert.Equal(t, "error when uploading image : errorchecking image", err.Error())

// Assert that the expected calls were made
mockCourseRepo.AssertExpectations(t)
mockCloudinaryUploader.AssertExpectations(t)
}

func TestUpdateImage_UpdateImageError(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CourseServiceImpl with the mock repositories and dependencies
	courseService := NewCourseService(mockCourseRepo, mockValidator, mockCloudinaryUploader)

// Define test data
courseID := uint(1)
userID := uint(2)
role := "admin"
request := &web.CourseUpdateImageRequest{
	// Set image update request fields
}

// Set up mock expectations for CekIdFromCourse
mockCourseRepo.On("CekIdFromCourse", userID, courseID, role).Return(&domain.Course{
	// Set existing course details
}, nil)

// Set up mock expectations for Uploader (CloudinaryUploader)
mockCloudinaryUploader.On("Uploader", nil, "file", "courses", true).Return("mockedImageUrl", nil)

// Set up mock expectations for UpdateImage (CourseRepository) error
mockCourseRepo.On("UpdateImage", mock.Anything).Return(fmt.Errorf("error updating image"))

// Call the function you want to test
err := courseService.UpdateImage(nil, courseID, userID, role, request)

// Assert the result
assert.Error(t, err)
assert.Equal(t, "error when updating image course : error updating image", err.Error())

// Assert that the expected calls were made
mockCourseRepo.AssertExpectations(t)
mockCloudinaryUploader.AssertExpectations(t)
}
