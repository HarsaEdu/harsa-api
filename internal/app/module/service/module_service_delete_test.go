package service

import (
	"errors"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteModule_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromModule with no error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(&domain.Module{}, nil)

	// Set up mock expectations for DeleteModule with no error
	mockRepo.On("DeleteModule", mock.Anything).Return(nil)

	// Call the function you want to test
	err := moduleService.DeleteModule(moduleID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDeleteModule_CekIdFromModuleError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromModule with an error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(nil, errors.New("error checking user from module"))

	// Call the function you want to test
	err := moduleService.DeleteModule(moduleID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when cek id user from course :error checking user from module")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDeleteModule_DeleteModuleError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromModule with no error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(&domain.Module{}, nil)

	// Set up mock expectations for DeleteModule with an error
	mockRepo.On("DeleteModule", mock.Anything).Return(errors.New("error deleting module"))

	// Call the function you want to test
	err := moduleService.DeleteModule(moduleID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when update module : error deleting module")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDeleteSection_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(&domain.Section{}, nil)

	// Set up mock expectations for DeleteSection with no error
	mockRepo.On("DeleteSection", mock.Anything).Return(nil)

	// Call the function you want to test
	err := moduleService.DeleteSection(sectionID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDeleteSection_CekIdFromSectionError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromSection with an error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(nil, errors.New("error checking user from section"))

	// Call the function you want to test
	err := moduleService.DeleteSection(sectionID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when cek id user from section :error checking user from section")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDeleteSection_DeleteSectionError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(&domain.Section{}, nil)

	// Set up mock expectations for DeleteSection with an error
	mockRepo.On("DeleteSection", mock.Anything).Return(errors.New("error deleting section"))

	// Call the function you want to test
	err := moduleService.DeleteSection(sectionID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when update section : error deleting section")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDeleteSubModule_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	subModuleID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromSubModule with no error
	mockRepo.On("CekIdFromSubModule", userID, subModuleID, role).Return(&domain.SubModule{}, nil)

	// Set up mock expectations for DeleteSubModule with no error
	mockRepo.On("DeleteSubModule", mock.Anything).Return(nil)

	// Call the function you want to test
	err := moduleService.DeleteSubModule(subModuleID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDeleteSubModule_CekIdFromSubModuleError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	subModuleID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromSubModule with an error
	mockRepo.On("CekIdFromSubModule", userID, subModuleID, role).Return(nil, errors.New("error checking user from submodule"))

	// Call the function you want to test
	err := moduleService.DeleteSubModule(subModuleID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when cek id user from course :error checking user from submodule")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestDeleteSubModule_DeleteSubModuleError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, nil)

	// Define test data
	subModuleID := uint(1)
	userID := uint(2)
	role := "admin"

	// Set up mock expectations for CekIdFromSubModule with no error
	mockRepo.On("CekIdFromSubModule", userID, subModuleID, role).Return(&domain.SubModule{}, nil)

	// Set up mock expectations for DeleteSubModule with an error
	mockRepo.On("DeleteSubModule", mock.Anything).Return(errors.New("error deleting submodule"))

	// Call the function you want to test
	err := moduleService.DeleteSubModule(subModuleID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when update section : error deleting submodule")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}