// Code generated by mocktail; DO NOT EDIT.

package user

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

// userMock mock of User.
type userMock struct{ mock.Mock }

// NewUserMock creates a new userMock.
func NewUserMock(tb testing.TB) *userMock {
	tb.Helper()

	m := &userMock{}
	m.Mock.Test(tb)

	tb.Cleanup(func() { m.AssertExpectations(tb) })

	return m
}

func (_m *userMock) CreateUser(_ context.Context, createUserRequest CreateUserRequest) (*CreateUserResponse, error) {
	_ret := _m.Called(createUserRequest)

	if _rf, ok := _ret.Get(0).(func(CreateUserRequest) (*CreateUserResponse, error)); ok {
		return _rf(createUserRequest)
	}

	_ra0, _ := _ret.Get(0).(*CreateUserResponse)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *userMock) OnCreateUser(createUserRequest CreateUserRequest) *userCreateUserCall {
	return &userCreateUserCall{Call: _m.Mock.On("CreateUser", createUserRequest), Parent: _m}
}

func (_m *userMock) OnCreateUserRaw(createUserRequest interface{}) *userCreateUserCall {
	return &userCreateUserCall{Call: _m.Mock.On("CreateUser", createUserRequest), Parent: _m}
}

type userCreateUserCall struct {
	*mock.Call
	Parent *userMock
}

func (_c *userCreateUserCall) Panic(msg string) *userCreateUserCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *userCreateUserCall) Once() *userCreateUserCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *userCreateUserCall) Twice() *userCreateUserCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *userCreateUserCall) Times(i int) *userCreateUserCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *userCreateUserCall) WaitUntil(w <-chan time.Time) *userCreateUserCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *userCreateUserCall) After(d time.Duration) *userCreateUserCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *userCreateUserCall) Run(fn func(args mock.Arguments)) *userCreateUserCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *userCreateUserCall) Maybe() *userCreateUserCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *userCreateUserCall) TypedReturns(a *CreateUserResponse, b error) *userCreateUserCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *userCreateUserCall) ReturnsFn(fn func(CreateUserRequest) (*CreateUserResponse, error)) *userCreateUserCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *userCreateUserCall) TypedRun(fn func(CreateUserRequest)) *userCreateUserCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_createUserRequest, _ := args.Get(0).(CreateUserRequest)
		fn(_createUserRequest)
	})
	return _c
}

func (_c *userCreateUserCall) OnCreateUser(createUserRequest CreateUserRequest) *userCreateUserCall {
	return _c.Parent.OnCreateUser(createUserRequest)
}

func (_c *userCreateUserCall) OnGetGroupChannelCount(userID string, getGroupChannelCountRequest GetGroupChannelCountRequest) *userGetGroupChannelCountCall {
	return _c.Parent.OnGetGroupChannelCount(userID, getGroupChannelCountRequest)
}

func (_c *userCreateUserCall) OnGetSessionToken(userID string, getSessionTokenRequest GetSessionTokenRequest) *userGetSessionTokenCall {
	return _c.Parent.OnGetSessionToken(userID, getSessionTokenRequest)
}

func (_c *userCreateUserCall) OnGetUnreadMessagesCount(userID string, getUnreadMessagesCountRequest GetUnreadMessagesCountRequest) *userGetUnreadMessagesCountCall {
	return _c.Parent.OnGetUnreadMessagesCount(userID, getUnreadMessagesCountRequest)
}

func (_c *userCreateUserCall) OnUpdateUser(userID string, updateUserRequest UpdateUserRequest) *userUpdateUserCall {
	return _c.Parent.OnUpdateUser(userID, updateUserRequest)
}

func (_c *userCreateUserCall) OnCreateUserRaw(createUserRequest interface{}) *userCreateUserCall {
	return _c.Parent.OnCreateUserRaw(createUserRequest)
}

func (_c *userCreateUserCall) OnGetGroupChannelCountRaw(userID interface{}, getGroupChannelCountRequest interface{}) *userGetGroupChannelCountCall {
	return _c.Parent.OnGetGroupChannelCountRaw(userID, getGroupChannelCountRequest)
}

func (_c *userCreateUserCall) OnGetSessionTokenRaw(userID interface{}, getSessionTokenRequest interface{}) *userGetSessionTokenCall {
	return _c.Parent.OnGetSessionTokenRaw(userID, getSessionTokenRequest)
}

