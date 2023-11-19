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

	var module = domain.Module{}

	if err := repository.DB.First(&module, quiz.ModuleId).Error; err != nil {
		return nil, err
	}

	var course = domain.Course{}

	if err := repository.DB.First(&course, module.CourseID).Error; err != nil {
		return nil, err
	}

	if course.UserID != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	

	return &quiz, nil
}

func (repository *QuizzesRepositoryImpl) CekIdFromModule(userId uint, moduleId uint, role string) error {


	var module = domain.Module{}

	if err := repository.DB.First(&module, moduleId).Error; err != nil {
		return err
	}

	var course = domain.Course{}

	if err := repository.DB.First(&course, module.CourseID).Error; err != nil {
		return err
	}

	if course.UserID != userId && role != "admin" {
		return fmt.Errorf("unauthorized")
	}
	

	return nil
}
