// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isHeader_Value is an autogenerated mock type for the isHeader_Value type
type isHeader_Value struct {
	mock.Mock
}

// isHeader_Value provides a mock function with given fields:
func (_m *isHeader_Value) isHeader_Value() {
	_m.Called()
}

type newIsHeader_ValueT interface {
	mock.TestingT
	Cleanup(func())
}

// newIsHeader_Value creates a new instance of isHeader_Value. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newIsHeader_Value(t newIsHeader_ValueT) *isHeader_Value {
	mock := &isHeader_Value{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
