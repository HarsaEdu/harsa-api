package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/history_sub_modules/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type HistorySubModuleService interface {
	GetHistorySubModuleByUserID(request *web.GetHistorySubModuleRequest) (*[]web.GetHistorySubModuleResponse, error)
	CreateHistorySubModule(request *web.CreateHistorySubModuleRequest, userID uint) error
	//		UpdateHistorySubModule(request *web.UpdateHistorySubModuleRequest) error
}

type HistorySubModuleServiceImpl struct {
	Repository repository.HistorySubModuleRepository
	Validator  *validator.Validate
}

func NewHistorySubModuleRepository(repository repository.HistorySubModuleRepository, validator *validator.Validate) HistorySubModuleService {
	return &HistorySubModuleServiceImpl{
		Repository: repository,
		Validator:  validator,
	}
}
