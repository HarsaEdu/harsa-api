package repository

import (
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
)

func (userRepository *UserRepositoryImpl) UserCreate(user *domain.User) (*domain.User, error) {
	// insert user into database
	result := userRepository.DB.Create(&user)

	// check if error when user inserted
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}


func (userRepository *UserRepositoryImpl) UserCreateMobile(user *domain.User) (*domain.User, error) {
	tx := userRepository.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			
			tx.Rollback()
		}
	}()
	
	result := tx.Create(&user)
	
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

	
	tx.Commit()

	return user, nil
}
