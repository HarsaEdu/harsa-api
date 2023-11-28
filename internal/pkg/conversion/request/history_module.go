package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func HistoryModuleCreateRequestToHistoryModuleDomain(historyModule *web.HistoryModuleCreateRequest) *domain.HistoryModule {
	return &domain.HistoryModule{UserID: historyModule.UserID}
}

func HistoryModuleUpdateRequestToHistoryModuleDomain(historyModule *web.HistoryModuleUpdateRequest) *domain.HistoryModule {
	return &domain.HistoryModule{UserID: historyModule.UserID}
}
