package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func SubModuleCreateRequestToSubModuleDomain(request *web.SubModuleCreateRequest, moduleId uint) *domain.SubModule {
	return &domain.SubModule{
		ModuleID:    moduleId,
		Title:       request.Title,
		Type:        domain.SubModuleType(request.Type),
		ContentUrl:  request.ContentUrl,
		ContentBody: request.ContentBody,
	}
}

	// func ModuleCreateRequestToModuleDomain(request *web.ModuleCreateRequest, courseId uint) *domain.Module {
	// 	var subModules []*domain.SubModule

	// 	module := &domain.Module{
	// 		CourseID:    courseId,
	// 		Title:       request.Title,
	// 		Description: request.Description,
	// 		Type:        request.Type,
	// 		Order:       request.Order,
	// 	}

	// 	for _, subModule := range request.SubModules {
	// 		subModules = append(subModules, SubModuleCreateRequestToSubModuleDomain(&subModule, module.ID))
	// 	}

	// 	module.SubModules = subModules

	// 	return module
	// }

func ModuleRequestToModuleDomain(request *web.ModuleRequest) *domain.Module {
	module := &domain.Module{
		CourseID:    request.CourseID,
		Title:       request.Title,
		Description: request.Description,
		Type:        request.Type,
		Order:       request.Order,
		SubModules:  request.SubModules,
	}

	return module
}
