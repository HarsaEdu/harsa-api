// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// DashboardRoutes is an autogenerated mock type for the DashboardRoutes type
type DashboardRoutes struct {
	mock.Mock
}

// DashboardWeb provides a mock function with given fields: apiGroup
func (_m *DashboardRoutes) DashboardWeb(apiGroup *echo.Group) {
	_m.Called(apiGroup)
}

// NewDashboardRoutes creates a new instance of DashboardRoutes. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDashboardRoutes(t interface {
	mock.TestingT
	Cleanup(func())
}) *DashboardRoutes {
	mock := &DashboardRoutes{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}