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

// TestUpdateImage_Success tests the case when updating the image is successful.
func TestUpdateImage_Success(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	// Mock the SubsPlanRepository FindById method with a valid result
	expectedPlan := &domain.SubsPlan{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedPlan, nil)

	// Mock the CloudinaryUploader Uploader method with valid results
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs_plan", true).Return("image.png", nil)

	// Mock the SubsPlanRepository UpdateImage method with no error
	mockRepository.On("UpdateImage", "image.png", 1).Return(nil)

	// Create a SubsPlanUpdateImage request with valid data
	updateImageRequest := &web.SubsPlanUpdateImage{
		// Provide valid data for the update request
	}

	// Call the function you want to test
	err := subsPlanService.UpdateImage(nil, updateImageRequest, 1)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

// TestUpdateImage_ErrorSubsPlanNotFound tests the case when the subscription plan is not found.
func TestUpdateImage_ErrorSubsPlanNotFound(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	// Mock the SubsPlanRepository FindById method with no plan found
	mockRepository.On("FindById", mock.Anything).Return(nil, fmt.Errorf("error"))

	// Create a SubsPlanUpdateImage request with valid data
	updateImageRequest := &web.SubsPlanUpdateImage{
		// Provide valid data for the update request
	}

	// Call the function you want to test
	err := subsPlanService.UpdateImage(nil, updateImageRequest, 1)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "subs plan not found")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

// TestUpdateImage_ErrorInvalidFileFormat tests the case when the uploaded image has an invalid file format.
func TestUpdateImage_ErrorInvalidFileFormat(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	// Mock the SubsPlanRepository FindById method with a valid result
	expectedPlan := &domain.SubsPlan{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedPlan, nil)

	// Mock the CloudinaryUploader Uploader method with an invalid file format
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs_plan", true).Return("document.pdf", nil)

	// Create a SubsPlanUpdateImage request with valid data
	updateImageRequest := &web.SubsPlanUpdateImage{
		// Provide valid data for the update request
	}

	// Call the function you want to test
	err := subsPlanService.UpdateImage(nil, updateImageRequest, 1)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid file format")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

// TestUpdateImage_ErrorUploadingImage tests the case when there's an error uploading the image.
func TestUpdateImage_ErrorUploadingImage(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	// Mock the SubsPlanRepository FindById method with a valid result
	expectedPlan := &domain.SubsPlan{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedPlan, nil)

	// Mock the CloudinaryUploader Uploader method with an error
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs_plan", true).Return("sdsd.png", fmt.Errorf("upload error"))

	// Create a SubsPlanUpdateImage request with valid data
	updateImageRequest := &web.SubsPlanUpdateImage{
		// Provide valid data for the update request
	}

	// Call the function you want to test
	err := subsPlanService.UpdateImage(nil, updateImageRequest, 1)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when uploading image")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}



// TestUpdateImage_ErrorUpdateImage tests the case when there's an error updating the image in the repository.
func TestUpdateImage_ErrorUpdateImage(t *testing.T) {
	mockRepository := new(mocks.SubsPlanRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockValidator := validator.New()

	// Create a SubsPlanServiceImpl with the mock repositories
	subsPlanService := NewsubsPlanService(mockRepository, mockValidator, mockCloudinaryUploader)

	// Mock the SubsPlanRepository FindById method with a valid result
	expectedPlan := &domain.SubsPlan{
		ID: uint(1),
	} // Replace with your actual struct
	mockRepository.On("FindById", mock.Anything).Return(expectedPlan, nil)

	// Mock the CloudinaryUploader Uploader method with valid results
	mockCloudinaryUploader.On("Uploader", mock.Anything, "image", "subs_plan", true).Return("image.png", nil)

	// Mock the SubsPlanRepository UpdateImage method with an error
	mockRepository.On("UpdateImage", "image.png", 1).Return(fmt.Errorf("update error"))

	// Create a SubsPlanUpdateImage request with valid data
	updateImageRequest := &web.SubsPlanUpdateImage{
		// Provide valid data for the update request
	}

	// Call the function you want to test
	err := subsPlanService.UpdateImage(nil, updateImageRequest, 1)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when updating image subs plan")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

