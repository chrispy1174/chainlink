// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	postgres "github.com/smartcontractkit/chainlink/core/services/postgres"
	mock "github.com/stretchr/testify/mock"
)

// Subscription is an autogenerated mock type for the Subscription type
type Subscription struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Subscription) Close() {
	_m.Called()
}

// Events provides a mock function with given fields:
func (_m *Subscription) Events() <-chan postgres.Event {
	ret := _m.Called()

	var r0 <-chan postgres.Event
	if rf, ok := ret.Get(0).(func() <-chan postgres.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan postgres.Event)
		}
	}

	return r0
}

// channelName provides a mock function with given fields:
func (_m *Subscription) channelName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// interestedIn provides a mock function with given fields: event
func (_m *Subscription) interestedIn(event postgres.Event) bool {
	ret := _m.Called(event)

	var r0 bool
	if rf, ok := ret.Get(0).(func(postgres.Event) bool); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// send provides a mock function with given fields: event
func (_m *Subscription) send(event postgres.Event) {
	_m.Called(event)
}
