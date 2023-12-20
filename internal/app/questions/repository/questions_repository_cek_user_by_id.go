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
	
	var moduleID uint

	if err := repositoryQuestions.DB.Model(&domain.Quizzes{}).Where("id = ?", question.QuizID).Select("module_id").Scan(&moduleID).Error; err != nil {
		return nil, err
	}

	var sectionID uint

	if err := repositoryQuestions.DB.Model(&domain.Module{}).Where("id = ?", moduleID).Select("section_id").Scan(&sectionID).Error; err != nil {
		return nil, err
	}

	var courseID uint

	if err := repositoryQuestions.DB.Model(&domain.Section{}).Where("id = ?", sectionID).Select("course_id").Scan(&courseID).Error; err != nil {
		return nil, err
	}
	
	var userID uint

	if err := repositoryQuestions.DB.Model(&domain.Course{}).Where("id = ?", courseID).Select("user_id").Scan(&userID).Error; err != nil {
		return nil, err
	}

	if userID != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	

	return &question, nil
}
