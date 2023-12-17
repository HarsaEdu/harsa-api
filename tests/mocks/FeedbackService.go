// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	web "github.com/HarsaEdu/harsa-api/internal/model/web"
)

// FeedbackService is an autogenerated mock type for the FeedbackService type
type FeedbackService struct {
	mock.Mock
}

// CreateByUserAndCourseId provides a mock function with given fields: feedback, userId, courseId
func (_m *FeedbackService) CreateByUserAndCourseId(feedback web.FeedbackCreateRequest, userId uint, courseId uint) error {
	ret := _m.Called(feedback, userId, courseId)

	var r0 error
	if rf, ok := ret.Get(0).(func(web.FeedbackCreateRequest, uint, uint) error); ok {
		r0 = rf(feedback, userId, courseId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByUserAndCourseId provides a mock function with given fields: userId, courseId
func (_m *FeedbackService) DeleteByUserAndCourseId(userId uint, courseId uint) error {
	ret := _m.Called(userId, courseId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userId, courseId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllByCourseId provides a mock function with given fields: courseId, offset, limit, search
func (_m *FeedbackService) GetAllByCourseId(courseId uint, offset int, limit int, search string) ([]web.FeedBackResponseForTracking, *web.Pagination, error) {
	ret := _m.Called(courseId, offset, limit, search)

	var r0 []web.FeedBackResponseForTracking
	var r1 *web.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(uint, int, int, string) ([]web.FeedBackResponseForTracking, *web.Pagination, error)); ok {
		return rf(courseId, offset, limit, search)
	}
	if rf, ok := ret.Get(0).(func(uint, int, int, string) []web.FeedBackResponseForTracking); ok {
		r0 = rf(courseId, offset, limit, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]web.FeedBackResponseForTracking)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, int, int, string) *web.Pagination); ok {
		r1 = rf(courseId, offset, limit, search)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*web.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(uint, int, int, string) error); ok {
		r2 = rf(courseId, offset, limit, search)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByIdAndCourseId provides a mock function with given fields: courseId, id
func (_m *FeedbackService) GetByIdAndCourseId(courseId uint, id uint) (*web.FeedBackResponseForTracking, error) {
	ret := _m.Called(courseId, id)

	var r0 *web.FeedBackResponseForTracking
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (*web.FeedBackResponseForTracking, error)); ok {
		return rf(courseId, id)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) *web.FeedBackResponseForTracking); ok {
		r0 = rf(courseId, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*web.FeedBackResponseForTracking)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(courseId, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIdUserAndCourseId provides a mock function with given fields: userId, courseId
func (_m *FeedbackService) GetByIdUserAndCourseId(userId uint, courseId uint) (*web.FeedBackResponseForTracking, error) {
	ret := _m.Called(userId, courseId)

	var r0 *web.FeedBackResponseForTracking
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (*web.FeedBackResponseForTracking, error)); ok {
		return rf(userId, courseId)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) *web.FeedBackResponseForTracking); ok {
		r0 = rf(userId, courseId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*web.FeedBackResponseForTracking)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userId, courseId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByUserAndCourseId provides a mock function with given fields: feedback, userId, courseId
func (_m *FeedbackService) UpdateByUserAndCourseId(feedback web.FeedbackUpdateRequest, userId uint, courseId uint) error {
	ret := _m.Called(feedback, userId, courseId)

	var r0 error
	if rf, ok := ret.Get(0).(func(web.FeedbackUpdateRequest, uint, uint) error); ok {
		r0 = rf(feedback, userId, courseId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFeedbackService creates a new instance of FeedbackService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFeedbackService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FeedbackService {
	mock := &FeedbackService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}