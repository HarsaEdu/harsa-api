package service

import (
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/request"
)

func (service *HistorySubModuleServiceImpl) UpdateHistorySubModule(request *web.UpdateHistorySubModuleRequest) error {
	err := service.Validator.Struct(request)
	if err != nil {
		return err
	}
	model := conversion.UpdateHistorySubModuleRequestToModel(request)
	err = service.Repository.UpdateHistorySubModule(model, request.ID)
	return nil
}
