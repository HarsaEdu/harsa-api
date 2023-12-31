package conversion

import "github.com/HarsaEdu/harsa-api/internal/model/web"

func ConvertToSubModuleTrackingResponse(module *web.ModuleResponseForTracking, history *web.HistorySubModuleTracking, subModule *web.SubModuleResponseModule, progress float32) *web.SubModuleTracking {
	return &web.SubModuleTracking{
		ID:               module.ID,
		Title:            module.Title,
		Description:      module.Description,
		Progress:         module.Progress,
		Order:            module.Order,
		HistorySubModule: *history,
		SubModule:        *subModule,
	}
}

func ConvertToModuleTrackingByID(module *web.ModuleResponseForTracking, moduleItems *web.CourseTrackingSub) *web.ModuleTrackingByID {
	return &web.ModuleTrackingByID{
		ID:          module.ID,
		Title:       module.Title,
		Description: module.Description,
		Progress:    module.Progress,
		Order:       module.Order,
		SubModules:  moduleItems.SubModules,
		Submissions: moduleItems.Submissions,
		Quizzes:     moduleItems.Quizzes,
	}
}
