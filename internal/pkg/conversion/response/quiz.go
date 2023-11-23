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
		ModuleID:    quiz.ModuleID,
		Durations:   quiz.Durations,
		Title:       quiz.Title,
		Description: quiz.Description,
		Questions: question,
	}
	return &quizRes
}

func ConvertQuizResMobile(quiz *domain.Quizzes) *web.QuizResponse {

	question:= ConvertAllQuestionsQuizMobile(quiz.Questions)

	quizRes := web.QuizResponse{
		Id:          quiz.ID,
		Created_at:  quiz.CreatedAt,
		Updatad_at:  quiz.UpdatedAt,
		ModuleID:    quiz.ModuleID,
		Durations:   quiz.Durations,
		Title:       quiz.Title,
		Description: quiz.Description,
		Questions: question,
	}
	return &quizRes
}

func ConvertAllQuizRes(quiz *domain.Quizzes) *web.GetAllQuizResponse {

	numberQuestions := len(quiz.Questions)

	quizRes := web.GetAllQuizResponse{
		Id:              quiz.ID,
		ModuleID:        quiz.ModuleID,
		Title:           quiz.Title,
		Description:     quiz.Description,
		Durations:       quiz.Durations,
		NumberQuestions: numberQuestions,
	}
	return &quizRes
}


func ConvertAllQuiz(Quizzes []domain.Quizzes) []web.GetAllQuizResponse {

	var quizRes []web.GetAllQuizResponse

	for i := range Quizzes {
		quizRes = append(quizRes, *ConvertAllQuizRes(&Quizzes[i]))
	}

	return quizRes
}