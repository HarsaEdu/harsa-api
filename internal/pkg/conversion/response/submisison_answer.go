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

func ConvertSubmissionAnswerResponseMobile(response *domain.SubmissionAnswer) *web.SubmissionAnswerResponseMobile {
	return &web.SubmissionAnswerResponseMobile{
		ID:         response.ID,
		Status:     response.Status,
		Feedback:   response.Feedback,
		Submission: response.SubmittedUrl,
	}
}

 func ConverstSubmissionAnswerToResponseWeb(answer []domain.SubmissionsAnswerDetail, submission *web.SubmissionsResponseModule) *web.SubmissionAnswerResponseWeb {
	return &web.SubmissionAnswerResponseWeb{
		Submission: *submission,
		SubmissionAnswer: answer,
	}


 }

// 	var response []*web.SubmissionAnswerResponseWeb

// 	for _, val := range responseSubmission {
// 		response = append(response, &web.SubmissionAnswerResponseWeb{
// 			Title:   val.Title,
// 			Content: val.Content,
// 			Peserta: &web.UserForCourseResponse{Name: val.Peserta},
// 			Answer:  &web.SubmissionAnswerResponseMobile{Submission: val.Content},
// 		},
// 		)
// 	}

// 	return response
// }

// func ConvertSubmissionAnswerResponseMobile(response *domain.SubmissionAnswer) *web.SubmissionAnswerResponseMobile {

// Submission:= ConvertSubmissionAnswerResponseTrackingMobile(&response.Submission)

// 	return &web.SubmissionAnswerResponseMobile{
// 		ID:          response.ID,
// 		Status: response.Status,
// 		Submission:   *Submission,
// 	}
// }

// func ConvertAllSubmissionAnswerResponseMobile(response []domain.SubmissionAnswer) []web.SubmissionAnswerResponseMobile {

// 	var historySubmissionAnswer []web.SubmissionAnswerResponseMobile

// 	for i := range response {
// 		historySubmissionAnswer = append(historySubmissionAnswer, *ConvertSubmissionAnswerResponseMobile(&response[i]))
// 	}

// 	return historySubmissionAnswer

// }
