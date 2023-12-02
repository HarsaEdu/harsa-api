package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	
)

func (courseTrackingrepository *CourseTrackingRepositoryImpl) CountProgressModule(moduleID uint, userID uint) (float32, error) {
	tx := courseTrackingrepository.DB.Begin()

	var progressModule float32

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	countItems := func(model interface{}, condition string, values []interface{}) (int64, error) {
		var count int64
		if err := tx.Model(model).Where(condition, values...).Count(&count).Error; err != nil {
			return 0, err
		}
		return count, nil
	}

	subModulesID, submissionsID, quizzesID := []uint{}, []uint{}, []uint{}

	if err := tx.Model(&domain.SubModule{}).Where("module_id = ?", moduleID).Pluck("id", &subModulesID).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Model(&domain.Submissions{}).Where("module_id = ?", moduleID).Pluck("id", &submissionsID).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Model(&domain.Quizzes{}).Where("module_id = ?", moduleID).Pluck("id", &quizzesID).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	countHistorySubModul, err := countItems(&domain.HistorySubModule{}, "sub_module_id IN (?) AND user_id = ?", []interface{}{subModulesID, userID})
	if err != nil {
		return 0, err
	}

	countSubmissionAnswer, err := countItems(&domain.SubmissionAnswer{}, "submission_id IN (?) AND user_id = ? AND status = 'accepted'", []interface{}{submissionsID, userID})
	if err != nil {
		return 0, err
	}

	countHistoryQuiz, err := countItems(&domain.HistoryQuiz{}, "quiz_id IN (?) AND user_id = ?", []interface{}{quizzesID, userID})
	if err != nil {
		return 0, err
	}

	totalTasksAndQuizzes := float32(len(submissionsID) + len(quizzesID) + len(submissionsID))
	if totalTasksAndQuizzes > 0 {
		progressModule = float32(countHistorySubModul + countSubmissionAnswer + countHistoryQuiz) / totalTasksAndQuizzes * 100
	}

	tx.Commit()
	return progressModule, nil
}


func (courseTrackingrepository *CourseTrackingRepositoryImpl) CountProgressCourse(courseID uint, userID uint) (float32, error) {
	var sectionIDs []uint

	if err := courseTrackingrepository.DB.Model(&domain.Section{}).Where("course_id = ?", courseID).Pluck("id", &sectionIDs).Error; err != nil {
		return 0, err
	}

	var totalModules float32
	var totalProgressModule float32

	var modulesID []uint

	for _, sectionID := range sectionIDs {
		if err := courseTrackingrepository.DB.Model(&domain.Module{}).Where("section_id = ?", sectionID).Pluck("id", &modulesID).Error; err != nil {
			return 0, err
		}

		for _, moduleID := range modulesID {
			progressModule, err := courseTrackingrepository.CountProgressModule(moduleID, userID)
			if err != nil {
				return 0, err
			}
			totalProgressModule += progressModule
		}

		totalModules += float32(len(modulesID))
	}

	if totalModules > 0 {
		averageProgress := totalProgressModule / totalModules
		return averageProgress, nil
	}

	return 0, nil
}