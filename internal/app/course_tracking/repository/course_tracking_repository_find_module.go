package repository

import (

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindAllModuleTracking(module []domain.Module, userID uint) ([]web.ModuleResponseForTracking, error) {

	var allModule []web.ModuleResponseForTracking

	for _, module := range module {

		progrees, err := courseTrackingRepository.CountProgressModule(module.ID, userID)
		if err != nil {
			return nil, err
		}

		convertModul := conversion.ConvertModuleResponseTrackingMobile(&module, progrees)		
		allModule = append(allModule, *convertModul)
	}

    return allModule, nil
}

func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindModuleTracking(moduleID uint, userID uint) (*web.ModuleResponseForTracking, error) {

	var module domain.Module
	if err := courseTrackingRepository.DB.First(&module, moduleID).Error; err != nil {
		return nil, err
	
	}

	progrees, err := courseTrackingRepository.CountProgressModule(module.ID, userID)
	if err != nil {
		return nil, err
	}

	convertModul := conversion.ConvertModuleResponseTrackingMobile(&module, progrees)		
	

    return convertModul, nil
}

func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindAllSubModule(moduleId uint, userID uint) ([]web.SubModuleResponseForTracking, error) {
    var subModule []domain.SubModule

    if err := courseTrackingRepository.DB.
        Where("module_id = ?", moduleId).
        Find(&subModule).
        Error; err != nil {
        return nil, err
		
    }

	var allSubModule []web.SubModuleResponseForTracking

	for _, subModule := range subModule {

		var countHistorySubModul int64

		if err := courseTrackingRepository.DB.Model(&domain.HistorySubModule{}).Where("sub_module_id = ? AND user_id  = ?", subModule.ID, userID).Count(&countHistorySubModul).Error; err != nil {
			return nil, err
		}

		var convertSubModul *web.SubModuleResponseForTracking
		if countHistorySubModul > 0 {
			convertSubModul = conversion.ConvertSubModuleResponseTrackingMobile(&subModule, true)
		}else{
			convertSubModul = conversion.ConvertSubModuleResponseTrackingMobile(&subModule, false)
		}
		
		allSubModule = append(allSubModule, *convertSubModul)
	}

    return allSubModule, nil
}

func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindAllSubmission(moduleId uint, userID uint) ([]web.SubmissionsResponseModuleMobile, error) {
    var submission []domain.Submissions

    if err := courseTrackingRepository.DB.
        Where("module_id = ?", moduleId).
        Find(&submission).
        Error; err != nil {
        return nil, err
    }

	var allSubmissisonAnswer []web.SubmissionsResponseModuleMobile

	for _, submission := range submission {

		var countSubmissisonAnswer int64

		submissionAnswer := &domain.SubmissionAnswer{}

		if err := courseTrackingRepository.DB.Find(&submissionAnswer).Where("submission_id = ? AND user_id  = ?", submission.ID, userID).Count(&countSubmissisonAnswer).Error; err != nil {
			return nil, err
		}

		var convertSubmissisonAnswer *web.SubmissionsResponseModuleMobile
		if countSubmissisonAnswer > 0 {
			convertSubmissisonAnswer = conversion.ConvertSubmissionAnswerResponseTrackingMobile(&submission, submissionAnswer, true)
		}else{
			convertSubmissisonAnswer = conversion.ConvertSubmissionAnswerResponseTrackingMobile(&submission, submissionAnswer,false)
		}
		
		allSubmissisonAnswer = append(allSubmissisonAnswer, *convertSubmissisonAnswer)
	}

    return allSubmissisonAnswer, nil
}

func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindAllQuiz(moduleId uint, userID uint) ([]web.QuizResponseForTracking, error) {
    var quizzes []domain.Quizzes

    if err := courseTrackingRepository.DB.
        Where("module_id = ?", moduleId).
        Find(&quizzes).
        Error; err != nil {		
        return nil, err
    }

	var allHistoryQuiz []web.QuizResponseForTracking

	for _, quiz := range quizzes {

		var countHistoryQuiz int64

		historyQuiz := &domain.HistoryQuiz{}

		if err := courseTrackingRepository.DB.Find(&historyQuiz).Where("quiz_id = ? AND user_id  = ?", quiz.ID, userID).Count(&countHistoryQuiz).Error; err != nil {
			return nil, err
		}

		var convertHistoryQuiz *web.QuizResponseForTracking
		if countHistoryQuiz > 0 {
			convertHistoryQuiz = conversion.ConvertQuizResponseTrackingMobile(&quiz, historyQuiz, true)
		}else{
			convertHistoryQuiz = conversion.ConvertQuizResponseTrackingMobile(&quiz, historyQuiz, false)
		}
		
		allHistoryQuiz = append(allHistoryQuiz, *convertHistoryQuiz)
	}

    return allHistoryQuiz, nil
}


func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindAllSub(moduleId uint, userID uint) (*web.CourseTrackingSub, error) {
    
	subModules ,err:= courseTrackingRepository.FindAllSubModule(moduleId,userID)
	if err != nil {
		return nil, err
	
	}

	submissions ,err:= courseTrackingRepository.FindAllSubmission(moduleId,userID)
	if err != nil {
		return nil, err
	
	}
    
	quizzes ,err:= courseTrackingRepository.FindAllQuiz(moduleId,userID)
	if err != nil {
		return nil, err
	
	}

	res := conversion.ConvertAllSubInModule(subModules, submissions, quizzes)

    return res, nil
}



