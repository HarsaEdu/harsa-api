package conversion

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func UserCreateRequestToUserModel(userRequest web.UserCreateRequest) *domain.User {
	return &domain.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		RoleID:   userRequest.RoleID,
	}
}
func UserCreateRequestToUserProfileModel(userRequest web.UserCreateRequest, userID uint, birthDate time.Time) *domain.UserProfile {
	return &domain.UserProfile{
		ID: userID,
		UserID:      userID,
		FirstName:   userRequest.FirstName,
		LastName:    userRequest.LastName,
		DateBirth:   birthDate,
		Bio:         userRequest.Bio,
		Gender:      userRequest.Gender,
		PhoneNumber: userRequest.PhoneNumber,
		City:        userRequest.City,
		Address:     userRequest.Address,
		Job:         userRequest.Job,
	}
}
func UserUpdateRequestToUserModel(userRequest web.UserUpdateRequest) *domain.User {
	return &domain.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		RoleID:   userRequest.RoleID,
	}
}
func UserProfileUpdateRequestToUserModel(userRequest web.UserProfileUpdateRequest, birthDate time.Time) *domain.UserProfile {
	return &domain.UserProfile{
		ID:          userRequest.ID,
		FirstName:   userRequest.FirstName,
		LastName:    userRequest.LastName,
		DateBirth:   birthDate,
		Bio:         userRequest.Bio,
		Gender:      userRequest.Gender,
		PhoneNumber: userRequest.PhoneNumber,
		City:        userRequest.City,
		Address:     userRequest.Address,
		Job:         userRequest.Job,
	}
}
