package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func InterestRequestToModel(profileID uint, request uint) *domain.UserInterest {
	return &domain.UserInterest{
		ProfileID:  profileID,
		CategoryID: request,
	}
}
