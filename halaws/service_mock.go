// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_halaws is a generated GoMock package.
package halaws

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// SendAlert mocks base method
func (m *MockService) SendAlert(ctx context.Context, destination, name string) {
	m.ctrl.Call(m, "SendAlert", ctx, destination, name)
}

// SendAlert indicates an expected call of SendAlert
func (mr *MockServiceMockRecorder) SendAlert(ctx, destination, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAlert", reflect.TypeOf((*MockService)(nil).SendAlert), ctx, destination, name)
}
