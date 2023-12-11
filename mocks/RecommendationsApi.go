// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	web "github.com/HarsaEdu/harsa-api/internal/model/web"
)

// RecommendationsApi is an autogenerated mock type for the RecommendationsApi type
type RecommendationsApi struct {
	mock.Mock
}

// GetRecommendations provides a mock function with given fields: request
func (_m *RecommendationsApi) GetRecommendations(request *web.GetRecommendationsRequest) (*web.GetRecommendationsResponse, error) {
	ret := _m.Called(request)

	var r0 *web.GetRecommendationsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*web.GetRecommendationsRequest) (*web.GetRecommendationsResponse, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(*web.GetRecommendationsRequest) *web.GetRecommendationsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*web.GetRecommendationsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*web.GetRecommendationsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRecommendationsApi creates a new instance of RecommendationsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRecommendationsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *RecommendationsApi {
	mock := &RecommendationsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
