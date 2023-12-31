package conversion

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/google/uuid"
)

func ProfileRequestToProfileModel(userID uint, request *web.UpdateProfileRequest) *domain.UserProfile {
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

func ProfileCreateRequestToModel(userID uint, request *web.CreateProfileRequest) *domain.UserProfile {
	if request.ImageUrl == "" {
		request.ImageUrl = fmt.Sprintf("https://robohash.org/%s.png?size=200x200", uuid.NewString())
	}
	
	return &domain.UserProfile{
		ID:          userID,
		UserID:      userID,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		PhoneNumber: request.PhoneNumber,
		DateBirth:   request.DateBirth,
		Gender: request.Gender,
		ImageUrl: request.ImageUrl,
	}
}

func ProfileModelToResponse(request *domain.UserProfile) *web.GetProfileResponse {
	return &web.GetProfileResponse{
		ID:          request.ID,
		ImageUrl:    request.ImageUrl,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateBirth:   request.DateBirth,
		Bio:         request.Bio,
		Gender:      request.Gender,
		PhoneNumber: request.PhoneNumber,
		City:        request.City,
		Address:     request.Address,
		Job:         request.Job,
	}
}
