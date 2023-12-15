package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDelete_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromCourse
	mockCourseRepo.On("CekIdFromCourse", userID, courseID, role).Return(&domain.Course{
		// Set course details
	}, nil)

	// Set up mock expectations for Delete
	mockCourseRepo.On("Delete", mock.Anything).Return(nil)

	// Call the function you want to test
	err := courseService.Delete(courseID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestDelete_CekIdFromCourseError(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromCourse
	mockCourseRepo.On("CekIdFromCourse", userID, courseID, role).Return(nil, fmt.Errorf("error checking ID"))

	// Call the function you want to test
	err := courseService.Delete(courseID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, "error when cek id user in course delete : error checking ID", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestDelete_DeleteError(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromCourse
	mockCourseRepo.On("CekIdFromCourse", userID, courseID, role).Return(&domain.Course{
		// Set course details
	}, nil)

	// Set up mock expectations for Delete
	mockCourseRepo.On("Delete", mock.Anything).Return(fmt.Errorf("error deleting"))

	// Call the function you want to test
	err := courseService.Delete(courseID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, "error when deleting : error deleting", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}
