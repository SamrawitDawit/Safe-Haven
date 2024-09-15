// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "backend/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UserRepositoryInterface is an autogenerated mock type for the UserRepositoryInterface type
type UserRepositoryInterface struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UserRepositoryInterface) CreateUser(user *domain.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByAnonymousDifferentiator provides a mock function with given fields: differentiator
func (_m *UserRepositoryInterface) GetUserByAnonymousDifferentiator(differentiator string) (*domain.User, error) {
	ret := _m.Called(differentiator)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByAnonymousDifferentiator")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(differentiator)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(differentiator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(differentiator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *UserRepositoryInterface) GetUserByEmail(email string) (*domain.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: id
func (_m *UserRepositoryInterface) GetUserByID(id uuid.UUID) (*domain.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*domain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByPhoneNumber provides a mock function with given fields: phoneNumber
func (_m *UserRepositoryInterface) GetUserByPhoneNumber(phoneNumber string) (*domain.User, error) {
	ret := _m.Called(phoneNumber)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByPhoneNumber")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(phoneNumber)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(phoneNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(phoneNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersCount provides a mock function with given fields:
func (_m *UserRepositoryInterface) GetUsersCount() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetUsersCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: user
func (_m *UserRepositoryInterface) UpdateUser(user *domain.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepositoryInterface creates a new instance of UserRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepositoryInterface {
	mock := &UserRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
