package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"

	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetById_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	courseID := uint(1)

	// Set up mock expectations for GetById
	mockCourseRepo.On("GetById", courseID).Return(&domain.Course{
		// Set course details
	}, nil)

	// Call the function you want to test
	result, err := courseService.GetById(courseID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetById_CourseNotFound(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	courseID := uint(1)

	// Set up mock expectations for GetById
	mockCourseRepo.On("GetById", courseID).Return(nil, fmt.Errorf("course not found"))

	// Call the function you want to test
	result, err := courseService.GetById(courseID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "course not found", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetById_Error(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	courseID := uint(1)

	// Set up mock expectations for GetById
	mockCourseRepo.On("GetById", courseID).Return(nil, fmt.Errorf("error getting course"))

	// Call the function you want to test
	result, err := courseService.GetById(courseID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "error getting course", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetDeatailCourse_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	courseID := uint(1)

	// Set up mock expectations for GetDetailDashBoardIntructur
	mockCourseRepo.On("GetDetailDashBoardIntructur", courseID).Return(&web.CourseResponseForIntructur{
		// Set course details
	}, nil)

	// Call the function you want to test
	result, err := courseService.GetDeatailCourse(courseID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetDeatailCourse_CourseNotFound(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	courseID := uint(1)

	// Set up mock expectations for GetDetailDashBoardIntructur
	mockCourseRepo.On("GetDetailDashBoardIntructur", courseID).Return(nil, fmt.Errorf("course not found"))

	// Call the function you want to test
	result, err := courseService.GetDeatailCourse(courseID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "course not found", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetDeatailCourse_Error(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	courseID := uint(1)

	// Set up mock expectations for GetDetailDashBoardIntructur
	mockCourseRepo.On("GetDetailDashBoardIntructur", courseID).Return(nil, fmt.Errorf("error getting course details"))

	// Call the function you want to test
	result, err := courseService.GetDeatailCourse(courseID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "error getting course details", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetByIdMobile_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	courseID := uint(1)

	// Set up mock expectations for GetByIdMobile
	mockCourseRepo.On("GetByIdMobile", courseID).Return(&domain.Course{
		ID: uint(1),
	}, int64(5), int64(10), nil)

	// Call the function you want to test
	result, err := courseService.GetByIdMobile(courseID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetByIdMobile_Error(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	courseID := uint(1)

	// Set up mock expectations for GetByIdMobile
	mockCourseRepo.On("GetByIdMobile", courseID).Return(nil, int64(0), int64(0), fmt.Errorf("error getting course details"))

	// Call the function you want to test
	result, err := courseService.GetByIdMobile(courseID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "error getting course details", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}