// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeReadSideManagerServiceServer is an autogenerated mock type for the UnsafeReadSideManagerServiceServer type
type UnsafeReadSideManagerServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedReadSideManagerServiceServer provides a mock function with given fields:
func (_m *UnsafeReadSideManagerServiceServer) mustEmbedUnimplementedReadSideManagerServiceServer() {
	_m.Called()
}

type NewUnsafeReadSideManagerServiceServerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeReadSideManagerServiceServer creates a new instance of UnsafeReadSideManagerServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeReadSideManagerServiceServer(t NewUnsafeReadSideManagerServiceServerT) *UnsafeReadSideManagerServiceServer {
	mock := &UnsafeReadSideManagerServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
