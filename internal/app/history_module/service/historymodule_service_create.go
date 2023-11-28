package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (historyModuleService *HistoryModuleServiceImpl) Create(historyModule web.HistoryModuleCreateRequest) error {
	err := historyModuleService.Validator.Struct(historyModule)
	if err != nil {
		return err
	}

	result := conversion.HistoryModuleCreateRequestToHistoryModuleDomain(&historyModule)

	err = historyModuleService.HistoryModuleRepository.Create(result)
	if err != nil {
		return fmt.Errorf("error when creating History Module %s", err.Error())
	}

	return nil

}
