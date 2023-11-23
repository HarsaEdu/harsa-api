package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (repositoryOptions *OptionsRepositoryImpl) CekIdFromOption(userId uint, optionId uint, role string) (*domain.Options, error) {

	var option = domain.Options{}

	if err := repositoryOptions.DB.First(&option, optionId).Error; err != nil {
		return nil, err
	}
	
	var question = domain.Questions{}

	if err := repositoryOptions.DB.First(&question, option.QuestionID).Error; err != nil {
		return nil, err
	}
	
	var quiz = domain.Quizzes{}

	if err := repositoryOptions.DB.First(&quiz, question.QuizID).Error; err != nil {
		return nil, err
	}

	var module = domain.Module{}

	if err := repositoryOptions.DB.First(&module, quiz.ModuleID).Error; err != nil {
		return nil, err
	}

	var course = domain.Course{}

	if err := repositoryOptions.DB.First(&course, module.CourseID).Error; err != nil {
		return nil, err
	}

	if course.UserID != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	

	return &option, nil
}