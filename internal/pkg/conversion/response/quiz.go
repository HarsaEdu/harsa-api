package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertQuizRes(quiz *domain.Quizzes) *web.QuizResponse {

	question:= ConvertAllQuestionsQuiz(quiz.Questions)

	quizRes := web.QuizResponse{
		Id:          quiz.ID,
		Created_at:  quiz.CreatedAt,
		Updatad_at:  quiz.UpdatedAt,
		User_id:     quiz.User_id,
		Title:       quiz.Title,
		Description: quiz.Description,
		Questions: question,
	}
	return &quizRes
}

func ConvertAllQuiz(Quizzes []domain.Quizzes) []web.QuizResponse {

	var quizRes []web.QuizResponse

	for i := range Quizzes {
		quizRes = append(quizRes, *ConvertQuizRes(&Quizzes[i]))
	}

	return quizRes
}