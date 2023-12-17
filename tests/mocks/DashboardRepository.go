// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	web "github.com/HarsaEdu/harsa-api/internal/model/web"
)

// DashboardRepository is an autogenerated mock type for the DashboardRepository type
type DashboardRepository struct {
	mock.Mock
}

// CountCourse provides a mock function with given fields:
func (_m *DashboardRepository) CountCourse() (int64, error) {
	ret := _m.Called()

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountInterest provides a mock function with given fields:
func (_m *DashboardRepository) CountInterest() ([]web.CountInterest, int64, error) {
	ret := _m.Called()

	var r0 []web.CountInterest
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func() ([]web.CountInterest, int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []web.CountInterest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]web.CountInterest)
		}
	}

	if rf, ok := ret.Get(1).(func() int64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CountIntructure provides a mock function with given fields:
func (_m *DashboardRepository) CountIntructure() (int64, error) {
	ret := _m.Called()

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountStudent provides a mock function with given fields:
func (_m *DashboardRepository) CountStudent() (int64, error) {
	ret := _m.Called()

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDashboardRepository creates a new instance of DashboardRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDashboardRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *DashboardRepository {
	mock := &DashboardRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}