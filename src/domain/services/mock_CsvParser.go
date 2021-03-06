// Code generated by mockery v2.4.0. DO NOT EDIT.

package services

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	model "github.com/nejtr0n/api_example/domain/model"
)

// MockCsvParser is an autogenerated mock type for the CsvParser type
type MockCsvParser struct {
	mock.Mock
}

// Parse provides a mock function with given fields: ctx, data, pipe
func (_m *MockCsvParser) Parse(ctx context.Context, data io.ReadCloser, pipe chan *model.Product) error {
	ret := _m.Called(ctx, data, pipe)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, io.ReadCloser, chan *model.Product) error); ok {
		r0 = rf(ctx, data, pipe)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
