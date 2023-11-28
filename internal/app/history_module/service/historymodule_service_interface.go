package service

import (
	"github.com/HarsaEdu/harsa-api/internal/app/history_module/repository"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/go-playground/validator"
)

type HistoryModuleService interface {
	Create(historyModule web.HistoryModuleCreateRequest) error
	Update(historyModule web.HistoryModuleUpdateRequest, id int) error
	FindById(id int) (*domain.HistoryModule, error)
	GetAll(offset, limit int, search string) ([]domain.HistoryModule, *web.Pagination, error)
}

type HistoryModuleServiceImpl struct {
	HistoryModuleRepository repository.HistoryModuleRepository
	Validator               *validator.Validate
}

func NewHistoryModuleService(cr repository.HistoryModuleRepository, validate *validator.Validate) HistoryModuleService {
	return &HistoryModuleServiceImpl{HistoryModuleRepository: cr, Validator: validate}
}
