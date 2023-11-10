package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func RegisterUserRequestToUserModel(r web.RegisterUserRequest) *domain.User {
	return &domain.User{
		Username:          r.Username,
		Email:             r.Email,
		Password:          r.Password,
		RoleID:            1,
		RegistrationToken: r.RegistrationToken,
	}
}
