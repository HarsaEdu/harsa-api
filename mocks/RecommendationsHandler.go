// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// RecommendationsHandler is an autogenerated mock type for the RecommendationsHandler type
type RecommendationsHandler struct {
	mock.Mock
}

// GetRecommendations provides a mock function with given fields: ctx
func (_m *RecommendationsHandler) GetRecommendations(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRecommendationsHandler creates a new instance of RecommendationsHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRecommendationsHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *RecommendationsHandler {
	mock := &RecommendationsHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
