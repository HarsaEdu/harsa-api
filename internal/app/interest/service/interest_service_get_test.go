package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestGetInterest(t *testing.T) {
	mockProfileRepo := new(mocks.ProfileRepository)
	mockRepo := new(mocks.InterestRepository)
	validate := validator.New()

	interestService := NewInterestService(mockRepo, validate, mockProfileRepo)

	profile := domain.ProfileDetailMobile{
		UserID:        1,
		UserProfileID: 1,
	}

	courses := []domain.CourseEntity{
		{ID: 1}, {ID: 2},
	}

	mockProfileRepo.On("FindByUserID", uint(1)).Return(&profile, nil)
	mockRepo.On("GetInterestRecommendation", profile.UserProfileID).
		Return(courses, int64(2), nil)
	res, err := interestService.GetInterestRecommendation(uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, res)

	mockRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestGetInterestInternalServerError(t *testing.T) {
	mockProfileRepo := new(mocks.ProfileRepository)
	mockRepo := new(mocks.InterestRepository)
	validate := validator.New()

	interestService := NewInterestService(mockRepo, validate, mockProfileRepo)

	profile := domain.ProfileDetailMobile{
		UserID:        1,
		UserProfileID: 1,
	}

	courses := []domain.CourseEntity{
		{ID: 1}, {ID: 2},
	}

	mockProfileRepo.On("FindByUserID", uint(1)).Return(&profile, nil)
	mockRepo.On("GetInterestRecommendation", profile.UserProfileID).
		Return(courses, int64(2), fmt.Errorf("error"))
	_, err := interestService.GetInterestRecommendation(uint(1))

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestGetInterestProfileNotFound(t *testing.T) {
	mockProfileRepo := new(mocks.ProfileRepository)
	mockRepo := new(mocks.InterestRepository)
	validate := validator.New()

	interestService := NewInterestService(mockRepo, validate, mockProfileRepo)

	mockProfileRepo.On("FindByUserID", uint(1)).Return(nil, fmt.Errorf("error"))
	_, err := interestService.GetInterestRecommendation(uint(1))

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestNoInterest(t *testing.T) {
	mockProfileRepo := new(mocks.ProfileRepository)
	mockRepo := new(mocks.InterestRepository)
	validate := validator.New()

	interestService := NewInterestService(mockRepo, validate, mockProfileRepo)

	profile := domain.ProfileDetailMobile{
		UserID:        1,
		UserProfileID: 1,
	}

	courses := []domain.CourseEntity{
		{ID: 1}, {ID: 2},
	}

	mockProfileRepo.On("FindByUserID", uint(1)).Return(&profile, nil)
	mockRepo.On("GetInterestRecommendation", profile.UserProfileID).
		Return(courses, int64(0), nil)
	_, err := interestService.GetInterestRecommendation(uint(1))

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}
