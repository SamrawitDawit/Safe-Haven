// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "backend/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// CaseRepositoryInterface is an autogenerated mock type for the CaseRepositoryInterface type
type CaseRepositoryInterface struct {
	mock.Mock
}

// CreateCase provides a mock function with given fields: Case
func (_m *CaseRepositoryInterface) CreateCase(Case *domain.Case) *domain.CustomError {
	ret := _m.Called(Case)

	if len(ret) == 0 {
		panic("no return value specified for CreateCase")
	}

	var r0 *domain.CustomError
	if rf, ok := ret.Get(0).(func(*domain.Case) *domain.CustomError); ok {
		r0 = rf(Case)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CustomError)
		}
	}

	return r0
}

// DeleteCase provides a mock function with given fields: CaseID
func (_m *CaseRepositoryInterface) DeleteCase(CaseID uuid.UUID) *domain.CustomError {
	ret := _m.Called(CaseID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCase")
	}

	var r0 *domain.CustomError
	if rf, ok := ret.Get(0).(func(uuid.UUID) *domain.CustomError); ok {
		r0 = rf(CaseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CustomError)
		}
	}

	return r0
}

// GetAllCases provides a mock function with given fields:
func (_m *CaseRepositoryInterface) GetAllCases() ([]*domain.Case, *domain.CustomError) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllCases")
	}

	var r0 []*domain.Case
	var r1 *domain.CustomError
	if rf, ok := ret.Get(0).(func() ([]*domain.Case, *domain.CustomError)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*domain.Case); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Case)
		}
	}

	if rf, ok := ret.Get(1).(func() *domain.CustomError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.CustomError)
		}
	}

	return r0, r1
}

// GetCaseByID provides a mock function with given fields: CaseID
func (_m *CaseRepositoryInterface) GetCaseByID(CaseID uuid.UUID) (*domain.Case, *domain.CustomError) {
	ret := _m.Called(CaseID)

	if len(ret) == 0 {
		panic("no return value specified for GetCaseByID")
	}

	var r0 *domain.Case
	var r1 *domain.CustomError
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*domain.Case, *domain.CustomError)); ok {
		return rf(CaseID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *domain.Case); ok {
		r0 = rf(CaseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Case)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) *domain.CustomError); ok {
		r1 = rf(CaseID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.CustomError)
		}
	}

	return r0, r1
}

// GetCasesByCounselorID provides a mock function with given fields: counselorID
func (_m *CaseRepositoryInterface) GetCasesByCounselorID(counselorID uuid.UUID) ([]*domain.Case, *domain.CustomError) {
	ret := _m.Called(counselorID)

	if len(ret) == 0 {
		panic("no return value specified for GetCasesByCounselorID")
	}

	var r0 []*domain.Case
	var r1 *domain.CustomError
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]*domain.Case, *domain.CustomError)); ok {
		return rf(counselorID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*domain.Case); ok {
		r0 = rf(counselorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Case)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) *domain.CustomError); ok {
		r1 = rf(counselorID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.CustomError)
		}
	}

	return r0, r1
}

// GetCasesByStatus provides a mock function with given fields: status
func (_m *CaseRepositoryInterface) GetCasesByStatus(status string) ([]*domain.Case, *domain.CustomError) {
	ret := _m.Called(status)

	if len(ret) == 0 {
		panic("no return value specified for GetCasesByStatus")
	}

	var r0 []*domain.Case
	var r1 *domain.CustomError
	if rf, ok := ret.Get(0).(func(string) ([]*domain.Case, *domain.CustomError)); ok {
		return rf(status)
	}
	if rf, ok := ret.Get(0).(func(string) []*domain.Case); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Case)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *domain.CustomError); ok {
		r1 = rf(status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.CustomError)
		}
	}

	return r0, r1
}

// GetCasesBySubmitterID provides a mock function with given fields: SubmitterID
func (_m *CaseRepositoryInterface) GetCasesBySubmitterID(SubmitterID uuid.UUID) ([]*domain.Case, *domain.CustomError) {
	ret := _m.Called(SubmitterID)

	if len(ret) == 0 {
		panic("no return value specified for GetCasesBySubmitterID")
	}

	var r0 []*domain.Case
	var r1 *domain.CustomError
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]*domain.Case, *domain.CustomError)); ok {
		return rf(SubmitterID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*domain.Case); ok {
		r0 = rf(SubmitterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Case)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) *domain.CustomError); ok {
		r1 = rf(SubmitterID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.CustomError)
		}
	}

	return r0, r1
}

// UpdateCaseFields provides a mock function with given fields: CaseID, fields
func (_m *CaseRepositoryInterface) UpdateCaseFields(CaseID uuid.UUID, fields map[string]interface{}) *domain.CustomError {
	ret := _m.Called(CaseID, fields)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCaseFields")
	}

	var r0 *domain.CustomError
	if rf, ok := ret.Get(0).(func(uuid.UUID, map[string]interface{}) *domain.CustomError); ok {
		r0 = rf(CaseID, fields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CustomError)
		}
	}

	return r0
}

// NewCaseRepositoryInterface creates a new instance of CaseRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCaseRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CaseRepositoryInterface {
	mock := &CaseRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
