// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/HarsaEdu/harsa-api/internal/model/domain"
	mock "github.com/stretchr/testify/mock"

	web "github.com/HarsaEdu/harsa-api/internal/model/web"
)

// ChatbotRepository is an autogenerated mock type for the ChatbotRepository type
type ChatbotRepository struct {
	mock.Mock
}

// CreateUserChatTopic provides a mock function with given fields: userChatTopic
func (_m *ChatbotRepository) CreateUserChatTopic(userChatTopic *domain.UserChatTopic) error {
	ret := _m.Called(userChatTopic)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserChatTopic) error); ok {
		r0 = rf(userChatTopic)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: topic
func (_m *ChatbotRepository) Delete(topic *domain.UserChatTopic) error {
	ret := _m.Called(topic)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.UserChatTopic) error); ok {
		r0 = rf(topic)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTopicByUserId provides a mock function with given fields: userId
func (_m *ChatbotRepository) GetAllTopicByUserId(userId uint) ([]domain.UserChatTopic, error) {
	ret := _m.Called(userId)

	var r0 []domain.UserChatTopic
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]domain.UserChatTopic, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(uint) []domain.UserChatTopic); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.UserChatTopic)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTopicById provides a mock function with given fields: id
func (_m *ChatbotRepository) GetTopicById(id string) (*domain.UserChatTopic, error) {
	ret := _m.Called(id)

	var r0 *domain.UserChatTopic
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.UserChatTopic, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.UserChatTopic); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserChatTopic)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTopic provides a mock function with given fields: topicId, request
func (_m *ChatbotRepository) UpdateTopic(topicId string, request *web.CreateThreadRequest) error {
	ret := _m.Called(topicId, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *web.CreateThreadRequest) error); ok {
		r0 = rf(topicId, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewChatbotRepository creates a new instance of ChatbotRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChatbotRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChatbotRepository {
	mock := &ChatbotRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}