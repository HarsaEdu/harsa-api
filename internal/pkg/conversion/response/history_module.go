package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertModuleResponseTrackingMobile(response *domain.Module, progress float32) *web.ModuleResponseForTracking {
	return &web.ModuleResponseForTracking{
		ID:          response.ID,
		Title:       response.Title,
		Description: response.Description,
		Order: response.Order,
		Progress: progress,
	}
}

func ConvertAllSubInModule(subModules []web.SubModuleResponseForTracking, submissions []web.SubmissionsResponseModuleMobile, quizzes []web.QuizResponseForTracking) *web.CourseTrackingSub{
	return &web.CourseTrackingSub{
		SubModules: subModules,
		Submissions: submissions,
		Quizzes: quizzes,
	}
}

// func ConvertHistoryModuleResponseMobile(response *domain.HistoryModule) *web.HistoryModuleResponseMobile {
	
// 	module := ConvertModuleResponseTrackingMobile(&response.Module)

// 	subModule := ConvertAllHistorySubmoduleResponseMobile(response.HistorySubModule)

// 	quiz := ConvertAllHistoryQuizResponseMobile(response.HistoryQuiz)

// 	submissionAnswer := ConvertAllSubmissionAnswerResponseMobile(response.SubmissionAnswer)
	
// 	return &web.HistoryModuleResponseMobile{
// 		ID:          response.ID,
// 		Module:   *module,
// 		Progress:    response.Progress,
// 		HistorySubModule: subModule,
// 		HistoryQuiz: quiz,
// 		SubmissionAnswer: submissionAnswer,
// 	}
// }

// func ConvertAllHistoryModuleResponseMobile(response []domain.HistoryModule) []web.HistoryModuleResponseMobile {
	
// 	var historyModule []web.HistoryModuleResponseMobile

// 	for i := range response {
// 		historyModule = append(historyModule, *ConvertHistoryModuleResponseMobile(&response[i]))
// 	}

// 	return historyModule
		
// }