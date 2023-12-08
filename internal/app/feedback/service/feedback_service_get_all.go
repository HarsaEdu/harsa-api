package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (feedbackService *FeedbackServiceImpl) GetAllByCourseId(courseId uint, offset, limit int, search string) ([]web.FeedBackResponseForTracking, *web.Pagination, error) {
	result, total, err := feedbackService.FeedbackRepository.GetAllByCourseId(courseId, offset, limit, search)

	if total == 0 {
		return nil, nil, fmt.Errorf("feedback not found")
	}

	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}

	response := conversion.ConvertAllFeedBackResponseMobile(result)

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return response, pagination, nil
}