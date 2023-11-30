package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	
)

func (courseTrackingrepository *CourseTrackingRepositoryImpl) CountProgressModule(moduleID uint, userID uint) (float32 ,error) { 

	var subModulesID []uint

	if err := courseTrackingrepository.DB.Model(&domain.SubModule{}).Where("module_id = ?", moduleID).Select("id").Scan(&subModulesID).Error; err != nil {
		return 0, err
	}
	
	var countHistorySubModul int64

	if err := courseTrackingrepository.DB.Model(&domain.HistorySubModule{}).Where("sub_module_id IN (?) AND user_id  = ?", subModulesID, userID).Count(&countHistorySubModul).Error; err != nil {
		return 0, err
	}

	var submissionsID []uint

	if err := courseTrackingrepository.DB.Model(&domain.Submissions{}).Where("module_id = ?", moduleID).Select("id").Scan(&submissionsID).Error; err != nil {
		return 0, err
	}
	
	var countSubmissisonAnswer int64

	if err := courseTrackingrepository.DB.Model(&domain.SubmissionAnswer{}).Where("submission_id IN (?) AND user_id = ? and status = 'accepted'", submissionsID, userID).Count(&countSubmissisonAnswer).Error; err != nil {
		return 0, err
	}

	var quizzesID []uint

	if err := courseTrackingrepository.DB.Model(&domain.Quizzes{}).Where("module_id = ?", moduleID).Select("id").Scan(&quizzesID).Error; err != nil {
		return 0, err
	}
	
	var countHistoryQuiz int64

	if err := courseTrackingrepository.DB.Model(&domain.HistoryQuiz{}).Where("quiz_id IN (?) AND user_id  = ?", quizzesID, userID).Count(&countHistoryQuiz).Error; err != nil {
		return 0, err
	}

	var progressModule float32

	totalTasksAndQuizzes := float32(len(submissionsID) + len(quizzesID) + len(submissionsID))
	if totalTasksAndQuizzes > 0 {
		progressModule = float32(countHistorySubModul+countSubmissisonAnswer+countHistoryQuiz) / totalTasksAndQuizzes * 100
		return progressModule , nil
	}

	return 0 , nil
}


func (courseTrackingrepository *CourseTrackingRepositoryImpl) CountProgressCourse(courseID uint, userID uint) (float32 ,error) { 

	var modulesID []uint

	if err := courseTrackingrepository.DB.Model(&domain.Course{}).Where("id = ?", courseID).Select("id").Scan(&modulesID).Error; err != nil {
		return 0, err
	}

	var totalProgressModule float32
	
	for _, moduleID := range modulesID {

		progressModul, err := courseTrackingrepository.CountProgressModule(moduleID, userID)
		if err != nil {
			return 0, err
		}
		totalProgressModule += progressModul
	}

	totalModules := float32(len(modulesID))
	if totalModules > 0 {
		averageProgress := totalProgressModule / totalModules
		return averageProgress, nil
	}

	return 0, nil
}