func (_c *userCreateUserCall) OnGetUnreadMessagesCountRaw(userID interface{}, getUnreadMessagesCountRequest interface{}) *userGetUnreadMessagesCountCall {
	return _c.Parent.OnGetUnreadMessagesCountRaw(userID, getUnreadMessagesCountRequest)
}

func (_c *userCreateUserCall) OnUpdateUserRaw(userID interface{}, updateUserRequest interface{}) *userUpdateUserCall {
	return _c.Parent.OnUpdateUserRaw(userID, updateUserRequest)
}

func (_m *userMock) GetGroupChannelCount(_ context.Context, userID string, getGroupChannelCountRequest GetGroupChannelCountRequest) (*GetGroupChannelCountResponse, error) {
	_ret := _m.Called(userID, getGroupChannelCountRequest)

	if _rf, ok := _ret.Get(0).(func(string, GetGroupChannelCountRequest) (*GetGroupChannelCountResponse, error)); ok {
		return _rf(userID, getGroupChannelCountRequest)
	}

	_ra0, _ := _ret.Get(0).(*GetGroupChannelCountResponse)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *userMock) OnGetGroupChannelCount(userID string, getGroupChannelCountRequest GetGroupChannelCountRequest) *userGetGroupChannelCountCall {
	return &userGetGroupChannelCountCall{Call: _m.Mock.On("GetGroupChannelCount", userID, getGroupChannelCountRequest), Parent: _m}
}

func (_m *userMock) OnGetGroupChannelCountRaw(userID interface{}, getGroupChannelCountRequest interface{}) *userGetGroupChannelCountCall {
	return &userGetGroupChannelCountCall{Call: _m.Mock.On("GetGroupChannelCount", userID, getGroupChannelCountRequest), Parent: _m}
}

type userGetGroupChannelCountCall struct {
	*mock.Call
	Parent *userMock
}

func (_c *userGetGroupChannelCountCall) Panic(msg string) *userGetGroupChannelCountCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *userGetGroupChannelCountCall) Once() *userGetGroupChannelCountCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *userGetGroupChannelCountCall) Twice() *userGetGroupChannelCountCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *userGetGroupChannelCountCall) Times(i int) *userGetGroupChannelCountCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *userGetGroupChannelCountCall) WaitUntil(w <-chan time.Time) *userGetGroupChannelCountCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *userGetGroupChannelCountCall) After(d time.Duration) *userGetGroupChannelCountCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *userGetGroupChannelCountCall) Run(fn func(args mock.Arguments)) *userGetGroupChannelCountCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *userGetGroupChannelCountCall) Maybe() *userGetGroupChannelCountCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *userGetGroupChannelCountCall) TypedReturns(a *GetGroupChannelCountResponse, b error) *userGetGroupChannelCountCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *userGetGroupChannelCountCall) ReturnsFn(fn func(string, GetGroupChannelCountRequest) (*GetGroupChannelCountResponse, error)) *userGetGroupChannelCountCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *userGetGroupChannelCountCall) TypedRun(fn func(string, GetGroupChannelCountRequest)) *userGetGroupChannelCountCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_userID := args.String(0)
		_getGroupChannelCountRequest, _ := args.Get(1).(GetGroupChannelCountRequest)
		fn(_userID, _getGroupChannelCountRequest)
	})
	return _c
}

func (_c *userGetGroupChannelCountCall) OnCreateUser(createUserRequest CreateUserRequest) *userCreateUserCall {
	return _c.Parent.OnCreateUser(createUserRequest)
}

func (_c *userGetGroupChannelCountCall) OnGetGroupChannelCount(userID string, getGroupChannelCountRequest GetGroupChannelCountRequest) *userGetGroupChannelCountCall {
	return _c.Parent.OnGetGroupChannelCount(userID, getGroupChannelCountRequest)
}

func (_c *userGetGroupChannelCountCall) OnGetSessionToken(userID string, getSessionTokenRequest GetSessionTokenRequest) *userGetSessionTokenCall {
	return _c.Parent.OnGetSessionToken(userID, getSessionTokenRequest)
}

func (_c *userGetGroupChannelCountCall) OnGetUnreadMessagesCount(userID string, getUnreadMessagesCountRequest GetUnreadMessagesCountRequest) *userGetUnreadMessagesCountCall {
	return _c.Parent.OnGetUnreadMessagesCount(userID, getUnreadMessagesCountRequest)
}

