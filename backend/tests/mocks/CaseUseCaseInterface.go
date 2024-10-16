// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "backend/domain"
	dto "backend/usecases/dto"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// CaseUseCaseInterface is an autogenerated mock type for the CaseUseCaseInterface type
type CaseUseCaseInterface struct {
	mock.Mock
}

// CreateCase provides a mock function with given fields: CaseDto
func (_m *CaseUseCaseInterface) CreateCase(CaseDto dto.CaseDto) (*domain.Case, *domain.CustomError) {
	ret := _m.Called(CaseDto)

	if len(ret) == 0 {
		panic("no return value specified for CreateCase")
	}

	var r0 *domain.Case
	var r1 *domain.CustomError
	if rf, ok := ret.Get(0).(func(dto.CaseDto) (*domain.Case, *domain.CustomError)); ok {
		return rf(CaseDto)
	}
	if rf, ok := ret.Get(0).(func(dto.CaseDto) *domain.Case); ok {
		r0 = rf(CaseDto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Case)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.CaseDto) *domain.CustomError); ok {
		r1 = rf(CaseDto)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.CustomError)
		}
	}

	return r0, r1
}

// Decrypt provides a mock function with given fields: Cases
func (_m *CaseUseCaseInterface) Decrypt(Cases []*domain.Case) ([]*domain.Case, *domain.CustomError) {
	ret := _m.Called(Cases)

	if len(ret) == 0 {
		panic("no return value specified for Decrypt")
	}

	var r0 []*domain.Case
	var r1 *domain.CustomError
	if rf, ok := ret.Get(0).(func([]*domain.Case) ([]*domain.Case, *domain.CustomError)); ok {
		return rf(Cases)
	}
	if rf, ok := ret.Get(0).(func([]*domain.Case) []*domain.Case); ok {
		r0 = rf(Cases)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Case)
		}
	}

	if rf, ok := ret.Get(1).(func([]*domain.Case) *domain.CustomError); ok {
		r1 = rf(Cases)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.CustomError)
		}
	}

	return r0, r1
}

// DecryptField provides a mock function with given fields: field
func (_m *CaseUseCaseInterface) DecryptField(field string) (string, *domain.CustomError) {
	ret := _m.Called(field)

	if len(ret) == 0 {
		panic("no return value specified for DecryptField")
	}

	var r0 string
	var r1 *domain.CustomError
	if rf, ok := ret.Get(0).(func(string) (string, *domain.CustomError)); ok {
		return rf(field)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(field)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) *domain.CustomError); ok {
		r1 = rf(field)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.CustomError)
		}
	}

	return r0, r1
}

// DeleteCase provides a mock function with given fields: CaseID
func (_m *CaseUseCaseInterface) DeleteCase(CaseID uuid.UUID) *domain.CustomError {
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
func (_m *CaseUseCaseInterface) GetAllCases() ([]*domain.Case, *domain.CustomError) {
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
func (_m *CaseUseCaseInterface) GetCaseByID(CaseID uuid.UUID) (*domain.Case, *domain.CustomError) {
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

// GetCasesByCounselorID provides a mock function with given fields: counselor
func (_m *CaseUseCaseInterface) GetCasesByCounselorID(counselor uuid.UUID) ([]*domain.Case, *domain.CustomError) {
	ret := _m.Called(counselor)

	if len(ret) == 0 {
		panic("no return value specified for GetCasesByCounselorID")
	}

	var r0 []*domain.Case
	var r1 *domain.CustomError
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]*domain.Case, *domain.CustomError)); ok {
		return rf(counselor)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*domain.Case); ok {
		r0 = rf(counselor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Case)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) *domain.CustomError); ok {
		r1 = rf(counselor)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.CustomError)
		}
	}

	return r0, r1
}

// GetCasesByStatus provides a mock function with given fields: status
func (_m *CaseUseCaseInterface) GetCasesByStatus(status string) ([]*domain.Case, *domain.CustomError) {
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
func (_m *CaseUseCaseInterface) GetCasesBySubmitterID(SubmitterID uuid.UUID) ([]*domain.Case, *domain.CustomError) {
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

// UpdateCase provides a mock function with given fields: caseID, CaseDto
func (_m *CaseUseCaseInterface) UpdateCase(caseID uuid.UUID, CaseDto dto.CaseDto) *domain.CustomError {
	ret := _m.Called(caseID, CaseDto)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCase")
	}

	var r0 *domain.CustomError
	if rf, ok := ret.Get(0).(func(uuid.UUID, dto.CaseDto) *domain.CustomError); ok {
		r0 = rf(caseID, CaseDto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CustomError)
		}
	}

	return r0
}

// ValidateCaseDto provides a mock function with given fields: CaseDto
func (_m *CaseUseCaseInterface) ValidateCaseDto(CaseDto dto.CaseDto) *domain.CustomError {
	ret := _m.Called(CaseDto)

	if len(ret) == 0 {
		panic("no return value specified for ValidateCaseDto")
	}

	var r0 *domain.CustomError
	if rf, ok := ret.Get(0).(func(dto.CaseDto) *domain.CustomError); ok {
		r0 = rf(CaseDto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CustomError)
		}
	}

	return r0
}

// NewCaseUseCaseInterface creates a new instance of CaseUseCaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCaseUseCaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CaseUseCaseInterface {
	mock := &CaseUseCaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
