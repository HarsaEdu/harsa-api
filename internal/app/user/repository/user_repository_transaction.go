package repository

import "github.com/labstack/echo/v4"

func (userRepository *UserRepositoryImpl) HandleTrx(ctx echo.Context, fn func(repo UserRepository) error) error {

	tx := userRepository.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	newRepo := &UserRepositoryImpl{
		DB: tx,
	}

	err := fn(newRepo)
	if err != nil {
		return err
	}

	err = tx.Commit().Error

	return err
}
