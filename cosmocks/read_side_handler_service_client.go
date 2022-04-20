// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	context "context"

	chief_of_statev1 "github.com/chief-of-state/cos-go-sdk/cospb/chief_of_state/v1"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ReadSideHandlerServiceClient is an autogenerated mock type for the ReadSideHandlerServiceClient type
type ReadSideHandlerServiceClient struct {
	mock.Mock
}

// HandleReadSide provides a mock function with given fields: ctx, in, opts
func (_m *ReadSideHandlerServiceClient) HandleReadSide(ctx context.Context, in *chief_of_statev1.HandleReadSideRequest, opts ...grpc.CallOption) (*chief_of_statev1.HandleReadSideResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chief_of_statev1.HandleReadSideResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.HandleReadSideRequest, ...grpc.CallOption) *chief_of_statev1.HandleReadSideResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.HandleReadSideResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.HandleReadSideRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewReadSideHandlerServiceClient creates a new instance of ReadSideHandlerServiceClient. It also registers a cleanup function to assert the mocks expectations.
func NewReadSideHandlerServiceClient(t testing.TB) *ReadSideHandlerServiceClient {
	mock := &ReadSideHandlerServiceClient{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}