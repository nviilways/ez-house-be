// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dto "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	entity "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"

	mock "github.com/stretchr/testify/mock"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// GetTransactionByWalletId provides a mock function with given fields: _a0, _a1
func (_m *TransactionRepository) GetTransactionByWalletId(_a0 uint, _a1 *dto.Pagination) ([]*entity.Transaction, int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*entity.Transaction
	if rf, ok := ret.Get(0).(func(uint, *dto.Pagination) []*entity.Transaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Transaction)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(uint, *dto.Pagination) int); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint, *dto.Pagination) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TopUp provides a mock function with given fields: _a0, _a1
func (_m *TransactionRepository) TopUp(_a0 int, _a1 *entity.Transaction) (*entity.Transaction, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.Transaction
	if rf, ok := ret.Get(0).(func(int, *entity.Transaction) *entity.Transaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *entity.Transaction) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionRepository(t mockConstructorTestingTNewTransactionRepository) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
