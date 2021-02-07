// Code generated by MockGen. DO NOT EDIT.
// Source: session.go

// Package session_test is a generated GoMock package.
package session_test

import (
	context "context"
	reflect "reflect"

	client "github.com/Meat-Hook/point-bank/internal/modules/session/client"
	gomock "github.com/golang/mock/gomock"
)

// MocksessionSvc is a mock of sessionSvc interface
type MocksessionSvc struct {
	ctrl     *gomock.Controller
	recorder *MocksessionSvcMockRecorder
}

// MocksessionSvcMockRecorder is the mock recorder for MocksessionSvc
type MocksessionSvcMockRecorder struct {
	mock *MocksessionSvc
}

// NewMocksessionSvc creates a new mock instance
func NewMocksessionSvc(ctrl *gomock.Controller) *MocksessionSvc {
	mock := &MocksessionSvc{ctrl: ctrl}
	mock.recorder = &MocksessionSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocksessionSvc) EXPECT() *MocksessionSvcMockRecorder {
	return m.recorder
}

// Session mocks base method
func (m *MocksessionSvc) Session(ctx context.Context, token string) (*client.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Session", ctx, token)
	ret0, _ := ret[0].(*client.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Session indicates an expected call of Session
func (mr *MocksessionSvcMockRecorder) Session(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*MocksessionSvc)(nil).Session), ctx, token)
}