func (_c *userGetGroupChannelCountCall) OnUpdateUser(userID string, updateUserRequest UpdateUserRequest) *userUpdateUserCall {
	return _c.Parent.OnUpdateUser(userID, updateUserRequest)
}

func (_c *userGetGroupChannelCountCall) OnCreateUserRaw(createUserRequest interface{}) *userCreateUserCall {
	return _c.Parent.OnCreateUserRaw(createUserRequest)
}

func (_c *userGetGroupChannelCountCall) OnGetGroupChannelCountRaw(userID interface{}, getGroupChannelCountRequest interface{}) *userGetGroupChannelCountCall {
	return _c.Parent.OnGetGroupChannelCountRaw(userID, getGroupChannelCountRequest)
}

func (_c *userGetGroupChannelCountCall) OnGetSessionTokenRaw(userID interface{}, getSessionTokenRequest interface{}) *userGetSessionTokenCall {
	return _c.Parent.OnGetSessionTokenRaw(userID, getSessionTokenRequest)
}

func (_c *userGetGroupChannelCountCall) OnGetUnreadMessagesCountRaw(userID interface{}, getUnreadMessagesCountRequest interface{}) *userGetUnreadMessagesCountCall {
	return _c.Parent.OnGetUnreadMessagesCountRaw(userID, getUnreadMessagesCountRequest)
}

func (_c *userGetGroupChannelCountCall) OnUpdateUserRaw(userID interface{}, updateUserRequest interface{}) *userUpdateUserCall {
	return _c.Parent.OnUpdateUserRaw(userID, updateUserRequest)
}

func (_m *userMock) GetSessionToken(_ context.Context, userID string, getSessionTokenRequest GetSessionTokenRequest) (*GetSessionTokenResponse, error) {
	_ret := _m.Called(userID, getSessionTokenRequest)

	if _rf, ok := _ret.Get(0).(func(string, GetSessionTokenRequest) (*GetSessionTokenResponse, error)); ok {
		return _rf(userID, getSessionTokenRequest)
	}

	_ra0, _ := _ret.Get(0).(*GetSessionTokenResponse)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *userMock) OnGetSessionToken(userID string, getSessionTokenRequest GetSessionTokenRequest) *userGetSessionTokenCall {
	return &userGetSessionTokenCall{Call: _m.Mock.On("GetSessionToken", userID, getSessionTokenRequest), Parent: _m}
}

func (_m *userMock) OnGetSessionTokenRaw(userID interface{}, getSessionTokenRequest interface{}) *userGetSessionTokenCall {
	return &userGetSessionTokenCall{Call: _m.Mock.On("GetSessionToken", userID, getSessionTokenRequest), Parent: _m}
}

type userGetSessionTokenCall struct {
	*mock.Call
	Parent *userMock
}

func (_c *userGetSessionTokenCall) Panic(msg string) *userGetSessionTokenCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *userGetSessionTokenCall) Once() *userGetSessionTokenCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *userGetSessionTokenCall) Twice() *userGetSessionTokenCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *userGetSessionTokenCall) Times(i int) *userGetSessionTokenCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *userGetSessionTokenCall) WaitUntil(w <-chan time.Time) *userGetSessionTokenCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *userGetSessionTokenCall) After(d time.Duration) *userGetSessionTokenCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *userGetSessionTokenCall) Run(fn func(args mock.Arguments)) *userGetSessionTokenCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *userGetSessionTokenCall) Maybe() *userGetSessionTokenCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *userGetSessionTokenCall) TypedReturns(a *GetSessionTokenResponse, b error) *userGetSessionTokenCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *userGetSessionTokenCall) ReturnsFn(fn func(string, GetSessionTokenRequest) (*GetSessionTokenResponse, error)) *userGetSessionTokenCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *userGetSessionTokenCall) TypedRun(fn func(string, GetSessionTokenRequest)) *userGetSessionTokenCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_userID := args.String(0)
		_getSessionTokenRequest, _ := args.Get(1).(GetSessionTokenRequest)
		fn(_userID, _getSessionTokenRequest)
	})
	return _c
}

func (_c *userGetSessionTokenCall) OnCreateUser(createUserRequest CreateUserRequest) *userCreateUserCall {
	return _c.Parent.OnCreateUser(createUserRequest)
}

