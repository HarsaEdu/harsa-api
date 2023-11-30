package web

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

type CourseTrackingRequest struct {
	UserID   uint                       `json:"user_id" validate:"required"`
	CourseID uint                       `json:"course_id" validate:"required"`
	Status   domain.StatusCourseTraking `json:"status"`
}
