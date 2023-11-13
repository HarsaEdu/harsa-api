package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertToAuthResponse(auth *domain.Auth) *web.AuthResponse {
	return &web.AuthResponse{
		ID:       auth.ID,
		Username: auth.Username,
		Email:    auth.Email,
		RoleName: web.Role(auth.RoleName),
	}
}
