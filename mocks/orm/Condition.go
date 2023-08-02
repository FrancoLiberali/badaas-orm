// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// Condition is an autogenerated mock type for the Condition type
type Condition[T interface{}] struct {
	mock.Mock
}

// ApplyTo provides a mock function with given fields: query, tableName
func (_m *Condition[T]) ApplyTo(query *gorm.DB, tableName string) (*gorm.DB, error) {
	ret := _m.Called(query, tableName)

	var r0 *gorm.DB
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) (*gorm.DB, error)); ok {
		return rf(query, tableName)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) *gorm.DB); ok {
		r0 = rf(query, tableName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string) error); ok {
		r1 = rf(query, tableName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// interfaceVerificationMethod provides a mock function with given fields: _a0
func (_m *Condition[T]) interfaceVerificationMethod(_a0 T) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewCondition interface {
	mock.TestingT
	Cleanup(func())
}

// NewCondition creates a new instance of Condition. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCondition[T interface{}](t mockConstructorTestingTNewCondition) *Condition[T] {
	mock := &Condition[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}