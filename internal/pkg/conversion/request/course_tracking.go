package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func CourseTrackingReqToDomain(request *web.CourseTrackingRequest) *domain.CourseTracking {
	return &domain.CourseTracking{
		CourseID: request.CourseID,
		UserID: request.UserID,
		Status: request.Status,
	}
}