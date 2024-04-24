// Code generated by mocktail; DO NOT EDIT.

package message

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

// clientMock mock of Client.
type clientMock struct{ mock.Mock }

// newClientMock creates a new clientMock.
func newClientMock(tb testing.TB) *clientMock {
	tb.Helper()

	m := &clientMock{}
	m.Mock.Test(tb)

	tb.Cleanup(func() { m.AssertExpectations(tb) })

	return m
}

func (_m *clientMock) Delete(_ context.Context, path string, obj any, resp any) (any, error) {
	_ret := _m.Called(path, obj, resp)

	if _rf, ok := _ret.Get(0).(func(string, any, any) (any, error)); ok {
		return _rf(path, obj, resp)
	}

	_ra0, _ := _ret.Get(0).(any)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *clientMock) OnDelete(path string, obj any, resp any) *clientDeleteCall {
	return &clientDeleteCall{Call: _m.Mock.On("Delete", path, obj, resp), Parent: _m}
}

func (_m *clientMock) OnDeleteRaw(path interface{}, obj interface{}, resp interface{}) *clientDeleteCall {
	return &clientDeleteCall{Call: _m.Mock.On("Delete", path, obj, resp), Parent: _m}
}

type clientDeleteCall struct {
	*mock.Call
	Parent *clientMock
}

func (_c *clientDeleteCall) Panic(msg string) *clientDeleteCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *clientDeleteCall) Once() *clientDeleteCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *clientDeleteCall) Twice() *clientDeleteCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *clientDeleteCall) Times(i int) *clientDeleteCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *clientDeleteCall) WaitUntil(w <-chan time.Time) *clientDeleteCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *clientDeleteCall) After(d time.Duration) *clientDeleteCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *clientDeleteCall) Run(fn func(args mock.Arguments)) *clientDeleteCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *clientDeleteCall) Maybe() *clientDeleteCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *clientDeleteCall) TypedReturns(a any, b error) *clientDeleteCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *clientDeleteCall) ReturnsFn(fn func(string, any, any) (any, error)) *clientDeleteCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *clientDeleteCall) TypedRun(fn func(string, any, any)) *clientDeleteCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_path := args.String(0)
		_obj, _ := args.Get(1).(any)
		_resp, _ := args.Get(2).(any)
		fn(_path, _obj, _resp)
	})
	return _c
}

func (_c *clientDeleteCall) OnDelete(path string, obj any, resp any) *clientDeleteCall {
	return _c.Parent.OnDelete(path, obj, resp)
}

func (_c *clientDeleteCall) OnGet(path string, obj any, resp any) *clientGetCall {
	return _c.Parent.OnGet(path, obj, resp)
}

func (_c *clientDeleteCall) OnPost(path string, obj any, resp any) *clientPostCall {
	return _c.Parent.OnPost(path, obj, resp)
}

func (_c *clientDeleteCall) OnPut(path string, obj any, resp any) *clientPutCall {
	return _c.Parent.OnPut(path, obj, resp)
}

func (_c *clientDeleteCall) OnDeleteRaw(path interface{}, obj interface{}, resp interface{}) *clientDeleteCall {
	return _c.Parent.OnDeleteRaw(path, obj, resp)
}

func (_c *clientDeleteCall) OnGetRaw(path interface{}, obj interface{}, resp interface{}) *clientGetCall {
	return _c.Parent.OnGetRaw(path, obj, resp)
}

func (_c *clientDeleteCall) OnPostRaw(path interface{}, obj interface{}, resp interface{}) *clientPostCall {
	return _c.Parent.OnPostRaw(path, obj, resp)
}

func (_c *clientDeleteCall) OnPutRaw(path interface{}, obj interface{}, resp interface{}) *clientPutCall {
	return _c.Parent.OnPutRaw(path, obj, resp)
}

func (_m *clientMock) Get(_ context.Context, path string, obj any, resp any) (any, error) {
	_ret := _m.Called(path, obj, resp)

	if _rf, ok := _ret.Get(0).(func(string, any, any) (any, error)); ok {
		return _rf(path, obj, resp)
	}

	_ra0, _ := _ret.Get(0).(any)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *clientMock) OnGet(path string, obj any, resp any) *clientGetCall {
	return &clientGetCall{Call: _m.Mock.On("Get", path, obj, resp), Parent: _m}
}

func (_m *clientMock) OnGetRaw(path interface{}, obj interface{}, resp interface{}) *clientGetCall {
	return &clientGetCall{Call: _m.Mock.On("Get", path, obj, resp), Parent: _m}
}

