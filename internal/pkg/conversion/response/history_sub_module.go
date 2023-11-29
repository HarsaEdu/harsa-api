package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertSubModuleResponseTrackingMobile(response *domain.SubModule) *web.SubModuleResponseForTracking {
	return &web.SubModuleResponseForTracking{
		ID:    response.ID,
		Title: response.Title,
	}
}

func ConvertHistorySubmoduleResponseMobile(response *domain.HistorySubModule) *web.HistorySubModuleResponseMobile {

	subModule := ConvertSubModuleResponseTrackingMobile(&response.SubModule)

	return &web.HistorySubModuleResponseMobile{
		ID:         response.ID,
		SubModule:  *subModule,
		IsComplete: response.IsComplete,
	}
}

func ConvertAllHistorySubmoduleResponseMobile(response []domain.HistorySubModule) []web.HistorySubModuleResponseMobile {

	var historySubModule []web.HistorySubModuleResponseMobile

	for i := range response {
		historySubModule = append(historySubModule, *ConvertHistorySubmoduleResponseMobile(&response[i]))
	}

	return historySubModule

}
