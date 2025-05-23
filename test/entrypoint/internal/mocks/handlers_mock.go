// Code generated by MockGen. DO NOT EDIT.
// Source: ../../handlers/handlers.go
//
// Generated by this command:
//
//	mockgen -package=mocks -destination=internal/mocks/handlers_mock.go -source=../../handlers/handlers.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	syscall "syscall"

	handlers "github.com/k-lb/entrypoint-framework/handlers"
	gomock "go.uber.org/mock/gomock"
)

// MockActivationHandler is a mock of ActivationHandler interface.
type MockActivationHandler struct {
	ctrl     *gomock.Controller
	recorder *MockActivationHandlerMockRecorder
	isgomock struct{}
}

// MockActivationHandlerMockRecorder is the mock recorder for MockActivationHandler.
type MockActivationHandlerMockRecorder struct {
	mock *MockActivationHandler
}

// NewMockActivationHandler creates a new mock instance.
func NewMockActivationHandler(ctrl *gomock.Controller) *MockActivationHandler {
	mock := &MockActivationHandler{ctrl: ctrl}
	mock.recorder = &MockActivationHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActivationHandler) EXPECT() *MockActivationHandlerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockActivationHandler) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockActivationHandlerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockActivationHandler)(nil).Close))
}

// GetWasChangedChannel mocks base method.
func (m *MockActivationHandler) GetWasChangedChannel() <-chan handlers.ActivationEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWasChangedChannel")
	ret0, _ := ret[0].(<-chan handlers.ActivationEvent)
	return ret0
}

// GetWasChangedChannel indicates an expected call of GetWasChangedChannel.
func (mr *MockActivationHandlerMockRecorder) GetWasChangedChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWasChangedChannel", reflect.TypeOf((*MockActivationHandler)(nil).GetWasChangedChannel))
}

// MockConfigurationHandler is a mock of ConfigurationHandler interface.
type MockConfigurationHandler[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationHandlerMockRecorder[T]
	isgomock struct{}
}

// MockConfigurationHandlerMockRecorder is the mock recorder for MockConfigurationHandler.
type MockConfigurationHandlerMockRecorder[T any] struct {
	mock *MockConfigurationHandler[T]
}

// NewMockConfigurationHandler creates a new mock instance.
func NewMockConfigurationHandler[T any](ctrl *gomock.Controller) *MockConfigurationHandler[T] {
	mock := &MockConfigurationHandler[T]{ctrl: ctrl}
	mock.recorder = &MockConfigurationHandlerMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationHandler[T]) EXPECT() *MockConfigurationHandlerMockRecorder[T] {
	return m.recorder
}

// Close mocks base method.
func (m *MockConfigurationHandler[T]) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockConfigurationHandlerMockRecorder[T]) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConfigurationHandler[T])(nil).Close))
}

// GetUpdateResultChannel mocks base method.
func (m *MockConfigurationHandler[T]) GetUpdateResultChannel() <-chan T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateResultChannel")
	ret0, _ := ret[0].(<-chan T)
	return ret0
}

// GetUpdateResultChannel indicates an expected call of GetUpdateResultChannel.
func (mr *MockConfigurationHandlerMockRecorder[T]) GetUpdateResultChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateResultChannel", reflect.TypeOf((*MockConfigurationHandler[T])(nil).GetUpdateResultChannel))
}

// GetWasChangedChannel mocks base method.
func (m *MockConfigurationHandler[T]) GetWasChangedChannel() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWasChangedChannel")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// GetWasChangedChannel indicates an expected call of GetWasChangedChannel.
func (mr *MockConfigurationHandlerMockRecorder[T]) GetWasChangedChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWasChangedChannel", reflect.TypeOf((*MockConfigurationHandler[T])(nil).GetWasChangedChannel))
}

// Update mocks base method.
func (m *MockConfigurationHandler[T]) Update() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update")
}

// Update indicates an expected call of Update.
func (mr *MockConfigurationHandlerMockRecorder[T]) Update() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockConfigurationHandler[T])(nil).Update))
}

// MockProcessHandler is a mock of ProcessHandler interface.
type MockProcessHandler struct {
	ctrl     *gomock.Controller
	recorder *MockProcessHandlerMockRecorder
	isgomock struct{}
}

// MockProcessHandlerMockRecorder is the mock recorder for MockProcessHandler.
type MockProcessHandlerMockRecorder struct {
	mock *MockProcessHandler
}

// NewMockProcessHandler creates a new mock instance.
func NewMockProcessHandler(ctrl *gomock.Controller) *MockProcessHandler {
	mock := &MockProcessHandler{ctrl: ctrl}
	mock.recorder = &MockProcessHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessHandler) EXPECT() *MockProcessHandlerMockRecorder {
	return m.recorder
}

// GetEndedChannel mocks base method.
func (m *MockProcessHandler) GetEndedChannel() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEndedChannel")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// GetEndedChannel indicates an expected call of GetEndedChannel.
func (mr *MockProcessHandlerMockRecorder) GetEndedChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEndedChannel", reflect.TypeOf((*MockProcessHandler)(nil).GetEndedChannel))
}

// GetStartedChannel mocks base method.
func (m *MockProcessHandler) GetStartedChannel() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStartedChannel")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// GetStartedChannel indicates an expected call of GetStartedChannel.
func (mr *MockProcessHandlerMockRecorder) GetStartedChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStartedChannel", reflect.TypeOf((*MockProcessHandler)(nil).GetStartedChannel))
}

// Kill mocks base method.
func (m *MockProcessHandler) Kill() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kill")
	ret0, _ := ret[0].(error)
	return ret0
}

// Kill indicates an expected call of Kill.
func (mr *MockProcessHandlerMockRecorder) Kill() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockProcessHandler)(nil).Kill))
}

// Signal mocks base method.
func (m *MockProcessHandler) Signal(arg0 syscall.Signal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signal", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal.
func (mr *MockProcessHandlerMockRecorder) Signal(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockProcessHandler)(nil).Signal), arg0)
}

// Start mocks base method.
func (m *MockProcessHandler) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockProcessHandlerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockProcessHandler)(nil).Start))
}

// Stop mocks base method.
func (m *MockProcessHandler) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockProcessHandlerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockProcessHandler)(nil).Stop))
}
