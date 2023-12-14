package repository

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (authRepository *AuthRepositoryImpl) RegisterUser(user *domain.User) (*domain.Auth, error) {
	// define domain auth
	auth := &domain.Auth{}

	// insert user into database
	result := authRepository.DB.Create(&user)

	// check if error when user inserted
	if result.Error != nil {
		return nil, result.Error
	}

	// get user data from database
	authRepository.DB.Model(&domain.User{}).
		Select("users.id as id, username, email, roles.name as role_name, users.created_at as created_at, users.registration_token as registration_token").
		Joins("left join roles on roles.id = users.role_id").
		Where("users.id = ?", user.ID).
		First(&auth)

	// check if error when get user data from database
	if result.Error != nil {
		return nil, result.Error
	}

	return auth, nil
}

func (authRepository *AuthRepositoryImpl) RegisterWithFreeSubscibe(user *domain.User) (*domain.Auth, error) {
	tx := authRepository.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Create(user)

	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	subscription := domain.Subscription{
		StartDate: time.Now(),
		UserID:    user.ID,
	}

	endDate := subscription.StartDate.Add(time.Hour * 24 * 7)
	subscription.EndDate = endDate.UTC()
	resultSubscribe := tx.Create(&subscription)

	if resultSubscribe.Error != nil {
		tx.Rollback()
		return nil, resultSubscribe.Error
	}

	auth := &domain.Auth{}
	err := tx.Model(&domain.User{}).
		Select("users.id, username, email, roles.name as role_name, users.created_at").
		Joins("left join roles on roles.id = users.role_id").
		Where("users.id = ?", user.ID).
		First(auth).Error

	// check if error when getting user data from the database
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return auth, nil
}
