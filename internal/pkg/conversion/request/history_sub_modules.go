package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func CreateHistorySubModuleRequestToModel(request *web.CreateHistorySubModuleRequest, userID uint) *domain.HistorySubModule {
	return &domain.HistorySubModule{
		SubModuleID: request.SubModuleID,
		UserID:      userID,
		IsCompleted: false,
	}
}
