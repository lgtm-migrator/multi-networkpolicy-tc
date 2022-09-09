// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	tc "github.com/Mellanox/multi-networkpolicy-tc/pkg/tc"
	mock "github.com/stretchr/testify/mock"
)

// Actuator is an autogenerated mock type for the Actuator type
type Actuator struct {
	mock.Mock
}

// Actuate provides a mock function with given fields: objects
func (_m *Actuator) Actuate(objects *tc.TCObjects) error {
	ret := _m.Called(objects)

	var r0 error
	if rf, ok := ret.Get(0).(func(*tc.TCObjects) error); ok {
		r0 = rf(objects)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewActuator interface {
	mock.TestingT
	Cleanup(func())
}

// NewActuator creates a new instance of Actuator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewActuator(t mockConstructorTestingTNewActuator) *Actuator {
	mock := &Actuator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}