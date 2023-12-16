package service

import (
	"errors"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSection_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionRequest{
		Title:    "Section 1",
		Order:    1,
	}

	// Set up mock expectations for CekIdFromCourse with no error
	mockRepo.On("CekIdFromCourse", userID, courseID, role).Return(nil)

	// Set up mock expectations for Validate.Struct with no error
	mockRepo.On("GetByTitleSectionAndCourseId", request.Title, courseID).Return(nil, nil)
	
	// Set up mock expectations for CreateSection with no error
	mockRepo.On("CreateSection",mock.Anything).Return(nil)

	// Call the function you want to test
	err := moduleService.CreateSection(request, courseID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCreateSection_CekIdFromCourseError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionRequest{
		Title:    "Section 1",
		Order:    1,
		
	}

	// Set up mock expectations for CekIdFromCourse with an error
	mockRepo.On("CekIdFromCourse", userID, courseID, role).Return(errors.New("error checking user from course"))

	// Call the function you want to test
	err := moduleService.CreateSection(request, courseID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when cek id user from course :error checking user from course")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCreateSection_ValidationFailed(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionRequest{
		// Add necessary fields for the SectionRequest object
	}

	// Set up mock expectations for CekIdFromCourse with no error
	mockRepo.On("CekIdFromCourse", userID, courseID, role).Return(nil)

	// Call the function you want to test
	err := moduleService.CreateSection(request, courseID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'SectionRequest.Title'")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCreateSection_SectionNameAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionRequest{
		Title:    "Section 1",
		Order:    1,
	}

	// Set up mock expectations for CekIdFromCourse with no error
	mockRepo.On("CekIdFromCourse", userID, courseID, role).Return(nil)

	// Set up mock expectations for Validate.Struct with no error
	mockRepo.On("GetByTitleSectionAndCourseId", request.Title, courseID).Return(&domain.Section{}, nil)

	// Call the function you want to test
	err := moduleService.CreateSection(request, courseID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "section name already exists")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCreateSection_CreateSectionError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	courseID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.SectionRequest{
		Title:    "Section 1",
		Order:    1,
	}

	// Set up mock expectations for CekIdFromCourse with no error
	mockRepo.On("CekIdFromCourse", userID, courseID, role).Return(nil)

	// Set up mock expectations for Validate.Struct with no error
	mockRepo.On("GetByTitleSectionAndCourseId", request.Title, courseID).Return(nil, nil)

	// Set up mock expectations for CreateSection with an error
	mockRepo.On("CreateSection", mock.Anything).Return(errors.New("error creating section"))

	// Call the function you want to test
	err := moduleService.CreateSection(request, courseID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when create module : error creating section")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}


func TestCreateModule_Success(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleRequest{
		Title:    "Module 1",
		Order:    1,
		Description: "ASDSA",
	}

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(nil,nil)

	// Set up mock expectations for Validate.Struct with no error
	mockRepo.On("GetByTitleAndSectionId", request.Title, sectionID).Return(nil, nil)
	
	// Set up mock expectations for CreateModule with no error
	mockRepo.On("CreateModule", mock.Anything).Return(nil)

	// Call the function you want to test
	err := moduleService.CreateModule(request, sectionID, userID, role)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCreateModule_CekIdFromSectionError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleRequest{
		Title:    "Module 1",
		Order:    1,
		// Add other necessary fields for the ModuleRequest object
	}

	// Set up mock expectations for CekIdFromSection with an error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(&domain.Section{}, errors.New("error checking user from section"))

	// Call the function you want to test
	err := moduleService.CreateModule(request, sectionID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when cek id user from course :error checking user from section")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCreateModule_ValidationFailed(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleRequest{
		// Add necessary fields for the ModuleRequest object
	}

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(&domain.Section{}, nil)

	// Call the function you want to test
	err := moduleService.CreateModule(request, sectionID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'ModuleRequest.Title'")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCreateModule_ModuleNameAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleRequest{
		Title:    "Module 1",
		Order:    1,
		Description: "aSDA",
		// Add other necessary fields for the ModuleRequest object
	}

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(nil, nil)

	// Set up mock expectations for Validate.Struct with no error
	mockRepo.On("GetByTitleAndSectionId", request.Title, sectionID).Return(&domain.Module{
		ID : uint(1),
	}, nil)

	// Call the function you want to test
	err := moduleService.CreateModule(request, sectionID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "module name already exists")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}

func TestCreateModule_CreateModuleError(t *testing.T) {
	mockRepo := new(mocks.ModuleRepository)
	mockValidator := validator.New()

	// Create a ModuleServiceImpl with the mock repository
	moduleService := NewModuleService(mockRepo, mockValidator)

	// Define test data
	sectionID := uint(1)
	userID := uint(2)
	role := "admin"
	request := &web.ModuleRequest{
		Title:    "Module 1",
		Order:    1,
		Description: "aSDA",
		// Add other necessary fields for the ModuleRequest object
	}

	// Set up mock expectations for CekIdFromSection with no error
	mockRepo.On("CekIdFromSection", userID, sectionID, role).Return(nil,nil)

	// Set up mock expectations for Validate.Struct with no error
	mockRepo.On("GetByTitleAndSectionId", request.Title, sectionID).Return(nil, nil)

	// Set up mock expectations for CreateModule with an error
	mockRepo.On("CreateModule", mock.Anything).Return(errors.New("error creating module"))

	// Call the function you want to test
	err := moduleService.CreateModule(request, sectionID, userID, role)

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "error when create module : error creating module")

	// Assert that the expected calls were made
	mockRepo.AssertExpectations(t)
}