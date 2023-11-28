package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func GetHistorySubModuleResultToResponse(result []domain.HistorySubModule) []web.GetHistorySubModuleResponse {
	response := []web.GetHistorySubModuleResponse{}
	for _, value := range result {
		history := web.GetHistorySubModuleResponse{
			ID:          value.ID,
			SubModuleID: value.SubModuleID,
			IsCompleted: value.IsCompleted,
		}
		response = append(response, history)
	}
	return response
}
