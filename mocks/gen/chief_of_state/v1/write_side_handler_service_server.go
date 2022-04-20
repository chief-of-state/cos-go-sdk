// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	context "context"

	chief_of_statev1 "github.com/chief-of-state/cos-go-sdk/gen/chief_of_state/v1"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// WriteSideHandlerServiceServer is an autogenerated mock type for the WriteSideHandlerServiceServer type
type WriteSideHandlerServiceServer struct {
	mock.Mock
}

// HandleCommand provides a mock function with given fields: _a0, _a1
func (_m *WriteSideHandlerServiceServer) HandleCommand(_a0 context.Context, _a1 *chief_of_statev1.HandleCommandRequest) (*chief_of_statev1.HandleCommandResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *chief_of_statev1.HandleCommandResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.HandleCommandRequest) *chief_of_statev1.HandleCommandResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.HandleCommandResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.HandleCommandRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleEvent provides a mock function with given fields: _a0, _a1
func (_m *WriteSideHandlerServiceServer) HandleEvent(_a0 context.Context, _a1 *chief_of_statev1.HandleEventRequest) (*chief_of_statev1.HandleEventResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *chief_of_statev1.HandleEventResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.HandleEventRequest) *chief_of_statev1.HandleEventResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.HandleEventResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.HandleEventRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewWriteSideHandlerServiceServer creates a new instance of WriteSideHandlerServiceServer. It also registers a cleanup function to assert the mocks expectations.
func NewWriteSideHandlerServiceServer(t testing.TB) *WriteSideHandlerServiceServer {
	mock := &WriteSideHandlerServiceServer{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
