package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func RegisterUserRequestToUserModel(userRequest web.RegisterUserRequest) *domain.User {
	return &domain.User{
		Username:          userRequest.Username,
		Email:             userRequest.Email,
		Password:          userRequest.Password,
		RoleID:            3,
		RegistrationToken: userRequest.RegistrationToken,
	}
}
