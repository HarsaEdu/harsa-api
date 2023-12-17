// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	web "github.com/HarsaEdu/harsa-api/internal/model/web"
)

// DashboardService is an autogenerated mock type for the DashboardService type
type DashboardService struct {
	mock.Mock
}

// CountAll provides a mock function with given fields:
func (_m *DashboardService) CountAll() (*web.AllCountDashboard, error) {
	ret := _m.Called()

	var r0 *web.AllCountDashboard
	var r1 error
	if rf, ok := ret.Get(0).(func() (*web.AllCountDashboard, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *web.AllCountDashboard); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*web.AllCountDashboard)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDashboardService creates a new instance of DashboardService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDashboardService(t interface {
	mock.TestingT
	Cleanup(func())
}) *DashboardService {
	mock := &DashboardService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}