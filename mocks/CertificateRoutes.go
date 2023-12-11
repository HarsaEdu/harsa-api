// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// CertificateRoutes is an autogenerated mock type for the CertificateRoutes type
type CertificateRoutes struct {
	mock.Mock
}

// CertificateMobile provides a mock function with given fields: apiGroup
func (_m *CertificateRoutes) CertificateMobile(apiGroup *echo.Group) {
	_m.Called(apiGroup)
}

// NewCertificateRoutes creates a new instance of CertificateRoutes. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCertificateRoutes(t interface {
	mock.TestingT
	Cleanup(func())
}) *CertificateRoutes {
	mock := &CertificateRoutes{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
