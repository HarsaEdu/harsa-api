package service

import (
	"errors"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCountAll(t *testing.T) {
	// Create a mock DashboardRepository
	mockRepo := new(mocks.DashboardRepository)

	// Create a DashboardServiceImpl with the mockRepo
	dashboardService := NewDashboardService(mockRepo)

	// Define expected values
	expectedCountCourse := int64(5)
	expectedCountInstructor := int64(10)
	expectedCountInterest := []web.CountInterest{{Name: "Interest1", Count: 3}, {Name: "Interest2", Count: 5}}
	expectedCountStudent := int64(20)

	// Set up mock expectations
	mockRepo.On("CountCourse").Return(expectedCountCourse, nil)
	mockRepo.On("CountIntructure").Return(expectedCountInstructor, nil)
	mockRepo.On("CountInterest").Return(expectedCountInterest, expectedCountStudent, nil)

	// Call the function you want to test
	result, err := dashboardService.CountAll()

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedCountCourse, result.CountCourse)
	assert.Equal(t, expectedCountInstructor, result.CountIntructure)
	assert.Equal(t, expectedCountInterest, result.CountInterest)
	assert.Equal(t, expectedCountStudent, result.CountStudent)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCountCourseWithError(t *testing.T) {
	// Create a mock DashboardRepository
	mockRepo := new(mocks.DashboardRepository)

	// Create a DashboardServiceImpl with the mockRepo
	dashboardService := NewDashboardService(mockRepo)

	// Define an expected error
	expectedError := errors.New("some error")

	// Set up mock expectations for CountCourse to return an error
	mockRepo.On("CountCourse").Return(int64(0), expectedError)

	// Call the function you want to test
	result, err := dashboardService.CountAll()

	// Assert that the result is nil (as there's an error)
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCountIntructureWithError(t *testing.T) {
	// Create a mock DashboardRepository
	mockRepo := new(mocks.DashboardRepository)

	// Create a DashboardServiceImpl with the mockRepo
	dashboardService := NewDashboardService(mockRepo)

	// Define an expected error
	expectedError := errors.New("some error")

	// Set up mock expectations for CountIntructure to return an error
	mockRepo.On("CountCourse").Return(int64(10), nil)
	mockRepo.On("CountIntructure").Return(int64(0), expectedError)

	// Call the function you want to test
	result, err := dashboardService.CountAll()

	// Assert that the result is nil (as there's an error)
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCountInterestWithError(t *testing.T) {
	// Create a mock DashboardRepository
	mockRepo := new(mocks.DashboardRepository)

	// Create a DashboardServiceImpl with the mockRepo
	dashboardService := NewDashboardService(mockRepo)

	// Define an expected error
	expectedError := errors.New("some error")

	// Set up mock expectations for CountIntructure to return an error
	mockRepo.On("CountCourse").Return(int64(10), nil)
	mockRepo.On("CountIntructure").Return(int64(0), nil)
	mockRepo.On("CountInterest").Return(nil, int64(0), expectedError)

	// Call the function you want to test
	result, err := dashboardService.CountAll()

	// Assert that the result is nil (as there's an error)
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}
