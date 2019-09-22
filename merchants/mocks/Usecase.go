package mocks

import (
	"github.com/Gustibimo/favetest/model"
	"github.com/stretchr/testify/mock"
)

type MerchantImpl struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *MerchantImpl) Delete(id int) error {
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

	return r1
}

// Fetch provides a mock function with given fields: cursor, limit
func (_m *MerchantImpl) Fetch(cursor string, limit int) ([]*model.Merchants, string, error) {
	ret := _m.Called(cursor, limit)

	var r0 []*model.Merchants
	if rf, ok := ret.Get(0).(func(string, int) []*model.Merchants); ok {
		r0 = rf(cursor, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Merchants)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, int) string); ok {
		r1 = rf(cursor, limit)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, int) error); ok {
		r2 = rf(cursor, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: id
func (_m *MerchantImpl) GetByID(id int) (*model.Merchants, error) {
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
