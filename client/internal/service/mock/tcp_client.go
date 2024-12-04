// Code generated by MockGen. DO NOT EDIT.
// Source: client/internal/service (interfaces: TCPClient)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTCPClient is a mock of TCPClient interface.
type MockTCPClient struct {
	ctrl     *gomock.Controller
	recorder *MockTCPClientMockRecorder
}

// MockTCPClientMockRecorder is the mock recorder for MockTCPClient.
type MockTCPClientMockRecorder struct {
	mock *MockTCPClient
}

// NewMockTCPClient creates a new mock instance.
func NewMockTCPClient(ctrl *gomock.Controller) *MockTCPClient {
	mock := &MockTCPClient{ctrl: ctrl}
	mock.recorder = &MockTCPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTCPClient) EXPECT() *MockTCPClientMockRecorder {
	return m.recorder
}

// GetChallenge mocks base method.
func (m *MockTCPClient) GetChallenge(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChallenge", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChallenge indicates an expected call of GetChallenge.
func (mr *MockTCPClientMockRecorder) GetChallenge(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChallenge", reflect.TypeOf((*MockTCPClient)(nil).GetChallenge), arg0, arg1)
}

// GetQuote mocks base method.
func (m *MockTCPClient) GetQuote(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuote", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuote indicates an expected call of GetQuote.
func (mr *MockTCPClientMockRecorder) GetQuote(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuote", reflect.TypeOf((*MockTCPClient)(nil).GetQuote), arg0)
}

// GetTokenBySolution mocks base method.
func (m *MockTCPClient) GetTokenBySolution(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenBySolution", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenBySolution indicates an expected call of GetTokenBySolution.
func (mr *MockTCPClientMockRecorder) GetTokenBySolution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenBySolution", reflect.TypeOf((*MockTCPClient)(nil).GetTokenBySolution), arg0)
}
