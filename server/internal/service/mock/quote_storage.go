// Code generated by MockGen. DO NOT EDIT.
// Source: server/internal/service (interfaces: QuoteStorage)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockQuoteStorage is a mock of QuoteStorage interface.
type MockQuoteStorage struct {
	ctrl     *gomock.Controller
	recorder *MockQuoteStorageMockRecorder
}

// MockQuoteStorageMockRecorder is the mock recorder for MockQuoteStorage.
type MockQuoteStorageMockRecorder struct {
	mock *MockQuoteStorage
}

// NewMockQuoteStorage creates a new mock instance.
func NewMockQuoteStorage(ctrl *gomock.Controller) *MockQuoteStorage {
	mock := &MockQuoteStorage{ctrl: ctrl}
	mock.recorder = &MockQuoteStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuoteStorage) EXPECT() *MockQuoteStorageMockRecorder {
	return m.recorder
}

// Quote mocks base method.
func (m *MockQuoteStorage) Quote() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quote")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Quote indicates an expected call of Quote.
func (mr *MockQuoteStorageMockRecorder) Quote() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quote", reflect.TypeOf((*MockQuoteStorage)(nil).Quote))
}
