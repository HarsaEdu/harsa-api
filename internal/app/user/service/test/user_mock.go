package test

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) UserAvailable(username, email string) (*domain.User, error) {
	args := m.Called(username, email)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.User), args.Error(1)
}

func (m *MockUserRepository) UserAvailableByID(id uint) (*domain.User, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.User), args.Error(1)
}

func (m *MockUserRepository) UserProfileCreate(userProfile *domain.UserProfile) error {
	args := m.Called(userProfile)
	return args.Error(0)
}

func (m *MockUserRepository) UserCreate(user *domain.User) (*domain.User, error) {
	args := m.Called(user)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.User), args.Error(1)
}

func (m *MockUserRepository) UserUpdate(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) UserProfileUpdate(userProfile *domain.UserProfile) error {
	args := m.Called(userProfile)
	return args.Error(0)
}

func (m *MockUserRepository) UserProfileAvailableByID(id uint) (*domain.UserProfile, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.UserProfile), args.Error(1)
}

func (m *MockUserRepository) UserDelete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) HandleTrx(ctx echo.Context, fn func(repo MockUserRepository) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}

func (m *MockUserRepository) UserGetAll(offset, limit int, search string, roleId int) ([]domain.UserEntity, int64, error) {
	args := m.Called(offset, limit, search, roleId)
	result := args.Get(0)
	if result == nil {
		return nil, 0, args.Error(2)
	}
	return result.([]domain.UserEntity), args.Get(1).(int64), args.Error(2)
}

func (m *MockUserRepository) GetUserByID(userID uint) (*domain.UserDetail, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.UserDetail), args.Error(1)
}

func (m *MockUserRepository) GetUserAccountByID(userID uint) (*domain.User, error) {
	args := m.Called(userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.User), args.Error(1)
}

func (m *MockUserRepository) UserAvailableUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.User), args.Error(1)
}

func (m *MockUserRepository) UserAvailableEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*domain.User), args.Error(1)
}

func (m *MockUserRepository) UserGetAllStudentSubscribe(offset, limit int, search string, courseId uint) ([]domain.UserEntity, int64, error) {
	args := m.Called(offset, limit, search, courseId)
	result := args.Get(0)
	if result == nil {
		return nil, 0, args.Error(2)
	}
	return result.([]domain.UserEntity), args.Get(1).(int64), args.Error(2)
}
