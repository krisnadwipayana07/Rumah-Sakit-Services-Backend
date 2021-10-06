// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	doctors "backend/business/doctors"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, domain
func (_m *Repository) Login(ctx context.Context, domain doctors.Domain) (doctors.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 doctors.Domain
	if rf, ok := ret.Get(0).(func(context.Context, doctors.Domain) doctors.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(doctors.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, doctors.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, domain
func (_m *Repository) Register(ctx context.Context, domain doctors.Domain) (doctors.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 doctors.Domain
	if rf, ok := ret.Get(0).(func(context.Context, doctors.Domain) doctors.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(doctors.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, doctors.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowAll provides a mock function with given fields: ctx
func (_m *Repository) ShowAll(ctx context.Context) ([]doctors.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []doctors.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []doctors.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]doctors.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, domain
func (_m *Repository) Update(ctx context.Context, domain doctors.Domain) (doctors.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 doctors.Domain
	if rf, ok := ret.Get(0).(func(context.Context, doctors.Domain) doctors.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(doctors.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, doctors.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}