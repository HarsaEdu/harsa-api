package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (submissionRepository *SubmissionRepositoryImpl) GetAll(moduleId int) ([]domain.Submissions, int64, error) {

	submission := []domain.Submissions{}
	var total int64

	query := submissionRepository.DB.Where("module_id=?", moduleId).Find(&submission)
	if query.Error != nil {
		return nil, 0, query.Error
	}

	query = query.Find(&submission).Count(&total)

	return submission, total, nil

}

func (submissionRepository *SubmissionRepositoryImpl) GetAllWeb(moduleId int) ([]web.SubmissionsResponseWeb, int64, error) {

	submission := []web.SubmissionsResponseWeb{}
	var total int64

	query := submissionRepository.DB.Model(&domain.Submissions{}).Where("submissions.module_id=?", moduleId).
	Select("submissions.id as id, submissions.title as submission_title, submissions.content as content, courses.title as course_title ").
	Joins("JOIN modules ON modules.id = submissions.module_id").
	Joins("JOIN sections ON sections.id = modules.section_id").
	Joins("JOIN courses ON courses.id = sections.course_id").
	Find(&submission)
	if query.Error != nil {
		return nil, 0, query.Error
	}

	query = query.Find(&submission).Count(&total)

	return submission, total, nil

}

