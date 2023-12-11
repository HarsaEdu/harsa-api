// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// SubmissionRoutes is an autogenerated mock type for the SubmissionRoutes type
type SubmissionRoutes struct {
	mock.Mock
}

// SubmissionMobile provides a mock function with given fields: apiGroup
func (_m *SubmissionRoutes) SubmissionMobile(apiGroup *echo.Group) {
	_m.Called(apiGroup)
}

// SubmissionWeb provides a mock function with given fields: apiGroup
func (_m *SubmissionRoutes) SubmissionWeb(apiGroup *echo.Group) {
	_m.Called(apiGroup)
}

// NewSubmissionRoutes creates a new instance of SubmissionRoutes. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubmissionRoutes(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubmissionRoutes {
	mock := &SubmissionRoutes{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
