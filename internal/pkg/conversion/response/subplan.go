package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertSubPlanResponse(subPlan *domain.SubsPlan) *web.SubsPlanResposne {
	submissionRes := web.SubsPlanResposne{
		ID:        subPlan.ID,
		Title: subPlan.Title,
		Description: subPlan.Description,
		Price: subPlan.Price,
		Image_url: subPlan.Image_url,
		Duration_days: subPlan.Duration_days,
	}
	return &submissionRes
}

func ConvertAllSubPlanResponse(subPlans []domain.SubsPlan) []web.SubsPlanResposne {
	var submissionRes []web.SubsPlanResposne
	for _, subPlan := range subPlans {
		submissionRes = append(submissionRes, *ConvertSubPlanResponse(&subPlan))
	}
	return submissionRes
}