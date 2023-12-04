package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (service *HistorySubModuleServiceImpl) CreateHistorySubModule(request *web.CreateHistorySubModuleRequest, userID uint) error {
	err := service.Validator.Struct(request)
	if err != nil {
		return err
	}

	model := conversion.CreateHistorySubModuleRequestToModel(request, userID)

	err = service.Repository.CreateHistorySubModule(model)
	if err != nil {
		return fmt.Errorf("failed to create history sub module, something happen")
	}
	return nil
}
