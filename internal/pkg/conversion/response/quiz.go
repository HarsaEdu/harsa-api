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

func ConvertAllQuizRes(quiz *domain.Quizzes, courseTitle string) *web.GetAllQuizResponse {

	numberQuestions := len(quiz.Questions)

	quizRes := web.GetAllQuizResponse{
		Id:              quiz.ID,
		ModuleID:        quiz.ModuleID,
		Title:           quiz.Title,
		CourseTitle:     courseTitle,
		Description:     quiz.Description,
		Durations:       quiz.Durations,
		NumberQuestions: numberQuestions,
	}
	return &quizRes
}


func ConvertAllQuiz(Quizzes []domain.Quizzes, courseTitle string) []web.GetAllQuizResponse {

	var quizRes []web.GetAllQuizResponse

	for i := range Quizzes {
		quizRes = append(quizRes, *ConvertAllQuizRes(&Quizzes[i], courseTitle))
	}

	return quizRes
}

func ConvertQuizResponseModule(quiz *domain.Quizzes) *web.QuizResponseModule{
	quizRes := web.QuizResponseModule{
		Id:          quiz.ID,
		Durations:   quiz.Durations,
		Title:       quiz.Title,
		Description: quiz.Description,
	}
	return &quizRes
}

func ConvertAllQuizModule(Quizzes []domain.Quizzes) []web.QuizResponseModule {

	var quizRes []web.QuizResponseModule

	for i := range Quizzes {
		quizRes = append(quizRes, *ConvertQuizResponseModule(&Quizzes[i]))
	}

	return quizRes
}

func ConvertQuizResponseModuleMobile(quiz *domain.Quizzes) *web.QuizResponseModuleMobile{
	quizRes := web.QuizResponseModuleMobile{
		Id:          quiz.ID,
		Title:       quiz.Title,
	}
	return &quizRes
}

func ConvertAllQuizModuleMobile(Quizzes []domain.Quizzes) []web.QuizResponseModuleMobile {

	var quizRes []web.QuizResponseModuleMobile

	for i := range Quizzes {
		quizRes = append(quizRes, *ConvertQuizResponseModuleMobile(&Quizzes[i]))
	}

	return quizRes
}