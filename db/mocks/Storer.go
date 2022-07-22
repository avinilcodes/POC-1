// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	db "taskmanager/db"

	mock "github.com/stretchr/testify/mock"
)

// Storer is an autogenerated mock type for the Storer type
type Storer struct {
	mock.Mock
}

// AssignTask provides a mock function with given fields: ctx, userId, taskId
func (_m *Storer) AssignTask(ctx context.Context, userId string, taskId string) error {
	ret := _m.Called(ctx, userId, taskId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userId, taskId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTask provides a mock function with given fields: ctx, task
func (_m *Storer) CreateTask(ctx context.Context, task db.Task) error {
	ret := _m.Called(ctx, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.Task) error); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Storer) CreateUser(ctx context.Context, user db.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUserByEmail provides a mock function with given fields: ctx, email
func (_m *Storer) FindUserByEmail(ctx context.Context, email string) (db.User, error) {
	ret := _m.Called(ctx, email)

	var r0 db.User
	if rf, ok := ret.Get(0).(func(context.Context, string) db.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTasks provides a mock function with given fields: ctx, email
func (_m *Storer) ListTasks(ctx context.Context, email string) ([]db.Task, error) {
	ret := _m.Called(ctx, email)

	var r0 []db.Task
	if rf, ok := ret.Get(0).(func(context.Context, string) []db.Task); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUserTask provides a mock function with given fields: ctx
func (_m *Storer) ListUserTask(ctx context.Context) ([]db.NameUserTask, error) {
	ret := _m.Called(ctx)

	var r0 []db.NameUserTask
	if rf, ok := ret.Get(0).(func(context.Context) []db.NameUserTask); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.NameUserTask)
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

// ListUsers provides a mock function with given fields: ctx, email
func (_m *Storer) ListUsers(ctx context.Context, email string) ([]db.User, error) {
	ret := _m.Called(ctx, email)

	var r0 []db.User
	if rf, ok := ret.Get(0).(func(context.Context, string) []db.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTaskStatus provides a mock function with given fields: ctx, id, status, userEmail
func (_m *Storer) UpdateTaskStatus(ctx context.Context, id string, status string, userEmail string) error {
	ret := _m.Called(ctx, id, status, userEmail)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, id, status, userEmail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStorer interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorer creates a new instance of Storer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorer(t mockConstructorTestingTNewStorer) *Storer {
	mock := &Storer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
