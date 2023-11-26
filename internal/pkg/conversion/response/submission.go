package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertSubmissionResponseModule(submission *domain.Submissions) *web.SubmissionsResponseModule{
	submissionRes := web.SubmissionsResponseModule{
		Id:            submission.ID,
		Title:         submission.Title,
		Description:   submission.Description,
	}
	return &submissionRes
}

func ConvertAllSubmissionModule(submission []domain.Submissions) []web.SubmissionsResponseModule {

	var submissionRes []web.SubmissionsResponseModule

	for i := range submission {
		submissionRes = append(submissionRes, *ConvertSubmissionResponseModule(&submission[i]))
	}

	return submissionRes
}