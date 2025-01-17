// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/smartcontractkit/chainlink/core/store/models"
	mock "github.com/stretchr/testify/mock"
)

// JobSubscriber is an autogenerated mock type for the JobSubscriber type
type JobSubscriber struct {
	mock.Mock
}

// AddJob provides a mock function with given fields: job, bn
func (_m *JobSubscriber) AddJob(job models.JobSpec, bn *models.Head) error {
	ret := _m.Called(job, bn)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.JobSpec, *models.Head) error); ok {
		r0 = rf(job, bn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connect provides a mock function with given fields: head
func (_m *JobSubscriber) Connect(head *models.Head) error {
	ret := _m.Called(head)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Head) error); ok {
		r0 = rf(head)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Disconnect provides a mock function with given fields:
func (_m *JobSubscriber) Disconnect() {
	_m.Called()
}

// Jobs provides a mock function with given fields:
func (_m *JobSubscriber) Jobs() []models.JobSpec {
	ret := _m.Called()

	var r0 []models.JobSpec
	if rf, ok := ret.Get(0).(func() []models.JobSpec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.JobSpec)
		}
	}

	return r0
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *JobSubscriber) OnNewLongestChain(ctx context.Context, head models.Head) {
	_m.Called(ctx, head)
}

// RemoveJob provides a mock function with given fields: ID
func (_m *JobSubscriber) RemoveJob(ID models.JobID) error {
	ret := _m.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.JobID) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *JobSubscriber) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
