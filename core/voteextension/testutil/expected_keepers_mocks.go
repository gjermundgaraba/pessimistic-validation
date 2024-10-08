// Code generated by MockGen. DO NOT EDIT.
// Source: voteextension/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	exported "github.com/cosmos/ibc-go/v9/modules/core/exported"
	gomock "github.com/golang/mock/gomock"
)

// MockClientKeeper is a mock of ClientKeeper interface.
type MockClientKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockClientKeeperMockRecorder
}

// MockClientKeeperMockRecorder is the mock recorder for MockClientKeeper.
type MockClientKeeperMockRecorder struct {
	mock *MockClientKeeper
}

// NewMockClientKeeper creates a new mock instance.
func NewMockClientKeeper(ctrl *gomock.Controller) *MockClientKeeper {
	mock := &MockClientKeeper{ctrl: ctrl}
	mock.recorder = &MockClientKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientKeeper) EXPECT() *MockClientKeeperMockRecorder {
	return m.recorder
}

// UpdateClient mocks base method.
func (m *MockClientKeeper) UpdateClient(ctx types.Context, clientID string, clientMsg exported.ClientMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClient", ctx, clientID, clientMsg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClient indicates an expected call of UpdateClient.
func (mr *MockClientKeeperMockRecorder) UpdateClient(ctx, clientID, clientMsg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClient", reflect.TypeOf((*MockClientKeeper)(nil).UpdateClient), ctx, clientID, clientMsg)
}
