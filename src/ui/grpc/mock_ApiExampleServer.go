// Code generated by mockery v2.4.0. DO NOT EDIT.

package api_example

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockApiExampleServer is an autogenerated mock type for the ApiExampleServer type
type MockApiExampleServer struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: _a0, _a1
func (_m *MockApiExampleServer) Fetch(_a0 context.Context, _a1 *FetchRequest) (*FetchResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *FetchResponse
	if rf, ok := ret.Get(0).(func(context.Context, *FetchRequest) *FetchResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*FetchResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *FetchRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *MockApiExampleServer) List(_a0 context.Context, _a1 *ListRequest) (*ListResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ListRequest) *ListResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ListRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}