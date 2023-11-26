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
		Section:       module.Section,
		Title:         module.Title,
		Description:   module.Description,
		Order:         module.Order,
		Type:          module.Type,
		CreatedAt:     module.CreatedAt,
		UpdatedAt:     module.UpdatedAt,
		SubModules:    subModules,
		Submissions:   submissions,
		Quizzes:       quizzes,

	}
	return &moduleRes
}

func ConvertAllModuleResponse(module []domain.Module) []web.ModuleResponse{
	
	var moduleRes []web.ModuleResponse

	for i := range module {
		moduleRes = append(moduleRes, *ConvertModuleResponse(&module[i]))
	}

	return moduleRes
}
