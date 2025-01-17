// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	assets "github.com/smartcontractkit/chainlink/core/assets"
	fluxmonitor "github.com/smartcontractkit/chainlink/core/services/fluxmonitor"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	orm "github.com/smartcontractkit/chainlink/core/store/orm"
)

// DeviationCheckerFactory is an autogenerated mock type for the DeviationCheckerFactory type
type DeviationCheckerFactory struct {
	mock.Mock
}

// New provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *DeviationCheckerFactory) New(_a0 models.Initiator, _a1 *assets.Link, _a2 fluxmonitor.RunManager, _a3 *orm.ORM, _a4 models.Duration) (fluxmonitor.DeviationChecker, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 fluxmonitor.DeviationChecker
	if rf, ok := ret.Get(0).(func(models.Initiator, *assets.Link, fluxmonitor.RunManager, *orm.ORM, models.Duration) fluxmonitor.DeviationChecker); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fluxmonitor.DeviationChecker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Initiator, *assets.Link, fluxmonitor.RunManager, *orm.ORM, models.Duration) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
