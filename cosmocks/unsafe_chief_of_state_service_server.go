// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeChiefOfStateServiceServer is an autogenerated mock type for the UnsafeChiefOfStateServiceServer type
type UnsafeChiefOfStateServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedChiefOfStateServiceServer provides a mock function with given fields:
func (_m *UnsafeChiefOfStateServiceServer) mustEmbedUnimplementedChiefOfStateServiceServer() {
	_m.Called()
}

type NewUnsafeChiefOfStateServiceServerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeChiefOfStateServiceServer creates a new instance of UnsafeChiefOfStateServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeChiefOfStateServiceServer(t NewUnsafeChiefOfStateServiceServerT) *UnsafeChiefOfStateServiceServer {
	mock := &UnsafeChiefOfStateServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
