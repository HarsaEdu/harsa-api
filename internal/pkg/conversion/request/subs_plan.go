package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func SubsPlanRequestToSubsPlanDomain(subsPlan *web.SubsPlanCreateRequest) *domain.SubsPlan {
	return &domain.SubsPlan{
		Title:         subsPlan.Title,
		Image_url:     subsPlan.Image_url,
		Duration_days: subsPlan.Duration_days,
		Description:   subsPlan.Description,
		Price:         subsPlan.Price,
	}
}

func ExistingSubsPlanToSubsPlanDomain(request *web.SubsPlanUpdateRequest, existingPlan *domain.SubsPlan) *domain.SubsPlan {
	if request.Title == "" {
		request.Title = existingPlan.Title
	}

	if request.Description == "" {
		request.Description = existingPlan.Description
	}

	if request.Price == 0 {
		request.Price = existingPlan.Price
	
	}

	if request.Duration_days == 0 {
		request.Duration_days = existingPlan.Duration_days
	}

	return &domain.SubsPlan{
		Title:         request.Title,
		Description:   request.Description,
		Price:         request.Price,
		Duration_days: request.Duration_days,
		Image_url:     existingPlan.Image_url,
	}
}

func SubsPlanUpdateImageToSubsPlanDomain(subsPlan *web.SubsPlanUpdateImage) *domain.SubsPlan {
	return &domain.SubsPlan{
		Image_url: subsPlan.Image_url,
	}
}
