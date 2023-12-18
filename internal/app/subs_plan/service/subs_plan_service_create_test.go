package service

import (
	"fmt"
	"mime/multipart"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate_ValidationError(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanCreateRequest{}

	var file *multipart.FileHeader

	// Call the function you want to test
	err := subsPlanService.Create(nil, file, request)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key: 'SubsPlanCreateRequest.Title'")
}

func TestCreate_Success(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanCreateRequest{
		// Valid request with all required fields
		Title:         "Test Plan",
		Price:         100,
		Duration_days: 10,
	}

	var file *multipart.FileHeader

	// Mock the SubsPlanRepository Create method with no error
	mockRepository.On("Create", mock.Anything).Return(nil)

	// Call the function you want to test
	err := subsPlanService.Create(nil, file, request)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockCloudinaryUploader.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}

func TestCreate_Success2(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanCreateRequest{
		// Valid request with all required fields
		Title:         "Test Plan",
		Price:         100,
		Duration_days: 10,
	}

	var file *multipart.FileHeader
	file = &multipart.FileHeader{
		Filename: "image.png",
	}

	// Mock the CloudinaryUploader method with no error
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs-plan", false).Return("image.png", nil)

	// Mock the SubsPlanRepository Create method with no error
	mockRepository.On("Create", mock.Anything).Return(nil)

	// Call the function you want to test
	err := subsPlanService.Create(nil, file, request)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockCloudinaryUploader.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}

func TestCreate_ErrorUploadingImage(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanCreateRequest{
		// Valid request with all required fields
		Title:         "Test Plan",
		Price:         100,
		Duration_days: 10,
	}

	var file *multipart.FileHeader
	file = &multipart.FileHeader{
		Filename: "invalid_image.txt", // Use an invalid image file
	}

	// Mock the CloudinaryUploader method with an error
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs-plan", false).Return("", fmt.Errorf("EROR"))

	// Call the function you want to test
	err := subsPlanService.Create(nil, file, request)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when uploading image")

	// Assert that the expected calls were made
	mockCloudinaryUploader.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}

// TestCreate_InvalidImageFormat tests the case when the uploaded image has an invalid format.
func TestCreate_InvalidImageFormat(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanCreateRequest{
		// Valid request with all required fields
		Title:         "Test Plan",
		Price:         100,
		Duration_days: 10,
	}

	var file *multipart.FileHeader
	file = &multipart.FileHeader{
		Filename: "invalid_image.txt", // Use an invalid image file
	}

	// Mock the CloudinaryUploader method with success
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs-plan", false).Return("invalid_image.txt", nil)

	// Call the function you want to test
	err := subsPlanService.Create(nil, file, request)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid file format")

	// Assert that the expected calls were made
	mockCloudinaryUploader.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}

func TestCreateFromExisting_Success(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanUpdateRequest{
		// Valid request with all required fields
		Title:         "Updated Plan",
		Price:         150,
		Duration_days: 15,
	}

	var file *multipart.FileHeader
	file = &multipart.FileHeader{
		Filename: "image.png",
	}

	// Mock the SubsPlanRepository FindById method with an existing plan
	existingPlan := &domain.SubsPlan{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(existingPlan, nil)

	// Mock the CloudinaryUploader method with success
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs-plan", false).Return("image.png", nil)

	// Mock the SubsPlanRepository Create method with no error
	mockRepository.On("Create", mock.Anything).Return(nil)

	// Mock the SubsPlanRepository UpdateStatus method with no error
	mockRepository.On("UpdateStatus", false, mock.Anything).Return(nil)

	// Call the function you want to test
	err := subsPlanService.CreateFromExisting(nil, file, request, 1) // Replace 1 with an actual ID

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestCreateFromExisting_ErrorSubsPlanNotFound(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanUpdateRequest{
		// Valid request with all required fields
		Title:         "Updated Plan",
		Price:         150,
		Duration_days: 15,
	}

	var file *multipart.FileHeader
	file = &multipart.FileHeader{
		Filename: "image.png",
	}

	// Mock the SubsPlanRepository FindById method with no plan found
	mockRepository.On("FindById", mock.Anything).Return(nil, fmt.Errorf("EROR"))

	// Call the function you want to test
	err := subsPlanService.CreateFromExisting(nil, file, request, 1) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "subs plan not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

// TestCreateFromExisting_ErrorUploadingImage tests the case when there's an error uploading an image.
func TestCreateFromExisting_ErrorUploadingImage(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockValidator := validator.New()
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	request := &web.SubsPlanUpdateRequest{
		// Valid request with all required fields
		Title:         "Updated Plan",
		Price:         150,
		Duration_days: 15,
	}

	var file *multipart.FileHeader
	file = &multipart.FileHeader{
		Filename: "invalid_image.txt", // Use an invalid image file
	}

	// Mock the SubsPlanRepository FindById method with an existing plan
	existingPlan := &domain.SubsPlan{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(existingPlan, nil)

	// Mock the CloudinaryUploader method with an error
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs-plan", false).Return("", fmt.Errorf("EROR"))

	// Call the function you want to test
	err := subsPlanService.CreateFromExisting(nil, file, request, 1) // Replace 1 with an actual ID

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when uploading image")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}
