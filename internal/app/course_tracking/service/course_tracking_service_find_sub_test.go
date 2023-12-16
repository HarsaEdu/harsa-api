package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFindSubByIdMobile_Success(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)

	// Set up mock expectations for FindAllSub
	mockTrackingRepo.On("FindAllSub", moduleID, userID).Return(&web.CourseTrackingSub{}, nil)

	// Call the function you want to test
	result, err := courseTrackingService.FindSubByIdMobile(moduleID, userID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}

func TestFindSubByIdMobile_Error(t *testing.T) {
	mockTrackingRepo := new(mocks.CourseTrackingRepository)

	// Create a CourseTrackingServiceImpl with the mock repository
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, nil, nil, nil, nil)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)

	// Set up mock expectations for FindAllSub returning an error
	expectedError := fmt.Errorf("some error")
	mockTrackingRepo.On("FindAllSub", moduleID, userID).Return(nil, expectedError)

	// Call the function you want to test
	result, err := courseTrackingService.FindSubByIdMobile(moduleID, userID)

	// Assert that an error occurred
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockTrackingRepo.AssertExpectations(t)
}

func TestFindModuleHistory_Success(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, nil, mockSubscription, nil)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)

	// Set up mock expectations for IsSubscription
	mockSubscription.On("IsSubscription", nil, userID).Return(true, nil)

	// Set up mock expectations for GetCourseIDbyModuleID
	mockTrackingRepo.On("GetCourseIDbyModuleID", moduleID).Return(uint(3), nil)

	// Set up mock expectations for FindByUserIdAndCourseID
	mockTrackingRepo.On("FindByUserIdAndCourseID", uint(3), userID).Return(&domain.CourseTracking{
		ID: uint(4),
	}, nil)

	// Set up mock expectations for FindModuleTracking
	expectedModule := &web.ModuleResponseForTracking{
		// Add necessary fields for the HistoryModule object
	}
	mockTrackingRepo.On("FindModuleTracking", moduleID, userID).Return(expectedModule, nil)

	// Set up mock expectations for FindAllSub
	expectedCourseTracking := &web.CourseTrackingSub{
		// Add necessary fields for the CourseTrackingSub object
	}
	mockTrackingRepo.On("FindAllSub", moduleID, userID).Return(expectedCourseTracking, nil)

	// Call the function you want to test
	result, err := courseTrackingService.FindModuleHistory(nil, moduleID, userID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockTrackingRepo.AssertExpectations(t)
	mockCourseRepo.AssertExpectations(t)
}

func TestFindModuleHistory_FindAllSubError(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, nil, mockSubscription, nil)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)

	// Set up mock expectations for IsSubscription
	mockSubscription.On("IsSubscription", nil, userID).Return(true, nil)

	// Set up mock expectations for GetCourseIDbyModuleID
	mockTrackingRepo.On("GetCourseIDbyModuleID", moduleID).Return(uint(3), nil)

	// Set up mock expectations for FindByUserIdAndCourseID
	mockTrackingRepo.On("FindByUserIdAndCourseID", uint(3), userID).Return(&domain.CourseTracking{
		ID: uint(4),
	}, nil)

	// Set up mock expectations for FindModuleTracking
	expectedModule := &web.ModuleResponseForTracking{
		// Add necessary fields for the HistoryModule object
	}
	mockTrackingRepo.On("FindModuleTracking", moduleID, userID).Return(expectedModule, nil)

	// Set up mock expectations for FindAllSub with an error
	mockTrackingRepo.On("FindAllSub", moduleID, userID).Return(nil, fmt.Errorf("error finding all submodules"))

	// Call the function you want to test
	result, err := courseTrackingService.FindModuleHistory(nil, moduleID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockTrackingRepo.AssertExpectations(t)
	mockCourseRepo.AssertExpectations(t)
}

