package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (service *HistorySubModuleServiceImpl) UpdateHistorySubModule(request *web.UpdateHistorySubModuleRequest) error {
	err := service.Validator.Struct(request)
	if err != nil {
		return err
	}
	_, err = service.Repository.GetHistorySubModuleByID(request.ID)
	if err != nil {
		return fmt.Errorf("not found")
	}
	model := conversion.UpdateHistorySubModuleRequestToModel(request)
	err = service.Repository.UpdateHistorySubModule(model, request.ID)
	return nil
}
