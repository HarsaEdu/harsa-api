package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
	"github.com/labstack/echo/v4"
)

func (courseTrackingService *CourseTrackingServiceImpl) FindSubByIdMobile(moduleID uint, userID uint) (*web.CourseTrackingSub, error) {
		
	res, err := courseTrackingService.CourseTrackingRepository.FindAllSub(moduleID, userID)
	if err != nil {
		return nil, fmt.Errorf("eror when find course tracking by id  :%s", err.Error())
	}

	return res, nil

}

func (courseTrackingService *CourseTrackingServiceImpl) FindModuleHistory(ctx echo.Context, courseID uint,moduleID uint, userID uint) (*web.ModuleTrackingByID, error) {
	
	isSubscription, err:= courseTrackingService.Subscription.IsSubscription(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("eror when cek subscription  :%s", err.Error())
	}

	if !isSubscription {
		return nil, fmt.Errorf("unauthorized")
	}
	
	courseTraking, err := courseTrackingService.CourseTrackingRepository.FindByUserIdAndCourseID(courseID, userID)
	if err != nil { 
		return nil, fmt.Errorf("eror when find course tracking by id  :%s", err.Error())
	}
	if courseTraking.ID == 0{
		return nil, fmt.Errorf("not enrolled")
	
	}

	module, err := courseTrackingService.CourseTrackingRepository.FindModuleTracking(moduleID, userID)
	if err != nil {
		return nil, fmt.Errorf("module not found")
	}

	courseTracking, err := courseTrackingService.CourseTrackingRepository.FindAllSub(moduleID, userID)
	if err != nil {
		return nil, fmt.Errorf("eror when find all sub  :%s", err.Error())
	}

	res := conversion.ConvertToModuleTrackingByID(module, courseTracking)
	return res, nil
}

func (courseTrackingService *CourseTrackingServiceImpl) FindSubModuleByID(ctx echo.Context, courseID uint, moduleID uint, subModuleID uint, userID uint) (*web.SubModuleTracking, error) {
	
	isSubscription, err:= courseTrackingService.Subscription.IsSubscription(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("eror when cek subscription  :%s", err.Error())
	}

	if !isSubscription {
		return nil, fmt.Errorf("unauthorized")
	}
	
	courseTraking, err := courseTrackingService.CourseTrackingRepository.FindByUserIdAndCourseID(courseID, userID)
	if err != nil { 
		return nil, fmt.Errorf("eror when find course tracking by id  :%s", err.Error())
	}
	if courseTraking.ID == 0{
		return nil, fmt.Errorf("not enrolled")
	
	}
	
	module, err := courseTrackingService.CourseTrackingRepository.FindModuleTracking(moduleID, userID)
	if err != nil {
		return nil, fmt.Errorf("module not found")
	}
	history, subModule, err := courseTrackingService.CourseTrackingRepository.FindSubModuleByID(moduleID, userID, subModuleID)
	if err != nil {
		return nil, fmt.Errorf("history not found")
	}
	progress, err := courseTrackingService.CourseTrackingRepository.CountProgressModule(moduleID, userID)
	if err != nil {
		return nil, fmt.Errorf("progress not found")
	}
	newHistory := conversion.ConvertHistorySubmoduleTracking(history)
	newSubModule := conversion.ConvertSubModuleResponseModule(*subModule)
	res := conversion.ConvertToSubModuleTrackingResponse(module, newHistory, newSubModule, progress)

	return res, nil
}

func (courseTrackingService *CourseTrackingServiceImpl) FindSubmissionByID(ctx echo.Context, courseID uint,moduleID uint, userID uint, submissionID uint) (*web.SubmissionAnswerTrackingByIDResponse, error) {
	
	isSubscription, err:= courseTrackingService.Subscription.IsSubscription(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("eror when cek subscription  :%s", err.Error())
	}

	if !isSubscription {
		return nil, fmt.Errorf("unauthorized")
	}
	
	courseTraking, err := courseTrackingService.CourseTrackingRepository.FindByUserIdAndCourseID(courseID, userID)
	if err != nil { 
		return nil, fmt.Errorf("eror when find course tracking by id  :%s", err.Error())
	}
	if courseTraking.ID == 0{
		return nil, fmt.Errorf("not enrolled")
	
	}
	
	module, err := courseTrackingService.CourseTrackingRepository.FindModuleTracking(moduleID, userID)
	if err != nil {
		return nil, fmt.Errorf("module not found")
	}

	submissionAnswer, submission, err := courseTrackingService.CourseTrackingRepository.FindSubmissionByID(moduleID, userID, submissionID)
	if err != nil {
		return nil, fmt.Errorf("history not found")
	}

	newAnswer := conversion.ConvertSubmissionAnswerTracking(submissionAnswer)
	newSubmission := conversion.ConvertSubmissionResponseModule(submission)

	res := conversion.ConvertSubmissionAnswerTrackingResponse(module, newAnswer, newSubmission)
	return res, nil
}

func (courseTrackingService *CourseTrackingServiceImpl) FindQuizzByID(ctx echo.Context, courseID uint,moduleID uint, userID uint, quizzID uint) (*web.HistoryQuizIDTracking, error) {
	isSubscription, err:= courseTrackingService.Subscription.IsSubscription(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("eror when cek subscription  :%s", err.Error())
	}

	if !isSubscription {
		return nil, fmt.Errorf("unauthorized")
	}
	
	courseTraking, err := courseTrackingService.CourseTrackingRepository.FindByUserIdAndCourseID(courseID, userID)
	if err != nil { 
		return nil, fmt.Errorf("eror when find course tracking by id  :%s", err.Error())
	}
	if courseTraking.ID == 0{
		return nil, fmt.Errorf("not enrolled")
	
	}
	
	module, err := courseTrackingService.CourseTrackingRepository.FindModuleTracking(moduleID, userID)
	if err != nil {
		return nil, fmt.Errorf("module not found")
	}

	historyQuizz, err := courseTrackingService.CourseTrackingRepository.FindQuizzByID(moduleID, userID, quizzID)
	if err != nil {
		return nil, fmt.Errorf("history not found")
	}

	quizz, err := courseTrackingService.QuizzService.FindByIdMobile(quizzID)
	if err != nil {
		return nil, err
	}

	newHistory := conversion.ConvertHistoryQuizForTracking(historyQuizz)
	res := conversion.ConvertHistoryQuizTrackingResponse(module, newHistory, quizz)

	return res, nil
}