type clientGetCall struct {
	*mock.Call
	Parent *clientMock
}

func (_c *clientGetCall) Panic(msg string) *clientGetCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *clientGetCall) Once() *clientGetCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *clientGetCall) Twice() *clientGetCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *clientGetCall) Times(i int) *clientGetCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *clientGetCall) WaitUntil(w <-chan time.Time) *clientGetCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *clientGetCall) After(d time.Duration) *clientGetCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *clientGetCall) Run(fn func(args mock.Arguments)) *clientGetCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *clientGetCall) Maybe() *clientGetCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *clientGetCall) TypedReturns(a any, b error) *clientGetCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *clientGetCall) ReturnsFn(fn func(string, any, any) (any, error)) *clientGetCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *clientGetCall) TypedRun(fn func(string, any, any)) *clientGetCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_path := args.String(0)
		_obj, _ := args.Get(1).(any)
		_resp, _ := args.Get(2).(any)
		fn(_path, _obj, _resp)
	})
	return _c
}

func (_c *clientGetCall) OnDelete(path string, obj any, resp any) *clientDeleteCall {
	return _c.Parent.OnDelete(path, obj, resp)
}

func (_c *clientGetCall) OnGet(path string, obj any, resp any) *clientGetCall {
	return _c.Parent.OnGet(path, obj, resp)
}

func (_c *clientGetCall) OnPost(path string, obj any, resp any) *clientPostCall {
	return _c.Parent.OnPost(path, obj, resp)
}

func (_c *clientGetCall) OnPut(path string, obj any, resp any) *clientPutCall {
	return _c.Parent.OnPut(path, obj, resp)
}

func (_c *clientGetCall) OnDeleteRaw(path interface{}, obj interface{}, resp interface{}) *clientDeleteCall {
	return _c.Parent.OnDeleteRaw(path, obj, resp)
}

func (_c *clientGetCall) OnGetRaw(path interface{}, obj interface{}, resp interface{}) *clientGetCall {
	return _c.Parent.OnGetRaw(path, obj, resp)
}

func (_c *clientGetCall) OnPostRaw(path interface{}, obj interface{}, resp interface{}) *clientPostCall {
	return _c.Parent.OnPostRaw(path, obj, resp)
}

func (_c *clientGetCall) OnPutRaw(path interface{}, obj interface{}, resp interface{}) *clientPutCall {
	return _c.Parent.OnPutRaw(path, obj, resp)
}

func (_m *clientMock) Post(_ context.Context, path string, obj any, resp any) (any, error) {
	_ret := _m.Called(path, obj, resp)

	if _rf, ok := _ret.Get(0).(func(string, any, any) (any, error)); ok {
		return _rf(path, obj, resp)
	}

	_ra0, _ := _ret.Get(0).(any)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *clientMock) OnPost(path string, obj any, resp any) *clientPostCall {
	return &clientPostCall{Call: _m.Mock.On("Post", path, obj, resp), Parent: _m}
}

func (_m *clientMock) OnPostRaw(path interface{}, obj interface{}, resp interface{}) *clientPostCall {
	return &clientPostCall{Call: _m.Mock.On("Post", path, obj, resp), Parent: _m}
}

type clientPostCall struct {
	*mock.Call
	Parent *clientMock
}

func (_c *clientPostCall) Panic(msg string) *clientPostCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *clientPostCall) Once() *clientPostCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *clientPostCall) Twice() *clientPostCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *clientPostCall) Times(i int) *clientPostCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *clientPostCall) WaitUntil(w <-chan time.Time) *clientPostCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *clientPostCall) After(d time.Duration) *clientPostCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *clientPostCall) Run(fn func(args mock.Arguments)) *clientPostCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *clientPostCall) Maybe() *clientPostCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *clientPostCall) TypedReturns(a any, b error) *clientPostCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *clientPostCall) ReturnsFn(fn func(string, any, any) (any, error)) *clientPostCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *clientPostCall) TypedRun(fn func(string, any, any)) *clientPostCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_path := args.String(0)
		_obj, _ := args.Get(1).(any)
		_resp, _ := args.Get(2).(any)
		fn(_path, _obj, _resp)
	})
	return _c
}

func (_c *clientPostCall) OnDelete(path string, obj any, resp any) *clientDeleteCall {
	return _c.Parent.OnDelete(path, obj, resp)
}

