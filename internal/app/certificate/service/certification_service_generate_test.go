package service

import (
	"embed"
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/HarsaEdu/harsa-api/web/template"
	
)



func TestGenerateCertificate(t *testing.T) {
	// Create mock repositories and dependencies
	mockCertificateRepo := new(mocks.CertificateRepository)
	mockCourseTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	blankCertificateTemplate := template.CertificateBlankContent

	// Create a CertificateServiceImpl with the mock repositories and dependencies
	certificateService := NewCertificateService(mockCertificateRepo, mockCourseTrackingRepo, mockCloudinaryUploader, blankCertificateTemplate)

	// Define test data
	courseTrackingID := uint(1)
	userID := uint(1)

	// Set up mock expectations for finding course tracking
	mockCourseTrackingRepo.On("FindById", courseTrackingID).Return(&domain.CourseTracking{ID: courseTrackingID, UserID: userID}, nil)

	// Set up mock expectations for generating the certificate
	mockCertificateRepo.On("Create", mock.AnythingOfType("*domain.Certificate")).Return(nil)

	// Call the function you want to test
	certificatePdf, certificate, err := certificateService.GenerateCertificate(courseTrackingID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, certificatePdf)
	assert.Nil(t, certificate)

	// Assert that the expected calls were made
	mockCourseTrackingRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)
}

func TestGenerateCertificate1(t *testing.T) {
	// Create mock repositories and dependencies
	mockCertificateRepo := new(mocks.CertificateRepository)
	mockCourseTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	blankCertificateTemplate := embed.FS{}

	// Create a CertificateServiceImpl with the mock repositories and dependencies
	certificateService := NewCertificateService(mockCertificateRepo, mockCourseTrackingRepo, mockCloudinaryUploader, blankCertificateTemplate)

	// Define test data
	courseTrackingID := uint(1)
	userID := uint(1)

	// Test case: Course tracking not found
	mockCourseTrackingRepo.On("FindById", courseTrackingID).Return(nil, fmt.Errorf("not found"))

	// Call the function you want to test
	certificatePdf, certificate, err := certificateService.GenerateCertificate(courseTrackingID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, certificatePdf)
	assert.Nil(t, certificate)

	// Assert that the expected calls were made
	mockCourseTrackingRepo.AssertExpectations(t)
	mockCertificateRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)

	// Reset the mocks for the next test case
	mockCourseTrackingRepo.ExpectedCalls = nil
	mockCertificateRepo.ExpectedCalls = nil
	mockCloudinaryUploader.ExpectedCalls = nil

	// Test case: Unauthorized user
	mockCourseTrackingRepo.On("FindById", courseTrackingID).Return(&domain.CourseTracking{ID: courseTrackingID, UserID: 2}, nil)

	// Call the function you want to test
	certificatePdf, certificate, err = certificateService.GenerateCertificate(courseTrackingID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, certificatePdf)
	assert.Nil(t, certificate)

	// Assert that the expected calls were made
	mockCourseTrackingRepo.AssertExpectations(t)
	mockCertificateRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)

	// Reset the mocks for the next test case
	mockCourseTrackingRepo.ExpectedCalls = nil
	mockCertificateRepo.ExpectedCalls = nil
	mockCloudinaryUploader.ExpectedCalls = nil

	// Test case: Certificate already exists
	mockCourseTrackingRepo.On("FindById", courseTrackingID).Return(&domain.CourseTracking{ID: courseTrackingID, UserID: userID}, nil)
	mockCertificateRepo.On("GetById", mock.AnythingOfType("uint")).Return(&domain.Certificate{ID: "id"}, nil)

	// Call the function you want to test
	certificatePdf, certificate, err = certificateService.GenerateCertificate(courseTrackingID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, certificatePdf)
	assert.Nil(t, certificate)

	// Assert that the expected calls were made
	mockCourseTrackingRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)

}

func TestGenerateCertificate2(t *testing.T) {
	// Create mock repositories and dependencies
	mockCertificateRepo := new(mocks.CertificateRepository)
	mockCourseTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	blankCertificateTemplate := template.CertificateBlankContent

	// Create a CertificateServiceImpl with the mock repositories and dependencies
	certificateService := NewCertificateService(mockCertificateRepo, mockCourseTrackingRepo, mockCloudinaryUploader, blankCertificateTemplate)

	// Define test data
	courseTrackingID := uint(1)
	userID := uint(1)

	// Test case: Course tracking not found
	mockCourseTrackingRepo.On("FindById", courseTrackingID).Return(nil, fmt.Errorf("not found"))

	// Call the function you want to test
	certificatePdf, certificate, err := certificateService.GenerateCertificate(courseTrackingID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, certificatePdf)
	assert.Nil(t, certificate)

	// Assert that the expected calls were made
	mockCourseTrackingRepo.AssertExpectations(t)
	mockCertificateRepo.AssertExpectations(t)
	mockCloudinaryUploader.AssertExpectations(t)

	// Reset the mocks for the next test case
	mockCourseTrackingRepo.ExpectedCalls = nil
	mockCertificateRepo.ExpectedCalls = nil
	mockCloudinaryUploader.ExpectedCalls = nil

	// ... (rest of the testing code)

	// Test case: Certificate creation success
	mockCourseTrackingRepo.On("FindById", courseTrackingID).Return(&domain.CourseTracking{ID: courseTrackingID, UserID: userID}, nil)
	mockCertificateRepo.On("GetById", mock.AnythingOfType("uint")).Return(nil, fmt.Errorf("not found"))
	mockCloudinaryUploader.On("Uploader", mock.Anything, mock.Anything).Return("uploaded_file_url", nil)
	mockCertificateRepo.On("Create", mock.AnythingOfType("*domain.Certificate")).Return(nil)

	// Call the function you want to test
	certificatePdf, certificate, err = certificateService.GenerateCertificate(courseTrackingID, userID)

	// Assert the result
	assert.NoError(t, err)
	assert.Nil(t, certificatePdf)
	assert.Nil(t, certificate)

	// Assert that the expected calls were made
	mockCourseTrackingRepo.AssertExpectations(t)
}