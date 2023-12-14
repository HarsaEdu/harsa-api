package conversion

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func CourseTrackingDomainToCertificateDomain(courseTracking *domain.CourseTracking) *domain.Certificate {
	certificate := domain.Certificate{
		ID: fmt.Sprintf("C%d%d%d", courseTracking.CourseID, courseTracking.UserID, courseTracking.UpdatedAt.Unix()),
		UserID: courseTracking.UserID,
		User: courseTracking.User,
		CourseID: courseTracking.CourseID,
		Course: courseTracking.Course,
		CreatedAt: courseTracking.UpdatedAt,
	}

	issuedYear, issuedMonth, issuedDay := certificate.CreatedAt.Date()

	certificate.IssuedAt = fmt.Sprintf("%d %s %d", issuedDay, issuedMonth, issuedYear)

	return &certificate
}