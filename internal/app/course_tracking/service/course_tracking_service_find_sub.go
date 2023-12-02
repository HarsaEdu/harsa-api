package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (courseTrackingService *CourseTrackingServiceImpl) FindSubByIdMobile(moduleID uint, userID uint) (*web.CourseTrackingSub, error) {

	res, err := courseTrackingService.CourseTrackingRepository.FindAllSub(moduleID, userID)
	if err != nil {
		return nil, fmt.Errorf("eror when find course tracking by id  :%s", err.Error())
	}

	return res, nil

}

func (courseTrackingService *CourseTrackingServiceImpl) FindSubModuleByID(moduleID uint, subModuleID uint, userID uint) (*web.SubModuleTrackingResponse, error) {
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
	newHistory := conversion.ConvertHistorySubmoduleResponseMobile(history)
	newSubModule := conversion.ConvertSubModuleResponseModule(*subModule)
	res := conversion.ConvertToSubModuleTrackingResponse(module, newHistory, newSubModule, progress)

	return res, nil
}

func (courseTrackingService *CourseTrackingServiceImpl) FindSubmissionByID(moduleID uint, userID uint, submissionID uint) (*web.SubmissionAnswerTrackingResponse, error) {
	module, err := courseTrackingService.CourseTrackingRepository.FindModuleTracking(moduleID, userID)
	if err != nil {
		return nil, fmt.Errorf("module not found")
	}

	submissionAnswer, submission, err := courseTrackingService.CourseTrackingRepository.FindSubmissionByID(moduleID, userID, submissionID)
	if err != nil {
		return nil, fmt.Errorf("history not found")
	}

	newAnswer := conversion.ConvertSubmissionAnswerTracking(submissionAnswer, submission.Title)
	newSubmission := conversion.ConvertSubmissionResponseModule(submission)

	res := conversion.ConvertSubmissionAnswerTrackingResponse(module, newAnswer, newSubmission)
	return res, nil
}
