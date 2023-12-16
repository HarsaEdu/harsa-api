package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCourseByUserIdMobile_Success(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	userID := uint(1)
	status := "active"

	// Set up mock expectations for GetAllCourseTrackingMobile
	expectedResult := []web.GetAllCourseForTraking{
		// Add necessary fields for the GetAllCourseForTraking object
	}
	expectedTotal := int64(20)
	mockTrackingRepo.On("GetAllCourseTrackingMobile", offset, limit, userID, search, status).Return(expectedResult, expectedTotal, nil)

	// Call the function you want to test
	result, pagination, err := courseTrackingService.GetAllCourseByUserIdMobile(offset, limit, search, userID, status)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}

func TestGetAllCourseByUserIdMobile_Error(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	userID := uint(1)
	status := "active"

	// Set up mock expectations for GetAllCourseTrackingMobile with an error
	mockTrackingRepo.On("GetAllCourseTrackingMobile", offset, limit, userID, search, status).Return(nil, int64(0), fmt.Errorf("some error"))

	// Call the function you want to test
	result, pagination, err := courseTrackingService.GetAllCourseByUserIdMobile(offset, limit, search, userID, status)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}

func TestGetAllCourseByUserIdMobile_ErrorNil(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	userID := uint(1)
	status := "active"

	// Set up mock expectations for GetAllCourseTrackingMobile with an error
	mockTrackingRepo.On("GetAllCourseTrackingMobile", offset, limit, userID, search, status).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := courseTrackingService.GetAllCourseByUserIdMobile(offset, limit, search, userID, status)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}

func TestGetAllCourseByUserIdWeb_Success(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	userID := uint(1)

	course1 := domain.CourseTracking{
		ID: uint(2),
	}
	
	// Set up mock expectations for GetAllCourseTrackingWeb with expected results
	expectedCourses := []domain.CourseTracking{
		course1,
	}
	mockTrackingRepo.On("GetAllCourseTrackingWeb", offset, limit, userID).Return(expectedCourses, int64(len(expectedCourses)), nil)

	// Call the function you want to test
	result, pagination, err := courseTrackingService.GetAllCourseByUserIdWeb(offset, limit, userID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}
func TestGetAllCourseByUserIdWeb_Error(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	userID := uint(1)

	// Set up mock expectations for GetAllCourseTrackingWeb with an error
	mockTrackingRepo.On("GetAllCourseTrackingWeb", offset, limit, userID).Return(nil, int64(0), fmt.Errorf("error fetching courses"))

	// Call the function you want to test
	result, pagination, err := courseTrackingService.GetAllCourseByUserIdWeb(offset, limit, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}

func TestGetAllCourseByUserIdWeb_ErrorNil(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	userID := uint(1)

	// Set up mock expectations for GetAllCourseTrackingWeb with an error
	mockTrackingRepo.On("GetAllCourseTrackingWeb", offset, limit, userID).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := courseTrackingService.GetAllCourseByUserIdWeb(offset, limit, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}

func TestGetAllUserCourseWeb_Success(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	courseID := uint(1)
	search := "example"

	userCourse1 := domain.CourseTracking{
		ID: uint(2),
	}
	
	// Set up mock expectations for GetAllCourseTrackingUserWeb with expected results
	expectedUserCourses := []domain.CourseTracking{
		userCourse1,
	}
	mockTrackingRepo.On("GetAllCourseTrackingUserWeb", offset, limit, courseID, search).Return(expectedUserCourses, int64(len(expectedUserCourses)), nil)

	// Call the function you want to test
	result, pagination, err := courseTrackingService.GetAllUserCourseWeb(offset, limit, courseID, search)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}

func TestGetAllUserCourseWeb_Error(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	courseID := uint(1)
	search := "example"

	// Set up mock expectations for GetAllCourseTrackingUserWeb with an error
	mockError := fmt.Errorf("error fetching user courses")
	mockTrackingRepo.On("GetAllCourseTrackingUserWeb", offset, limit, courseID, search).Return(nil, int64(0), mockError)

	// Call the function you want to test
	result, pagination, err := courseTrackingService.GetAllUserCourseWeb(offset, limit, courseID, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}

func TestGetAllUserCourseWeb_ErrorNil(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	courseID := uint(1)
	search := "example"

	// Set up mock expectations for GetAllCourseTrackingUserWeb with an error
	
	mockTrackingRepo.On("GetAllCourseTrackingUserWeb", offset, limit, courseID, search).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := courseTrackingService.GetAllUserCourseWeb(offset, limit, courseID, search)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}
