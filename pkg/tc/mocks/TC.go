// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/Mellanox/multi-networkpolicy-tc/pkg/tc/types"
)

// TC is an autogenerated mock type for the TC type
type TC struct {
	mock.Mock
}

// ChainAdd provides a mock function with given fields: qdisc, chain
func (_m *TC) ChainAdd(qdisc types.QDisc, chain types.Chain) error {
	ret := _m.Called(qdisc, chain)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.QDisc, types.Chain) error); ok {
		r0 = rf(qdisc, chain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChainDel provides a mock function with given fields: qdisc, chain
func (_m *TC) ChainDel(qdisc types.QDisc, chain types.Chain) error {
	ret := _m.Called(qdisc, chain)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.QDisc, types.Chain) error); ok {
		r0 = rf(qdisc, chain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChainList provides a mock function with given fields: qdisc
func (_m *TC) ChainList(qdisc types.QDisc) ([]types.Chain, error) {
	ret := _m.Called(qdisc)

	var r0 []types.Chain
	if rf, ok := ret.Get(0).(func(types.QDisc) []types.Chain); ok {
		r0 = rf(qdisc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Chain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.QDisc) error); ok {
		r1 = rf(qdisc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterAdd provides a mock function with given fields: qdisc, filter
func (_m *TC) FilterAdd(qdisc types.QDisc, filter types.Filter) error {
	ret := _m.Called(qdisc, filter)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.QDisc, types.Filter) error); ok {
		r0 = rf(qdisc, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterDel provides a mock function with given fields: qdisc, filterAttr
func (_m *TC) FilterDel(qdisc types.QDisc, filterAttr *types.FilterAttrs) error {
	ret := _m.Called(qdisc, filterAttr)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.QDisc, *types.FilterAttrs) error); ok {
		r0 = rf(qdisc, filterAttr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterList provides a mock function with given fields: qdisc
func (_m *TC) FilterList(qdisc types.QDisc) ([]types.Filter, error) {
	ret := _m.Called(qdisc)

	var r0 []types.Filter
	if rf, ok := ret.Get(0).(func(types.QDisc) []types.Filter); ok {
		r0 = rf(qdisc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Filter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.QDisc) error); ok {
		r1 = rf(qdisc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QDiscAdd provides a mock function with given fields: qdisc
func (_m *TC) QDiscAdd(qdisc types.QDisc) error {
	ret := _m.Called(qdisc)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.QDisc) error); ok {
		r0 = rf(qdisc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QDiscDel provides a mock function with given fields: qdisc
func (_m *TC) QDiscDel(qdisc types.QDisc) error {
	ret := _m.Called(qdisc)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.QDisc) error); ok {
		r0 = rf(qdisc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QDiscList provides a mock function with given fields:
func (_m *TC) QDiscList() ([]types.QDisc, error) {
	ret := _m.Called()

	var r0 []types.QDisc
	if rf, ok := ret.Get(0).(func() []types.QDisc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.QDisc)
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

type mockConstructorTestingTNewTC interface {
	mock.TestingT
	Cleanup(func())
}

// NewTC creates a new instance of TC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTC(t mockConstructorTestingTNewTC) *TC {
	mock := &TC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}