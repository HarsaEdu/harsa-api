package repository

import (

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (courseTrackingrepository *CourseTrackingRepositoryImpl) GetCourseIDbyModuleID(moduleId uint) (uint,error) {

	var courseID uint

	if err := courseTrackingrepository.DB.Model(&domain.Module{}).Where("modules.id = ?", moduleId).Select("courses.id").
	Joins("JOIN sections ON sections.id = modules.section_id").
	Joins("JOIN courses ON courses.id = sections.course_id").
	Scan(&courseID).Error; err != nil {
		return 0, err
	}

	return courseID, nil
}

func (courseTrackingrepository *CourseTrackingRepositoryImpl) GetCourseIDbySubmssionID(id uint) (uint,error) {

	var courseID uint

	if err := courseTrackingrepository.DB.Model(&domain.Submissions{}).Where("id = ?", id).Select("courses.id").
	Joins("JOIN modules ON modules.id = submissions.module_id").
	Joins("JOIN sections ON sections.id = modules.section_id").
	Joins("JOIN courses ON courses.id = sections.course_id").
	Scan(&courseID).Error; err != nil {
		return 0, err
	}

	return courseID, nil
}

func (courseTrackingrepository *CourseTrackingRepositoryImpl) GetCourseIDbySubModuleID(id uint) (uint,error) {

	var courseID uint

	if err := courseTrackingrepository.DB.Model(&domain.SubModule{}).Where("id = ?", id).Select("courses.id").
	Joins("JOIN modules ON modules.id = sub_modules.module_id").
	Joins("JOIN sections ON sections.id = modules.section_id").
	Joins("JOIN courses ON courses.id = sections.course_id").
	Scan(&courseID).Error; err != nil {
		return 0, err
	}

	return courseID, nil
}

func (courseTrackingrepository *CourseTrackingRepositoryImpl) GetCourseIDbyQuizzesID(id uint) (uint,error) {

	var courseID uint

	if err := courseTrackingrepository.DB.Model(&domain.Quizzes{}).Where("id = ?", id).Select("courses.id").
	Joins("JOIN modules ON modules.id = quizzes.module_id").
	Joins("JOIN sections ON sections.id = modules.section_id").
	Joins("JOIN courses ON courses.id = sections.course_id").
	Scan(&courseID).Error; err != nil {
		return 0, err
	}

	return courseID, nil
}

