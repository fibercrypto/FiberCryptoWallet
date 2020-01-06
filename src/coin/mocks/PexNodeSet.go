// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	core "github.com/fibercrypto/fibercryptowallet/src/core"
	mock "github.com/stretchr/testify/mock"
)

// PexNodeSet is an autogenerated mock type for the PexNodeSet type
type PexNodeSet struct {
	mock.Mock
}

// ListPeers provides a mock function with given fields:
func (_m *PexNodeSet) ListPeers() core.PexNodeIterator {
	ret := _m.Called()

	var r0 core.PexNodeIterator
	if rf, ok := ret.Get(0).(func() core.PexNodeIterator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.PexNodeIterator)
		}
	}

	return r0
}
