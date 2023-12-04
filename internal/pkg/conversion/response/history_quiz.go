package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertQuizResponseTrackingMobile(responseQuiz *domain.Quizzes, responseHistoryQuiz *domain.HistoryQuiz, complete bool) *web.QuizResponseForTracking {
	return &web.QuizResponseForTracking{
		ID:            responseQuiz.ID,
		Title:         responseQuiz.Title,
		HistoryQuizID: responseHistoryQuiz.ID,
		IsComplete:    complete,
	}
}

func ConvertHistoryQuizForTracking(quizz *domain.HistoryQuiz) *web.HistoryQuizTracking {
	return &web.HistoryQuizTracking{
		HistoryQuizID: quizz.ID,
		IsComplete:    quizz.IsComplete,
	}
}

func ConvertHistoryQuizTrackingResponse(module *web.ModuleResponseForTracking, history *web.HistoryQuizTracking, quizz *web.QuizResponse) *web.HistoryQuizIDTracking {
	return &web.HistoryQuizIDTracking{
		ID:          module.ID,
		Title:       module.Title,
		Description: module.Description,
		Progress:    module.Progress,
		Order:       module.Order,
		HistoryQuiz: *history,
		Quizz:       *quizz,
	}

}

// func ConvertHistoryQuizResponseMobile(response *domain.HistoryQuiz) *web.HistoryQuizResponseMobile {

// 	quiz:= ConvertQuizResponseTrackingMobile(&response.Quiz)

// 	return &web.HistoryQuizResponseMobile{
// 		ID:          response.ID,
// 		Quiz:   *quiz,
// 		IsComplete:  response.IsComplete,
// 	}
// }

// func ConvertAllHistoryQuizResponseMobile(response []domain.HistoryQuiz) []web.HistoryQuizResponseMobile {

// 	var historyQuiz []web.HistoryQuizResponseMobile

// 	for i := range response {
// 		historyQuiz = append(historyQuiz, *ConvertHistoryQuizResponseMobile(&response[i]))
// 	}

// 	return historyQuiz

// }
