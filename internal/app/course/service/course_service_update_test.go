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
	// Create mock repositories and dependencies
	mockCourseRepo := new(mocks.CourseRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CourseServiceImpl with the mock repositories and dependencies
	courseService := NewCourseService(mockCourseRepo, mockValidator, mockCloudinaryUploader)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.CourseUpdateRequest{
		Title:       "UpdatedTitle",
		Description: "UpdatedDescription",
		CategoryID:  3,
		UserId: userID,
	}

	// Set up mock expectations for CekIdFromCourse
	mockCourseRepo.On("CekIdFromCourse", userID, courseID, role).Return(&domain.Course{
		// Set existing course details
	}, nil)

	// Set up mock expectations for Uploader (CloudinaryUploader)
	mockCloudinaryUploader.On("Uploader", nil, "file", "courses", false).Return("mockedImageUrl", nil)

	// Set up mock expectations for Update (CourseRepository)
	mockCourseRepo.On("Update", courseID, mock.Anything).Return(nil)

	// Call the function you want to test
	err := courseService.Update(nil, courseID, userID, role, request)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestUpdate_ValidatorError(t *testing.T) {
	// Create mock repositories and dependencies
	mockCourseRepo := new(mocks.CourseRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a CourseServiceImpl with the mock repositories and dependencies
	courseService := NewCourseService(mockCourseRepo, mockValidator, mockCloudinaryUploader)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.CourseUpdateRequest{
		
	}


	// Call the function you want to test
	err := courseService.Update(nil, courseID, userID, role, request)

	// Assert the result
	assert.Error(t, err)
	assert.NotNil(t, err)

	mockCloudinaryUploader.AssertExpectations(t)
}
