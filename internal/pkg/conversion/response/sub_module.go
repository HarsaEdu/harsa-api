package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertSubModuleResponseModule(subModule *domain.SubModule) *web.SubModuleResponseModule{
	subModuleRes := web.SubModuleResponseModule{
		ID:            subModule.ID,
		Title:         subModule.Title,
		ContentUrl:    subModule.ContentUrl,
		ContentBody:   subModule.ContentBody,
		Type:          subModule.Type,
	}
	return &subModuleRes
}

func ConvertAllSubModuleModule(subModule []*domain.SubModule) []web.SubModuleResponseModule {

	var subModuleRes []web.SubModuleResponseModule

	for i := range subModule {
		subModuleRes = append(subModuleRes, *ConvertSubModuleResponseModule(subModule[i]))
	}

	return subModuleRes
}