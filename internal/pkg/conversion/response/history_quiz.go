package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertQuizResponseTrackingMobile(response *domain.Quizzes) *web.QuizResponseForTracking {
	return &web.QuizResponseForTracking{
		ID:          response.ID,
		Title:       response.Title,
	}
}

func ConvertHistoryQuizResponseMobile(response *domain.HistoryQuiz) *web.HistoryQuizResponseMobile {
	
	quiz:= ConvertQuizResponseTrackingMobile(&response.Quiz)
	
	return &web.HistoryQuizResponseMobile{
		ID:          response.ID,
		Quiz:   *quiz,
		IsComplete:  response.IsComplete,
	}
}

func ConvertAllHistoryQuizResponseMobile(response []domain.HistoryQuiz) []web.HistoryQuizResponseMobile {
	
	var historyQuiz []web.HistoryQuizResponseMobile

	for i := range response {
		historyQuiz = append(historyQuiz, *ConvertHistoryQuizResponseMobile(&response[i]))
	}

	return historyQuiz
		
}