func TestFindModuleHistory_FindModuleTrackingError(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, nil, mockSubscription, nil)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)

	// Set up mock expectations for IsSubscription
	mockSubscription.On("IsSubscription", nil, userID).Return(true, nil)

	// Set up mock expectations for GetCourseIDbyModuleID
	mockTrackingRepo.On("GetCourseIDbyModuleID", moduleID).Return(uint(3), nil)

	// Set up mock expectations for FindByUserIdAndCourseID
	mockTrackingRepo.On("FindByUserIdAndCourseID", uint(3), userID).Return(&domain.CourseTracking{
		ID: uint(4),
	}, nil)

	// Set up mock expectations for FindModuleTracking with an error
	mockTrackingRepo.On("FindModuleTracking", moduleID, userID).Return(nil, fmt.Errorf("error finding module tracking"))

	// Call the function you want to test
	result, err := courseTrackingService.FindModuleHistory(nil, moduleID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockTrackingRepo.AssertExpectations(t)
	mockCourseRepo.AssertExpectations(t)
}



func TestFindModuleHistory_FindCourseIDByModuleIDError(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, nil, mockSubscription, nil)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)

	// Set up mock expectations for IsSubscription
	mockSubscription.On("IsSubscription", nil, userID).Return(true, nil)

	// Set up mock expectations for GetCourseIDbyModuleID with an error
	mockTrackingRepo.On("GetCourseIDbyModuleID", moduleID).Return(uint(0), fmt.Errorf("error getting course ID"))

	// Call the function you want to test
	result, err := courseTrackingService.FindModuleHistory(nil, moduleID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockTrackingRepo.AssertExpectations(t)
	mockCourseRepo.AssertExpectations(t)
}

