// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DataSource is an autogenerated mock type for the DataSource type
type DataSource struct {
	mock.Mock
}

// Get provides a mock function with given fields: key
func (_m *DataSource) Get(key string) {
	_m.Called(key)
}

// Set provides a mock function with given fields: key, value
func (_m *DataSource) Set(key string, value interface{}) {
	_m.Called(key, value)
}

type mockConstructorTestingTNewDataSource interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataSource creates a new instance of DataSource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataSource(t mockConstructorTestingTNewDataSource) *DataSource {
	mock := &DataSource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
