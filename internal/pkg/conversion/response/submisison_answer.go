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

func ConvertSubissionAnswerUser(answer *domain.SubmissionsAnswerDetail)*web.SubmissionAnswerList{
	
	name := answer.FirstName + " " + answer.LastName
	
	return &web.SubmissionAnswerList{
		ID: answer.ID,
		Name: name,
		Status: answer.Status,
	}

}

func ConvertAllSubissionAnswerUser(answer []domain.SubmissionsAnswerDetail)[]web.SubmissionAnswerList{
	var response []web.SubmissionAnswerList
	for i := range answer {
		response = append(response, *ConvertSubissionAnswerUser(&answer[i]))
	}
	return response

}

 func ConverstSubmissionAnswerToResponseWeb(answer []domain.SubmissionsAnswerDetail, submission *web.SubmissionsResponseModule) *web.SubmissionAnswerResponseWeb {
	
	answers:= ConvertAllSubissionAnswerUser(answer)
	
	return &web.SubmissionAnswerResponseWeb{
		Submission: *submission,
		SubmissionAnswer: answers,
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
