// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/account.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	object "yatter-backend-go/app/domain/object"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockAccount is a mock of Account interface.
type MockAccount struct {
	ctrl     *gomock.Controller
	recorder *MockAccountMockRecorder
}

// MockAccountMockRecorder is the mock recorder for MockAccount.
type MockAccountMockRecorder struct {
	mock *MockAccount
}

// NewMockAccount creates a new mock instance.
func NewMockAccount(ctrl *gomock.Controller) *MockAccount {
	mock := &MockAccount{ctrl: ctrl}
	mock.recorder = &MockAccountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccount) EXPECT() *MockAccountMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccount) Create(ctx context.Context, tx *sqlx.Tx, acc *object.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, tx, acc)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAccountMockRecorder) Create(ctx, tx, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccount)(nil).Create), ctx, tx, acc)
}

// FindByUsername mocks base method.
func (m *MockAccount) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUsername", ctx, username)
	ret0, _ := ret[0].(*object.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUsername indicates an expected call of FindByUsername.
func (mr *MockAccountMockRecorder) FindByUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockAccount)(nil).FindByUsername), ctx, username)
}
