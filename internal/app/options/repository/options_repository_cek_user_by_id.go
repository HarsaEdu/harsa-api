package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (repositoryOptions *OptionsRepositoryImpl) CekIdFromOption(userId uint, optionId uint, role string) (*domain.Options, error) {
	tx := repositoryOptions.DB.Begin()

	var option domain.Options
	if err := tx.First(&option, optionId).Error; err != nil {
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

	var userID, quizID, moduleID, sectionID, courseID uint

	if err := getScannedValue(&domain.Questions{}, "id = ?", option.QuestionID, &quizID); err != nil {
		return nil, err
	}

	if err := getScannedValue(&domain.Quizzes{}, "id = ?", quizID, &moduleID); err != nil {
		return nil, err
	}

	if err := getScannedValue(&domain.Module{}, "id = ?", moduleID, &sectionID); err != nil {
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

	return &option, nil
}