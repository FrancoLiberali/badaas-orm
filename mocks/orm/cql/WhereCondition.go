// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	cql "github.com/ditrit/badaas/orm/cql"
	mock "github.com/stretchr/testify/mock"

	model "github.com/ditrit/badaas/orm/model"
)

// WhereCondition is an autogenerated mock type for the WhereCondition type
type WhereCondition[T model.Model] struct {
	mock.Mock
}

// AffectsDeletedAt provides a mock function with given fields:
func (_m *WhereCondition[T]) AffectsDeletedAt() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ApplyTo provides a mock function with given fields: _a0, _a1
func (_m *WhereCondition[T]) ApplyTo(_a0 *cql.GormQuery, _a1 cql.Table) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cql.GormQuery, cql.Table) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSQL provides a mock function with given fields: query, table
func (_m *WhereCondition[T]) GetSQL(query *cql.GormQuery, table cql.Table) (string, []interface{}, error) {
	ret := _m.Called(query, table)

	var r0 string
	var r1 []interface{}
	var r2 error
	if rf, ok := ret.Get(0).(func(*cql.GormQuery, cql.Table) (string, []interface{}, error)); ok {
		return rf(query, table)
	}
	if rf, ok := ret.Get(0).(func(*cql.GormQuery, cql.Table) string); ok {
		r0 = rf(query, table)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*cql.GormQuery, cql.Table) []interface{}); ok {
		r1 = rf(query, table)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	if rf, ok := ret.Get(2).(func(*cql.GormQuery, cql.Table) error); ok {
		r2 = rf(query, table)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InterfaceVerificationMethod provides a mock function with given fields: _a0
func (_m *WhereCondition[T]) InterfaceVerificationMethod(_a0 T) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewWhereCondition interface {
	mock.TestingT
	Cleanup(func())
}

// NewWhereCondition creates a new instance of WhereCondition. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWhereCondition[T model.Model](t mockConstructorTestingTNewWhereCondition) *WhereCondition[T] {
	mock := &WhereCondition[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
