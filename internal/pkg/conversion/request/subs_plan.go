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

func SubsPlanUpdateRequestToSubsPlanDomain(subsPlan *web.SubsPlanUpdateRequest) *domain.SubsPlan {
	return &domain.SubsPlan{
		Title:         subsPlan.Title,
		Duration_days: subsPlan.Duration_days,
		Description:   subsPlan.Description,
		Price:         subsPlan.Price,
	}
}

func SubsPlanUpdateImageToSubsPlanDomain(subsPlan *web.SubsPlanUpdateImage) *domain.SubsPlan {
	return &domain.SubsPlan{
		Image_url: subsPlan.Image_url,
	}
}
