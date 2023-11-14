package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func QuizCreateRequestToQuizzesModel(request web.QuizRequest) *domain.Quizzes {
	return &domain.Quizzes{
		Module_id: request.Module_id,
		User_id: request.User_id,
		Title: request.Title,
		Description: request.Description,
		Durations: request.Durations,
		Questions: request.Questions,
	}
}