func (_c *userGetSessionTokenCall) OnGetGroupChannelCount(userID string, getGroupChannelCountRequest GetGroupChannelCountRequest) *userGetGroupChannelCountCall {
	return _c.Parent.OnGetGroupChannelCount(userID, getGroupChannelCountRequest)
}

func (_c *userGetSessionTokenCall) OnGetSessionToken(userID string, getSessionTokenRequest GetSessionTokenRequest) *userGetSessionTokenCall {
	return _c.Parent.OnGetSessionToken(userID, getSessionTokenRequest)
}

func (_c *userGetSessionTokenCall) OnGetUnreadMessagesCount(userID string, getUnreadMessagesCountRequest GetUnreadMessagesCountRequest) *userGetUnreadMessagesCountCall {
	return _c.Parent.OnGetUnreadMessagesCount(userID, getUnreadMessagesCountRequest)
}

func (_c *userGetSessionTokenCall) OnUpdateUser(userID string, updateUserRequest UpdateUserRequest) *userUpdateUserCall {
	return _c.Parent.OnUpdateUser(userID, updateUserRequest)
}

func (_c *userGetSessionTokenCall) OnCreateUserRaw(createUserRequest interface{}) *userCreateUserCall {
	return _c.Parent.OnCreateUserRaw(createUserRequest)
}

func (_c *userGetSessionTokenCall) OnGetGroupChannelCountRaw(userID interface{}, getGroupChannelCountRequest interface{}) *userGetGroupChannelCountCall {
	return _c.Parent.OnGetGroupChannelCountRaw(userID, getGroupChannelCountRequest)
}

func (_c *userGetSessionTokenCall) OnGetSessionTokenRaw(userID interface{}, getSessionTokenRequest interface{}) *userGetSessionTokenCall {
	return _c.Parent.OnGetSessionTokenRaw(userID, getSessionTokenRequest)
}

func (_c *userGetSessionTokenCall) OnGetUnreadMessagesCountRaw(userID interface{}, getUnreadMessagesCountRequest interface{}) *userGetUnreadMessagesCountCall {
	return _c.Parent.OnGetUnreadMessagesCountRaw(userID, getUnreadMessagesCountRequest)
}

func (_c *userGetSessionTokenCall) OnUpdateUserRaw(userID interface{}, updateUserRequest interface{}) *userUpdateUserCall {
	return _c.Parent.OnUpdateUserRaw(userID, updateUserRequest)
}

func (_m *userMock) GetUnreadMessagesCount(_ context.Context, userID string, getUnreadMessagesCountRequest GetUnreadMessagesCountRequest) (*GetUnreadMessagesCountResponse, error) {
	_ret := _m.Called(userID, getUnreadMessagesCountRequest)

	if _rf, ok := _ret.Get(0).(func(string, GetUnreadMessagesCountRequest) (*GetUnreadMessagesCountResponse, error)); ok {
		return _rf(userID, getUnreadMessagesCountRequest)
	}

	_ra0, _ := _ret.Get(0).(*GetUnreadMessagesCountResponse)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *userMock) OnGetUnreadMessagesCount(userID string, getUnreadMessagesCountRequest GetUnreadMessagesCountRequest) *userGetUnreadMessagesCountCall {
	return &userGetUnreadMessagesCountCall{Call: _m.Mock.On("GetUnreadMessagesCount", userID, getUnreadMessagesCountRequest), Parent: _m}
}

func (_m *userMock) OnGetUnreadMessagesCountRaw(userID interface{}, getUnreadMessagesCountRequest interface{}) *userGetUnreadMessagesCountCall {
	return &userGetUnreadMessagesCountCall{Call: _m.Mock.On("GetUnreadMessagesCount", userID, getUnreadMessagesCountRequest), Parent: _m}
}

type userGetUnreadMessagesCountCall struct {
	*mock.Call
	Parent *userMock
}

