package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func CourseCreateRequestToCourseDomain(request *web.CourseCreateRequest, instructorId uint) *domain.Course {
	return &domain.Course{
		Title:       request.Title,
		Description: request.Description,
		ImageUrl: "https://placewaifu.com/image/200/200",
		UserID: instructorId,
		CategoryID: uint(request.CategoryID),
	}
}

func CourseUpdateRequestToCourseDomain(request *web.CourseUpdateRequest, courseId uint) *domain.Course {
	return &domain.Course{
		ID:          courseId,
		UserID:      request.UserId,
		Title:       request.Title,
		Description: request.Description,
		CategoryID:  uint(request.CategoryID),
	}
}

func CourseUpdateImageRequestToCourseDomain(request *web.CourseUpdateImageRequest, courseId uint) *domain.Course {
	return &domain.Course{
		ID: courseId,
		ImageUrl:    request.ImageUrl,
	}
}