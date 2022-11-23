// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dto "github.com/ezequiel-uala/contacts/lambda/go/get-contact-lambda/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// Processor is an autogenerated mock type for the Processor type
type Processor struct {
	mock.Mock
}

// Process provides a mock function with given fields: _a0
func (_m *Processor) Process(_a0 string) (dto.ContactResponse, error) {
	ret := _m.Called(_a0)

	var r0 dto.ContactResponse
	if rf, ok := ret.Get(0).(func(string) dto.ContactResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(dto.ContactResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProcessor interface {
	mock.TestingT
	Cleanup(func())
}

// NewProcessor creates a new instance of Processor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProcessor(t mockConstructorTestingTNewProcessor) *Processor {
	mock := &Processor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}