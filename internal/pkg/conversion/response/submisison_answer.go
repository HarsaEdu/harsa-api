package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertSubmissionAnswerResponseTrackingMobile(response *domain.Submissions) *web.SubmissionsResponseModuleMobile{
	return &web.SubmissionsResponseModuleMobile{
		Id:          response.ID,
		Title:       response.Title,
	}
}

func ConvertSubmissionAnswerResponseMobile(response *domain.SubmissionAnswer) *web.SubmissionAnswerResponseMobile {
	
	Submission:= ConvertSubmissionAnswerResponseTrackingMobile(&response.Submission)
	
	return &web.SubmissionAnswerResponseMobile{
		ID:          response.ID,
		Submission:   *Submission,
		Status: response.Status,
	}
}

func ConvertAllSubmissionAnswerResponseMobile(response []domain.SubmissionAnswer) []web.SubmissionAnswerResponseMobile {
	
	var historySubmissionAnswer []web.SubmissionAnswerResponseMobile

	for i := range response {
		historySubmissionAnswer = append(historySubmissionAnswer, *ConvertSubmissionAnswerResponseMobile(&response[i]))
	}

	return historySubmissionAnswer
		
}