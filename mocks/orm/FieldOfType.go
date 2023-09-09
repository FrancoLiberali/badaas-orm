// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	orm "github.com/ditrit/badaas/orm"
	mock "github.com/stretchr/testify/mock"

	reflect "reflect"
)

// FieldOfType is an autogenerated mock type for the FieldOfType type
type FieldOfType[T interface{}] struct {
	mock.Mock
}

// ColumnName provides a mock function with given fields: query, table
func (_m *FieldOfType[T]) ColumnName(query *orm.GormQuery, table orm.Table) string {
	ret := _m.Called(query, table)

	var r0 string
	if rf, ok := ret.Get(0).(func(*orm.GormQuery, orm.Table) string); ok {
		r0 = rf(query, table)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ColumnSQL provides a mock function with given fields: query, table
func (_m *FieldOfType[T]) ColumnSQL(query *orm.GormQuery, table orm.Table) string {
	ret := _m.Called(query, table)

	var r0 string
	if rf, ok := ret.Get(0).(func(*orm.GormQuery, orm.Table) string); ok {
		r0 = rf(query, table)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// FieldName provides a mock function with given fields:
func (_m *FieldOfType[T]) FieldName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetModelType provides a mock function with given fields:
func (_m *FieldOfType[T]) GetModelType() reflect.Type {
	ret := _m.Called()

	var r0 reflect.Type
	if rf, ok := ret.Get(0).(func() reflect.Type); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.Type)
		}
	}

	return r0
}

// GetType provides a mock function with given fields:
func (_m *FieldOfType[T]) GetType() T {
	ret := _m.Called()

	var r0 T
	if rf, ok := ret.Get(0).(func() T); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(T)
	}

	return r0
}

type mockConstructorTestingTNewFieldOfType interface {
	mock.TestingT
	Cleanup(func())
}

// NewFieldOfType creates a new instance of FieldOfType. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFieldOfType[T interface{}](t mockConstructorTestingTNewFieldOfType) *FieldOfType[T] {
	mock := &FieldOfType[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
