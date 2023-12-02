package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertSubmissionAnswerResponseTrackingMobile(responseSubmission *domain.Submissions, responseAnswer *domain.SubmissionAnswer, complete bool) *web.SubmissionsResponseModuleMobile {
	return &web.SubmissionsResponseModuleMobile{
		Id:    responseSubmission.ID,
		Title: responseSubmission.Title,
		SubmissionAnswer: web.SubmissionAnswerRes{
			Id:     responseAnswer.ID,
			Status: responseAnswer.Status,
		},
		Is_complete: complete,
	}
}

func ConvertSubmissionAnswerTracking(response *domain.SubmissionAnswer, title string) *web.SubmissionsResponseModuleMobile {
	var completed bool = false
	if response.Status == "accepted" {
		completed = true
	}
	return &web.SubmissionsResponseModuleMobile{
		Id:    response.ID,
		Title: title,
		SubmissionAnswer: web.SubmissionAnswerRes{
			Id:     response.ID,
			Status: response.Status,
		},
		Is_complete: completed,
	}
}

func ConvertSubmissionAnswerTrackingResponse(module *web.ModuleResponseForTracking, answer *web.SubmissionsResponseModuleMobile, submission *web.SubmissionsResponseModule) *web.SubmissionAnswerTrackingResponse {
	return &web.SubmissionAnswerTrackingResponse{
		ID:               module.ID,
		Title:            module.Title,
		Description:      module.Description,
		Progress:         module.Progress,
		Order:            module.Order,
		SubmissionAnswer: *answer,
		Submission:       *submission,
	}
}

// func ConvertAllSubmissionAnswerResponseMobile(response []domain.SubmissionAnswer) []web.SubmissionAnswerResponseMobile {

// 	var historySubmissionAnswer []web.SubmissionAnswerResponseMobile

// 	for i := range response {
// 		historySubmissionAnswer = append(historySubmissionAnswer, *ConvertSubmissionAnswerResponseMobile(&response[i]))
// 	}

// 	return historySubmissionAnswer

// }
