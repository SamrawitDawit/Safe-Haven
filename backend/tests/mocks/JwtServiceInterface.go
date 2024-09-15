// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "backend/domain"

	jwt "github.com/dgrijalva/jwt-go"

	mock "github.com/stretchr/testify/mock"
)

// JwtServiceInterface is an autogenerated mock type for the JwtServiceInterface type
type JwtServiceInterface struct {
	mock.Mock
}

// ExtractTokenClaims provides a mock function with given fields: token
func (_m *JwtServiceInterface) ExtractTokenClaims(token *jwt.Token) (jwt.MapClaims, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ExtractTokenClaims")
	}

	var r0 jwt.MapClaims
	var r1 error
	if rf, ok := ret.Get(0).(func(*jwt.Token) (jwt.MapClaims, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(*jwt.Token) jwt.MapClaims); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jwt.MapClaims)
		}
	}

	if rf, ok := ret.Get(1).(func(*jwt.Token) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateResetToken provides a mock function with given fields: email, code
func (_m *JwtServiceInterface) GenerateResetToken(email string, code int64) (string, error) {
	ret := _m.Called(email, code)

	if len(ret) == 0 {
		panic("no return value specified for GenerateResetToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64) (string, error)); ok {
		return rf(email, code)
	}
	if rf, ok := ret.Get(0).(func(string, int64) string); ok {
		r0 = rf(email, code)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(email, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateToken provides a mock function with given fields: user
func (_m *JwtServiceInterface) GenerateToken(user *domain.User) (string, string, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for GenerateToken")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(*domain.User) (string, string, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*domain.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*domain.User) string); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(*domain.User) error); ok {
		r2 = rf(user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ValidateToken provides a mock function with given fields: token
func (_m *JwtServiceInterface) ValidateToken(token string) (*jwt.Token, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateToken")
	}

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJwtServiceInterface creates a new instance of JwtServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJwtServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JwtServiceInterface {
	mock := &JwtServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}