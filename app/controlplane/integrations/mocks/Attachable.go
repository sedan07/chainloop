// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	integrations "github.com/chainloop-dev/chainloop/app/controlplane/integrations"
	mock "github.com/stretchr/testify/mock"
)

// Attachable is an autogenerated mock type for the Attachable type
type Attachable struct {
	mock.Mock
}

// PreAttach provides a mock function with given fields: ctx, c
func (_m *Attachable) PreAttach(ctx context.Context, c *integrations.BundledConfig) (*integrations.PreAttachment, error) {
	ret := _m.Called(ctx, c)

	var r0 *integrations.PreAttachment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *integrations.BundledConfig) (*integrations.PreAttachment, error)); ok {
		return rf(ctx, c)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *integrations.BundledConfig) *integrations.PreAttachment); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*integrations.PreAttachment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *integrations.BundledConfig) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAttachable interface {
	mock.TestingT
	Cleanup(func())
}

// NewAttachable creates a new instance of Attachable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAttachable(t mockConstructorTestingTNewAttachable) *Attachable {
	mock := &Attachable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
