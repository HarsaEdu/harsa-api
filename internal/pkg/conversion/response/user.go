package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertUserAccountResponse(userAccount *domain.User) *web.UserAccountResponse {
	return &web.UserAccountResponse{
		ID:        userAccount.ID,
		Username:  userAccount.Username,
		Email:     userAccount.Email,
		CreatedAt: userAccount.CreatedAt,
		UpdatedAt: userAccount.UpdatedAt,
		Role: web.UserRole{
			ID:   userAccount.Role.ID,
			Name: userAccount.Role.Name,
		},
	}
}