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

func TestUpdateModule_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleRequest{
		Title:       "Updated Module",
		Order:       1,
		Description: "asDSAD",
	}

	existingModule := &domain.Module{
		ID:        moduleID,
		Title:     "Existing Module",
		OrderBy:   1,
		SectionID: 1,
		// Add other necessary fields for Module
	}

	// Set up mock expectations for CekIdFromModule with no error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(existingModule, nil)

	// Set up mock expectations for GetByTitleAndSectionId with no error
	mockRepo.On("GetByTitleAndSectionId", request.Title, existingModule.SectionID).Return(nil, nil)

	// Set up mock expectations for UpdateModule with no error
	mockRepo.On("UpdateModule", mock.Anything, existingModule).Return(nil)

	// Call the function you want to test
	err := moduleService.UpdateModule(request, moduleID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateModule_CekIdFromModuleError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleRequest{
		// Add necessary fields for ModuleRequest
	}

	// Set up mock expectations for CekIdFromModule with an error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(nil, fmt.Errorf("error checking user from module"))

	// Call the function you want to test
	err := moduleService.UpdateModule(request, moduleID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when cek id user from course :error checking user from module")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)

}

func TestUpdateModule_ValidationFailed(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleRequest{
		// Add necessary fields for ModuleRequest
	}

	existingModule := &domain.Module{
		ID: moduleID,
		// Add other necessary fields for Module
	}

	// Set up mock expectations for CekIdFromModule with no error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(existingModule, nil)

	// Call the function you want to test
	err := moduleService.UpdateModule(request, moduleID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'ModuleRequest.SectionID'")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)

}

func TestUpdateModule_ModuleNameAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleRequest{
		Title:       "Updated Module",
		Order:       1,
		Description: "asDSAD",
	}

	existingModule := &domain.Module{
		ID:        moduleID,
		Title:     "Existing Module",
		OrderBy:   1,
		SectionID: 1,
		// Add other necessary fields for Module
	}

	// Set up mock expectations for CekIdFromModule with no error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(existingModule, nil)

	// Set up mock expectations for GetByTitleAndSectionId with no error and existing module
	mockRepo.On("GetByTitleAndSectionId", request.Title, existingModule.SectionID).Return(existingModule, nil)

	// Call the function you want to test
	err := moduleService.UpdateModule(request, moduleID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "module name already exists")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateModule_UpdateModuleError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleRequest{
		Title:       "Updated Module",
		Order:       1,
		Description: "asDSAD",
	}

	existingModule := &domain.Module{
		ID:        moduleID,
		Title:     "Existing Module",
		OrderBy:   1,
		SectionID: 1,
		// Add other necessary fields for Module
	}

	// Set up mock expectations for CekIdFromModule with no error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(existingModule, nil)

	// Set up mock expectations for GetByTitleAndSectionId with no error and no existing module
	mockRepo.On("GetByTitleAndSectionId", request.Title, existingModule.SectionID).Return(nil, nil)

	// Set up mock expectations for UpdateModule with an error
	mockRepo.On("UpdateModule", mock.Anything, existingModule).Return(fmt.Errorf("error updating module"))

	// Call the function you want to test
	err := moduleService.UpdateModule(request, moduleID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when update module : error updating module")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateModuleOrder_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleOrderRequest{
		Order: 2,
	}

	existingModule := &domain.Module{
		ID:        moduleID,
		Title:     "Existing Module",
		OrderBy:   1,
		SectionID: 1,
		// Add other necessary fields for Module
	}

	// Set up mock expectations for CekIdFromModule with no error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(existingModule, nil)

	// Set up mock expectations for UpdateOrderModule with no error
	mockRepo.On("UpdateOrderModule", request.Order, existingModule).Return(nil)

	// Call the function you want to test
	err := moduleService.UpdateModuleOrder(request, moduleID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)

}

func TestUpdateModuleOrder_CekIdFromModuleError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleOrderRequest{
		// Add necessary fields for ModuleOrderRequest
	}

	// Set up mock expectations for CekIdFromModule with an error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(nil, fmt.Errorf("error checking user from module"))

	// Call the function you want to test
	err := moduleService.UpdateModuleOrder(request, moduleID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when cek id user from course :error checking user from module")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateModuleOrder_ValidationFailed(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleOrderRequest{
		// Add necessary fields for ModuleOrderRequest
	}

	// Set up mock expectations for CekIdFromModule with no error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(nil, nil)

	// Call the function you want to test
	err := moduleService.UpdateModuleOrder(request, moduleID, userID, role)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'ModuleOrderRequest.Order'")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateModuleOrder_UpdateOrderModuleError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleOrderRequest{
		Order: 2,
	}

	existingModule := &domain.Module{
		ID:        moduleID,
		Title:     "Existing Module",
		OrderBy:   1,
		SectionID: 1,
		// Add other necessary fields for Module
	}

	// Set up mock expectations for CekIdFromModule with no error
	mockRepo.On("CekIdFromModule", userID, moduleID, role).Return(existingModule, nil)

	// Set up mock expectations for UpdateOrderModule with an error
	mockRepo.On("UpdateOrderModule", request.Order, existingModule).Return(fmt.Errorf("error updating module order"))

	// Call the function you want to test
	err := moduleService.UpdateModuleOrder(request, moduleID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when create module : error updating module order")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateSectionOrder_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionOrderRequest{
		Order: 2,
	}

	existingSection := &domain.Section{
		ID:    sectionID,
		Title: "Existing Section",
		OrderBy: 1,
		// Add other necessary fields for Section
	}

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(existingSection, nil)

	// Set up mock expectations for UpdateOrderSection with no error
	mockRepo.On("UpdateOrderSection", request.Order, existingSection).Return(nil)

	// Call the function you want to test
	err := moduleService.UpdateSectionOrder(request, sectionID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateSectionOrder_CekIdFromSectionError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionOrderRequest{
		// Add necessary fields for SectionOrderRequest
	}

	// Set up mock expectations for CekIdFromSection with an error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(nil, fmt.Errorf("error checking user from section"))

	// Call the function you want to test
	err := moduleService.UpdateSectionOrder(request, sectionID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when cek id user from course :error checking user from section")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateSectionOrder_ValidationFailed(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionOrderRequest{
		// Add necessary fields for SectionOrderRequest
	}

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(nil, nil)

	// Call the function you want to test
	err := moduleService.UpdateSectionOrder(request, sectionID, userID, role)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'SectionOrderRequest.Order'")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateSectionOrder_UpdateOrderSectionError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionOrderRequest{
		Order: 2,
	}

	existingSection := &domain.Section{
		ID:    sectionID,
		Title: "Existing Section",
		OrderBy: 1,
		// Add other necessary fields for Section
	}

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(existingSection, nil)

	// Set up mock expectations for UpdateOrderSection with an error
	mockRepo.On("UpdateOrderSection", request.Order, existingSection).Return(fmt.Errorf("error updating section order"))

	// Call the function you want to test
	err := moduleService.UpdateSectionOrder(request, sectionID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when create module : error updating section order")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateSection_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionUpdateRequest{
		Title:    "Updated Section",
		
	}

	existingSection := &domain.Section{
		ID:       sectionID,
		Title:    "Existing Section",
		// Add other necessary fields for Section
	}

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(existingSection, nil)

	// Set up mock expectations for GetByTitleSectionAndCourseId with no error
	mockRepo.On("GetByTitleSectionAndCourseId", request.Title, existingSection.CourseID).Return(nil, nil)

	// Set up mock expectations for UpdateSection with no error
	mockRepo.On("UpdateSection", mock.Anything, existingSection).Return(nil)

	// Call the function you want to test
	err := moduleService.UpdateSection(request, sectionID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateSection_CekIdFromSectionError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionUpdateRequest{
		// Add necessary fields for SectionUpdateRequest
	}

	// Set up mock expectations for CekIdFromSection with an error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(nil, fmt.Errorf("error checking user from section"))

	// Call the function you want to test
	err := moduleService.UpdateSection(request, sectionID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when cek id user from course :error checking user from section")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateSection_ValidationFailed(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionUpdateRequest{
		// Add necessary fields for SectionUpdateRequest
	}

	existingSection := &domain.Section{
		ID: sectionID,
		// Add other necessary fields for Section
	}

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(existingSection, nil)

	// Call the function you want to test
	err := moduleService.UpdateSection(request, sectionID, userID, role)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'SectionUpdateRequest.Title'")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateSection_SectionNameAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionUpdateRequest{
		Title:    "Updated Section",
	}

	existingSection := &domain.Section{
		ID:       sectionID,
		Title:    "Existing Section",
		// Add other necessary fields for Section
	}

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(existingSection, nil)

	// Set up mock expectations for GetByTitleSectionAndCourseId with no error and existing section
	mockRepo.On("GetByTitleSectionAndCourseId", request.Title, existingSection.CourseID).Return(existingSection, nil)

	// Call the function you want to test
	err := moduleService.UpdateSection(request, sectionID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "section name already exists")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestUpdateSection_UpdateSectionError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionUpdateRequest{
		Title:    "Updated Section",
	}

	existingSection := &domain.Section{
		ID:       sectionID,
		Title:    "Existing Section",
		// Add other necessary fields for Section
	}

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(existingSection, nil)

	// Set up mock expectations for GetByTitleSectionAndCourseId with no error and no existing section
	mockRepo.On("GetByTitleSectionAndCourseId", request.Title, existingSection.CourseID).Return(nil, nil)

	// Set up mock expectations for UpdateSection with an error
	mockRepo.On("UpdateSection", mock.Anything, existingSection).Return(fmt.Errorf("error updating section"))

	// Call the function you want to test
	err := moduleService.UpdateSection(request, sectionID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when update section : error updating section")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}