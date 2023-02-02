// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dto "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	entity "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// HouseUsecase is an autogenerated mock type for the HouseUsecase type
type HouseUsecase struct {
	mock.Mock
}

// AddHouse provides a mock function with given fields: _a0, _a1
func (_m *HouseUsecase) AddHouse(_a0 *entity.House, _a1 []*multipart.FileHeader) (*entity.House, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.House
	if rf, ok := ret.Get(0).(func(*entity.House, []*multipart.FileHeader) *entity.House); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.House)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.House, []*multipart.FileHeader) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddPhotoHouse provides a mock function with given fields: _a0, _a1
func (_m *HouseUsecase) AddPhotoHouse(_a0 *entity.Photo, _a1 []*multipart.FileHeader) (*entity.Photo, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.Photo
	if rf, ok := ret.Get(0).(func(*entity.Photo, []*multipart.FileHeader) *entity.Photo); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Photo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Photo, []*multipart.FileHeader) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteHouse provides a mock function with given fields: _a0, _a1
func (_m *HouseUsecase) DeleteHouse(_a0 uint, _a1 uint) (*entity.House, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.House
	if rf, ok := ret.Get(0).(func(uint, uint) *entity.House); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.House)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCityList provides a mock function with given fields:
func (_m *HouseUsecase) GetCityList() ([]*entity.City, error) {
	ret := _m.Called()

	var r0 []*entity.City
	if rf, ok := ret.Get(0).(func() []*entity.City); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.City)
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

// GetHouseByHost provides a mock function with given fields: _a0
func (_m *HouseUsecase) GetHouseByHost(_a0 uint) ([]*entity.House, error) {
	ret := _m.Called(_a0)

	var r0 []*entity.House
	if rf, ok := ret.Get(0).(func(uint) []*entity.House); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.House)
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

// GetHouseByID provides a mock function with given fields: _a0
func (_m *HouseUsecase) GetHouseByID(_a0 uint) (*entity.House, error) {
	ret := _m.Called(_a0)

	var r0 *entity.House
	if rf, ok := ret.Get(0).(func(uint) *entity.House); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.House)
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

// GetHouseListByVacancy provides a mock function with given fields: _a0, _a1
func (_m *HouseUsecase) GetHouseListByVacancy(_a0 *dto.FilterHouse, _a1 *dto.Pagination) ([]*entity.House, int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*entity.House
	if rf, ok := ret.Get(0).(func(*dto.FilterHouse, *dto.Pagination) []*entity.House); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.House)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*dto.FilterHouse, *dto.Pagination) int); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*dto.FilterHouse, *dto.Pagination) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateHouse provides a mock function with given fields: _a0, _a1
func (_m *HouseUsecase) UpdateHouse(_a0 uint, _a1 *entity.House) (*entity.House, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.House
	if rf, ok := ret.Get(0).(func(uint, *entity.House) *entity.House); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.House)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *entity.House) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHouseUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewHouseUsecase creates a new instance of HouseUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHouseUsecase(t mockConstructorTestingTNewHouseUsecase) *HouseUsecase {
	mock := &HouseUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}