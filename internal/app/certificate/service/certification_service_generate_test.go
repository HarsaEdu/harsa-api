package service

import (
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateCertificate_Success(t *testing.T) {
	// Create mock repositories and dependencies
	mockCertificateRepo := new(mocks.CertificateRepository)
	mockCourseTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockPDFGenerator := new(mocks.PDFGenerator)

	// Create a CertificateServiceImpl with the mock repositories and dependencies
	certificateService := NewCertificateService(mockCertificateRepo, mockCourseTrackingRepo, mockCloudinaryUploader, mockPDFGenerator)

	// Define test data
	courseTrackingID := uint(1)
	userID := uint(1)

	// Set up mock expectations for FindById
	mockCourseTrackingRepo.On("FindById", courseTrackingID).Return(&domain.CourseTracking{
		ID:     courseTrackingID,
		Status: "completed",
		UserID: userID,
	}, nil)

	// Set up mock expectations for GenerateCertificate
	mockPDFGenerator.On("GenerateCertificate", mock.Anything).Return([]byte("certificate"), nil)

	// Set up mock expectations for GetById
	mockCertificateRepo.On("GetById", mock.Anything).Return(nil, nil)

	// Set up mock expectations for Create
	mockCertificateRepo.On("Create", mock.AnythingOfType("*domain.Certificate")).Return(nil)

	// Call the function you want to test
	certificate, certificateData, err := certificateService.GenerateCertificate(courseTrackingID, userID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, certificate)
	assert.Nil(t, certificateData)
	assert.Equal(t, []byte("certificate"), certificate)

	// Assert that the expected calls were made
	mockCourseTrackingRepo.AssertExpectations(t)
	mockPDFGenerator.AssertExpectations(t)
	mockCertificateRepo.AssertExpectations(t)
}

func TestGenerateCertificate_CourseTrackingNotFound(t *testing.T) {
	// Create mock repositories and dependencies
	mockCertificateRepo := new(mocks.CertificateRepository)
	mockCourseTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockPDFGenerator := new(mocks.PDFGenerator)

	// Create a CertificateServiceImpl with the mock repositories and dependencies
	certificateService := NewCertificateService(mockCertificateRepo, mockCourseTrackingRepo, mockCloudinaryUploader, mockPDFGenerator)

	// Define test data
	courseTrackingID := uint(1)
	userID := uint(1)

	// Set up mock expectations for FindById
	mockCourseTrackingRepo.On("FindById", courseTrackingID).Return(nil, nil)

	// Call the function you want to test
	certificate, certificateData, err := certificateService.GenerateCertificate(courseTrackingID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, certificate)
	assert.Nil(t, certificateData)
	assert.Contains(t, err.Error(), "course tracking not found")

	// Assert that the expected calls were made
	mockCourseTrackingRepo.AssertExpectations(t)
	mockPDFGenerator.AssertExpectations(t)
	mockCertificateRepo.AssertExpectations(t)
}


func TestGenerateCertificate_CourseNotCompleted(t *testing.T) {
	// Create mock repositories and dependencies
	mockCertificateRepo := new(mocks.CertificateRepository)
	mockCourseTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockPDFGenerator := new(mocks.PDFGenerator)

	// Create a CertificateServiceImpl with the mock repositories and dependencies
	certificateService := NewCertificateService(mockCertificateRepo, mockCourseTrackingRepo, mockCloudinaryUploader, mockPDFGenerator)

	// Define test data
	courseTrackingID := uint(1)
	userID := uint(1)

	// Set up mock expectations for FindById
	mockCourseTrackingRepo.On("FindById", courseTrackingID).Return(&domain.CourseTracking{
		ID:     courseTrackingID,
		Status: "in_progress", // Assuming the course is not completed
		UserID: userID,
	}, nil)

	// Call the function you want to test
	certificate, certificateData, err := certificateService.GenerateCertificate(courseTrackingID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, certificate)
	assert.Nil(t, certificateData)
	assert.Contains(t, err.Error(), "course not completed")

	// Assert that the expected calls were made
	mockCourseTrackingRepo.AssertExpectations(t)
	mockPDFGenerator.AssertExpectations(t)
	mockCertificateRepo.AssertExpectations(t)
}

func TestGenerateCertificate_UnauthorizedAccess(t *testing.T) {
	// Create mock repositories and dependencies
	mockCertificateRepo := new(mocks.CertificateRepository)
	mockCourseTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCloudinaryUploader := new(mocks.CloudinaryUploader)
	mockPDFGenerator := new(mocks.PDFGenerator)

	// Create a CertificateServiceImpl with the mock repositories and dependencies
	certificateService := NewCertificateService(mockCertificateRepo, mockCourseTrackingRepo, mockCloudinaryUploader, mockPDFGenerator)

	// Define test data
	courseTrackingID := uint(1)
	userID := uint(2) // Different user ID

	// Set up mock expectations for FindById
	mockCourseTrackingRepo.On("FindById", courseTrackingID).Return(&domain.CourseTracking{
		ID:     courseTrackingID,
		Status: "completed",
		UserID: 1,
	}, nil)

	// Call the function you want to test
	certificate, certificateData, err := certificateService.GenerateCertificate(courseTrackingID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, certificate)
	assert.Nil(t, certificateData)
	assert.Contains(t, err.Error(), "unauthorized")

	// Assert that the expected calls were made
	mockCourseTrackingRepo.AssertExpectations(t)
	mockPDFGenerator.AssertExpectations(t)
	mockCertificateRepo.AssertExpectations(t)
}
