package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertSubmissionAnswerRequestToSubmissionAnswerDomain(request *web.SubmissionAnswerRequest) *domain.SubmissionAnswer {
	return &domain.SubmissionAnswer{SubmittedUrl: request.SubmittedUrl, Status: request.Status}
}

func ConvertSubmissionAnswerUpdateToSubmissionAnswerDomain(request *web.SubmissionAnswerUpdateRequest) *domain.SubmissionAnswer {
	return &domain.SubmissionAnswer{SubmittedUrl: request.SubmittedUrl}
}
