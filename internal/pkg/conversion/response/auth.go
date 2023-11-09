package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertToAuthResponse(r *domain.Auth) *web.AuthResponse {
	return &web.AuthResponse{
		ID:       r.ID,
		Username: r.Username,
		Email:    r.Email,
		RoleName: web.Role(r.RoleName),
	}
}
