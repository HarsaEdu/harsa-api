// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/HarsaEdu/harsa-api/internal/model/domain"
	mock "github.com/stretchr/testify/mock"
)

// PaymentRepository is an autogenerated mock type for the PaymentRepository type
type PaymentRepository struct {
	mock.Mock
}

// CreatePaymentHistory provides a mock function with given fields: paymentHistory
func (_m *PaymentRepository) CreatePaymentHistory(paymentHistory *domain.PaymentHistory) error {
	ret := _m.Called(paymentHistory)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.PaymentHistory) error); ok {
		r0 = rf(paymentHistory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPaymentHistory provides a mock function with given fields: offset, limit, search, status
func (_m *PaymentRepository) GetAllPaymentHistory(offset int, limit int, search string, status string) ([]domain.PaymentHistory, int64, error) {
	ret := _m.Called(offset, limit, search, status)

	var r0 []domain.PaymentHistory
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]domain.PaymentHistory, int64, error)); ok {
		return rf(offset, limit, search, status)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []domain.PaymentHistory); ok {
		r0 = rf(offset, limit, search, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PaymentHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) int64); ok {
		r1 = rf(offset, limit, search, status)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string, string) error); ok {
		r2 = rf(offset, limit, search, status)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllPaymentHistoryByUserId provides a mock function with given fields: userId, offset, limit, search, status
func (_m *PaymentRepository) GetAllPaymentHistoryByUserId(userId uint, offset int, limit int, search string, status string) ([]domain.PaymentHistory, int64, error) {
	ret := _m.Called(userId, offset, limit, search, status)

	var r0 []domain.PaymentHistory
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(uint, int, int, string, string) ([]domain.PaymentHistory, int64, error)); ok {
		return rf(userId, offset, limit, search, status)
	}
	if rf, ok := ret.Get(0).(func(uint, int, int, string, string) []domain.PaymentHistory); ok {
		r0 = rf(userId, offset, limit, search, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PaymentHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, int, int, string, string) int64); ok {
		r1 = rf(userId, offset, limit, search, status)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(uint, int, int, string, string) error); ok {
		r2 = rf(userId, offset, limit, search, status)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPaymentHistoryById provides a mock function with given fields: paymentHistoryId
func (_m *PaymentRepository) GetPaymentHistoryById(paymentHistoryId string) (*domain.PaymentHistory, error) {
	ret := _m.Called(paymentHistoryId)

	var r0 *domain.PaymentHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.PaymentHistory, error)); ok {
		return rf(paymentHistoryId)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.PaymentHistory); ok {
		r0 = rf(paymentHistoryId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(paymentHistoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentHistoryByUserIdAndPaymentId provides a mock function with given fields: userId, paymentId
func (_m *PaymentRepository) GetPaymentHistoryByUserIdAndPaymentId(userId uint, paymentId string) (*domain.PaymentHistory, error) {
	ret := _m.Called(userId, paymentId)

	var r0 *domain.PaymentHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string) (*domain.PaymentHistory, error)); ok {
		return rf(userId, paymentId)
	}
	if rf, ok := ret.Get(0).(func(uint, string) *domain.PaymentHistory); ok {
		r0 = rf(userId, paymentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(userId, paymentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatusPaymentHistory provides a mock function with given fields: status, transactionResult
func (_m *PaymentRepository) UpdateStatusPaymentHistory(status string, transactionResult *domain.PaymentTransactionStatus) error {
	ret := _m.Called(status, transactionResult)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *domain.PaymentTransactionStatus) error); ok {
		r0 = rf(status, transactionResult)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPaymentRepository creates a new instance of PaymentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentRepository {
	mock := &PaymentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}