package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (SubmissionRepository *SubmissionRepositoryImpl) CekUserIDfromModuleID(id uint, userId uint, role string) error {

	var userID uint

	if err := SubmissionRepository.DB.Model(&domain.Module{}).Where("modules.id = ?", id).Select("courses.user_id").
	Joins("JOIN sections ON sections.id = modules.section_id").
	Joins("JOIN courses ON courses.id = sections.course_id").
	Scan(&userID).Error; err != nil {
		return err
	}

	if userID != userId && role != "admin" {
		return fmt.Errorf("unauthorized")
	}

	return nil
}

func (SubmissionRepository *SubmissionRepositoryImpl) CekUserIDfromSubmission(id uint, userId uint, role string)  error {

	var userID uint

	if err := SubmissionRepository.DB.Model(&domain.SubModule{}).Where("id = ?", id).Select("courses.user_id").
		Joins("JOIN modules ON modules.id = sub_modules.module_id").
		Joins("JOIN sections ON sections.id = modules.section_id").
		Joins("JOIN courses ON courses.id = sections.course_id").
		Scan(&userID).Error; err != nil {
		return  err
	}

	if userID != userId && role != "admin" {
		return fmt.Errorf("unauthorized")
	}

	return  nil
}