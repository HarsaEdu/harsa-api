package service

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	conversion "github.com/HarsaEdu/harsa-api/internal/pkg/conversion/response"
)

func (feedbackService *FeedbackServiceImpl) GetAll(courseid, offset, limit int) ([]domain.Feedback, *web.Pagination, error) {
	result, total, err := feedbackService.FeedbackRepository.GetAll(courseid, offset, limit)

	if total == 0 {
		return nil, nil, fmt.Errorf("feedback not found")
	}

	if err != nil {
		return nil, nil, fmt.Errorf("internal Server Error")
	}

	pagination := conversion.RecordToPaginationResponse(offset, limit, total)

	return result, pagination, nil
}
