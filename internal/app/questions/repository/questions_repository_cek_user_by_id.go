package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (repositoryQuestions *QuestionsRepositoryImpl) CekIdFromQuestion(userId uint, questionId uint, role string) (*domain.Questions, error) {

	var question = domain.Questions{}

	if err := repositoryQuestions.DB.First(&question, questionId).Error; err != nil {
		return nil, err
	}
	
	var quiz = domain.Quizzes{}

	if err := repositoryQuestions.DB.First(&quiz, question.QuizID).Error; err != nil {
		return nil, err
	}

	var module = domain.Module{}

	if err := repositoryQuestions.DB.First(&module, quiz.ModuleID).Error; err != nil {
		return nil, err
	}

	var course = domain.Course{}

	if err := repositoryQuestions.DB.First(&course, module.CourseID).Error; err != nil {
		return nil, err
	}

	if course.UserID != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	

	return &question, nil
}