// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package ssh is a generated GoMock package.
package ssh

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

// ExecuteRemoteCommand mocks base method
func (m *MockService) ExecuteRemoteCommand(ctx context.Context, chatId uint32, commandName, address string) error {
	ret := m.ctrl.Call(m, "ExecuteRemoteCommand", ctx, chatId, commandName, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteRemoteCommand indicates an expected call of ExecuteRemoteCommand
func (mr *MockServiceMockRecorder) ExecuteRemoteCommand(ctx, chatId, commandName, address interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteRemoteCommand", reflect.TypeOf((*MockService)(nil).ExecuteRemoteCommand), ctx, chatId, commandName, address)
}

// addCommand mocks base method
func (m *MockService) addCommand(chatId uint32, name, command string) error {
	ret := m.ctrl.Call(m, "addCommand", chatId, name, command)
	ret0, _ := ret[0].(error)
	return ret0
}

// addCommand indicates an expected call of addCommand
func (mr *MockServiceMockRecorder) addCommand(chatId, name, command interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addCommand", reflect.TypeOf((*MockService)(nil).addCommand), chatId, name, command)
}

// addKey mocks base method
func (m *MockService) addKey(chatId uint32, userName, base64Key string) error {
	ret := m.ctrl.Call(m, "addKey", chatId, userName, base64Key)
	ret0, _ := ret[0].(error)
	return ret0
}

// addKey indicates an expected call of addKey
func (mr *MockServiceMockRecorder) addKey(chatId, userName, base64Key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addKey", reflect.TypeOf((*MockService)(nil).addKey), chatId, userName, base64Key)
}

// addServer mocks base method
func (m *MockService) addServer(chatId uint32, address, description string) error {
	ret := m.ctrl.Call(m, "addServer", chatId, address, description)
	ret0, _ := ret[0].(error)
	return ret0
}

// addServer indicates an expected call of addServer
func (mr *MockServiceMockRecorder) addServer(chatId, address, description interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addServer", reflect.TypeOf((*MockService)(nil).addServer), chatId, address, description)
}
