package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (repository *QuizzesRepositoryImpl) CekIdFromQuiz(userId uint, quizId uint, role string) (*domain.Quizzes, error) {
	tx := repository.DB.Begin()

	var quiz domain.Quizzes
	if err := tx.First(&quiz, quizId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	getScannedValue := func(model interface{}, condition string, value interface{}, dest interface{}) error {
		if err := tx.Model(model).Where(condition, value).Select(dest).Scan(dest).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}

	var sectionID, courseID, userID uint

	if err := getScannedValue(&domain.Module{}, "id = ?", quiz.ModuleID, &sectionID); err != nil {
		return nil, err
	}

	if err := getScannedValue(&domain.Section{}, "id = ?", sectionID, &courseID); err != nil {
		return nil, err
	}

	if err := getScannedValue(&domain.Course{}, "id = ?", courseID, &userID); err != nil {
		return nil, err
	}

	tx.Commit()

	if userID != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}

	return &quiz, nil
}


func (repository *QuizzesRepositoryImpl) CekIdFromModule(userId uint, moduleId uint, role string) error {
	tx := repository.DB.Begin()

	var sectionID, courseID, userID uint

	if err := repository.DB.Model(&domain.Module{}).Where("id = ?", moduleId).Select("section_id").Scan(&sectionID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := repository.DB.Model(&domain.Section{}).Where("id = ?", sectionID).Select("course_id").Scan(&courseID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := repository.DB.Model(&domain.Course{}).Where("id = ?", courseID).Select("user_id").Scan(&userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	if userID != userId && role != "admin" {
		return fmt.Errorf("unauthorized")
	}

	return nil
}
