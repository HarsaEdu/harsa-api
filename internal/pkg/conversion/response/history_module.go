package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertModuleResponseTrackingMobile(response *domain.Module) *web.ModuleResponseForTracking {
	return &web.ModuleResponseForTracking{
		ID:          response.ID,
		Section:     response.Section,
		Title:       response.Title,
		Description: response.Description,
		Order: response.Order,
	}
}

func ConvertHistoryModuleResponseMobile(response *domain.HistoryModule) *web.HistoryModuleResponseMobile {
	
	module := ConvertModuleResponseTrackingMobile(&response.Module)

	subModule := ConvertAllHistorySubmoduleResponseMobile(response.HistorySubModule)

	quiz := ConvertAllHistoryQuizResponseMobile(response.HistoryQuiz)

	submissionAnswer := ConvertAllSubmissionAnswerResponseMobile(response.SubmissionAnswer)
	
	return &web.HistoryModuleResponseMobile{
		ID:          response.ID,
		Module:   *module,
		Progress:    response.Progress,
		HistorySubModule: subModule,
		HistoryQuiz: quiz,
		SubmissionAnswer: submissionAnswer,
	}
}

func ConvertAllHistoryModuleResponseMobile(response []domain.HistoryModule) []web.HistoryModuleResponseMobile {
	
	var historyModule []web.HistoryModuleResponseMobile

	for i := range response {
		historyModule = append(historyModule, *ConvertHistoryModuleResponseMobile(&response[i]))
	}

	return historyModule
		
}