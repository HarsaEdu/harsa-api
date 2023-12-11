package repository

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindAllModuleTracking(sections []domain.Section, userID uint) ([]web.SectionResponseMobile, error) {

	allSection:= []web.SectionResponseMobile{}

	for _,section := range sections {

		var allModule []web.ModuleResponseForTracking
		for _, module := range section.Modules {

			progrees, err := courseTrackingRepository.CountProgressModule(module.ID, userID)
			if err != nil {
				return nil, err
			}

			convertModul := conversion.ConvertModuleResponseTrackingMobile(&module, progrees)		
			allModule = append(allModule, *convertModul)
		}
		sectionRes:= conversion.ConvertSectionResponseMobile(&section, allModule)
		allSection = append(allSection, *sectionRes)
	}

    return allSection, nil
}

func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindAllModuleTrackingWithProgress(sections []domain.Section, userID uint, courseID uint) ([]web.SectionResponseMobile, float32 ,error) {

	var status string

	if err := courseTrackingRepository.DB.Model(&domain.CourseTracking{}).Where("course_id = ? and user_id = ?", courseID, userID).Pluck("status", &status).Error; err != nil {
		return nil,0, err
	}


	allSection:= []web.SectionResponseMobile{}

	var totalProgress float32
	var totalModule float32

	for _,section := range sections {

		var allModule []web.ModuleResponseForTracking
		for _, module := range section.Modules {

			progrees, err := courseTrackingRepository.CountProgressModule(module.ID, userID)
			if err != nil {
				return nil,0, err
			}

			totalProgress += progrees

			convertModul := conversion.ConvertModuleResponseTrackingMobile(&module, progrees)		
			allModule = append(allModule, *convertModul)
		}
		totalModule += float32(len(section.Modules))
		sectionRes:= conversion.ConvertSectionResponseMobile(&section, allModule)
		allSection = append(allSection, *sectionRes)
	}

	var progress float32
	if totalModule == 0{
		progress = 0
	}else{
		progress = totalProgress / totalModule
	}
	if status == "in progress" && progress == 100{
		newStatus := "completed"
		result := courseTrackingRepository.DB.Model(&domain.CourseTracking{}).Where("course_id = ? and user_id = ?", courseID, userID).Update("status", newStatus)
		if result.Error != nil {
			return nil, 0,result.Error
		}
	}
    return allSection, progress, nil
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
		} else {
			convertSubModul = conversion.ConvertSubModuleResponseTrackingMobile(&subModule, false)
		}

		allSubModule = append(allSubModule, *convertSubModul)
	}

	return allSubModule, nil
}

func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindSubModuleByID(moduleID uint, userID uint, subModuleID uint) (*domain.HistorySubModule, *domain.SubModule, error) {
	historySubModule := domain.HistorySubModule{}
	err := courseTrackingRepository.DB.
		Where("sub_module_id = ? AND user_id = ?", subModuleID, userID).Preload("SubModule").
		Find(&historySubModule).Error
	if err != nil {
		return nil, nil, err
	}

	subModule := domain.SubModule{}
	err = courseTrackingRepository.DB.Where("module_id = ? AND id = ?", moduleID, subModuleID).First(&subModule).Error
	if err != nil {
		return nil, nil, err
	}
	return &historySubModule, &subModule, nil
}

func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindSubmissionByID(moduleID uint, userID uint, submissionID uint) (*domain.SubmissionAnswer, *domain.Submissions, error) {
	historySubmission := domain.SubmissionAnswer{}
	err := courseTrackingRepository.DB.
		Where("submission_id = ? AND user_id = ?", submissionID, userID).Preload("Submission").
		Find(&historySubmission).Error
	if err != nil {
		return nil, nil, err
	}

	submission := domain.Submissions{}
	err = courseTrackingRepository.DB.Where("module_id = ? AND id = ?", moduleID, submissionID).First(&submission).Error
	if err != nil {
		return nil, nil, err
	}
	return &historySubmission, &submission, nil
}

func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindQuizzByID(moduleID uint, userID uint, quizID uint) (*domain.HistoryQuiz, error) {
	historyQuizz := domain.HistoryQuiz{}
	err := courseTrackingRepository.DB.
		Where("quiz_id = ? AND user_id = ?", quizID, userID).Preload("HistoryQuizAnswer").
		Preload("Quiz").
		Find(&historyQuizz).Error
	if err != nil {
		return nil, err
	}
	return &historyQuizz, nil
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
		} else {
			convertSubmissisonAnswer = conversion.ConvertSubmissionAnswerResponseTrackingMobile(&submission, submissionAnswer, false)
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
		} else {
			convertHistoryQuiz = conversion.ConvertQuizResponseTrackingMobile(&quiz, historyQuiz, false)
		}

		allHistoryQuiz = append(allHistoryQuiz, *convertHistoryQuiz)
	}

	return allHistoryQuiz, nil
}

func (courseTrackingRepository *CourseTrackingRepositoryImpl) FindAllSub(moduleId uint, userID uint) (*web.CourseTrackingSub, error) {

	subModules, err := courseTrackingRepository.FindAllSubModule(moduleId, userID)
	if err != nil {
		return nil, err

	}

	submissions, err := courseTrackingRepository.FindAllSubmission(moduleId, userID)
	if err != nil {
		return nil, err

	}

	quizzes, err := courseTrackingRepository.FindAllQuiz(moduleId, userID)
	if err != nil {
		return nil, err

	}

	res := conversion.ConvertAllSubInModule(subModules, submissions, quizzes)

	return res, nil
}
