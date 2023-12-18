package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestCreateInterest(t *testing.T) {
	mockProfileRepo := new(mocks.ProfileRepository)
	mockRepo := new(mocks.InterestRepository)
	validate := validator.New()

	interestService := NewInterestService(mockRepo, validate, mockProfileRepo)

	interest := web.InterestRequest{
		CategoryID: []uint{1},
	}

	profile := domain.ProfileDetailMobile{
		UserID:        1,
		UserProfileID: 1,
	}

	userInterest := domain.UserInterest{
		ProfileID:  1,
		CategoryID: 1,
	}

	mockProfileRepo.On("FindByUserID", uint(1)).Return(&profile, nil)
	mockRepo.On("FindByProfileID", profile.UserProfileID).Return(nil, nil)
	mockRepo.On("CreateInterest", &userInterest).Return(nil, nil)
	err := interestService.CreateInterest(uint(1), &interest)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestCreateInterestValidationError(t *testing.T) {
	mockProfileRepo := new(mocks.ProfileRepository)
	mockRepo := new(mocks.InterestRepository)
	validate := validator.New()

	interestService := NewInterestService(mockRepo, validate, mockProfileRepo)

	interest := web.InterestRequest{}
	err := interestService.CreateInterest(uint(1), &interest)

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestCreateInterestExists(t *testing.T) {
	mockProfileRepo := new(mocks.ProfileRepository)
	mockRepo := new(mocks.InterestRepository)
	validate := validator.New()

	interestService := NewInterestService(mockRepo, validate, mockProfileRepo)

	interest := web.InterestRequest{
		CategoryID: []uint{1},
	}

	profile := domain.ProfileDetailMobile{
		UserID:        1,
		UserProfileID: 1,
	}

	userInterest := domain.UserInterest{
		ProfileID:  1,
		CategoryID: 1,
	}

	mockProfileRepo.On("FindByUserID", uint(1)).Return(&profile, nil)
	mockRepo.On("FindByProfileID", profile.UserProfileID).Return(&userInterest, nil)
	err := interestService.CreateInterest(uint(1), &interest)

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestCreateInterestInternalServerError(t *testing.T) {
	mockProfileRepo := new(mocks.ProfileRepository)
	mockRepo := new(mocks.InterestRepository)
	validate := validator.New()

	interestService := NewInterestService(mockRepo, validate, mockProfileRepo)

	interest := web.InterestRequest{
		CategoryID: []uint{1},
	}

	profile := domain.ProfileDetailMobile{
		UserID:        1,
		UserProfileID: 1,
	}

	userInterest := domain.UserInterest{
		ProfileID:  1,
		CategoryID: 1,
	}

	mockProfileRepo.On("FindByUserID", uint(1)).Return(&profile, nil)
	mockRepo.On("FindByProfileID", profile.UserProfileID).Return(nil, nil)
	mockRepo.On("CreateInterest", &userInterest).Return(fmt.Errorf("error"))
	err := interestService.CreateInterest(uint(1), &interest)

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}
