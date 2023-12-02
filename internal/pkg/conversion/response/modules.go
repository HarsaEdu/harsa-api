package conversion

import (

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertModuleResponse(module *domain.Module) *web.ModuleResponse{
	
	subModules := ConvertAllSubModuleModule(module.SubModules)

	submissions := ConvertAllSubmissionModule(module.Submissions)
	
	quizzes := ConvertAllQuizModule(module.Quizzes)

	moduleRes := web.ModuleResponse{
		ID:            module.ID,
		CourseID:      module.CourseID,
		Title:         module.Title,
		Description:   module.Description,
		Order:         module.OrderBy,
		CreatedAt:     module.CreatedAt,
		UpdatedAt:     module.UpdatedAt,
		SubModules:    subModules,
		Submissions:   submissions,
		Quizzes:       quizzes,

	}
	return &moduleRes
}

func ConvertModuleWithTitle(module domain.Module) *web.ModuleResponseWithTitle{
	return &web.ModuleResponseWithTitle{
		ID : module.ID,
		Title: module.Title,
		Description: module.Description,
		Order : module.OrderBy,
	}
}

func ConvertAllModuleWithTitle(module []domain.Module) []web.ModuleResponseWithTitle{
	var moduleRes []web.ModuleResponseWithTitle

	for i := range module {
		moduleRes = append(moduleRes, *ConvertModuleWithTitle(module[i]))
	}

	return moduleRes
}

func ConvertSectionResponse(section *domain.Section) *web.SectionResponse{
	
	module := ConvertAllModuleWithTitle(section.Modules)

	sectionRes := web.SectionResponse{
		ID:            section.ID,
		CourseID:      section.CourseID,
		Title:         section.Title,
		Order:         section.OrderBy,
		Modules:       module,
	}
	return &sectionRes
}

func ConvertAllSectionResponse(section []domain.Section) []web.SectionResponse{
	
	var sectionRes []web.SectionResponse

	for i := range section {
		sectionRes = append(sectionRes, *ConvertSectionResponse(&section[i]))
	}

	return sectionRes

}

func ConvertAllModuleResponse(module []domain.Module) []web.ModuleResponse{
	
	var moduleRes []web.ModuleResponse

	for i := range module {
		moduleRes = append(moduleRes, *ConvertModuleResponse(&module[i]))
	}

	return moduleRes
}

func ConvertModuleRequest(module *web.ModuleRequest) *domain.Module{
	
	moduleRes := domain.Module{
		CourseID:      module.CourseID,
		Title:         module.Title,
		Description:   module.Description,
		OrderBy:       module.Order,
	}
	return &moduleRes
}

func ConvertModuleResponseMobile(module *domain.Module) *web.ModuleResponseMobile{
	
	subModules := ConvertAllSubModuleModuleMobile(module.SubModules)

	submissions := ConvertAllSubmissionModuleMobile(module.Submissions)
	
	quizzes := ConvertAllQuizModuleMobile(module.Quizzes)

	moduleRes := web.ModuleResponseMobile{
		ID:            module.ID,
		CourseID:      module.CourseID,
		Title:         module.Title,
		Description:   module.Description,
		Order:         module.OrderBy,
		CreatedAt:     module.CreatedAt,
		UpdatedAt:     module.UpdatedAt,
		SubModules:    subModules,
		Submissions:   submissions,
		Quizzes:       quizzes,

	}
	return &moduleRes
}

func ConvertAllModuleResponseMobile(module []domain.Module) []web.ModuleResponseMobile{
	
	var moduleRes []web.ModuleResponseMobile

	for i := range module {
		moduleRes = append(moduleRes, *ConvertModuleResponseMobile(&module[i]))
	}

	return moduleRes
}
