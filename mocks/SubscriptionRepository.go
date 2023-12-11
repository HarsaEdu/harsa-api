// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/HarsaEdu/harsa-api/internal/model/domain"
	mock "github.com/stretchr/testify/mock"
)

// SubscriptionRepository is an autogenerated mock type for the SubscriptionRepository type
type SubscriptionRepository struct {
	mock.Mock
}

// AddSubscription provides a mock function with given fields: subscription
func (_m *SubscriptionRepository) AddSubscription(subscription *domain.Subscription) error {
	ret := _m.Called(subscription)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Subscription) error); ok {
		r0 = rf(subscription)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindSubscription provides a mock function with given fields: userID
func (_m *SubscriptionRepository) FindSubscription(userID uint) (*domain.Subscription, error) {
	ret := _m.Called(userID)

	var r0 *domain.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*domain.Subscription, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) *domain.Subscription); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSubscription provides a mock function with given fields: subscription
func (_m *SubscriptionRepository) UpdateSubscription(subscription *domain.Subscription) error {
	ret := _m.Called(subscription)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Subscription) error); ok {
		r0 = rf(subscription)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSubscriptionRepository creates a new instance of SubscriptionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubscriptionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubscriptionRepository {
	mock := &SubscriptionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
