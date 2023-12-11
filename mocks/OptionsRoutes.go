// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// OptionsRoutes is an autogenerated mock type for the OptionsRoutes type
type OptionsRoutes struct {
	mock.Mock
}

// OptionsWeb provides a mock function with given fields: apiGroup
func (_m *OptionsRoutes) OptionsWeb(apiGroup *echo.Group) {
	_m.Called(apiGroup)
}

// NewOptionsRoutes creates a new instance of OptionsRoutes. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOptionsRoutes(t interface {
	mock.TestingT
	Cleanup(func())
}) *OptionsRoutes {
	mock := &OptionsRoutes{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
