// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	core "github.com/fibercrypto/fibercryptowallet/src/core"
	mock "github.com/stretchr/testify/mock"
)

// TransactionInputIterator is an autogenerated mock type for the TransactionInputIterator type
type TransactionInputIterator struct {
	mock.Mock
}

// HasNext provides a mock function with given fields:
func (_m *TransactionInputIterator) HasNext() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *TransactionInputIterator) Next() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Value provides a mock function with given fields:
func (_m *TransactionInputIterator) Value() core.TransactionInput {
	ret := _m.Called()

	var r0 core.TransactionInput
	if rf, ok := ret.Get(0).(func() core.TransactionInput); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.TransactionInput)
		}
	}

	return r0
}
