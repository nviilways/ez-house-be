// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	entity "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetUserByID provides a mock function with given fields: _a0
func (_m *UserRepository) GetUserByID(_a0 uint) (*entity.User, error) {
	ret := _m.Called(_a0)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(uint) *entity.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignIn provides a mock function with given fields: _a0
func (_m *UserRepository) SignIn(_a0 *entity.User) (*entity.User, error) {
	ret := _m.Called(_a0)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(*entity.User) *entity.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignOut provides a mock function with given fields: _a0
func (_m *UserRepository) SignOut(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignUp provides a mock function with given fields: _a0
func (_m *UserRepository) SignUp(_a0 *entity.User) (*entity.User, error) {
	ret := _m.Called(_a0)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(*entity.User) *entity.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokenCheck provides a mock function with given fields: _a0
func (_m *UserRepository) TokenCheck(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *UserRepository) Update(_a0 *entity.User) (*entity.User, error) {
	ret := _m.Called(_a0)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(*entity.User) *entity.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