func (_c *userGetUnreadMessagesCountCall) Panic(msg string) *userGetUnreadMessagesCountCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *userGetUnreadMessagesCountCall) Once() *userGetUnreadMessagesCountCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *userGetUnreadMessagesCountCall) Twice() *userGetUnreadMessagesCountCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *userGetUnreadMessagesCountCall) Times(i int) *userGetUnreadMessagesCountCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *userGetUnreadMessagesCountCall) WaitUntil(w <-chan time.Time) *userGetUnreadMessagesCountCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *userGetUnreadMessagesCountCall) After(d time.Duration) *userGetUnreadMessagesCountCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *userGetUnreadMessagesCountCall) Run(fn func(args mock.Arguments)) *userGetUnreadMessagesCountCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *userGetUnreadMessagesCountCall) Maybe() *userGetUnreadMessagesCountCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *userGetUnreadMessagesCountCall) TypedReturns(a *GetUnreadMessagesCountResponse, b error) *userGetUnreadMessagesCountCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *userGetUnreadMessagesCountCall) ReturnsFn(fn func(string, GetUnreadMessagesCountRequest) (*GetUnreadMessagesCountResponse, error)) *userGetUnreadMessagesCountCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *userGetUnreadMessagesCountCall) TypedRun(fn func(string, GetUnreadMessagesCountRequest)) *userGetUnreadMessagesCountCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_userID := args.String(0)
		_getUnreadMessagesCountRequest, _ := args.Get(1).(GetUnreadMessagesCountRequest)
		fn(_userID, _getUnreadMessagesCountRequest)
	})
	return _c
}

func (_c *userGetUnreadMessagesCountCall) OnCreateUser(createUserRequest CreateUserRequest) *userCreateUserCall {
	return _c.Parent.OnCreateUser(createUserRequest)
}

func (_c *userGetUnreadMessagesCountCall) OnGetGroupChannelCount(userID string, getGroupChannelCountRequest GetGroupChannelCountRequest) *userGetGroupChannelCountCall {
	return _c.Parent.OnGetGroupChannelCount(userID, getGroupChannelCountRequest)
}

func (_c *userGetUnreadMessagesCountCall) OnGetSessionToken(userID string, getSessionTokenRequest GetSessionTokenRequest) *userGetSessionTokenCall {
	return _c.Parent.OnGetSessionToken(userID, getSessionTokenRequest)
}

func (_c *userGetUnreadMessagesCountCall) OnGetUnreadMessagesCount(userID string, getUnreadMessagesCountRequest GetUnreadMessagesCountRequest) *userGetUnreadMessagesCountCall {
	return _c.Parent.OnGetUnreadMessagesCount(userID, getUnreadMessagesCountRequest)
}

func (_c *userGetUnreadMessagesCountCall) OnUpdateUser(userID string, updateUserRequest UpdateUserRequest) *userUpdateUserCall {
	return _c.Parent.OnUpdateUser(userID, updateUserRequest)
}

func (_c *userGetUnreadMessagesCountCall) OnCreateUserRaw(createUserRequest interface{}) *userCreateUserCall {
	return _c.Parent.OnCreateUserRaw(createUserRequest)
}

func (_c *userGetUnreadMessagesCountCall) OnGetGroupChannelCountRaw(userID interface{}, getGroupChannelCountRequest interface{}) *userGetGroupChannelCountCall {
	return _c.Parent.OnGetGroupChannelCountRaw(userID, getGroupChannelCountRequest)
}

func (_c *userGetUnreadMessagesCountCall) OnGetSessionTokenRaw(userID interface{}, getSessionTokenRequest interface{}) *userGetSessionTokenCall {
	return _c.Parent.OnGetSessionTokenRaw(userID, getSessionTokenRequest)
}

func (_c *userGetUnreadMessagesCountCall) OnGetUnreadMessagesCountRaw(userID interface{}, getUnreadMessagesCountRequest interface{}) *userGetUnreadMessagesCountCall {
	return _c.Parent.OnGetUnreadMessagesCountRaw(userID, getUnreadMessagesCountRequest)
}

func (_c *userGetUnreadMessagesCountCall) OnUpdateUserRaw(userID interface{}, updateUserRequest interface{}) *userUpdateUserCall {
	return _c.Parent.OnUpdateUserRaw(userID, updateUserRequest)
}

func (_m *userMock) UpdateUser(_ context.Context, userID string, updateUserRequest UpdateUserRequest) (*UpdateUserResponse, error) {
	_ret := _m.Called(userID, updateUserRequest)

	if _rf, ok := _ret.Get(0).(func(string, UpdateUserRequest) (*UpdateUserResponse, error)); ok {
		return _rf(userID, updateUserRequest)
	}

	_ra0, _ := _ret.Get(0).(*UpdateUserResponse)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *userMock) OnUpdateUser(userID string, updateUserRequest UpdateUserRequest) *userUpdateUserCall {
	return &userUpdateUserCall{Call: _m.Mock.On("UpdateUser", userID, updateUserRequest), Parent: _m}
}

