// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// Tx is an autogenerated mock type for the Tx type
type Tx struct {
	mock.Mock
}

type Tx_Expecter struct {
	mock *mock.Mock
}

func (_m *Tx) EXPECT() *Tx_Expecter {
	return &Tx_Expecter{mock: &_m.Mock}
}

// Commit provides a mock function with given fields:
func (_m *Tx) Commit() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Tx_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type Tx_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
func (_e *Tx_Expecter) Commit() *Tx_Commit_Call {
	return &Tx_Commit_Call{Call: _e.mock.On("Commit")}
}

func (_c *Tx_Commit_Call) Run(run func()) *Tx_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tx_Commit_Call) Return(_a0 error) *Tx_Commit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Tx_Commit_Call) RunAndReturn(run func() error) *Tx_Commit_Call {
	_c.Call.Return(run)
	return _c
}

// ExecContext provides a mock function with given fields: ctx, query, args
func (_m *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecContext")
	}

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (sql.Result, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tx_ExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecContext'
type Tx_ExecContext_Call struct {
	*mock.Call
}

// ExecContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *Tx_Expecter) ExecContext(ctx interface{}, query interface{}, args ...interface{}) *Tx_ExecContext_Call {
	return &Tx_ExecContext_Call{Call: _e.mock.On("ExecContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *Tx_ExecContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *Tx_ExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Tx_ExecContext_Call) Return(_a0 sql.Result, _a1 error) *Tx_ExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Tx_ExecContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (sql.Result, error)) *Tx_ExecContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryContext provides a mock function with given fields: ctx, query, args
func (_m *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryContext")
	}

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*sql.Rows, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tx_QueryContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryContext'
type Tx_QueryContext_Call struct {
	*mock.Call
}

// QueryContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *Tx_Expecter) QueryContext(ctx interface{}, query interface{}, args ...interface{}) *Tx_QueryContext_Call {
	return &Tx_QueryContext_Call{Call: _e.mock.On("QueryContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *Tx_QueryContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *Tx_QueryContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Tx_QueryContext_Call) Return(_a0 *sql.Rows, _a1 error) *Tx_QueryContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Tx_QueryContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (*sql.Rows, error)) *Tx_QueryContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRowContext provides a mock function with given fields: ctx, query, args
func (_m *Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRowContext")
	}

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Row); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// Tx_QueryRowContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRowContext'
type Tx_QueryRowContext_Call struct {
	*mock.Call
}

// QueryRowContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *Tx_Expecter) QueryRowContext(ctx interface{}, query interface{}, args ...interface{}) *Tx_QueryRowContext_Call {
	return &Tx_QueryRowContext_Call{Call: _e.mock.On("QueryRowContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *Tx_QueryRowContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *Tx_QueryRowContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Tx_QueryRowContext_Call) Return(_a0 *sql.Row) *Tx_QueryRowContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Tx_QueryRowContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) *sql.Row) *Tx_QueryRowContext_Call {
	_c.Call.Return(run)
	return _c
}

// Rollback provides a mock function with given fields:
func (_m *Tx) Rollback() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Tx_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type Tx_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
func (_e *Tx_Expecter) Rollback() *Tx_Rollback_Call {
	return &Tx_Rollback_Call{Call: _e.mock.On("Rollback")}
}

func (_c *Tx_Rollback_Call) Run(run func()) *Tx_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tx_Rollback_Call) Return(_a0 error) *Tx_Rollback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Tx_Rollback_Call) RunAndReturn(run func() error) *Tx_Rollback_Call {
	_c.Call.Return(run)
	return _c
}

// NewTx creates a new instance of Tx. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTx(t interface {
	mock.TestingT
	Cleanup(func())
}) *Tx {
	mock := &Tx{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
