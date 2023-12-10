package conversion

import (
	"fmt"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func CourseTrackingDomainToCertificateDomain(courseTracking *domain.CourseTracking) *domain.Certificate {
	certificate := domain.Certificate{
		ID: fmt.Sprintf("C%d%d%d", courseTracking.CourseID, courseTracking.UserID, courseTracking.UpdatedAt.Day()),
		UserID: courseTracking.UserID,
		User: courseTracking.User,
		CourseID: courseTracking.CourseID,
		Course: courseTracking.Course,
		CreatedAt: courseTracking.UpdatedAt,
	}

	tookDate := strings.Split(certificate.CreatedAt.String(), " ")
	certificate.IssuedAt = tookDate[0]

	return &certificate
}