func TestFindModuleHistory_FindByUserIdAndCourseIDError(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, nil, mockSubscription, nil)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)

	// Set up mock expectations for IsSubscription
	mockSubscription.On("IsSubscription", nil, userID).Return(true, nil)

	// Set up mock expectations for GetCourseIDbyModuleID
	mockTrackingRepo.On("GetCourseIDbyModuleID", moduleID).Return(uint(3), nil)

	// Set up mock expectations for FindByUserIdAndCourseID with an error
	mockTrackingRepo.On("FindByUserIdAndCourseID", uint(3), userID).Return(nil, fmt.Errorf("error finding course tracking by ID"))

	// Call the function you want to test
	result, err := courseTrackingService.FindModuleHistory(nil, moduleID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockTrackingRepo.AssertExpectations(t)
	mockCourseRepo.AssertExpectations(t)
}
func TestFindModuleHistory_FindByUserIdAndCourseIDErrorNil(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, nil, mockSubscription, nil)

	// Define test data
	moduleID := uint(1)
	userID := uint(2)

	course:=&domain.CourseTracking{
		ID : uint(0),
	}

	// Set up mock expectations for IsSubscription
	mockSubscription.On("IsSubscription", nil, userID).Return(true, nil)

	// Set up mock expectations for GetCourseIDbyModuleID
	mockTrackingRepo.On("GetCourseIDbyModuleID", moduleID).Return(uint(3), nil)

	// Set up mock expectations for FindByUserIdAndCourseID with an error
	mockTrackingRepo.On("FindByUserIdAndCourseID", uint(3), userID).Return(course, nil)

	// Call the function you want to test
	result, err := courseTrackingService.FindModuleHistory(nil, moduleID, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockTrackingRepo.AssertExpectations(t)
	mockCourseRepo.AssertExpectations(t)
}
func TestFindSubModuleByID_Success(t *testing.T) {
	mockSubscription := new(mocks.SubscriptionService)
	mockTrackingRepo := new(mocks.CourseTrackingRepository)
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
	courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, nil, mockSubscription, nil)

	// Define test data
	moduleID := uint(1)
	subModuleID := uint(2)
	userID := uint(3)

	// Set up mock expectations for IsSubscription
	mockSubscription.On("IsSubscription", nil, userID).Return(true, nil)

	// Set up mock expectations for GetCourseIDbyModuleID
	mockTrackingRepo.On("GetCourseIDbyModuleID", moduleID).Return(uint(4), nil)

	// Set up mock expectations for FindByUserIdAndCourseID
	mockTrackingRepo.On("FindByUserIdAndCourseID", uint(4), userID).Return(&domain.CourseTracking{
		ID: uint(5),
	}, nil)

	// Set up mock expectations for FindModuleTracking
	expectedModule := &web.ModuleResponseForTracking{
		// Add necessary fields for the ModuleResponseForTracking object
	}
	mockTrackingRepo.On("FindModuleTracking", moduleID, userID).Return(expectedModule, nil)

	// Set up mock expectations for FindSubModuleByID with expected results
	expectedHistory := &domain.HistorySubModule{
		// Add necessary fields for the ModuleTrackingHistory object
	}
	expectedSubModule := &domain.SubModule{
		// Add necessary fields for the SubModule object
	}
	mockTrackingRepo.On("FindSubModuleByID", moduleID, userID, subModuleID).Return(expectedHistory, expectedSubModule, nil)

	// Set up mock expectations for CountProgressModule
	mockTrackingRepo.On("CountProgressModule", moduleID, userID).Return(float32(75), nil)

	// Call the function you want to test
	result, err := courseTrackingService.FindSubModuleByID(nil, moduleID, subModuleID, userID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockTrackingRepo.AssertExpectations(t)
	mockCourseRepo.AssertExpectations(t)
	assert.Equal(t, float32(0), result.Progress)


	// Assert that the expected calls were made
	mockSubscription.AssertExpectations(t)
	mockTrackingRepo.AssertExpectations(t)
	mockCourseRepo.AssertExpectations(t)
	}

	func TestFindSubModuleByID_Error(t *testing.T) {
		mockSubscription := new(mocks.SubscriptionService)
		mockTrackingRepo := new(mocks.CourseTrackingRepository)
		mockCourseRepo := new(mocks.CourseRepository)
	
		// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
		courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, nil, mockSubscription, nil)
	
		// Define test data
		moduleID := uint(1)
		subModuleID := uint(2)
		userID := uint(3)
	
		// Set up mock expectations for IsSubscription
		mockSubscription.On("IsSubscription", nil, userID).Return(true, nil)
	
		// Set up mock expectations for GetCourseIDbyModuleID
		mockTrackingRepo.On("GetCourseIDbyModuleID", moduleID).Return(uint(4), nil)
	
		// Set up mock expectations for FindByUserIdAndCourseID
		mockTrackingRepo.On("FindByUserIdAndCourseID", uint(4), userID).Return(&domain.CourseTracking{
			ID: uint(5),
		}, nil)
	
		// Set up mock expectations for FindModuleTracking
		expectedModule := &web.ModuleResponseForTracking{
			// Add necessary fields for the ModuleResponseForTracking object
		}
		mockTrackingRepo.On("FindModuleTracking", moduleID, userID).Return(expectedModule, nil)
	
		// Set up mock expectations for FindSubModuleByID with an error
		mockTrackingRepo.On("FindSubModuleByID", moduleID, userID, subModuleID).Return(nil, nil, fmt.Errorf("error finding sub-module"))
	
		// Call the function you want to test
		result, err := courseTrackingService.FindSubModuleByID(nil, moduleID, subModuleID, userID)
	
		// Assert the result
		assert.Error(t, err)
		assert.Nil(t, result)
	
		// Assert that the expected calls were made
		mockSubscription.AssertExpectations(t)
		mockTrackingRepo.AssertExpectations(t)
		mockCourseRepo.AssertExpectations(t)
	}

	func TestFindSubmissionByID_Success(t *testing.T) {
		mockSubscription := new(mocks.SubscriptionService)
		mockTrackingRepo := new(mocks.CourseTrackingRepository)
		mockCourseRepo := new(mocks.CourseRepository)
	
		// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
		courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, nil, mockSubscription, nil)
	
		// Define test data
		moduleID := uint(1)
		userID := uint(2)
		submissionID := uint(3)
	
		// Set up mock expectations for IsSubscription
		mockSubscription.On("IsSubscription", nil, userID).Return(true, nil)
	
		// Set up mock expectations for GetCourseIDbyModuleID
		mockTrackingRepo.On("GetCourseIDbyModuleID", moduleID).Return(uint(4), nil)
	
		// Set up mock expectations for FindByUserIdAndCourseID
		mockTrackingRepo.On("FindByUserIdAndCourseID", uint(4), userID).Return(&domain.CourseTracking{
			ID: uint(5),
		}, nil)
	
		// Set up mock expectations for FindModuleTracking
		expectedModule := &web.ModuleResponseForTracking{
			// Add necessary fields for the ModuleResponseForTracking object
		}
		mockTrackingRepo.On("FindModuleTracking", moduleID, userID).Return(expectedModule, nil)
	
		// Set up mock expectations for FindSubmissionByID with expected results
		expectedSubmissionAnswer := &domain.SubmissionAnswer{
			// Add necessary fields for the SubmissionAnswer object
		}
		expectedSubmission := &domain.Submissions{
			// Add necessary fields for the Submission object
		}
		mockTrackingRepo.On("FindSubmissionByID", moduleID, userID, submissionID).Return(expectedSubmissionAnswer, expectedSubmission, nil)
	
		// Call the function you want to test
		result, err := courseTrackingService.FindSubmissionByID(nil, moduleID, userID, submissionID)
	
		// Assert the result
		assert.NoError(t, err)
		assert.NotNil(t, result)
	
		// Assert that the expected calls were made
		mockSubscription.AssertExpectations(t)
		mockTrackingRepo.AssertExpectations(t)
		mockCourseRepo.AssertExpectations(t)
	}
	
	func TestFindSubmissionByID_Unauthorized(t *testing.T) {
		mockSubscription := new(mocks.SubscriptionService)
		mockTrackingRepo := new(mocks.CourseTrackingRepository)
		mockCourseRepo := new(mocks.CourseRepository)
	
		// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
		courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, nil, mockSubscription, nil)
	
		// Define test data
		moduleID := uint(1)
		userID := uint(2)
		submissionID := uint(3)
	
		// Set up mock expectations for IsSubscription returning unauthorized
		mockSubscription.On("IsSubscription", nil, userID).Return(false, nil)
	
		// Call the function you want to test
		result, err := courseTrackingService.FindSubmissionByID(nil, moduleID, userID, submissionID)
	
		// Assert the result
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "unauthorized")
	
		// Assert that the expected calls were made
		mockSubscription.AssertExpectations(t)
		mockTrackingRepo.AssertNotCalled(t, "GetCourseIDbyModuleID")
		mockTrackingRepo.AssertNotCalled(t, "FindByUserIdAndCourseID")
		mockTrackingRepo.AssertNotCalled(t, "FindModuleTracking")
		mockTrackingRepo.AssertNotCalled(t, "FindSubmissionByID")
		mockCourseRepo.AssertNotCalled(t, "GetByIdMobile")
	}


	func TestFindQuizzByID_Success(t *testing.T) {
		mockSubscription := new(mocks.SubscriptionService)
		mockTrackingRepo := new(mocks.CourseTrackingRepository)
		mockQuizzService := new(mocks.QuizzesService)
		mockCourseRepo := new(mocks.CourseRepository)
	
		// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
		courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, mockQuizzService, mockSubscription, nil)
	
		// Define test data
		moduleID := uint(1)
		userID := uint(2)
		quizzID := uint(3)
	
		// Set up mock expectations for IsSubscription
		mockSubscription.On("IsSubscription", nil, userID).Return(true, nil)
	
		// Set up mock expectations for GetCourseIDbyModuleID
		mockTrackingRepo.On("GetCourseIDbyModuleID", moduleID).Return(uint(4), nil)
	
		// Set up mock expectations for FindByUserIdAndCourseID
		mockTrackingRepo.On("FindByUserIdAndCourseID", uint(4), userID).Return(&domain.CourseTracking{
			ID: uint(5),
		}, nil)
	
		// Set up mock expectations for FindModuleTracking
		expectedModule := &web.ModuleResponseForTracking{
			// Add necessary fields for the ModuleResponseForTracking object
		}
		mockTrackingRepo.On("FindModuleTracking", moduleID, userID).Return(expectedModule, nil)
	
		// Set up mock expectations for FindQuizzByID with expected results
		expectedHistoryQuizz := &domain.HistoryQuiz{
			// Add necessary fields for the HistoryQuiz object
		}
		mockTrackingRepo.On("FindQuizzByID", moduleID, userID, quizzID).Return(expectedHistoryQuizz, nil)
	
		// Set up mock expectations for FindByIdMobile (QuizzService)
		expectedQuizz := &web.QuizResponse{
			// Add necessary fields for the QuizzResponseMobile object
		}
		mockQuizzService.On("FindByIdMobile", quizzID).Return(expectedQuizz, nil)
	
		// Call the function you want to test
		result, err := courseTrackingService.FindQuizzByID(nil, moduleID, userID, quizzID)
	
		// Assert the result
		assert.NoError(t, err)
		assert.NotNil(t, result)
	
		// Assert that the expected calls were made
		mockSubscription.AssertExpectations(t)
		mockTrackingRepo.AssertExpectations(t)
		mockQuizzService.AssertExpectations(t)
		mockCourseRepo.AssertExpectations(t)
	}

	func TestFindQuizzByID_Error(t *testing.T) {
		mockSubscription := new(mocks.SubscriptionService)
		mockTrackingRepo := new(mocks.CourseTrackingRepository)
		mockQuizzService := new(mocks.QuizzesService)
		mockCourseRepo := new(mocks.CourseRepository)
	
		// Create a CourseTrackingServiceImpl with the mock repositories and dependencies
		courseTrackingService := NewCourseTrackingService(mockTrackingRepo, nil, mockCourseRepo, mockQuizzService, mockSubscription, nil)
	
		// Define test data
		moduleID := uint(1)
		userID := uint(2)
		quizzID := uint(3)
	
		// Set up mock expectations for IsSubscription
		mockSubscription.On("IsSubscription", nil, userID).Return(true, nil)
	
		// Set up mock expectations for GetCourseIDbyModuleID
		mockTrackingRepo.On("GetCourseIDbyModuleID", moduleID).Return(uint(4), nil)
	
		// Set up mock expectations for FindByUserIdAndCourseID
		mockTrackingRepo.On("FindByUserIdAndCourseID", uint(4), userID).Return(&domain.CourseTracking{
			ID: uint(5),
		}, nil)
	
		// Set up mock expectations for FindModuleTracking
		expectedModule := &web.ModuleResponseForTracking{
			// Add necessary fields for the ModuleResponseForTracking object
		}
		mockTrackingRepo.On("FindModuleTracking", moduleID, userID).Return(expectedModule, nil)
	
		// Set up mock expectations for FindQuizzByID with an error
		mockError := fmt.Errorf("history not found")
		mockTrackingRepo.On("FindQuizzByID", moduleID, userID, quizzID).Return(nil, mockError)
	
		// Call the function you want to test
		result, err := courseTrackingService.FindQuizzByID(nil, moduleID, userID, quizzID)
	
		// Assert the result
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, mockError.Error())
	
		// Assert that the expected calls were made
		mockSubscription.AssertExpectations(t)
		mockTrackingRepo.AssertExpectations(t)
		mockQuizzService.AssertExpectations(t)
		mockCourseRepo.AssertExpectations(t)
	}
	
	