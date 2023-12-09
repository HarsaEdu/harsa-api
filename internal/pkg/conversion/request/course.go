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

func CourseCreateRequestToCourseDomainIntructure(request *web.CourseCreateRequestIntructure, instructorId uint) *domain.Course {
	return &domain.Course{
		UserID: instructorId,
		Title:       request.Title,
		Description: request.Description,
		ImageUrl: "https://placewaifu.com/image/200/200",
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

func CourseUpdateRequestToCourseDomainIntructure(request *web.CourseUpdateRequestIntructure, courseId uint) *domain.Course {
	return &domain.Course{
		ID:          courseId,
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