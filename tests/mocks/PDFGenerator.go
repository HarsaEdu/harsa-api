// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/HarsaEdu/harsa-api/internal/model/domain"
	mock "github.com/stretchr/testify/mock"
)

// PDFGenerator is an autogenerated mock type for the PDFGenerator type
type PDFGenerator struct {
	mock.Mock
}

// GenerateCertificate provides a mock function with given fields: certificate
func (_m *PDFGenerator) GenerateCertificate(certificate *domain.Certificate) ([]byte, error) {
	ret := _m.Called(certificate)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.Certificate) ([]byte, error)); ok {
		return rf(certificate)
	}
	if rf, ok := ret.Get(0).(func(*domain.Certificate) []byte); ok {
		r0 = rf(certificate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.Certificate) error); ok {
		r1 = rf(certificate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPDFGenerator creates a new instance of PDFGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPDFGenerator(t interface {
	mock.TestingT
	Cleanup(func())
}) *PDFGenerator {
	mock := &PDFGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
