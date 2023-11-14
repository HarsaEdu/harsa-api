package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func CourseDomainToCourseGetAllResponse(courseDomain []domain.Course) []web.GetCourseResponse {
	var courseGetAllResponse []web.GetCourseResponse

	for _, course := range courseDomain {

		courseGetAllResponse = append(courseGetAllResponse, web.GetCourseResponse{
			ID:        course.ID,
			Title: course.Title,
			Description: course.Description,
			ImageUrl: course.ImageUrl,
			Enrolled: course.Enrolled,
			Rating: course.Rating,
			CreatedAt: course.CreatedAt,
			UpdatedAt: course.UpdatedAt,
			Category: web.CategoryForCourseResponse{
				ID: course.Category.ID,
				Name: course.Category.Name,
			},
			User: web.UserForCourseResponse{
				ID: course.User.ID,
				Email: course.User.Email,
				Role: web.RoleForUserForCourseResponse{
					ID: course.User.Role.ID,
					Name: course.User.Role.Name,
				},
			},
		})
	}

	return courseGetAllResponse
}