func (_c *clientPostCall) OnGet(path string, obj any, resp any) *clientGetCall {
	return _c.Parent.OnGet(path, obj, resp)
}

func (_c *clientPostCall) OnPost(path string, obj any, resp any) *clientPostCall {
	return _c.Parent.OnPost(path, obj, resp)
}

func (_c *clientPostCall) OnPut(path string, obj any, resp any) *clientPutCall {
	return _c.Parent.OnPut(path, obj, resp)
}

func (_c *clientPostCall) OnDeleteRaw(path interface{}, obj interface{}, resp interface{}) *clientDeleteCall {
	return _c.Parent.OnDeleteRaw(path, obj, resp)
}

func (_c *clientPostCall) OnGetRaw(path interface{}, obj interface{}, resp interface{}) *clientGetCall {
	return _c.Parent.OnGetRaw(path, obj, resp)
}

func (_c *clientPostCall) OnPostRaw(path interface{}, obj interface{}, resp interface{}) *clientPostCall {
	return _c.Parent.OnPostRaw(path, obj, resp)
}

func (_c *clientPostCall) OnPutRaw(path interface{}, obj interface{}, resp interface{}) *clientPutCall {
	return _c.Parent.OnPutRaw(path, obj, resp)
}

func (_m *clientMock) Put(_ context.Context, path string, obj any, resp any) (any, error) {
	_ret := _m.Called(path, obj, resp)

	if _rf, ok := _ret.Get(0).(func(string, any, any) (any, error)); ok {
		return _rf(path, obj, resp)
	}

	_ra0, _ := _ret.Get(0).(any)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *clientMock) OnPut(path string, obj any, resp any) *clientPutCall {
	return &clientPutCall{Call: _m.Mock.On("Put", path, obj, resp), Parent: _m}
}

func (_m *clientMock) OnPutRaw(path interface{}, obj interface{}, resp interface{}) *clientPutCall {
	return &clientPutCall{Call: _m.Mock.On("Put", path, obj, resp), Parent: _m}
}

type clientPutCall struct {
	*mock.Call
	Parent *clientMock
}

func (_c *clientPutCall) Panic(msg string) *clientPutCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *clientPutCall) Once() *clientPutCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *clientPutCall) Twice() *clientPutCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *clientPutCall) Times(i int) *clientPutCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *clientPutCall) WaitUntil(w <-chan time.Time) *clientPutCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *clientPutCall) After(d time.Duration) *clientPutCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *clientPutCall) Run(fn func(args mock.Arguments)) *clientPutCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *clientPutCall) Maybe() *clientPutCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *clientPutCall) TypedReturns(a any, b error) *clientPutCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *clientPutCall) ReturnsFn(fn func(string, any, any) (any, error)) *clientPutCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *clientPutCall) TypedRun(fn func(string, any, any)) *clientPutCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_path := args.String(0)
		_obj, _ := args.Get(1).(any)
		_resp, _ := args.Get(2).(any)
		fn(_path, _obj, _resp)
	})
	return _c
}

func (_c *clientPutCall) OnDelete(path string, obj any, resp any) *clientDeleteCall {
	return _c.Parent.OnDelete(path, obj, resp)
}

func (_c *clientPutCall) OnGet(path string, obj any, resp any) *clientGetCall {
	return _c.Parent.OnGet(path, obj, resp)
}

func (_c *clientPutCall) OnPost(path string, obj any, resp any) *clientPostCall {
	return _c.Parent.OnPost(path, obj, resp)
}

func (_c *clientPutCall) OnPut(path string, obj any, resp any) *clientPutCall {
	return _c.Parent.OnPut(path, obj, resp)
}

func (_c *clientPutCall) OnDeleteRaw(path interface{}, obj interface{}, resp interface{}) *clientDeleteCall {
	return _c.Parent.OnDeleteRaw(path, obj, resp)
}

func (_c *clientPutCall) OnGetRaw(path interface{}, obj interface{}, resp interface{}) *clientGetCall {
	return _c.Parent.OnGetRaw(path, obj, resp)
}

func (_c *clientPutCall) OnPostRaw(path interface{}, obj interface{}, resp interface{}) *clientPostCall {
	return _c.Parent.OnPostRaw(path, obj, resp)
}

func (_c *clientPutCall) OnPutRaw(path interface{}, obj interface{}, resp interface{}) *clientPutCall {
	return _c.Parent.OnPutRaw(path, obj, resp)
}
