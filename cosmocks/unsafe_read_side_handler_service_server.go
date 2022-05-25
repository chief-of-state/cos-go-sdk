// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeReadSideHandlerServiceServer is an autogenerated mock type for the UnsafeReadSideHandlerServiceServer type
type UnsafeReadSideHandlerServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedReadSideHandlerServiceServer provides a mock function with given fields:
func (_m *UnsafeReadSideHandlerServiceServer) mustEmbedUnimplementedReadSideHandlerServiceServer() {
	_m.Called()
}

type NewUnsafeReadSideHandlerServiceServerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeReadSideHandlerServiceServer creates a new instance of UnsafeReadSideHandlerServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeReadSideHandlerServiceServer(t NewUnsafeReadSideHandlerServiceServerT) *UnsafeReadSideHandlerServiceServer {
	mock := &UnsafeReadSideHandlerServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
