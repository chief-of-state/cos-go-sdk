// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	context "context"

	chief_of_statev1 "github.com/chief-of-state/cos-go-sdk/cospb/chief_of_state/v1"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// ReadSideManagerServiceClient is an autogenerated mock type for the ReadSideManagerServiceClient type
type ReadSideManagerServiceClient struct {
	mock.Mock
}

// GetLatestOffset provides a mock function with given fields: ctx, in, opts
func (_m *ReadSideManagerServiceClient) GetLatestOffset(ctx context.Context, in *chief_of_statev1.GetLatestOffsetRequest, opts ...grpc.CallOption) (*chief_of_statev1.GetLatestOffsetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chief_of_statev1.GetLatestOffsetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.GetLatestOffsetRequest, ...grpc.CallOption) *chief_of_statev1.GetLatestOffsetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.GetLatestOffsetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.GetLatestOffsetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestOffsetByShard provides a mock function with given fields: ctx, in, opts
func (_m *ReadSideManagerServiceClient) GetLatestOffsetByShard(ctx context.Context, in *chief_of_statev1.GetLatestOffsetByShardRequest, opts ...grpc.CallOption) (*chief_of_statev1.GetLatestOffsetByShardResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chief_of_statev1.GetLatestOffsetByShardResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.GetLatestOffsetByShardRequest, ...grpc.CallOption) *chief_of_statev1.GetLatestOffsetByShardResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.GetLatestOffsetByShardResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.GetLatestOffsetByShardRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PauseReadSide provides a mock function with given fields: ctx, in, opts
func (_m *ReadSideManagerServiceClient) PauseReadSide(ctx context.Context, in *chief_of_statev1.PauseReadSideRequest, opts ...grpc.CallOption) (*chief_of_statev1.PauseReadSideResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chief_of_statev1.PauseReadSideResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.PauseReadSideRequest, ...grpc.CallOption) *chief_of_statev1.PauseReadSideResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.PauseReadSideResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.PauseReadSideRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PauseReadSideByShard provides a mock function with given fields: ctx, in, opts
func (_m *ReadSideManagerServiceClient) PauseReadSideByShard(ctx context.Context, in *chief_of_statev1.PauseReadSideByShardRequest, opts ...grpc.CallOption) (*chief_of_statev1.PauseReadSideByShardResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chief_of_statev1.PauseReadSideByShardResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.PauseReadSideByShardRequest, ...grpc.CallOption) *chief_of_statev1.PauseReadSideByShardResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.PauseReadSideByShardResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.PauseReadSideByShardRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestartReadSide provides a mock function with given fields: ctx, in, opts
func (_m *ReadSideManagerServiceClient) RestartReadSide(ctx context.Context, in *chief_of_statev1.RestartReadSideRequest, opts ...grpc.CallOption) (*chief_of_statev1.RestartReadSideResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chief_of_statev1.RestartReadSideResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.RestartReadSideRequest, ...grpc.CallOption) *chief_of_statev1.RestartReadSideResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.RestartReadSideResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.RestartReadSideRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestartReadSideByShard provides a mock function with given fields: ctx, in, opts
func (_m *ReadSideManagerServiceClient) RestartReadSideByShard(ctx context.Context, in *chief_of_statev1.RestartReadSideByShardRequest, opts ...grpc.CallOption) (*chief_of_statev1.RestartReadSideByShardResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chief_of_statev1.RestartReadSideByShardResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.RestartReadSideByShardRequest, ...grpc.CallOption) *chief_of_statev1.RestartReadSideByShardResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.RestartReadSideByShardResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.RestartReadSideByShardRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResumeReadSide provides a mock function with given fields: ctx, in, opts
func (_m *ReadSideManagerServiceClient) ResumeReadSide(ctx context.Context, in *chief_of_statev1.ResumeReadSideRequest, opts ...grpc.CallOption) (*chief_of_statev1.ResumeReadSideResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chief_of_statev1.ResumeReadSideResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.ResumeReadSideRequest, ...grpc.CallOption) *chief_of_statev1.ResumeReadSideResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.ResumeReadSideResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.ResumeReadSideRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResumeReadSideByShard provides a mock function with given fields: ctx, in, opts
func (_m *ReadSideManagerServiceClient) ResumeReadSideByShard(ctx context.Context, in *chief_of_statev1.ResumeReadSideByShardRequest, opts ...grpc.CallOption) (*chief_of_statev1.ResumeReadSideByShardResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chief_of_statev1.ResumeReadSideByShardResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.ResumeReadSideByShardRequest, ...grpc.CallOption) *chief_of_statev1.ResumeReadSideByShardResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.ResumeReadSideByShardResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.ResumeReadSideByShardRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SkipOffset provides a mock function with given fields: ctx, in, opts
func (_m *ReadSideManagerServiceClient) SkipOffset(ctx context.Context, in *chief_of_statev1.SkipOffsetRequest, opts ...grpc.CallOption) (*chief_of_statev1.SkipOffsetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chief_of_statev1.SkipOffsetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.SkipOffsetRequest, ...grpc.CallOption) *chief_of_statev1.SkipOffsetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.SkipOffsetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.SkipOffsetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SkipOffsetByShard provides a mock function with given fields: ctx, in, opts
func (_m *ReadSideManagerServiceClient) SkipOffsetByShard(ctx context.Context, in *chief_of_statev1.SkipOffsetByShardRequest, opts ...grpc.CallOption) (*chief_of_statev1.SkipOffsetByShardResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chief_of_statev1.SkipOffsetByShardResponse
	if rf, ok := ret.Get(0).(func(context.Context, *chief_of_statev1.SkipOffsetByShardRequest, ...grpc.CallOption) *chief_of_statev1.SkipOffsetByShardResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chief_of_statev1.SkipOffsetByShardResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *chief_of_statev1.SkipOffsetByShardRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewReadSideManagerServiceClientT interface {
	mock.TestingT
	Cleanup(func())
}

// NewReadSideManagerServiceClient creates a new instance of ReadSideManagerServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReadSideManagerServiceClient(t NewReadSideManagerServiceClientT) *ReadSideManagerServiceClient {
	mock := &ReadSideManagerServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
