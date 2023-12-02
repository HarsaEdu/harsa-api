package conversion

import "github.com/HarsaEdu/harsa-api/internal/model/web"

func ConvertToSubModuleTrackingResponse(module *web.ModuleResponseForTracking, history *web.HistorySubModuleResponseMobile, subModule *web.SubModuleResponseModule, progress float32) *web.SubModuleTrackingResponse {
	return &web.SubModuleTrackingResponse{
		ID:               module.ID,
		Title:            module.Title,
		Description:      module.Description,
		Progress:         module.Progress,
		Order:            module.Order,
		HistorySubModule: *history,
		SubModule:        *subModule,
	}
}
