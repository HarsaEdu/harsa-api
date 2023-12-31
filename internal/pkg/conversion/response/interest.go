package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func InterestResultToResponse(interest []domain.CourseEntity) *[]web.InterestResponse {
	response := []web.InterestResponse{}

	for _, value := range interest {
		name := value.FirstName + " " + value.LastName
		result := &web.InterestResponse{
			CourseID:   value.ID,
			Title:      value.Title,
			ImageUrl:   value.ImageUrl,
			Rating:     value.Rating,
			Instructor: name,
		}

		response = append(response, *result)
	}
	return &response
}
