package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (historyModuleService *HistoryModuleServiceImpl) FindById(id int) (*domain.HistoryModule, error) {
	result, _ := historyModuleService.HistoryModuleRepository.GetById(id)

	if result == nil {
		return nil, fmt.Errorf("history module not found")
	}

	return result, nil
}
