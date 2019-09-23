package mocks

import (
	"github.com/Gustibimo/favetest/model"
	"github.com/stretchr/testify/mock"
)

type MerchantRepository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: cursor, limit
func (_m *MerchantRepository) Fetch(cursor string, limit int) ([]*model.Merchants, error) {
	ret := _m.Called(cursor, limit)

	var r0 []*model.Merchants
	if rf, ok := ret.Get(0).(func(string, int) []*model.Merchants); ok {
		r0 = rf(cursor, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Merchants)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(cursor, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MerchantRepository) Delete(id int) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *MerchantRepository) GetByID(id int) (*model.Merchants, error) {
	ret := _m.Called(id)

	var r0 *model.Merchants
	if rf, ok := ret.Get(0).(func(int) *model.Merchants); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Merchants)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: a
func (_m *MerchantRepository) Create(a *model.Merchants) (int, error) {
	ret := _m.Called(a)

	var r0 int
	if rf, ok := ret.Get(0).(func(*model.Merchants) int); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Merchants) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *MerchantRepository) Update(_a0 *model.Merchants) (*model.Merchants, error) {
	ret := _m.Called(_a0)

	var r0 *model.Merchants
	if rf, ok := ret.Get(0).(func(*model.Merchants) *model.Merchants); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Merchants)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Merchants) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
