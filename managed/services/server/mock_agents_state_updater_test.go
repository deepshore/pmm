// Code generated by mockery v1.0.0. DO NOT EDIT.

package server

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockAgentsStateUpdater is an autogenerated mock type for the agentsStateUpdater type
type mockAgentsStateUpdater struct {
	mock.Mock
}

// UpdateAgentsState provides a mock function with given fields: ctx
func (_m *mockAgentsStateUpdater) UpdateAgentsState(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}