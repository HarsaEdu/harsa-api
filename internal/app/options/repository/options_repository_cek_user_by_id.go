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
	
	var quizID uint

	if err := repositoryOptions.DB.Model(&domain.Questions{}).Where("id = ?", option.QuestionID).Select("section_id").Scan(&quizID).Error; err != nil {
		return nil, err
	}
	
	var moduleID uint

	if err := repositoryOptions.DB.Model(&domain.Quizzes{}).Where("id = ?", quizID).Select("section_id").Scan(&moduleID).Error; err != nil {
		return nil, err
	}

	var sectionID uint

	if err := repositoryOptions.DB.Model(&domain.Module{}).Where("id = ?", moduleID).Select("section_id").Scan(&sectionID).Error; err != nil {
		return nil, err
	}

	var courseID uint

	if err := repositoryOptions.DB.Model(&domain.Section{}).Where("id = ?", sectionID).Select("course_id").Scan(&courseID).Error; err != nil {
		return nil, err
	}
	
	var userID uint

	if err := repositoryOptions.DB.Model(&domain.Course{}).Where("id = ?", courseID).Select("user_id").Scan(&userID).Error; err != nil {
		return nil, err
	}

	if userID != userId && role != "admin" {
		return nil, fmt.Errorf("unauthorized")
	}
	

	return &option, nil
}