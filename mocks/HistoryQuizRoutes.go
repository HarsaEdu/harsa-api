// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// HistoryQuizRoutes is an autogenerated mock type for the HistoryQuizRoutes type
type HistoryQuizRoutes struct {
	mock.Mock
}

// HistoryQuizMobile provides a mock function with given fields: apiGroup
func (_m *HistoryQuizRoutes) HistoryQuizMobile(apiGroup *echo.Group) {
	_m.Called(apiGroup)
}

// HistoryQuizWeb provides a mock function with given fields: apiGroup
func (_m *HistoryQuizRoutes) HistoryQuizWeb(apiGroup *echo.Group) {
	_m.Called(apiGroup)
}

// NewHistoryQuizRoutes creates a new instance of HistoryQuizRoutes. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHistoryQuizRoutes(t interface {
	mock.TestingT
	Cleanup(func())
}) *HistoryQuizRoutes {
	mock := &HistoryQuizRoutes{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
