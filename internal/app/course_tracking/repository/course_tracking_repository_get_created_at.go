package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseTrackingrepository *CourseTrackingRepositoryImpl) GetCreatedAt(id uint) (int64, error) {
	var user = domain.User{}

	if err := courseTrackingrepository.DB.First(&user, id).Error; err != nil {
		return 0, err
	}

	createdAt := int64(user.CreatedAt.Unix())

	return createdAt, nil
}