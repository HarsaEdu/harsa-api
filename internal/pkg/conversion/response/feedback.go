package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertUserFeedback(userCourse *domain.UserProfile) *web.UserForFeedBack {
	name := userCourse.FirstName + " " + userCourse.LastName
	return &web.UserForFeedBack{
		ID:   userCourse.ID,
		Name: name,
		ImageUrl: userCourse.ImageUrl,
	}
}

func ConvertFeedbackForTracking(feedback *domain.Feedback) *web.FeedBackResponseForTracking {
	
	user := ConvertUserFeedback(&feedback.User.UserProfile)
	
	return &web.FeedBackResponseForTracking{
		ID:   feedback.ID,
		Rating: feedback.Rating,
		Content: feedback.Content,
		CreatedAt: feedback.CreatedAt,
		User: *user,
		CreatedAt: feedback.CreatedAt,
		UpdatedAt: feedback.UpdatedAt,
	}
}

func ConvertAllFeedBackResponseMobile(response []domain.Feedback) []web.FeedBackResponseForTracking {
	
	var historyQuiz []web.FeedBackResponseForTracking

	for i := range response {
		historyQuiz = append(historyQuiz, *ConvertFeedbackForTracking(&response[i]))
	}

	return historyQuiz
		
}

