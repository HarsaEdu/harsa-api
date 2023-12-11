package mocks

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

type MockAuthRepository struct {
	RegisterUserFunc func(user *domain.User) (*domain.Auth, error)
	LoginUserFunc    func(id uint) (*domain.Auth, error)
}

func (m *MockAuthRepository) RegisterUser(user *domain.User) (*domain.Auth, error) {
	return m.RegisterUser(user)
}

func (m *MockAuthRepository) LoginUser(id uint) (*domain.Auth, error) {
	return m.LoginUser(id)
}
