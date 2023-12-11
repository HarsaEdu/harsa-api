// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/HarsaEdu/harsa-api/internal/model/domain"
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"

	web "github.com/HarsaEdu/harsa-api/internal/model/web"
)

// SubsPlanService is an autogenerated mock type for the SubsPlanService type
type SubsPlanService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, subsPlan
func (_m *SubsPlanService) Create(ctx echo.Context, subsPlan *web.SubsPlanCreateRequest) error {
	ret := _m.Called(ctx, subsPlan)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, *web.SubsPlanCreateRequest) error); ok {
		r0 = rf(ctx, subsPlan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *SubsPlanService) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: id
func (_m *SubsPlanService) FindById(id int) (*domain.SubsPlan, error) {
	ret := _m.Called(id)

	var r0 *domain.SubsPlan
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*domain.SubsPlan, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *domain.SubsPlan); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.SubsPlan)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: search
func (_m *SubsPlanService) GetAll(search string) ([]domain.SubsPlan, *web.Pagination, error) {
	ret := _m.Called(search)

	var r0 []domain.SubsPlan
	var r1 *web.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(string) ([]domain.SubsPlan, *web.Pagination, error)); ok {
		return rf(search)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.SubsPlan); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SubsPlan)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *web.Pagination); ok {
		r1 = rf(search)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*web.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(search)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: subsPlan, id
func (_m *SubsPlanService) Update(subsPlan *web.SubsPlanUpdateRequest, id int) error {
	ret := _m.Called(subsPlan, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*web.SubsPlanUpdateRequest, int) error); ok {
		r0 = rf(subsPlan, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateImage provides a mock function with given fields: ctx, subsPlan, id
func (_m *SubsPlanService) UpdateImage(ctx echo.Context, subsPlan *web.SubsPlanUpdateImage, id int) error {
	ret := _m.Called(ctx, subsPlan, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, *web.SubsPlanUpdateImage, int) error); ok {
		r0 = rf(ctx, subsPlan, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSubsPlanService creates a new instance of SubsPlanService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubsPlanService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubsPlanService {
	mock := &SubsPlanService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}