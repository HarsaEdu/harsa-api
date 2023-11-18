package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ProfileRequestToProfileModel(userID uint, request *web.ProfileRequest) *domain.UserProfile {
	return &domain.UserProfile{
		UserID:      userID,
		ImageUrl:    request.ImageUrl,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateBirth:   request.DateBirth,
		Bio:         request.Bio,
		Gender:      domain.Gender(request.Gender),
		PhoneNumber: request.PhoneNumber,
		City:        request.City,
		Address:     request.Address,
		Job:         request.Job,
	}
}
