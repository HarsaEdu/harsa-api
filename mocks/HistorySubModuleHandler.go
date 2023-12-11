// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// HistorySubModuleHandler is an autogenerated mock type for the HistorySubModuleHandler type
type HistorySubModuleHandler struct {
	mock.Mock
}

// CreateHistoryModule provides a mock function with given fields: ctx
func (_m *HistorySubModuleHandler) CreateHistoryModule(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewHistorySubModuleHandler creates a new instance of HistorySubModuleHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHistorySubModuleHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *HistorySubModuleHandler {
	mock := &HistorySubModuleHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
