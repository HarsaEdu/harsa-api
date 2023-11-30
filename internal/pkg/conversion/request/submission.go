package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func SubmissionRequestToSubmissionDomain(request *web.SubmissionRequest) *domain.Submissions {
	return &domain.Submissions{
		Title:       request.Title,
		Description: request.Description,
		ModuleID:    request.ModuleID,
	}
}

func SubmissionUpdateRequestToSubmissionDomain(request *web.SubmissionUpdateRequest) *domain.Submissions {
	return &domain.Submissions{
		Title:       request.Title,
		Description: request.Description,
	}
}
