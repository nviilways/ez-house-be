// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dto "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	entity "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"

	mock "github.com/stretchr/testify/mock"
)

// PickupRepository is an autogenerated mock type for the PickupRepository type
type PickupRepository struct {
	mock.Mock
}

// GetPickupById provides a mock function with given fields: _a0
func (_m *PickupRepository) GetPickupById(_a0 uint) (*entity.Pickup, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Pickup
	if rf, ok := ret.Get(0).(func(uint) *entity.Pickup); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Pickup)
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

// GetPickupList provides a mock function with given fields:
func (_m *PickupRepository) GetPickupList() ([]*entity.Pickup, error) {
	ret := _m.Called()

	var r0 []*entity.Pickup
	if rf, ok := ret.Get(0).(func() []*entity.Pickup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Pickup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPickupPrice provides a mock function with given fields: _a0
func (_m *PickupRepository) GetPickupPrice(_a0 *entity.Reservation) (*dto.PickupPrice, error) {
	ret := _m.Called(_a0)

	var r0 *dto.PickupPrice
	if rf, ok := ret.Get(0).(func(*entity.Reservation) *dto.PickupPrice); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.PickupPrice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Reservation) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestPickup provides a mock function with given fields: _a0
func (_m *PickupRepository) RequestPickup(_a0 *entity.Pickup) (*entity.Pickup, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Pickup
	if rf, ok := ret.Get(0).(func(*entity.Pickup) *entity.Pickup); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Pickup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Pickup) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: _a0
func (_m *PickupRepository) UpdateStatus(_a0 uint) (*entity.Pickup, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Pickup
	if rf, ok := ret.Get(0).(func(uint) *entity.Pickup); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Pickup)
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

type mockConstructorTestingTNewPickupRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPickupRepository creates a new instance of PickupRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPickupRepository(t mockConstructorTestingTNewPickupRepository) *PickupRepository {
	mock := &PickupRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