func (_m *userMock) OnUpdateUserRaw(userID interface{}, updateUserRequest interface{}) *userUpdateUserCall {
	return &userUpdateUserCall{Call: _m.Mock.On("UpdateUser", userID, updateUserRequest), Parent: _m}
}

type userUpdateUserCall struct {
	*mock.Call
	Parent *userMock
}

func (_c *userUpdateUserCall) Panic(msg string) *userUpdateUserCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *userUpdateUserCall) Once() *userUpdateUserCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *userUpdateUserCall) Twice() *userUpdateUserCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *userUpdateUserCall) Times(i int) *userUpdateUserCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *userUpdateUserCall) WaitUntil(w <-chan time.Time) *userUpdateUserCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *userUpdateUserCall) After(d time.Duration) *userUpdateUserCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *userUpdateUserCall) Run(fn func(args mock.Arguments)) *userUpdateUserCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *userUpdateUserCall) Maybe() *userUpdateUserCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *userUpdateUserCall) TypedReturns(a *UpdateUserResponse, b error) *userUpdateUserCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *userUpdateUserCall) ReturnsFn(fn func(string, UpdateUserRequest) (*UpdateUserResponse, error)) *userUpdateUserCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *userUpdateUserCall) TypedRun(fn func(string, UpdateUserRequest)) *userUpdateUserCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_userID := args.String(0)
		_updateUserRequest, _ := args.Get(1).(UpdateUserRequest)
		fn(_userID, _updateUserRequest)
	})
	return _c
}

func (_c *userUpdateUserCall) OnCreateUser(createUserRequest CreateUserRequest) *userCreateUserCall {
	return _c.Parent.OnCreateUser(createUserRequest)
}

func (_c *userUpdateUserCall) OnGetGroupChannelCount(userID string, getGroupChannelCountRequest GetGroupChannelCountRequest) *userGetGroupChannelCountCall {
	return _c.Parent.OnGetGroupChannelCount(userID, getGroupChannelCountRequest)
}

func (_c *userUpdateUserCall) OnGetSessionToken(userID string, getSessionTokenRequest GetSessionTokenRequest) *userGetSessionTokenCall {
	return _c.Parent.OnGetSessionToken(userID, getSessionTokenRequest)
}

func (_c *userUpdateUserCall) OnGetUnreadMessagesCount(userID string, getUnreadMessagesCountRequest GetUnreadMessagesCountRequest) *userGetUnreadMessagesCountCall {
	return _c.Parent.OnGetUnreadMessagesCount(userID, getUnreadMessagesCountRequest)
}

func (_c *userUpdateUserCall) OnUpdateUser(userID string, updateUserRequest UpdateUserRequest) *userUpdateUserCall {
	return _c.Parent.OnUpdateUser(userID, updateUserRequest)
}

func (_c *userUpdateUserCall) OnCreateUserRaw(createUserRequest interface{}) *userCreateUserCall {
	return _c.Parent.OnCreateUserRaw(createUserRequest)
}

func (_c *userUpdateUserCall) OnGetGroupChannelCountRaw(userID interface{}, getGroupChannelCountRequest interface{}) *userGetGroupChannelCountCall {
	return _c.Parent.OnGetGroupChannelCountRaw(userID, getGroupChannelCountRequest)
}

func (_c *userUpdateUserCall) OnGetSessionTokenRaw(userID interface{}, getSessionTokenRequest interface{}) *userGetSessionTokenCall {
	return _c.Parent.OnGetSessionTokenRaw(userID, getSessionTokenRequest)
}

func (_c *userUpdateUserCall) OnGetUnreadMessagesCountRaw(userID interface{}, getUnreadMessagesCountRequest interface{}) *userGetUnreadMessagesCountCall {
	return _c.Parent.OnGetUnreadMessagesCountRaw(userID, getUnreadMessagesCountRequest)
}

func (_c *userUpdateUserCall) OnUpdateUserRaw(userID interface{}, updateUserRequest interface{}) *userUpdateUserCall {
	return _c.Parent.OnUpdateUserRaw(userID, updateUserRequest)
}
