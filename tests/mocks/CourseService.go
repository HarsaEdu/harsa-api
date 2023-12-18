// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"

	web "github.com/HarsaEdu/harsa-api/internal/model/web"
)

// CourseService is an autogenerated mock type for the CourseService type
type CourseService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request, instructorId
func (_m *CourseService) Create(ctx echo.Context, request *web.CourseCreateRequest, instructorId uint) error {
	ret := _m.Called(ctx, request, instructorId)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, *web.CourseCreateRequest, uint) error); ok {
		r0 = rf(ctx, request, instructorId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id, userId, role
func (_m *CourseService) Delete(id uint, userId uint, role string) error {
	ret := _m.Called(id, userId, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, string) error); ok {
		r0 = rf(id, userId, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: offset, limit, search, category
func (_m *CourseService) GetAll(offset int, limit int, search string, category string) ([]web.GetCourseResponseMobile, *web.Pagination, error) {
	ret := _m.Called(offset, limit, search, category)

	var r0 []web.GetCourseResponseMobile
	var r1 *web.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]web.GetCourseResponseMobile, *web.Pagination, error)); ok {
		return rf(offset, limit, search, category)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []web.GetCourseResponseMobile); ok {
		r0 = rf(offset, limit, search, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]web.GetCourseResponseMobile)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) *web.Pagination); ok {
		r1 = rf(offset, limit, search, category)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*web.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(int, int, string, string) error); ok {
		r2 = rf(offset, limit, search, category)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllByCategory provides a mock function with given fields: offset, limit, search, category
func (_m *CourseService) GetAllByCategory(offset int, limit int, search string, category uint) ([]web.GetCourseResponseMobileNew, *web.Pagination, error) {
	ret := _m.Called(offset, limit, search, category)

	var r0 []web.GetCourseResponseMobileNew
	var r1 *web.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string, uint) ([]web.GetCourseResponseMobileNew, *web.Pagination, error)); ok {
		return rf(offset, limit, search, category)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, uint) []web.GetCourseResponseMobileNew); ok {
		r0 = rf(offset, limit, search, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]web.GetCourseResponseMobileNew)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, uint) *web.Pagination); ok {
		r1 = rf(offset, limit, search, category)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*web.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(int, int, string, uint) error); ok {
		r2 = rf(offset, limit, search, category)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllByRating provides a mock function with given fields: offset, limit, search, category
func (_m *CourseService) GetAllByRating(offset int, limit int, search string, category string) ([]web.GetCourseResponseMobileNew, *web.Pagination, error) {
	ret := _m.Called(offset, limit, search, category)

	var r0 []web.GetCourseResponseMobileNew
	var r1 *web.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]web.GetCourseResponseMobileNew, *web.Pagination, error)); ok {
		return rf(offset, limit, search, category)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []web.GetCourseResponseMobileNew); ok {
		r0 = rf(offset, limit, search, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]web.GetCourseResponseMobileNew)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) *web.Pagination); ok {
		r1 = rf(offset, limit, search, category)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*web.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(int, int, string, string) error); ok {
		r2 = rf(offset, limit, search, category)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllByUserId provides a mock function with given fields: offset, limit, search, userID
func (_m *CourseService) GetAllByUserId(offset int, limit int, search string, userID uint) (*web.DashboardIntructur, *web.Pagination, error) {
	ret := _m.Called(offset, limit, search, userID)

	var r0 *web.DashboardIntructur
	var r1 *web.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string, uint) (*web.DashboardIntructur, *web.Pagination, error)); ok {
		return rf(offset, limit, search, userID)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, uint) *web.DashboardIntructur); ok {
		r0 = rf(offset, limit, search, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*web.DashboardIntructur)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, uint) *web.Pagination); ok {
		r1 = rf(offset, limit, search, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*web.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(int, int, string, uint) error); ok {
		r2 = rf(offset, limit, search, userID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllCourseByUserId provides a mock function with given fields: offset, limit, search, userID
func (_m *CourseService) GetAllCourseByUserId(offset int, limit int, search string, userID uint) (*web.DashboardAllCourseIntructur, *web.Pagination, error) {
	ret := _m.Called(offset, limit, search, userID)

	var r0 *web.DashboardAllCourseIntructur
	var r1 *web.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string, uint) (*web.DashboardAllCourseIntructur, *web.Pagination, error)); ok {
		return rf(offset, limit, search, userID)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, uint) *web.DashboardAllCourseIntructur); ok {
		r0 = rf(offset, limit, search, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*web.DashboardAllCourseIntructur)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, uint) *web.Pagination); ok {
		r1 = rf(offset, limit, search, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*web.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(int, int, string, uint) error); ok {
		r2 = rf(offset, limit, search, userID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllMobile provides a mock function with given fields: offset, limit, search, category
func (_m *CourseService) GetAllMobile(offset int, limit int, search string, category string) ([]web.GetCourseResponseMobileNew, *web.Pagination, error) {
	ret := _m.Called(offset, limit, search, category)

	var r0 []web.GetCourseResponseMobileNew
	var r1 *web.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]web.GetCourseResponseMobileNew, *web.Pagination, error)); ok {
		return rf(offset, limit, search, category)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []web.GetCourseResponseMobileNew); ok {
		r0 = rf(offset, limit, search, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]web.GetCourseResponseMobileNew)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) *web.Pagination); ok {
		r1 = rf(offset, limit, search, category)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*web.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(int, int, string, string) error); ok {
		r2 = rf(offset, limit, search, category)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllMyCourse provides a mock function with given fields: offset, limit, search, category, userID
func (_m *CourseService) GetAllMyCourse(offset int, limit int, search string, category string, userID uint) ([]web.GetCourseResponseMobile, *web.Pagination, error) {
	ret := _m.Called(offset, limit, search, category, userID)

	var r0 []web.GetCourseResponseMobile
	var r1 *web.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string, string, uint) ([]web.GetCourseResponseMobile, *web.Pagination, error)); ok {
		return rf(offset, limit, search, category, userID)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string, uint) []web.GetCourseResponseMobile); ok {
		r0 = rf(offset, limit, search, category, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]web.GetCourseResponseMobile)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string, uint) *web.Pagination); ok {
		r1 = rf(offset, limit, search, category, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*web.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(int, int, string, string, uint) error); ok {
		r2 = rf(offset, limit, search, category, userID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetById provides a mock function with given fields: id
func (_m *CourseService) GetById(id uint) (*web.GetCourseResponseByIdWeb, error) {
	ret := _m.Called(id)

	var r0 *web.GetCourseResponseByIdWeb
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*web.GetCourseResponseByIdWeb, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *web.GetCourseResponseByIdWeb); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*web.GetCourseResponseByIdWeb)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIdMobile provides a mock function with given fields: id
func (_m *CourseService) GetByIdMobile(id uint) (*web.CourseForTraking, error) {
	ret := _m.Called(id)

	var r0 *web.CourseForTraking
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*web.CourseForTraking, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *web.CourseForTraking); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*web.CourseForTraking)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeatailCourse provides a mock function with given fields: id
func (_m *CourseService) GetDeatailCourse(id uint) (*web.CourseResponseForIntructur, error) {
	ret := _m.Called(id)

	var r0 *web.CourseResponseForIntructur
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*web.CourseResponseForIntructur, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *web.CourseResponseForIntructur); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*web.CourseResponseForIntructur)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, userId, role, request
func (_m *CourseService) Update(ctx echo.Context, id uint, userId uint, role string, request *web.CourseUpdateRequest) error {
	ret := _m.Called(ctx, id, userId, role, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, uint, uint, string, *web.CourseUpdateRequest) error); ok {
		r0 = rf(ctx, id, userId, role, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateImage provides a mock function with given fields: ctx, id, userId, role, request
func (_m *CourseService) UpdateImage(ctx echo.Context, id uint, userId uint, role string, request *web.CourseUpdateImageRequest) error {
	ret := _m.Called(ctx, id, userId, role, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, uint, uint, string, *web.CourseUpdateImageRequest) error); ok {
		r0 = rf(ctx, id, userId, role, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCourseService creates a new instance of CourseService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCourseService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CourseService {
	mock := &CourseService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
