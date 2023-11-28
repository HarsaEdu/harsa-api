package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (service *HistorySubModuleServiceImpl) GetHistorySubModuleByUserID(request *web.GetHistorySubModuleRequest) (*[]web.GetHistorySubModuleResponse, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, err
	}

	result, total, err := service.Repository.GetHistorySubModuleByUserID(request.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get history sub modules, something wrong")
	}

	if total == 0 {
		return nil, fmt.Errorf("not found")
	}
	response := conversion.GetHistorySubModuleResultToResponse(result)
	return &response, nil
}
