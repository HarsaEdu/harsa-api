package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (historyModuleService *HistoryModuleServiceImpl) Update(historyModule web.HistoryModuleUpdateRequest, id int) error {
	err := historyModuleService.Validator.Struct(historyModule)
	if err != nil {
		return err
	}

	ifExist, _ := historyModuleService.HistoryModuleRepository.GetById(id)
	if ifExist == nil {
		return fmt.Errorf("history Module not found")
	}

	result := conversion.HistoryModuleUpdateRequestToHistoryModuleDomain(&historyModule)

	err = historyModuleService.HistoryModuleRepository.Update(id, result)
	if err != nil {
		return fmt.Errorf("error when updating history Module %s", err.Error())
	}

	return nil
}
