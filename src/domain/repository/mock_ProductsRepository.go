// Code generated by mockery v2.4.0. DO NOT EDIT.

package repository

import (
	context "context"

	model "github.com/nejtr0n/api_example/domain/model"
	mock "github.com/stretchr/testify/mock"
)

// MockProductsRepository is an autogenerated mock type for the ProductsRepository type
type MockProductsRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, limit, offset, sortBy, sortOrder
func (_m *MockProductsRepository) Find(ctx context.Context, limit int64, offset int64, sortBy string, sortOrder int) ([]*model.Product, error) {
	ret := _m.Called(ctx, limit, offset, sortBy, sortOrder)

	var r0 []*model.Product
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string, int) []*model.Product); ok {
		r0 = rf(ctx, limit, offset, sortBy, sortOrder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, string, int) error); ok {
		r1 = rf(ctx, limit, offset, sortBy, sortOrder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAll provides a mock function with given fields: ctx, pipe
func (_m *MockProductsRepository) SaveAll(ctx context.Context, pipe chan *model.Product) (int64, error) {
	ret := _m.Called(ctx, pipe)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, chan *model.Product) int64); ok {
		r0 = rf(ctx, pipe)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, chan *model.Product) error); ok {
		r1 = rf(ctx, pipe)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
