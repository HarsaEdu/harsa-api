// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/HarsaEdu/harsa-api/internal/model/domain"
	mock "github.com/stretchr/testify/mock"

	web "github.com/HarsaEdu/harsa-api/internal/model/web"
)

// HistoryQuizRepository is an autogenerated mock type for the HistoryQuizRepository type
type HistoryQuizRepository struct {
	mock.Mock
}

// Answering provides a mock function with given fields: userId, answers, quizId
func (_m *HistoryQuizRepository) Answering(userId uint, answers []web.AnswersRequest, quizId uint) error {
	ret := _m.Called(userId, answers, quizId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, []web.AnswersRequest, uint) error); ok {
		r0 = rf(userId, answers, quizId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Cek provides a mock function with given fields: userId, quizID
func (_m *HistoryQuizRepository) Cek(userId uint, quizID uint) (*domain.HistoryQuiz, error) {
	ret := _m.Called(userId, quizID)

	var r0 *domain.HistoryQuiz
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (*domain.HistoryQuiz, error)); ok {
		return rf(userId, quizID)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) *domain.HistoryQuiz); ok {
		r0 = rf(userId, quizID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.HistoryQuiz)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userId, quizID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllHistoryQuizByQuizId provides a mock function with given fields: offset, limit, quizID, search
func (_m *HistoryQuizRepository) GetAllHistoryQuizByQuizId(offset int, limit int, quizID uint, search string) ([]domain.HistoryQuiz, int64, error) {
	ret := _m.Called(offset, limit, quizID, search)

	var r0 []domain.HistoryQuiz
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, uint, string) ([]domain.HistoryQuiz, int64, error)); ok {
		return rf(offset, limit, quizID, search)
	}
	if rf, ok := ret.Get(0).(func(int, int, uint, string) []domain.HistoryQuiz); ok {
		r0 = rf(offset, limit, quizID, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.HistoryQuiz)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, uint, string) int64); ok {
		r1 = rf(offset, limit, quizID, search)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, uint, string) error); ok {
		r2 = rf(offset, limit, quizID, search)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetById provides a mock function with given fields: id
func (_m *HistoryQuizRepository) GetById(id uint) (*domain.HistoryQuiz, error) {
	ret := _m.Called(id)

	var r0 *domain.HistoryQuiz
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*domain.HistoryQuiz, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *domain.HistoryQuiz); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.HistoryQuiz)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMyHistoryQuizByQuizId provides a mock function with given fields: userID, quizID
func (_m *HistoryQuizRepository) GetMyHistoryQuizByQuizId(userID uint, quizID uint) (*domain.HistoryQuiz, error) {
	ret := _m.Called(userID, quizID)

	var r0 *domain.HistoryQuiz
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (*domain.HistoryQuiz, error)); ok {
		return rf(userID, quizID)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) *domain.HistoryQuiz); ok {
		r0 = rf(userID, quizID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.HistoryQuiz)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, quizID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewHistoryQuizRepository creates a new instance of HistoryQuizRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHistoryQuizRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *HistoryQuizRepository {
	mock := &HistoryQuizRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}