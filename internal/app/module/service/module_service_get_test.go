package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllSectionByCourseId_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "section"
	courseID := uint(1)

	section:= domain.Section{
		ID : uint(1),
	}

	// Set up mock expectations for GetAllSectionByCourseId with no error
	mockRepo.On("GetAllSectionByCourseId", offset, limit, search, courseID).Return([]domain.Section{section}, int64(2), nil)

	// Call the function you want to test
	result, pagination, err := moduleService.GetAllSectionByCourseId(offset, limit, search, courseID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetAllSectionByCourseId_RepositoryError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "section"
	courseID := uint(1)

	// Set up mock expectations for GetAllSectionByCourseId with an error
	mockRepo.On("GetAllSectionByCourseId", offset, limit, search, courseID).Return(nil, int64(0), fmt.Errorf("error getting sections"))

	// Call the function you want to test
	result, pagination, err := moduleService.GetAllSectionByCourseId(offset, limit, search, courseID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetAllSectionByCourseId_SectionNotFound(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "section"
	courseID := uint(1)

	// Set up mock expectations for GetAllSectionByCourseId with no error and empty result
	mockRepo.On("GetAllSectionByCourseId", offset, limit, search, courseID).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := moduleService.GetAllSectionByCourseId(offset, limit, search, courseID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}


func TestGetAllModuleBySectionId_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	sectionID := uint(1)

	module := &domain.Section{
		ID: sectionID,
		// Add necessary fields for the Section object
	}

	// Set up mock expectations for GetAllModuleBySecsionId with no error
	mockRepo.On("GetAllModuleBySecsionId", sectionID).Return(module, nil)

	// Call the function you want to test
	result, err := moduleService.GetAllModuleBySectionId(sectionID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetAllModuleBySectionId_RepositoryError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	sectionID := uint(1)

	// Set up mock expectations for GetAllModuleBySecsionId with an error
	mockRepo.On("GetAllModuleBySecsionId", sectionID).Return(nil, fmt.Errorf("error getting modules"))

	// Call the function you want to test
	result, err := moduleService.GetAllModuleBySectionId(sectionID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetModuleById_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	moduleID := uint(1)

	module := &domain.Module{
		ID: moduleID,
		// Add necessary fields for the Module object
	}

	// Set up mock expectations for GetModuleById with no error
	mockRepo.On("GetModuleById", moduleID).Return(module, nil)

	// Call the function you want to test
	result, err := moduleService.GetModuleById(moduleID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestGetModuleById_RepositoryError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	moduleID := uint(1)

	// Set up mock expectations for GetModuleById with an error
	mockRepo.On("GetModuleById", moduleID).Return(nil, fmt.Errorf("error getting module"))

	// Call the function you want to test
	result, err := moduleService.GetModuleById(moduleID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}