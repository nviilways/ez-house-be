// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dto "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	entity "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"

	mock "github.com/stretchr/testify/mock"
)

// AdminUsecase is an autogenerated mock type for the AdminUsecase type
type AdminUsecase struct {
	mock.Mock
}

// SignIn provides a mock function with given fields: _a0, _a1
func (_m *AdminUsecase) SignIn(_a0 string, _a1 *entity.Admin) (*dto.Token, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dto.Token
	if rf, ok := ret.Get(0).(func(string, *entity.Admin) *dto.Token); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *entity.Admin) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: _a0
func (_m *AdminUsecase) SignUp(_a0 *entity.Admin) (*entity.Admin, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Admin
	if rf, ok := ret.Get(0).(func(*entity.Admin) *entity.Admin); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Admin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Admin) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAdminUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAdminUsecase creates a new instance of AdminUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAdminUsecase(t mockConstructorTestingTNewAdminUsecase) *AdminUsecase {
	mock := &AdminUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
