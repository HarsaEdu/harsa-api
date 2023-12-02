package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func FeedbackCreateRequestToCategoriesModel(request web.FeedbackCreateRequest, userId uint, courseId uint) *domain.Feedback {
	return &domain.Feedback{
		UserID:  userId,
		CourseID: courseId,
		Rating:  request.Rating,
		Content: request.Content,
	}
}

func FeedbackUpdateRequestToCategoriesModel(request web.FeedbackUpdateRequest) *domain.Feedback {
	return &domain.Feedback{
		Rating:  request.Rating,
		Content: request.Content,
	}
}
