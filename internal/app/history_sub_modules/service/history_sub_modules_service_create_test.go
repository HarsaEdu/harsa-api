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
	"github.com/stretchr/testify/require"
)

func TestCreateHistorySubModules(t *testing.T) {
	mockRepo := new(mocks.HistorySubModuleRepository)
	validate := validator.New()
	historySubModuleService := NewHistorySubModuleRepository(mockRepo, validate)

	historyRequest := web.CreateHistorySubModuleRequest{
		SubModuleID: 1,
	}

	historySubModule := domain.HistorySubModule{
		SubModuleID: 1,
		UserID:      1,
		IsComplete:  true,
	}

	mockRepo.On("CreateHistorySubModule", &historySubModule).Return(nil)

	err := historySubModuleService.CreateHistorySubModule(&historyRequest, 1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestFail(t *testing.T) {
	mockRepo := new(mocks.HistorySubModuleRepository)
	validate := validator.New()
	historySubModuleService := NewHistorySubModuleRepository(mockRepo, validate)

	historyRequest := web.CreateHistorySubModuleRequest{
		SubModuleID: 1,
	}

	expectedError := fmt.Errorf("error")
	mockRepo.On("CreateHistorySubModule", mock.AnythingOfType("*domain.HistorySubModule")).Return(expectedError)

	err := historySubModuleService.CreateHistorySubModule(&historyRequest, 0)

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}

func TestFailValidation(t *testing.T) {
	mockRepo := new(mocks.HistorySubModuleRepository)
	validate := validator.New()
	historySubModuleService := NewHistorySubModuleRepository(mockRepo, validate)

	historyRequest := web.CreateHistorySubModuleRequest{}
	err := historySubModuleService.CreateHistorySubModule(&historyRequest, 0)

	require.Error(t, err)

	mockRepo.AssertExpectations(t)
}
