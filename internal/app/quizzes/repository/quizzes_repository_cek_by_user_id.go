package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (repository *QuizzesRepositoryImpl) CekIdFromQuiz(userId uint, quizId uint, role string) (*domain.Quizzes, error) {

	var quiz = domain.Quizzes{}

	if err := repository.DB.First(&quiz, quizId).Error; err != nil {
		return nil, err
	}

	var sectionID uint

	if err := repository.DB.Model(&domain.Module{}).Where("id = ?", quiz.ModuleID).Select("section_id").Scan(&sectionID).Error; err != nil {
		return nil, err
	}

	var courseID uint

	if err := repository.DB.Model(&domain.Section{}).Where("id = ?", sectionID).Select("course_id").Scan(&courseID).Error; err != nil {
		return nil, err
	}
	
	var userID uint

	if err := repository.DB.Model(&domain.Course{}).Where("id = ?", courseID).Select("user_id").Scan(&userID).Error; err != nil {
		return nil, err
	}

	if userID != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	

	return &quiz, nil
}

func (repository *QuizzesRepositoryImpl) CekIdFromModule(userId uint, moduleId uint, role string) error {

	var sectionID uint

	if err := repository.DB.Model(&domain.Module{}).Where("id = ?", moduleId).Select("section_id").Scan(&sectionID).Error; err != nil {
		return err
	}

	var courseID uint

	if err := repository.DB.Model(&domain.Section{}).Where("id = ?", sectionID).Select("course_id").Scan(&courseID).Error; err != nil {
		return err
	}
	
	var userID uint

	if err := repository.DB.Model(&domain.Course{}).Where("id = ?", courseID).Select("user_id").Scan(&userID).Error; err != nil {
		return err
	}

	if userID != userId && role != "admin" {
		return fmt.Errorf("unauthorized")
	}
	

	return nil
} 
